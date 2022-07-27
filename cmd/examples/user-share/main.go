package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

const (
	// Seeded secp256k1 private key of the current owner, used to decrypt the initial encrypted data blob.
	testPrivkeyHex = "12de257b783b96ce90012a6c45f3ce61216dd60f22159d2f5cb9e17f3126bbe5"
	// Seeded secp256k1 public key of the owner to be added.
	testPubkeyHex = "0237dcc127cd98c4080fbb82f148e4bb4281153276acfa718b3366b74c4a1039bb"
)

// Get the running directory of the current file.
func fileInRuntimeDir(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename) + file
}

func main() {
	// --
	// Initialize *ecdsa.PrivateKey, using seeded hex private key.
	// In real environment the user is expected to input the data.
	fmt.Printf("Using private key %v for decryption", testPrivkeyHex)
	alicePrivateKey, err := ethereum.HexToECDSA(testPrivkeyHex)
	if err != nil {
		panic(err)
	}
	// Initialize *ecdsa.PublicKey, using seeded hex public key.
	// In real environment the user is expected to input the data.
	fmt.Printf("Using public key %v for new shared owner", testPrivkeyHex)
	bobPublicKeyHex, err := hex.DecodeString(testPubkeyHex)
	if err != nil {
		panic(err)
	}
	bobPublicKey, err := ethereum.DecompressPubkey(bobPublicKeyHex)
	if err != nil {
		panic(err)
	}

	// Load the data the user wants to encrypt from a JSON file.
	dataBlobPath := "/data_blob.json"
	fmt.Printf("Using %v as example encrypted data blob", dataBlobPath)
	blobBytes, err := ioutil.ReadFile(fileInRuntimeDir(dataBlobPath))
	if err != nil {
		panic("could not find blob file")
	}
	var dataBlob protods.DataBlob
	err = json.Unmarshal(blobBytes, &dataBlob)
	if err != nil {
		panic(err)
	}

	// Add the new shared owner, using its public key.
	// Create new encrypted data blob.
	fmt.Println("Decrypting encrypted data blob and adding new owner's key")
	newDataBlob, err := ds.AddKey(&dataBlob, bobPublicKey, alicePrivateKey)
	if err != nil {
		panic(err)
	}
	// Marshal the new encrypted data blob to JSON byte slice.
	newDataBlobJson, err := json.Marshal(newDataBlob)
	if err != nil {
		panic(err)
	}
	// Write the new encrypted data blob to a JSON file.
	newDataBlobPath := "/new_data_blob.json"
	fmt.Printf("Saving the new encrpyed data blob to %v", newDataBlobPath)
	err = ioutil.WriteFile(fileInRuntimeDir(newDataBlobPath), newDataBlobJson, 0644)
	if err != nil {
		log.Fatal("could not create data file")
	}

	// Decrypt the data from the encrypted data blob, using initial owner's private key.
	retrievedData, err := ds.ExtractData(newDataBlob, alicePrivateKey)
	if err != nil {
		panic(err)
	}
	// Print the new data blob
	fmt.Println("Decrypted new encrpyed data blob, using initial owner's private key")
	fmt.Println(string(retrievedData))
}
