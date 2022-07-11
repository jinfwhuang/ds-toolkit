package ds

import (
	"bytes"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func createTestUser(name string) *User {
	privKey, err := ethereum.GenerateKey()
	if err != nil {
		panic("Could not generate ecdsa private key")
	}

	return &User{
		Userid:  common.BigToAddress(privKey.D),
		Name:    name,
		privkey: privKey,
		pubkey:  &privKey.PublicKey,
	}
}

func TestCreateDataBlob(t *testing.T) {
	data := []byte("test")
	alice := createTestUser("Alice")
	dataBlob, err := CreateDataBlob(data, alice.pubkey)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))
}

func TestExtractData(t *testing.T) {
	data := []byte("test")
	alice := createTestUser("Alice")
	dataBlob, err := CreateDataBlob(data, alice.pubkey)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	decryptedData, err := ExtractData(dataBlob, alice.privkey)
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(data, decryptedData))
}

func TestCheckPerm(t *testing.T) {
	data := []byte("test")
	alice := createTestUser("Alice")
	dataBlob, err := CreateDataBlob(data, alice.pubkey)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	assert.True(t, CheckPerm(dataBlob, alice.pubkey))

	bob := createTestUser("Bob")
	assert.False(t, CheckPerm(dataBlob, bob.pubkey))
}

func TestAddKey(t *testing.T) {
	data := []byte("test")
	alice := createTestUser("Alice")
	dataBlob, err := CreateDataBlob(data, alice.pubkey)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	bob := createTestUser("Bob")

	assert.False(t, CheckPerm(dataBlob, bob.pubkey))
	_, err = ExtractData(dataBlob, bob.privkey)
	assert.ErrorContains(t, err, "could not find public key")

	dataBlob, err = AddKey(dataBlob, bob.pubkey, alice.privkey)
	assert.NoError(t, err)

	assert.True(t, CheckPerm(dataBlob, bob.pubkey))
	decryptedData, err := ExtractData(dataBlob, bob.privkey)
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(data, decryptedData))
	assert.False(t, bytes.Equal([]byte{4, 2}, decryptedData))
}
