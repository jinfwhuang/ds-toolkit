package ds

import (
	"bytes"
	"log"
	"testing"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/encrypt"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestGenerateHiddenKey(t *testing.T) {
	alice := createTestUser("Alice")
	dataKey := encrypt.GenAes128Key()
	hiddenDataKey, err := generateHiddenDataKey(dataKey, ethereum.CompressPubkey(alice.pubkey))
	assert.NoError(t, err)

	assert.Equal(t, 32, len(hiddenDataKey.EncryptedDataKey))

	recoveredKey, err := recoverHiddenDataKey(hiddenDataKey, alice.privkey.D.Bytes())
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(dataKey, recoveredKey))
}
