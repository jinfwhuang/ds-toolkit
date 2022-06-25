package ds

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func createUser() *User {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		panic("Could not generate ecdsa private key")
	}

	return &User{
		Userid:  common.HexToAddress("0x29e4Af06632c0eAA6e5b8377d1561E0950B7E963"),
		Name:    "Alice",
		privkey: privKey,
		pubkey:  &privKey.PublicKey,
	}
}

//TODO: Add more complete testing (end to end flow) once the full permission package is implemented

func TestCreateDataBlob(t *testing.T) {
	data := []byte("test")
	user := createUser()
	dataBlob, err := user.createDataBlob(data)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))
}

func TestGenerateHiddenKey(t *testing.T) {
	user := createUser()
	hiddenDataKey, _, err := user.generateHiddenKey()
	assert.NoError(t, err)

	assert.Equal(t, 32, len(hiddenDataKey.EncryptedDataKey))
}
