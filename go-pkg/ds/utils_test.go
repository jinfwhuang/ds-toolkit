package ds

import (
	"log"
	"reflect"
	"testing"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/encrypt"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestGenerateHiddenKey(t *testing.T) {
	user := createTestUser()
	dataKey := encrypt.GenAes128Key()
	hiddenDataKey, err := generateHiddenDataKey(dataKey, ethereum.CompressPubkey(user.pubkey))
	assert.NoError(t, err)

	assert.Equal(t, 32, len(hiddenDataKey.EncryptedDataKey))

	recoveredKey, err := recoverHiddenDataKey(hiddenDataKey, user.privkey.D.Bytes())
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(dataKey, recoveredKey))
}
