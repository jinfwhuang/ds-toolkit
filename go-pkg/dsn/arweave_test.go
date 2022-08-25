package dsn

import (
	b64 "encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/stretchr/testify/assert"
)

var test_wallet string = "./test_wallet.json"

func TestGenerateWallet(t *testing.T) {
	w, err := GenerateWallet()
	assert.NoError(t, err)
	assert.NotNil(t, w)

	walletAssertions(t, w)
}

func TestGenerateWalletFromPath(t *testing.T) {
	w, err := GenerateWalletFromPath(test_wallet)
	assert.NoError(t, err)
	assert.NotNil(t, w)
	walletAssertions(t, w)
}

func TestGenerateWalletFromJWK(t *testing.T) {
	b, err := ioutil.ReadFile(test_wallet)
	assert.NoError(t, err)
	assert.NotNil(t, b)

	w, err := GenerateWalletFromJWK(b)
	assert.NoError(t, err)
	assert.NotNil(t, w)

	walletAssertions(t, w)
}

func walletAssertions(t *testing.T, w *goar.Wallet) {
	data := []byte{2}
	dataBase64 := b64.RawURLEncoding.EncodeToString(data)

	item, err := w.CreateAndSignBundleItem(data, 1, "dGVzdA", "anchor", []types.Tag{})
	assert.NoError(t, err)
	assert.NotNil(t, item.Data)
	assert.NotNil(t, item.Signature)
	assert.Equal(t, dataBase64, item.Data)
}
