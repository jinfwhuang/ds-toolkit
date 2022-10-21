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
	// Generate ethereum ECDSA private key.
	// This key will be used for decrypting the data later on.
	privKey, err := ethereum.GenerateKey()
	if err != nil {
		panic("Could not generate ecdsa private key")
	}
	// Derive the public key.
	// This key will be used for encrypting the data.
	pubKey := &privKey.PublicKey

	// Example data to be stored and/or shared later on.
	// It can be any form of data represented in bytes.
	mySecretData := []byte("password")

	// Create a data blob by encrypting the data with the public key.
	// The data blob contains encrypted version of the pubkey and an AES symmetric secret key.
	// The AES key is used for encrypting the data and destroyed afterwards.
	// The AES key can be obtained again to decrypt the data by using the private key.
	// The data blob also contains the encrypted data and metadata for data integrity.
	dataBlob, err := ds.CreateDataBlob(mySecretData, pubKey)
	if err != nil {
		panic("Could not create data blob")
	}

	// Create a JSON representation of the data blob.
	dataBlobJson, err := json.Marshal(dataBlob)
	if err != nil {
		panic("Could not marshal data blob to JSON")
	}

	// Generate a wallet for Arweave DLT.
	// Arweave is the decentralised network of storage, used to store our data blob.
	// This generates a JWK, from which a wallet is derived.
	// Alternatively, if JWK is already present, `dsn.GenerateWalletFromJWK(jwk []byte)` can be used for generation of a wallet.
	// `dsn.GenerateWalletFromPath(path string)` generates wallet from a file containing JWK.
	//
	// JWK can also be created from RSA private key using `dsn.RSAToJWK(privatekey *rsa.PrivateKey)`.
	wallet, err := dsn.GenerateWallet()
	if err != nil {
		panic("Could not create wallet")
	}

	// Write the JSON data blob represented in bytes to Arweave.
	// Note that the wallet should be funded before this step.
	// Funding a wallet with AR (native currency of Arweave) happens outside of the toolkit.
	// The address of the wallet can be obtained by `wallet.Signer.Address`.
	id, err := dsn.Write(dataBlobJson, wallet)
	if err != nil {
		panic("Could not write to Arweave DLT")
	}

	println("Transaction ID of data blob on Arweave DLT: %v", id)

	// Read the data uploaded to Arweave, using ID of transaction.
	// Note that this will not be available immediately after writing.
	// Arweave block generation takes roughly 2 minutes.
	// This means that in around 2 minutes, data will be readable from the DLT.
	payload, err := dsn.Read(id)
	if err != nil {
		panic("Could not read Arweave transaction")
	}

	// Unmarshal retrieved data blob to go struct.
	var dataBlobFetched protods.DataBlob
	err = json.Unmarshal(payload, &dataBlobFetched)
	if err != nil {
		panic("Could not unmarshal data blob")
	}

	// Extract original data (in this case "password") using the private key.
	extractedData, err := ds.ExtractData(&dataBlobFetched, privKey)
	if err != nil {
		panic("Could not extract data from fetched Arweave data")
	}

	if !bytes.Equal(extractedData, mySecretData) {
		panic(fmt.Sprintf("Extracted data: '%v' does not match initial data: '%v'", extractedData, mySecretData))
	}

	// Generate new private key.
	privKeyNew, err := ethereum.GenerateKey()
	if err != nil {
		panic("Could not generate second ecdsa private key")
	}

	// Derive new public key from the new private key.
	pubKeyNew := &privKeyNew.PublicKey

	// Create new data blob from the previous one, which now includes both keys, using the initial private key and the new public key.
	// The data in the new data blob will be decryptable by both private keys.
	// Original data blob is not modified.
	dataBlobNew, err := ds.AddKey(&dataBlobFetched, pubKeyNew, privKey)
	if err != nil {
		panic("Could not add key to data blob")
	}

	// Extract data from the new data blob, using the new private key.
	extractedDataNew, err := ds.ExtractData(dataBlobNew, privKeyNew)
	if err != nil {
		panic("Could not extract data using newly added key")
	}

	if !bytes.Equal(extractedData, mySecretData) {
		panic(fmt.Sprintf("Data from original blob: '%v' does not match extracted data from new user: '%v'", extractedData, extractedDataNew))
	}

}
