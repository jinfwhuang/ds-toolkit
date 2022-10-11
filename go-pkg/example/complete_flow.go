package example

import (
	"bytes"
	"encoding/json"
	"fmt"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/dsn"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

func main() {
	privKey, err := ethereum.GenerateKey()
	if err != nil {
		panic("Could not generate ecdsa private key")
	}
	pubKey := &privKey.PublicKey

	mySecretData := []byte("password")

	dataBlob, err := ds.CreateDataBlob(mySecretData, pubKey)
	if err != nil {
		panic("Could not create data blob")
	}

	dataBlobJson, err := json.Marshal(dataBlob)
	if err != nil {
		panic("Could not marshal data blob to JSON")
	}

	wallet, err := dsn.GenerateWallet()
	if err != nil {
		panic("Could not create wallet")
	}

	id, err := dsn.Write(dataBlobJson, wallet)
	if err != nil {
		panic("Could not write to Arweave DLT")
	}

	println("Transaction ID of data blob on Arweave DLT: %v", id)

	payload, err := dsn.Read(id)
	if err != nil {
		panic("Could not read Arweave transaction")
	}

	var dataBlobFetched protods.DataBlob
	err = json.Unmarshal(payload, &dataBlobFetched)
	if err != nil {
		panic("Could not unmarshal data blob")
	}

	extractedData, err := ds.ExtractData(&dataBlobFetched, privKey)
	if err != nil {
		panic("Could not extract data from fetched Arweave data")
	}

	if !bytes.Equal(extractedData, mySecretData) {
		panic(fmt.Sprintf("Extracted data: '%v' does not match initial data: '%v'", extractedData, mySecretData))
	}

	privKeyNew, err := ethereum.GenerateKey()
	if err != nil {
		panic("Could not generate second ecdsa private key")
	}
	pubKeyNew := &privKeyNew.PublicKey

	dataBlobNew, err := ds.AddKey(&dataBlobFetched, pubKeyNew, privKey)
	if err != nil {
		panic("Could not add key to data blob")
	}

	extractedDataNew, err := ds.ExtractData(dataBlobNew, privKeyNew)
	if err != nil {
		panic("Could not extract data using newly added key")
	}

	if !bytes.Equal(extractedData, mySecretData) {
		panic(fmt.Sprintf("Data from original blob: '%v' does not match extracted data from new user: '%v'", extractedData, extractedDataNew))
	}

}
