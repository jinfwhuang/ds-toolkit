package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/bytesutil"
	cmd_utils "github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	protoId "github.com/jinfwhuang/ds-toolkit/proto/identity"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"os"
	"sync"
	"time"
)

const (
	SignatureLength = 65
	SignInMsgFmt = "Sign this message to prove you have access to the public key. " +
		"pubkey=%s rand-string=%s timestamp= %d"

)

var AppFlags = []cli.Flag{
	cmd_utils.GrpcPort,
	cmd_utils.LogLevel,
	cmd_utils.LogCaller,
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
	server := NewIdServer()
	protoId.RegisterIdentityServer(grpcServer, server)
	reflection.Register(grpcServer)  // Enable reflection
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

	// Store temporary login msg, key by address
	// Login msg has to be invalidated as soon as it is used
	loginMsgStore map[string]string

	mutexLock sync.Mutex
}

func NewIdServer() *IdServer {
	return &IdServer{
		loginMsgStore: make(map[string]string),
	}
}

func GenerateSignInMessage(key []byte) string {
	keyB16 := base64.StdEncoding.EncodeToString(key)
	randStr := base64.StdEncoding.EncodeToString(bytesutil.RandBytes(17))
	t := time.Now().UnixMilli()
	return fmt.Sprintf(SignInMsgFmt, keyB16, randStr, t)
}

func (s *IdServer) RequestLogin(ctx context.Context, loginMsg *protoId.LoginMessage) (*protoId.LoginMessage, error) {
	if loginMsg.PubKey == nil {
		return nil, fmt.Errorf("pub key is not provided")
	}
	signInMsg := GenerateSignInMessage(loginMsg.PubKey)
	pubkey := base64.StdEncoding.EncodeToString(loginMsg.PubKey)

	// Keep the unsigned message
	s.mutexLock.Lock()
	s.loginMsgStore[pubkey] = signInMsg
	s.mutexLock.Unlock()

	return &protoId.LoginMessage{
		PubKey:      loginMsg.PubKey,
		UnsignedMsg: signInMsg,
	}, nil
}

func (s *IdServer) Login(ctx context.Context, msg *protoId.LoginMessage) (*protoId.LoginResp, error) {
	pubkey := msg.PubKey
	addr, err := ecdsa_util.ToAddress(pubkey)
	if err != nil {
		return nil, err
	}
	pubkeyStr := base64.StdEncoding.EncodeToString(pubkey)

	// Retrieve the unsigned message
	unSignMsg := []byte(s.loginMsgStore[pubkeyStr])
	msgHash := crypto.Keccak256Hash(unSignMsg)
	digestHash := msgHash.Bytes()

	// Validate the signature
	sigWithoutID := msg.Signature[:SignatureLength-1] // remove recovery id
	validated := crypto.VerifySignature(pubkey, digestHash, sigWithoutID)
	status := "failed"
	if validated {
		status = "ok"
		delete(s.loginMsgStore, pubkeyStr) // Remove the message from store

		logrus.Infof("login successful: address=%s", addr.Hex())
	}

	return &protoId.LoginResp{
		PubKey: msg.PubKey,
		Status: status,
	}, nil
}

func (s *IdServer) Debug(context.Context, *emptypb.Empty) (*protoId.LoginMessage, error) {
	pubKeyStr := "4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12"
	pubKey, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		logrus.Fatal(err)
	}
	signInMsg := GenerateSignInMessage(pubKey)

	logrus.Info("pubkey=", pubKey, "len=", len(pubKey))
	logrus.Info("sign-in-message=", signInMsg)

	return &protoId.LoginMessage{
		PubKey: pubKey,
		UnsignedMsg: signInMsg,
	}, nil
}
