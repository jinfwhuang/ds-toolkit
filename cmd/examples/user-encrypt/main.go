package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
)

const (
	// Seeded secp256k1 private key.
	testPrivkeyHex = "12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
)

// Get the running directory of the current file.
func fileInRuntimeDir(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename) + file
}

func main() {
	// -ENCRYPTING-
	// Initialize *ecdsa.PrivateKey, using seeded hex private key.
	// In real environment the user is expected to input the data.
	alicePrivateKey, err := ethereum.HexToECDSA(testPrivkeyHex)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Using public key %v of initial owner", &alicePrivateKey.PublicKey)

	// Load the data the user wants to encrypt.
	dataPath := "/data.json"
	fmt.Printf("Using %v as example data blob", dataPath)
	exampleData, err := ioutil.ReadFile(fileInRuntimeDir(dataPath))
	if err != nil {
		log.Fatal("could not find data file")
	}

	// Create encrypted data blob for the user, using user's public key.
	fmt.Println("Encrypting data blob")
	dataBlob, err := ds.CreateDataBlob(exampleData, &alicePrivateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	// Create JSON struct from the encrypted data blob.
	dataBlobPath := "/data_blob.json"
	fmt.Printf("Saving the encrpyed data blob to %v", dataBlobPath)
	dataBlobJson, err := json.Marshal(dataBlob)
	if err != nil {
		panic(err)
	}

	// Write the encrypted data blob to a JSON file.
	err = ioutil.WriteFile(fileInRuntimeDir(dataBlobPath), dataBlobJson, 0644)
	if err != nil {
		log.Fatal("could not find data file")
	}

	// Decrypt the data from the encrypted data blob, using user's private key.
	retrievedData, err := ds.ExtractData(dataBlob, alicePrivateKey)
	if err != nil {
		panic(err)
	}
	// Print the initial data blob
	fmt.Printf("Decrypted encrpyed data blob, using owner's private key: %v", testPrivkeyHex)
	fmt.Println(string(retrievedData))
}
