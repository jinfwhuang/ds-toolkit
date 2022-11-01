package dsn

import (
	"crypto/sha256"
	"fmt"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
)

const (
	arNode = "https://arweave.net"
)

// Generate RSA keys and an Arweave wallet.
func GenerateWallet() (*goar.Wallet, error) {
	jwk, err := GenerateJWK()
	if err != nil {
		return nil, err
	}
	return goar.NewWallet(jwk, arNode)
}

// Generate an Arweave wallet using pre-existing byte representation of a JSON Web Key (JWK).
func GenerateWalletFromJWK(jwk []byte) (*goar.Wallet, error) {
	return goar.NewWallet(jwk, arNode)
}

// Generate an Arweave wallet from JWK in a file. The format can be verified with test_wallet.json.
func GenerateWalletFromPath(path string) (*goar.Wallet, error) {
	return goar.NewWalletFromPath(path, arNode)
}

// Write data to the Arweave protocol using an Arweave wallet.
func Write(data []byte, wallet *goar.Wallet) (string, error) {
	tags := []types.Tag{{Name: "Content-Type", Value: "application/pdf"}, {Name: "goar", Value: "testdata"}}
	tx, err := assemblyDataTx(data, wallet, tags)
	if err != nil {
		return "", err
	}

	// uploader Transaction
	uploader, err := goar.CreateUploader(wallet.Client, tx, nil)
	if err != nil {
		return "", err
	}

	err = uploader.Once()
	if err != nil {
		return "", err
	}

	return tx.ID, nil
}

// Read data from the Arweave protocol using an id, retrieved by the result of Write([]byte, *goar.Wallet).
func Read(id string) ([]byte, error) {
	arCli := goar.NewClient(arNode)
	txDataEncoded, err := arCli.GetTransactionData(id)
	if err != nil {
		return nil, err
	}
	return utils.Base64Decode(string(txDataEncoded))
}

func assemblyDataTx(bigData []byte, wallet *goar.Wallet, tags []types.Tag) (*types.Transaction, error) {
	reward, err := wallet.Client.GetTransactionPrice(bigData, nil)
	if err != nil {
		return nil, err
	}
	tx := &types.Transaction{
		Format:   2,
		Target:   "",
		Quantity: "0",
		Tags:     utils.TagsEncode(tags),
		Data:     utils.Base64Encode(bigData),
		DataSize: fmt.Sprintf("%d", len(bigData)),
		Reward:   fmt.Sprintf("%d", reward),
	}
	anchor, err := wallet.Client.GetTransactionAnchor()
	if err != nil {
		return nil, err
	}
	tx.LastTx = anchor
	tx.Owner = wallet.Owner()

	signData, err := utils.GetSignatureData(tx)
	if err != nil {
		return nil, err
	}

	sign, err := wallet.Signer.SignMsg(signData)
	if err != nil {
		return nil, err
	}

	txHash := sha256.Sum256(sign)
	tx.ID = utils.Base64Encode(txHash[:])

	tx.Signature = utils.Base64Encode(sign)
	return tx, nil
}
