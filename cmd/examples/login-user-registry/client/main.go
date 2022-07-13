package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	cmd_utils "github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	protoId "github.com/jinfwhuang/ds-toolkit/proto/identity"
)

var (
	userName = &cli.StringFlag{
		Name:  "user-name",
		Usage: "TODO: xxx",
		Value: "",
	}

	privKey = &cli.StringFlag{
		Name:  "priv-key",
		Usage: "TODO: xxx",
		Value: "",
	}

	appFlags = []cli.Flag{
		cmd_utils.GrpcPort,
		cmd_utils.LogLevel,
		userName,
		privKey,
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
	logrus.Infof("connecting to grpc server: %s", addr)
	conn, err := grpc.DialContext(ctx, addr, dialOpts...)
	if err != nil {
		return err
	}

	userRegistryClient := protoId.NewUserRegistryLoginClient(conn)

	if cliCtx.String(userName.Name) != "" {
		user := &protoId.User{
			UserName: cliCtx.String(userName.Name),
			PrivKey:  cliCtx.String(privKey.Name),
		}
		addUser(ctx, userRegistryClient, user)
	}

	if err := listAllUsers(ctx, userRegistryClient); err != nil {
		return err
	}

	// The following is for testing Get functions.
	privkey, err := ecdsa_util.RecoverPrivkey(testPrivkeyHex)
	pubkey := crypto.FromECDSAPub(&privkey.PublicKey)
	userName := &protoId.UserName{
		UserName: testUserName,
	}
	pubKey := &protoId.PubKey{
		PubKey: pubkey,
	}
	debug(ctx, userRegistryClient, userName, pubKey)

	// The following is for testing the login flow.
	loginFlow(ctx, userRegistryClient, userName, testPrivkeyHex)

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

func addUser(ctx context.Context, userRegistryClient protoId.UserRegistryLoginClient, user *protoId.User) {
	userRegistryClient.AddUser(ctx, user)
}

func loginFlow(ctx context.Context, userRegistryClient protoId.UserRegistryLoginClient, userName *protoId.UserName, privkeyHex string) {
	privateKey, err := ecdsa_util.RecoverPrivkey(privkeyHex)
	logrus.Info("recovered ", privateKey)
	if err != nil {
		logrus.Fatal(err)
	}

	// Step 1: Request login
	loginInfo, err := userRegistryClient.RequestLogin(ctx, userName)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Server requested login client to sign: ", loginInfo.UnsignedMsg)

	// Step 2: Sign message
	unsignedMsg := []byte(loginInfo.UnsignedMsg)
	msgHash := crypto.Keccak256Hash(unsignedMsg)
	digestHash := msgHash.Bytes()
	sig, err := crypto.Sign(digestHash, privateKey)
	if err != nil {
		logrus.Fatal(sig)
	}

	// Submit signed msg
	loginInfo.Signature = sig
	loginResp, err := userRegistryClient.Login(ctx, loginInfo)
	logrus.Info("status=", loginResp.Status)

}

func debug(ctx context.Context, userRegistryClient protoId.UserRegistryLoginClient, userName *protoId.UserName, pubKey *protoId.PubKey) {
	u, err := userRegistryClient.GetUserByPubKey(ctx, pubKey)
	if err != nil {
		logrus.Fatalf("Failed to call GetUserByPubKey: %v", err)
	}
	logrus.Printf("Username: %s, pubkey=0x%x\n", u.UserName, u.PubKey)

	u, err = userRegistryClient.GetUserByUserName(ctx, userName)
	if err != nil {
		logrus.Fatalf("Failed to call GetUserByUserName: %v", err)
	}
	logrus.Printf("Username: %s, pubkey=0x%x\n", u.UserName, u.PubKey)
}
