package main

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

const (
	// secp256k1 keys
	testPrivkeyHex = "12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
	testPubkeyHex  = "0237dcc127cd98c4080fbb82f148e4bb4281153276acfa718b3366b74c4a1039bb"
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
	bobPublicKeyHex, err := hex.DecodeString(testPubkeyHex)
	if err != nil {
		panic(err)
	}
	bobPublicKey, err := ethereum.DecompressPubkey(bobPublicKeyHex)
	if err != nil {
		panic(err)
	}

	blobBytes, err := ioutil.ReadFile(fileInRuntimeDir("/data_blob.json"))
	if err != nil {
		panic("could not find blob file")
	}
	var dataBlob protods.DataBlob
	err = json.Unmarshal(blobBytes, &dataBlob)
	if err != nil {
		panic(err)
	}

	newDataBlob, err := ds.AddKey(&dataBlob, bobPublicKey, alicePrivateKey)
	if err != nil {
		panic(err)
	}
	newDataBlobJson, err := json.Marshal(newDataBlob)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileInRuntimeDir("/new_data_blob.json"), newDataBlobJson, 0644)
	if err != nil {
		log.Fatal("could not create data file")
	}

	retrievedData, err := ds.ExtractData(newDataBlob, alicePrivateKey)
	if err != nil {
		panic(err)
	}
	println(string(retrievedData))
}
