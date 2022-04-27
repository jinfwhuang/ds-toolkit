package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/pkg/bytesutil"
	cmd_utils "github.com/jinfwhuang/ds-toolkit/pkg/cmd-utils"
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


var AppFlags = []cli.Flag{
	cmd_utils.GrpcPort,
	cmd_utils.LogLevel,
	cmd_utils.LogCaller,
}

func main() {
	app := cli.App{}
	//app.Name = "xxx"
	//app.Usage = "yyy"
	app.Action = start

	app.Flags = AppFlags

	if err := app.Run(os.Args); err != nil {
		logrus.Info(err)
	}
}



func testLogrus() {
	logrus.Debug("debuge")
	logrus.Info("abbc")
	logrus.Warn("abbc")
}


//type MainApp struct {
//	cliCtx   *cli.Context
//	ctx      context.Context
//	cancel   context.CancelFunc
//	lock     sync.RWMutex
//	stop     chan struct{} // Channel to wait for termination notifications.
//	app *cli.App,
//}


func start(cliCtx *cli.Context) error {
	cmd_utils.SetupLogrus(cliCtx)

	stop := make(chan struct{})

	grpcPort := cliCtx.Int(cmd_utils.GrpcPort.Name)
	address := fmt.Sprintf("%s:%d", "127.0.0.1", grpcPort)
	logrus.Info("grpc addres: %s", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(fmt.Errorf("could not listen to port in Start() %s: %v", address, err))
	}

	// GRPC Server
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	// Register endpoints
	server := NewServer()
	protoId.RegisterIdentityServer(grpcServer, server)

	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(fmt.Errorf("could not serve gRPC: %v", err))
		}
	}()

	logrus.Info("Wait for stop channel to be closed.")
	<-stop
	return nil
}

type Server struct {
	protoId.UnsafeIdentityServer

	// Store temporary login msg, key by address
	// Login msg has to be invalidated as soon as it is used
	loginMsgStore map[string]string

	mutexLock sync.Mutex
}

func NewServer() *Server {
	return &Server{
		loginMsgStore: make(map[string]string),
	}
}

func (s *Server) RequestLogin(ctx context.Context, loginMsg *protoId.LoginMessage) (*protoId.LoginMessage, error) {
	logrus.Info(loginMsg)
	if loginMsg.PubKey == nil {
		return nil, fmt.Errorf("pub key is not provided")
	}

	b64 := base64.StdEncoding.EncodeToString(bytesutil.RandBytes(17))
	t := time.Now().UnixMilli()
	msgFmt := "Sign this message to prove you have access to the public key. \n" +
		"Random string: %s \n" +
		"Timestamp: %d"
	msg := fmt.Sprintf(msgFmt, b64, t)

	pubkey := base64.StdEncoding.EncodeToString(loginMsg.PubKey)
	s.mutexLock.Lock()
	s.loginMsgStore[pubkey] = msg
	s.mutexLock.Unlock()

	return &protoId.LoginMessage{
		PubKey: loginMsg.PubKey,
		UnsignedMsg: msg,
	}, nil
}

//func validateMsg(msg *protoId.LoginMessage) (bool, error) {
//	//curve := secp256k1.S256()
//
//	pubkey, err := crypto.UnmarshalPubkey(msg.PubKey)
//	if err != nil {
//		return false, err
//	}
//
//
//
//
//
//}

const (
	SignatureLength = 65
)

func (s *Server) Login(ctx context.Context, msg *protoId.LoginMessage) (*protoId.LoginResp, error) {
	pubkey := msg.PubKey
	pubkeyStr := base64.StdEncoding.EncodeToString(pubkey)

	// Use unsigned msg stored on server
	unSignMsg := []byte(s.loginMsgStore[pubkeyStr])
	msgHash := crypto.Keccak256Hash(unSignMsg)
	digestHash := msgHash.Bytes()

	// Validate the signature
	sigWithoutID := msg.Signature[:SignatureLength-1] // remove recovery id
	validated := crypto.VerifySignature(pubkey, digestHash, sigWithoutID)
	status := "failed"
	if validated {
		status = "ok"
		// Remove the message from store
		delete(s.loginMsgStore, pubkeyStr)
	}

	return &protoId.LoginResp{
		PubKey: msg.PubKey,
		Status: status,
	}, nil
}

func (s *Server) Debug(context.Context, *emptypb.Empty) (*protoId.LoginMessage, error) {
	pubKeyStr := "4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12"
	pubKey, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		panic(nil)
	}
	logrus.Info(pubKey, len(pubKey))

	return &protoId.LoginMessage{
		PubKey: pubKey,
	}, nil
}


