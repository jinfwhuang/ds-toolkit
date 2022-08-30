package dsn

import (
	b64 "encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var testWalletPath string = "./testdata/wallet.json"

func TestGenerateWallet(t *testing.T) {
	w, err := GenerateWallet()
	assert.NoError(t, err)
	assert.NotNil(t, w)

	walletAssertions(t, w)
}

func TestGenerateWalletFromPath(t *testing.T) {
	w, err := GenerateWalletFromPath(testWalletPath)
	assert.NoError(t, err)
	assert.NotNil(t, w)
	walletAssertions(t, w)
}

func TestGenerateWalletFromJWK(t *testing.T) {
	b, err := ioutil.ReadFile(testWalletPath)
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

func TestWrite(t *testing.T) {
	b, err := ioutil.ReadFile(testWalletPath)
	assert.NoError(t, err)
	assert.NotNil(t, b)

	w, err := GenerateWalletFromJWK(b)
	assert.NoError(t, err)
	assert.NotNil(t, w)

	walletAssertions(t, w)

	// Mock HTTP calls from Arweave SDK
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"POST",
		"https://arweave.net/chunk",
		httpmock.NewStringResponder(200, ""),
	)
	httpmock.RegisterResponder(
		"GET",
		"https://arweave.net/price/4",
		httpmock.NewStringResponder(200, "5"),
	)
	httpmock.RegisterResponder(
		"GET",
		"https://arweave.net/tx_anchor",
		httpmock.NewStringResponder(200, "test_anchor"),
	)

	res, err := Write([]byte("test"), w)

	// txID is based on the return of the tx_anchor call, it is deterministic
	txID := "9aFl1WV_94G4KNCVsC2EnI_utHrmcRv5sy_xoPpqp5Q"
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, txID, res)
}

func TestRead(t *testing.T) {
	b, err := ioutil.ReadFile(testWalletPath)
	assert.NoError(t, err)
	assert.NotNil(t, b)

	w, err := GenerateWalletFromJWK(b)
	assert.NoError(t, err)
	assert.NotNil(t, w)

	walletAssertions(t, w)

	expectedRes := "test"

	// Mock HTTP calls from Arweave SDK
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://arweave.net/tx/9aFl1WV_94G4KNCVsC2EnI_utHrmcRv5sy_xoPpqp5Q/data",
		httpmock.NewStringResponder(200, expectedRes),
	)

	txID := "9aFl1WV_94G4KNCVsC2EnI_utHrmcRv5sy_xoPpqp5Q"
	res, err := Read(txID)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, expectedRes, string(res))
}
