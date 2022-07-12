package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
)

const (
	// secp256k1 key
	testPrivkeyHex = "12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
)

func fileInRuntimeDir(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename) + file
}

func main() {
	alicePrivateKey, err := ethereum.HexToECDSA(testPrivkeyHex)
	if err != nil {
		panic(err)
	}

	exampleData, err := ioutil.ReadFile(fileInRuntimeDir("/data.json"))
	if err != nil {
		log.Fatal("could not find data file")
	}

	dataBlob, err := ds.CreateDataBlob(exampleData, &alicePrivateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	dataBlobJson, err := json.Marshal(dataBlob)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileInRuntimeDir("/data_blob.json"), dataBlobJson, 0644)
	if err != nil {
		log.Fatal("could not find data file")
	}

	retrievedData, err := ds.ExtractData(dataBlob, alicePrivateKey)
	if err != nil {
		panic(err)
	}
	println(string(retrievedData))
}
