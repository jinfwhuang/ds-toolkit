package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	cmd_utils "github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	protoId "github.com/jinfwhuang/ds-toolkit/proto/identity"
)

var (
	appFlags = []cli.Flag{
		cmd_utils.GrpcPort,
		cmd_utils.LogLevel,
	}
)

const (
	// secp256k1 keys
	testPrivkeyHex  = "0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
	testUserName    = "jinhuang001"
	testPrivkeyHex2 = "0x89df41bce61452e4bed3e6325873022ba877a1683cfca9d03c0a5bf542944321"
	testUserName2   = "stella021"
	testUserName3   = "user3"
	testPrivkeyHex3 = "0xa329248cb59ccd5429a67e11529a831147bf1193bafcc45255129288362095cb"
)

func main() {
	app := cli.App{}
	app.Name = "UserRegistryLoginClient"
	app.Flags = appFlags
	app.Action = start
	if err := app.Run(os.Args); err != nil {
		logrus.Info(err)
	}
}

func start(cliCtx *cli.Context) error {
	cmd_utils.SetupLogrus(cliCtx)

	ctx := context.Background()
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	grpcPort := cliCtx.Int(cmd_utils.GrpcPort.Name)
	addr := fmt.Sprintf("%s:%d", "localhost", grpcPort)
	logrus.Info("connecting to grpc server: %s", addr)
	conn, err := grpc.DialContext(ctx, addr, dialOpts...)
	if err != nil {
		return err
	}

	userRegistryClient := protoId.NewUserRegistryLoginClient(conn)

	if err := listAllUsers(ctx, userRegistryClient); err != nil {
		return err
	}

	return nil
}

func listAllUsers(ctx context.Context, userRegistryClient protoId.UserRegistryLoginClient) error {
	users, err := userRegistryClient.ListAllUsers(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}
	for i, u := range users.Users {
		logrus.Printf("User %d: name=%s, pubkey=0x%x\n", i, u.UserName, u.PubKey)
	}
	return nil
}
