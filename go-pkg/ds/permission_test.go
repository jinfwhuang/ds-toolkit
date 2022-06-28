package ds

import (
	"log"
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
