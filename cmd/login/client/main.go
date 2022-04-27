package main

import (
	"context"
	"github.com/jinfwhuang/ds-toolkit/pkg/ecdsa-util"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/jinfwhuang/ds-toolkit/pkg/cmd-utils"
	protoId "github.com/jinfwhuang/ds-toolkit/proto/identity"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//logrus "log"
	"github.com/sirupsen/logrus"
)



var AppFlags = []cli.Flag{
	cmd_utils.GrpcPort,
	cmd_utils.LogLevel,
}

const (
	// secp256k1 public key
	privkeyStr = "0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
	pubkeyStr = "0x040c1ca15b1ee87e5c493b85d4f2db6b13bc3aadb61f7af5b84ad30451074ad500b95b745a6600326d91bd4323da514b4b81d5d76f0973b66d6cf8e3b131525d41"

	/**
	Private Key: 0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5
	Public Key: 0x040c1ca15b1ee87e5c493b85d4f2db6b13bc3aadb61f7af5b84ad30451074ad500b95b745a6600326d91bd4323da514b4b81d5d76f0973b66d6cf8e3b131525d41
	Address: 0x67F72BcD03F63A448A0B5cFFe7DfA34C6f9382eD
	 */
)

func main() {
	setupLogrus()

	// Testing a login
	loginFlow()

}


func ZeroXToByte(s string) []byte {
	b, err := hexutil.Decode(s)
	if err != nil {
		panic(err)
	}
	return b
}

func setupLogrus()  {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)
}


func loginFlow() {
	setupLogrus()

	ctx := context.Background()
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithInsecure(),
	}
	addr := "localhost:4000"
	logrus.Info("connecting to grpc server: %s", addr)
	conn, err := grpc.DialContext(ctx, "localhost:4000", dialOpts...)
	if err != nil {
		panic(err)
	}
	server := protoId.NewIdentityClient(conn)


	// Keys
	privateKey, err := ecdsa_util.RecoverPrivkey(privkeyStr)
	if err != nil {
		panic(err)
	}
	pubkey := crypto.FromECDSAPub(&privateKey.PublicKey)

	//publicKey0x := "0x040c1ca15b1ee87e5c493b85d4f2db6b13bc3aadb61f7af5b84ad30451074ad500b95b745a6600326d91bd4323da514b4b81d5d76f0973b66d6cf8e3b131525d41"
	//pubkey := RecoverPubkey(publicKey0x)
	//pubkeyByte := crypto.FromECDSAPub(pubkey)

	//pubKey, err := base64.StdEncoding.DecodeString("4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12") // TODO: not a pub key, it is 20 byte pub address according to eth spec

	// Step 1: request login
	loginMsg := &protoId.LoginMessage {
		PubKey: pubkey,
	}
	loginMsg, err = server.RequestLogin(ctx, loginMsg)
	if err != nil {
		panic(err)
	}
	logrus.Info("step 1: ", loginMsg)

	// Step 2: Sign message
	unSignMsg := []byte(loginMsg.UnsignedMsg)
	msgHash := crypto.Keccak256Hash(unSignMsg)
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
	loginResp, err := server.Login(ctx, loginMsg)
	logrus.Info(loginResp)

}

