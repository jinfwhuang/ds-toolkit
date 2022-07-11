package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	cmd_utils "github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
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
	testPrivkeyHex  = "12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
	testPrivkeyHex2 = "894755216337ea35abdef4e4f7f172bda9f6039120d046bb4e14c29b751e4e84"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func getFullPath(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename) + file
}

func main() {
	aliceKey, err := ethereum.HexToECDSA(testPrivkeyHex)
	if err != nil {
		panic(err)
	}
	bobKey, err := ethereum.HexToECDSA(testPrivkeyHex2)
	if err != nil {
		panic(err)
	}

	exampleData, err := ioutil.ReadFile(getFullPath("/data.json"))
	if err != nil {
		log.Fatal("could not find data file")
	}

	dataBlob, err := ds.CreateDataBlob(exampleData, &aliceKey.PublicKey)
	if err != nil {
		panic(err)
	}
	newDataBlob, err := ds.AddKey(dataBlob, &bobKey.PublicKey, aliceKey)
	if err != nil {
		panic(err)
	}
	newDataBlobJson, err := json.Marshal(newDataBlob)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(getFullPath("/data_blob.json"), newDataBlobJson, 0644)
	if err != nil {
		log.Fatal("could not find data file")
	}

	retrievedData, err := ds.ExtractData(newDataBlob, bobKey)
	if err != nil {
		panic(err)
	}
	println(string(retrievedData))
}
