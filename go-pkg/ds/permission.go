package ds

import (
	"errors"
	"fmt"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/encrypt"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

// Do I have permission to the DataBlob
func (u *User) checkPerm(data *protods.DataBlob) bool {
	panic("not implemented")
}

// Get the decrypted data in EncryptedData
func (u *User) extractData(data *protods.DataBlob) ([]byte, error) {
	pubKey := ethereum.CompressPubkey(u.pubkey)
	userKey, err := findUserKey(data.Keys, pubKey)
	if err != nil {
		return nil, err
	}
	dataKey, err := recoverHiddenDataKey(userKey, u.privkey.D.Bytes())
	if err != nil {
		return nil, err
	}

	decryptedData, err := encrypt.Decrypt(dataKey, data.Iv, data.EncryptedData)
	if err != nil {
		return nil, errors.New("failed to decrypt data")
	}

	// AES blocks are padded, we need to get rid of the padding
	unpaddedDecryptedData := decryptedData[:(len(decryptedData) - int(decryptedData[len(decryptedData)-1]))]

	return unpaddedDecryptedData, nil
}

// User add key to an existing Blob and creat a new DataBlob
func (u *User) addKey(blob *protods.DataBlob, newDataOwner *User) *protods.DataBlob {
	panic("not implemented")
}

// 1. Generate an AES-key
// 2. Add an entry to "Secrets"
// 3. Encrypt data with AES key
func createDataBlob(data []byte, pubKey []byte) (*protods.DataBlob, error) {
	if len(pubKey) != 33 {
		return nil, fmt.Errorf("incorrect length %v, public key has to be compressed secp256k1 key", len(pubKey))
	}
	dataKey := encrypt.GenAes128Key()
	hiddenDataKey, err := generateHiddenDataKey(dataKey, pubKey[:])
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
