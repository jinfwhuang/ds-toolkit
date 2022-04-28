package main

import (
	"context"
	"encoding/hex"
	"flag"
	"github.com/jinfwhuang/ds-toolkit/pkg/ecdsa-util"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/pkg/cmd-utils"
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
)

const (
	// secp256k1 keys
	testPrivkeyHex = "0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
)

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
	if err != nil {
		logrus.Fatal(err)
	}
	pubkey := crypto.FromECDSAPub(&privateKey.PublicKey)
	logrus.Infof("loggin in with pubkey=0x%s", hex.EncodeToString(pubkey))

	// Step 1: Request login
	loginMsg := &protoId.LoginMessage{
		PubKey: pubkey,
	}
	loginMsg, err = idClient.RequestLogin(ctx, loginMsg)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Server requested login client to sign: ", loginMsg.UnsignedMsg)

	// Step 2: Sign message
	unsignedMsg := []byte(loginMsg.UnsignedMsg)
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
	loginMsg.Signature = sig
	loginResp, err := idClient.Login(ctx, loginMsg)
	logrus.Info("status=", loginResp.Status)
}
