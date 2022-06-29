package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	//eth_client "github.com/jinfwhuang/ds-toolkit/experimental/eth-client"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/bytesutil"
	cmd_utils "github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	eth_client "github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	protoId "github.com/jinfwhuang/ds-toolkit/proto/identity"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	SignatureLength = 65
	SignInMsgFmt    = "Sign this message to prove you have access to the public key. " +
		"pubkey=%s rand-string=%s timestamp= %d"
)

var ethAddr = &cli.StringFlag{
	Name:  "eth-address",
	Usage: "TODO: xxx",
	Value: "http://127.0.0.1:8545",
}

var contractAddr = &cli.StringFlag{
	Name:  "contract-address",
	Usage: "TODO: xxx",
	Value: "0xC26f4289DFB6138C5f12e7D52D283F0Ce15FF985",
}

var AppFlags = []cli.Flag{
	cmd_utils.GrpcPort,
	cmd_utils.LogLevel,
	cmd_utils.LogCaller,
	ethAddr,
	contractAddr,
}

var callOpt = &bind.CallOpts{
	Pending: true,
}

func main() {
	app := cli.App{}
	app.Name = "Identity-IdServer"
	app.Flags = AppFlags
	app.Action = start
	if err := app.Run(os.Args); err != nil {
		logrus.Info(err)
	}
}

func start(cliCtx *cli.Context) error {
	cmd_utils.SetupLogrus(cliCtx)

	stop := make(chan struct{})

	// Tcp Listener
	grpcPort := cliCtx.Int(cmd_utils.GrpcPort.Name)
	address := fmt.Sprintf("%s:%d", "0.0.0.0", grpcPort)
	logrus.Infof("grpc address: %s", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatalf("could not listen to port in Start() %s: %v", address, err)
	}
	// GRPC
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	idServer := NewIdServer()
	userRegistryServer := NewUserRegistryServer(ethAddr.GetValue(), contractAddr.GetValue())
	protoId.RegisterIdentityServer(grpcServer, idServer)
	protoId.RegisterUserRegistryLoginServer(grpcServer, userRegistryServer)
	reflection.Register(grpcServer) // Enable reflection
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalf("could not serve gRPC: %v", err)
		}
	}()

	logrus.Info("wait for stop channel to be closed.")
	<-stop
	return nil
}

type IdServer struct {
	protoId.UnsafeIdentityServer
	//protoId defined above
	// Store temporary login msg, key by address
	// Login msg has to be invalidated as soon as it is used
	loginMsgStore map[string]string

	mutexLock sync.Mutex
}

type UserRegistryServer struct {
	protoId.UnimplementedUserRegistryLoginServer //unimplemented vs unsafe
	loginMsgStore                                map[string]string
	//userName                                     string
	mutexLock sync.Mutex //what is this doing?

	userRegistry *eth_client.UserRegistry
}

func NewIdServer() *IdServer {
	return &IdServer{
		loginMsgStore: make(map[string]string), //what is this v.s. whats defined in message?
	}
}

func NewUserRegistryServer(ethAddr, contractAddr string) *UserRegistryServer {
	ethconn, err := ethclient.Dial(ethAddr)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	userRegistryAcc := common.HexToAddress(contractAddr)
	userRegistry, err := eth_client.NewUserRegistry(userRegistryAcc, ethconn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	return &UserRegistryServer{
		loginMsgStore: make(map[string]string),
		userRegistry:  userRegistry,
	}
}

func GenerateSignInMessage(key []byte) string {
	keyB16 := base64.StdEncoding.EncodeToString(key)
	randStr := base64.StdEncoding.EncodeToString(bytesutil.RandBytes(17))
	t := time.Now().UnixMilli()
	return fmt.Sprintf(SignInMsgFmt, keyB16, randStr, t)
}

func (s *UserRegistryServer) RequestLogin(ctx context.Context, userLogin *protoId.UserLogin) (*protoId.UserLogin, error) {
	userName := userLogin.UserName
	userAddr, err := s.userRegistry.GetUser(callOpt, userName)
	logrus.Info("USER ADDRESS:", userAddr)
	if err != nil {
		log.Fatalf("Username not registered")
	}
	//retrieves the first registered key
	ethPubKey, err := s.userRegistry.GetKey(callOpt, userAddr, uint8(userLogin.PubKeyId))
	//length, err := s.userRegistry.GetLenKeys(callOpt, userAddr)
	//logrus.Info("Key ", ethPubKey.Key)

	if err != nil {
		log.Fatalf("No keys registered")
	}

	/*
		if userLogin.PubKey == nil {
			return nil, fmt.Errorf("pub key is not provided")
		}
		signInMsg := GenerateSignInMessage(userLogin.PubKey)
		pubkey := base64.StdEncoding.EncodeToString(userLogin.PubKey)
	*/

	// Keep the unsigned message
	signInMsg := GenerateSignInMessage(ethPubKey.Key)
	pubKey := base64.StdEncoding.EncodeToString(ethPubKey.Key)
	s.mutexLock.Lock()
	s.loginMsgStore[pubKey] = signInMsg //retrieving data from database
	s.mutexLock.Unlock()
	//logrus.Info("pubkey:", pubKey, "username:", userName, []byte(pubKey))

	return &protoId.UserLogin{
		UserName:    userName,
		PubKey:      ethPubKey.Key,
		UnsignedMsg: signInMsg,
	}, nil
}

func (s *UserRegistryServer) Login(ctx context.Context, msg *protoId.UserLogin) (*protoId.LoginResp, error) {
	pubkey := msg.PubKey
	logrus.Info("PUBLICKEY: ", msg.PubKey)
	addr, err := ecdsa_util.ToAddress(pubkey) //invalid secp public key
	//logrus.Infof("hihihi", addr, err, "MSG.PUBKEY", msg.PubKey)
	if err != nil {
		return nil, err
	}
	pubkeyStr := base64.StdEncoding.EncodeToString(pubkey)
	// Retrieve the unsigned message
	unSignMsg := []byte(s.loginMsgStore[pubkeyStr])
	msgHash := crypto.Keccak256Hash(unSignMsg)
	digestHash := msgHash.Bytes()
	logrus.Info("PUBKEYSTR, UNSIGNMSG, msgHASH, digestHASH: ", pubkeyStr, unSignMsg, msgHash, digestHash)

	// Validate the signature
	sigWithoutID := msg.Signature[:SignatureLength-1] // remove recovery id
	validated := crypto.VerifySignature(pubkey, digestHash, sigWithoutID)
	status := "failed"
	if validated {
		status = "ok"
		delete(s.loginMsgStore, pubkeyStr) // Remove the message from store
		//defined in IdStruct above
		logrus.Infof("login successful: address=%s", addr.Hex())
	}

	return &protoId.LoginResp{
		PubKey: msg.PubKey,
		Status: status,
	}, nil
}

func (s *IdServer) RequestLogin(ctx context.Context, loginMsg *protoId.LoginMessage) (*protoId.LoginMessage, error) {
	//from main_test.go
	//utilizing smart contract for UserRegistry data
	if loginMsg.PubKey == nil {
		return nil, fmt.Errorf("pub key is not provided")
	}
	logrus.Info("PubKEy request login: ", loginMsg.PubKey)
	signInMsg := GenerateSignInMessage(loginMsg.PubKey)
	pubkey := base64.StdEncoding.EncodeToString(loginMsg.PubKey)

	// Keep the unsigned message
	s.mutexLock.Lock()
	s.loginMsgStore[pubkey] = signInMsg //retrieving data from database
	s.mutexLock.Unlock()

	return &protoId.LoginMessage{
		PubKey:      loginMsg.PubKey,
		UnsignedMsg: signInMsg,
	}, nil
}

func (s *IdServer) Login(ctx context.Context, msg *protoId.LoginMessage) (*protoId.LoginResp, error) {
	pubkey := msg.PubKey
	logrus.Info("PUBLICKEY: ", msg.PubKey)
	addr, err := ecdsa_util.ToAddress(pubkey)
	if err != nil {
		return nil, err
	}
	pubkeyStr := base64.StdEncoding.EncodeToString(pubkey)

	// Retrieve the unsigned message
	unSignMsg := []byte(s.loginMsgStore[pubkeyStr])
	msgHash := crypto.Keccak256Hash(unSignMsg)
	digestHash := msgHash.Bytes()
	logrus.Info("CORRECT PUBKEYSTR, UNSIGNMSG, msgHASH, digestHASH: ", pubkeyStr, unSignMsg, msgHash, digestHash)

	// Validate the signature
	sigWithoutID := msg.Signature[:SignatureLength-1] // remove recovery id
	logrus.Info("CORRECT SIG W/O ID:", sigWithoutID)
	validated := crypto.VerifySignature(pubkey, digestHash, sigWithoutID)
	status := "failed"
	if validated {
		status = "ok"
		delete(s.loginMsgStore, pubkeyStr) // Remove the message from store
		//defined in IdStruct above
		logrus.Infof("login successful: address=%s", addr.Hex())
	}

	return &protoId.LoginResp{
		PubKey: msg.PubKey,
		Status: status,
	}, nil
}

func (s *IdServer) Debug(context.Context, *emptypb.Empty) (*protoId.LoginMessage, error) {

	//logrus.Info("here")

	pubKeyStr := "4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12"
	pubKey, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		logrus.Fatal(err)
	}
	signInMsg := GenerateSignInMessage(pubKey)

	logrus.Info("pubkey=", pubKey, "len=", len(pubKey))
	logrus.Info("sign-in-message=", signInMsg)

	return &protoId.LoginMessage{
		PubKey:      pubKey,
		UnsignedMsg: signInMsg,
	}, nil
}
