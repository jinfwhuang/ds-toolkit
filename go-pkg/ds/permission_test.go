package ds

import (
	"log"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func createTestUser() *User {
	privKey, err := ethereum.GenerateKey()
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

func TestCreateDataBlob(t *testing.T) {
	data := []byte("test")
	alice := createTestUser()
	dataBlob, err := createDataBlob(data, alice.pubkey)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))
}

func TestExtractData(t *testing.T) {
	data := []byte("test")
	alice := createTestUser()
	dataBlob, err := createDataBlob(data, alice.pubkey)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	decryptedData, err := extractData(dataBlob, alice.privkey)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(data, decryptedData))
}

func TestCheckPerm(t *testing.T) {
	data := []byte("test")
	alice := createTestUser()
	dataBlob, err := createDataBlob(data, alice.pubkey)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	assert.True(t, checkPerm(dataBlob, alice.pubkey))

	bob := createTestUser()
	assert.False(t, checkPerm(dataBlob, bob.pubkey))
}

func TestAddKey(t *testing.T) {
	data := []byte("test")
	alice := createTestUser()
	dataBlob, err := createDataBlob(data, alice.pubkey)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	bob := createTestUser()

	assert.False(t, checkPerm(dataBlob, bob.pubkey))
	_, err = extractData(dataBlob, bob.privkey)
	assert.ErrorContains(t, err, "could not find public key")

	dataBlob, err = addKey(dataBlob, bob.pubkey, alice.privkey)
	assert.NoError(t, err)

	assert.True(t, checkPerm(dataBlob, bob.pubkey))
	decryptedData, err := extractData(dataBlob, bob.privkey)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(data, decryptedData))
	assert.False(t, reflect.DeepEqual([]byte{4, 2}, decryptedData))
}
