package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	cmd_utils "github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	eth_client "github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
	protoId "github.com/jinfwhuang/ds-toolkit/proto/identity"
)

var (
	ethAddr = &cli.StringFlag{
		Name:  "eth-address",
		Usage: "TODO: xxx",
		Value: "http://127.0.0.1:8545",
	}

	contractAddr = &cli.StringFlag{
		Name:  "contract-address",
		Usage: "TODO: xxx",
		Value: "0xC26f4289DFB6138C5f12e7D52D283F0Ce15FF985",
	}

	appFlags = []cli.Flag{
		cmd_utils.GrpcPort,
		cmd_utils.LogLevel,
		cmd_utils.LogCaller,
		ethAddr,
		contractAddr,
	}

	callOpt = &bind.CallOpts{
		Pending: true,
	}
)

func main() {
	app := cli.App{}
	app.Name = "UserRegistryLogin"
	app.Flags = appFlags
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
	grpcServer := grpc.NewServer()
	userRegistryServer := NewUserRegistryServer(ethAddr.GetValue(), contractAddr.GetValue())
	protoId.RegisterUserRegistryLoginServer(grpcServer, userRegistryServer)
	reflection.Register(grpcServer) // Enable reflection
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	logrus.Info("wait for stop channel to be closed.")
	<-stop
	return nil
}

type UserRegistryServer struct {
	protoId.UnimplementedUserRegistryLoginServer

	userRegistry *eth_client.UserRegistry
}

func NewUserRegistryServer(ethAddr, contractAddr string) *UserRegistryServer {
	ethconn, err := ethclient.Dial(ethAddr)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	userRegistryAcc := common.HexToAddress(contractAddr)
	userRegistry, err := eth_client.NewUserRegistry(userRegistryAcc, ethconn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	return &UserRegistryServer{
		userRegistry: userRegistry,
	}
}

func (s *UserRegistryServer) ListAllUsers(context.Context, *emptypb.Empty) (*protoId.UserList, error) {
	// Gathers all users by userID
	addrs, err := s.userRegistry.GetAllUsers(callOpt)
	if err != nil {
		log.Fatalf("Failed to retrieve all users")
	}

	userList := make([]*protoId.User, len(addrs))
	for i, addr := range addrs {
		// Retrieving username with userId
		userName, err := s.userRegistry.GetName(callOpt, addr)
		if err != nil {
			log.Fatalf("Failed to find user with userId")
		}
		// Retrieving first key with userId
		pubKey, err := s.userRegistry.GetKey(callOpt, addr, 0)
		if err != nil {
			log.Fatalf("Failed to find keys with userId")
		}

		userList[i] = &protoId.User{
			UserName: userName,
			PubKey:   pubKey.Key,
		}
	}

	return &protoId.UserList{
		Users: userList,
	}, nil
}
