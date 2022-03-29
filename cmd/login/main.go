package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/jinfwhuang/ds-sdk/pkg/bytesutil"
	"github.com/jinfwhuang/ds-sdk/pkg/ds"
	protoId "github.com/jinfwhuang/ds-sdk/proto/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	tmplog "log"
	"net"
	"os"
	"time"
	"github.com/urfave/cli/v2"
)

func init() {
	tmplog.SetFlags(tmplog.Llongfile)
}

var (
	GrpcPort = &cli.IntFlag{
		Name:  "grpc-port",
		Usage: "TODO: xxx",
		Value: 4000,
	}
)

var AppFlags = []cli.Flag{
	GrpcPort,
}

func main() {
	app := cli.App{}
	app.Name = "beacon-chain-light-client"
	app.Usage = "Beacon Chain Light Client"
	app.Action = start

	app.Flags = AppFlags

	if err := app.Run(os.Args); err != nil {
		tmplog.Println(err)
	}
}

func main2() {
	ds.Put()

	//StartServer()

	tmplog.Println("finish debug main")
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
	stop := make(chan struct{})

	grpcPort := cliCtx.Int(GrpcPort.Name)
	address := fmt.Sprintf("%s:%d", "127.0.0.1", grpcPort)
	tmplog.Printf("grpc addres: %s", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(fmt.Errorf("could not listen to port in Start() %s: %v", address, err))
	}

	// GRPC Server
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	// Register endpoints
	protoId.RegisterIdentityServer(grpcServer, &Server{})

	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(fmt.Errorf("could not serve gRPC: %v", err))
		}
	}()

	tmplog.Println("Wait for stop channel to be closed.")
	<-stop
	return nil
}

type Server struct {
	protoId.UnsafeIdentityServer
}

func (s *Server) RequestLogin(ctx context.Context, loginMsg *protoId.LoginMessage) (*protoId.LoginMessage, error) {
	tmplog.Println(loginMsg)

	if loginMsg.PubKey == nil {
		return nil, fmt.Errorf("pub key is not provided")
	}

	b := bytesutil.RandBytes(17)
	b64 := base64.StdEncoding.EncodeToString(b)
	t := time.Now().UnixMilli()

	msgFmt := "Sign this message to prove you have access to this wallet and we will sign you in. \n" +
		"This won't cost you any crypto. \n " +
		"Random string: %s" +
		"Timestamp: %d"
	msg := fmt.Sprintf(msgFmt, b64, t)

	return &protoId.LoginMessage{
		PubKey: loginMsg.PubKey,
		UnsignedMsg: msg,
	}, nil
}

func (s *Server) Login(context.Context, *protoId.LoginMessage) (*protoId.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (s *Server) Debug(context.Context, *emptypb.Empty) (*protoId.LoginMessage, error) {
	pubKeyStr := "4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12"
	pubKey, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		panic(nil)
	}
	tmplog.Println(pubKey, len(pubKey))

	return &protoId.LoginMessage{
		PubKey: pubKey,
	}, nil
}


