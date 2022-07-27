package ds

import (
	"crypto/ecdsa"
	"errors"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/encrypt"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

// Does user have permission to the DataBlob
func CheckPerm(data *protods.DataBlob, pubKey *ecdsa.PublicKey) bool {
	pubKeyBytes := ethereum.CompressPubkey(pubKey)
	_, err := findUserKey(data.Keys, pubKeyBytes)
	return err == nil
}

// Get the decrypted data in DataBlob
func ExtractData(data *protods.DataBlob, privKey *ecdsa.PrivateKey) ([]byte, error) {
	pubKeyBytes := ethereum.CompressPubkey(&privKey.PublicKey)
	userKey, err := findUserKey(data.Keys, pubKeyBytes)
	if err != nil {
		return nil, err
	}
	dataKey, err := recoverHiddenDataKey(userKey, privKey.D.Bytes())
	if err != nil {
		return nil, err
	}

	decryptedData, err := encrypt.Decrypt(dataKey, data.Iv, data.EncryptedData)
	if err != nil {
		return nil, errors.New("failed to decrypt data")
	}

	// AES blocks are padded, we need to get rid of the padding
	unpaddedDecryptedData := decryptedData[:data.DataLen]

	return unpaddedDecryptedData, nil
}

// Add key to an existing DataBlob and create a new DataBlob
func AddKey(blob *protods.DataBlob, newPubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey) (*protods.DataBlob, error) {
	pubKeyBytes := ethereum.CompressPubkey(&privKey.PublicKey)
	userKey, err := findUserKey(blob.Keys, pubKeyBytes)
	if err != nil {
		return nil, err
	}
	dataKey, err := recoverHiddenDataKey(userKey, privKey.D.Bytes())
	if err != nil {
		return nil, err
	}

	newPubKeyBytes := ethereum.CompressPubkey(newPubKey)
	newDataKey, err := generateHiddenDataKey(dataKey, newPubKeyBytes)
	if err != nil {
		return nil, err
	}

	blob.Keys = append(blob.Keys, newDataKey)

	return blob, nil
}

// 1. Generate an AES-key
// 2. Add an entry to "Secrets"
// 3. Encrypt data with AES key
func CreateDataBlob(data []byte, pubKey *ecdsa.PublicKey) (*protods.DataBlob, error) {
	dataKey := encrypt.GenAes128Key()
	pubKeyBytes := ethereum.CompressPubkey(pubKey)
	hiddenDataKey, err := generateHiddenDataKey(dataKey, pubKeyBytes)
	if err != nil {
		return nil, err
	}
	dataLen := len(data)
	iv := encrypt.GenCBCIv()

	encryptedData, err := encrypt.Encrypt(dataKey, iv, data)
	if err != nil {
		return nil, errors.New("encryption of data with data key failed")
	}

	encryptedDataHash := ethereum.Keccak256Hash(encryptedData)

	dataBlob := protods.DataBlob{
		DataLen:           uint64(dataLen),
		Iv:                iv,
		EncryptedDataHash: encryptedDataHash[:],
		EncryptedData:     encryptedData,
		Keys:              []*protods.HiddenDataKey{hiddenDataKey},
	}

	return &dataBlob, nil
}
