package ds

import (
	"log"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/encrypt"
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
