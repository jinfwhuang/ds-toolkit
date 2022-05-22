package main

import (
	"flag"
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
	testPrivkeyHex = "0x12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	//ctx := context.Background()
	flag.Parse()
	if *privateKey == "" {
		*privateKey = testPrivkeyHex
	}

	log.Println("fff")
}



