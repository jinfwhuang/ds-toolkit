package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	protoId "github.com/jinfwhuang/ds-sdk/proto/identity"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	tmplog "log"
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

const (
	// secp256k1 public key
	privkeyStr = "0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
	pubkey = "0x040c1ca15b1ee87e5c493b85d4f2db6b13bc3aadb61f7af5b84ad30451074ad500b95b745a6600326d91bd4323da514b4b81d5d76f0973b66d6cf8e3b131525d41"

	/**
	SAVE BUT DO NOT SHARE THIS (Private Key): 0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5
	Public Key: 0x040c1ca15b1ee87e5c493b85d4f2db6b13bc3aadb61f7af5b84ad30451074ad500b95b745a6600326d91bd4323da514b4b81d5d76f0973b66d6cf8e3b131525d41
	Address: 0x67F72BcD03F63A448A0B5cFFe7DfA34C6f9382eD
	 */
)

func ZeroXToByte(s string) []byte {
	b, err := hexutil.Decode(s)
	if err != nil {
		panic(err)
	}
	return b
}

// Use secp256k1
// See https://jinsnotes.com/2020-12-30-elliptical-curve-cryptography
func RecoverPrivkey(s string)  *ecdsa.PrivateKey {
	curve := secp256k1.S256()
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	d, _ := hexutil.DecodeBig(s)
	priv.D = d

	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(priv.D.Bytes())

	// DEBUG MSG
	//tmplog.Println("Private Key:", hexutil.Encode(crypto.FromECDSA(priv)))
	//pub := priv.Public()
	//publicKeyECDSA, ok := pub.(*ecdsa.PublicKey)
	//if !ok {
	//	tmplog.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//tmplog.Println("Public Key:", hexutil.Encode(publicKeyBytes))
	//
	//address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//tmplog.Println("Address:", address)

	return priv
}

func main() {



	RecoverPrivkey(privkeyStr)




}

func loginFlow() {
	ctx := context.Background()
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithInsecure(),
	}
	addr := "localhost:4000"
	tmplog.Printf("connecting to grpc server: %s", addr)
	conn, err := grpc.DialContext(ctx, "localhost:4000", dialOpts...)
	if err != nil {
		panic(err)
	}
	server := protoId.NewIdentityClient(conn)

	pubKey, err := base64.StdEncoding.DecodeString("4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12") // TODO: not a pub key, it is 20 byte pub address according to eth spec

	loginMsg := &protoId.LoginMessage{
		PubKey: pubKey,
	}
	tmplog.Println("step 1", loginMsg)

	// Request a login
	loginMsg, err = server.RequestLogin(ctx, loginMsg)
	if err != nil {
		panic(err)
	}
	tmplog.Println("step 2", loginMsg)

	// Sign message
}

