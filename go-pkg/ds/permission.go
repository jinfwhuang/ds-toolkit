package ds

import (
	"crypto/elliptic"
	"crypto/sha256"
	"errors"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/encrypt"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

// Do I have permission to the DataBlob
func (u *User) checkPerm(data *protods.DataBlob) bool {
	panic("not implemented")
}

// Get the decrypted data in EncryptedData
func (u *User) extractData(data *protods.DataBlob) []byte {
	panic("not implemented")
}

// User add key to an existing Blob and creat a new DataBlob
func (u *User) addKey(blob *protods.DataBlob, newDataOwner *User) *protods.DataBlob {
	panic("not implemented")
}

// 1. Generate an AES-key
// 2. Add an entry to "Secrets"
// 3. Encrypt data with AES key
func (u *User) createDataBlob(data []byte) (*protods.DataBlob, error) {
	//HiddenDataKey creation
	dataKey := encrypt.GenAes128Key()

	ephemeralPrivKey, err := ethereum.GenerateKey()
	if err != nil {
		return nil, errors.New("creation of ephemeral private key failed")
	}
	sharedSecretX, sharedSecretY := ephemeralPrivKey.PublicKey.Add(
		u.privkey.X,
		u.privkey.Y,
		ephemeralPrivKey.PublicKey.X,
		ephemeralPrivKey.PublicKey.Y,
	)

	sharedSecret := elliptic.Marshal(ethereum.S256(), sharedSecretX, sharedSecretY)
	sharedSecretHash := sha256.Sum256(sharedSecret)

	iv := encrypt.GenCBCIv()

	encryptedHiddenKey, err := encrypt.Encrypt(sharedSecretHash[:], iv, dataKey)
	if err != nil {
		return nil, errors.New("signing failed")
	}

	hiddenDataKey := protods.SharedKey{
		Pubkey:          ethereum.CompressPubkey(&u.privkey.PublicKey),
		EphemeralPubkey: ethereum.CompressPubkey(&ephemeralPrivKey.PublicKey),
		EncryptedKey:    encryptedHiddenKey,
		Iv:              iv,
	}

	//DataBlob creation
	dataLen := len(data)

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
		Secrets:           []*protods.SharedKey{&hiddenDataKey},
	}

	return &dataBlob, nil
}
