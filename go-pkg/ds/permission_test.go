package ds

import (
	"log"
	"reflect"
	"testing"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestCreateDataBlob(t *testing.T) {
	data := []byte("test")
	user := createTestUser()
	compressedPubKeyBytes := ethereum.CompressPubkey(user.pubkey)
	dataBlob, err := createDataBlob(data, compressedPubKeyBytes)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))
}

func TestExtractData(t *testing.T) {
	data := []byte("test")
	user := createTestUser()
	compressedPubKeyBytes := ethereum.CompressPubkey(user.pubkey)
	dataBlob, err := createDataBlob(data, compressedPubKeyBytes)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	decryptedData, err := user.extractData(dataBlob)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(data, decryptedData))
}

func TestCheckPerm(t *testing.T) {
	data := []byte("test")
	user := createTestUser()
	compressedPubKeyBytes := ethereum.CompressPubkey(user.pubkey)
	dataBlob, err := createDataBlob(data, compressedPubKeyBytes)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	assert.True(t, user.checkPerm(dataBlob))

	user2 := createTestUser()
	assert.False(t, user2.checkPerm(dataBlob))
}

func TestAddKey(t *testing.T) {
	data := []byte("test")
	user := createTestUser()
	compressedPubKeyBytes := ethereum.CompressPubkey(user.pubkey)
	dataBlob, err := createDataBlob(data, compressedPubKeyBytes)
	assert.NoError(t, err)

	assert.Equal(t, 16, len(dataBlob.EncryptedData))

	user2 := createTestUser()

	assert.False(t, user2.checkPerm(dataBlob))
	_, err = user2.extractData(dataBlob)
	assert.Error(t, err)

	user.addKey(dataBlob, user2)

	assert.True(t, user2.checkPerm(dataBlob))
	decryptedData, err := user2.extractData(dataBlob)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(data, decryptedData))
}
