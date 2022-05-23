package main

import (
	"encoding/base64"
	"flag"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/bytesutil"
	"log"
	"os"

	"github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	"github.com/urfave/cli/v2"
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
	user1PrivkeyHex = "0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"

fdasf
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	//ctx := context.Background()
	flag.Parse()
	if *privateKey == "" {
		*privateKey = user1PrivkeyHex
	}

	log.Println("fff")
	b := bytesutil.RandBytes(64)
	log.Println(hexutil.Encode(b))
	log.Println(base64.StdEncoding.EncodeToString(b))

	log.Println(hexutil.Encode(bytesutil.RandBytes(64)))
	log.Println(base64.StdEncoding.EncodeToString(bytesutil.RandBytes(64)))

}



