package main

import (
	"context"
	//"encoding/hex"
	"flag"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	cmd_utils "github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	protoId "github.com/jinfwhuang/ds-toolkit/proto/identity"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AppFlags = []cli.Flag{
	cmd_utils.GrpcPort,
	cmd_utils.LogLevel,
}

var (
	privateKey = flag.String("private_key", os.Getenv("PRIVATE_KEY"), "A secp256k1 private key in hex form, prefixed with 0x")
	userName   = flag.String("user_name", "", "username in the UserRegistry")
	pubKeyId   = flag.Int64("pubKey_id", 0, "public key ID for user")
)

const (
	// secp256k1 keys
	testPrivkeyHex  = "0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
	testUserName    = "jinhuang001"
	testPrivkeyHex2 = "0x89df41bce61452e4bed3e6325873022ba877a1683cfca9d03c0a5bf542944321"
	testUserName2   = "stella021"
	testUserName3   = "user3"
	testPrivkeyHex3 = "0xa329248cb59ccd5429a67e11529a831147bf1193bafcc45255129288362095cb"

	//pubkey = "0x040c1ca15b1ee87e5c493b85d4f2db6b13bc3aadb61f7af5b84ad30451074ad500b95b745a6600326d91bd4323da514b4b81d5d76f0973b66d6cf8e3b131525d41"
)

func main() {
	setupLogrus()

	flag.Parse()

	if *privateKey == "" {
		*privateKey = testPrivkeyHex
	}

	if *userName == "" {
		*userName = testUserName
	}

	logrus.Info("*private key ", *privateKey)
	ctx := context.Background()
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := "localhost:4000"
	logrus.Info("connecting to grpc server: %s", addr)
	conn, err := grpc.DialContext(ctx, "localhost:4000", dialOpts...)
	if err != nil {
		panic(err)
	}
	//idClient := protoId.NewIdentityClient(conn) //client stub to perform the RPCs
	userRegistryClient := protoId.NewUserRegistryLoginClient(conn)

	// Testing a login
	loginFlow(ctx, userRegistryClient, *userName, *privateKey, *pubKeyId)
}

func setupLogrus() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)
}

func loginFlow(ctx context.Context, userRegistryClient protoId.UserRegistryLoginClient, userName string, privkeyHex string, pubKeyId int64) {
	// Keys
	privateKey, err := ecdsa_util.RecoverPrivkey(privkeyHex)
	logrus.Info("recovered ", privateKey)
	if err != nil {
		logrus.Fatal(err)
	}
	//pubkey := crypto.FromECDSAPub(&privateKey.PublicKey)
	//logrus.Infof("logging in with pubkey=0x%s", hex.EncodeToString(pubkey))

	// Step 1: Request login
	loginInfo := &protoId.UserLogin{
		UserName: userName,
		PubKeyId: pubKeyId,
	}
	loginInfo, err = userRegistryClient.RequestLogin(ctx, loginInfo)
	logrus.Info("login info: ", loginInfo)

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

	//// Debug
	//sigWithoutID := sig[:len(sig)-1] // remove recovery id
	//validated := crypto.VerifySignature(pubkey, digestHash, sigWithoutID)
	//logrus.Info("hash string ", msgHash.String())
	//logrus.Info("hash hex    ", msgHash.Hex())
	//logrus.Info("hash byte    ", msgHash.Bytes())
	//logrus.Info("digest    ", digestHash)
	//logrus.Info("sig        ", base64.StdEncoding.EncodeToString(sig))
	//logrus.Info("validation  ", validated)

	// Submit signed msg
	loginInfo.Signature = sig
	loginResp, err := userRegistryClient.Login(ctx, loginInfo)
	logrus.Info("status=", loginResp.Status)

}

/*
func main() {
	setupLogrus()

	flag.Parse()
	if *privateKey == "" {
		*privateKey = testPrivkeyHex
	}

	ctx := context.Background()
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := "localhost:4000"
	logrus.Info("connecting to grpc server: %s", addr)
	conn, err := grpc.DialContext(ctx, "localhost:4000", dialOpts...)
	if err != nil {
		panic(err)
	}
	idClient := protoId.NewIdentityClient(conn)

	// Testing a login
	loginFlow(ctx, idClient, *privateKey)
}
func setupLogrus() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)
}

func loginFlow(ctx context.Context, idClient protoId.IdentityClient, privkeyHex string) {
	// Keys
	privateKey, err := ecdsa_util.RecoverPrivkey(privkeyHex)
	logrus.Info("recovered ", privateKey)
	if err != nil {
		logrus.Fatal(err)
	}
	pubkey := crypto.FromECDSAPub(&privateKey.PublicKey)
	logrus.Infof("logging in with pubkey=0x%s", hex.EncodeToString(pubkey))

	// Step 1: Request login
	loginMsg := &protoId.LoginMessage{
		PubKey: pubkey,
	}
	loginMsg, err = idClient.RequestLogin(ctx, loginMsg)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Server requested login client to sign: ", loginMsg.UnsignedMsg)
	logrus.Info("CORRECT LOGIN INFO:", loginMsg)

	// Step 2: Sign message
	unsignedMsg := []byte(loginMsg.UnsignedMsg)
	msgHash := crypto.Keccak256Hash(unsignedMsg)
	digestHash := msgHash.Bytes()
	sig, err := crypto.Sign(digestHash, privateKey)
	if err != nil {
		logrus.Fatal(sig)
	}
	logrus.Info("CORRECT HASH AND SIG: ", digestHash, sig, privateKey)
	//// Debug
	//sigWithoutID := sig[:len(sig)-1] // remove recovery id
	//validated := crypto.VerifySignature(pubkey, digestHash, sigWithoutID)
	//logrus.Info("hash string ", msgHash.String())
	//logrus.Info("hash hex    ", msgHash.Hex())
	//logrus.Info("hash byte    ", msgHash.Bytes())
	//logrus.Info("digest    ", digestHash)
	//logrus.Info("sig        ", base64.StdEncoding.EncodeToString(sig))
	//logrus.Info("validation  ", validated)

	// Submit signed msg
	loginMsg.Signature = sig
	loginResp, err := idClient.Login(ctx, loginMsg)
	logrus.Info("status=", loginResp.Status)
}
*/
