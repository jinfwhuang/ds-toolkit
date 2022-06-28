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

//TODO: Add more complete testing (end to end flow) once the full permission package is implemented

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
