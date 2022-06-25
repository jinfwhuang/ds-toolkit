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
	hiddenDataKey, AESKey, err := u.generateHiddenKey()
	if err != nil {
		return nil, err
	}
	dataLen := len(data)
	iv := encrypt.GenCBCIv()

	encryptedData, err := encrypt.Encrypt(AESKey, iv, data)
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

func (u *User) generateHiddenKey() (*protods.HiddenDataKey, []byte, error) {
	//HiddenDataKey creation
	dataAESKey := encrypt.GenAes128Key()

	ephemeralPrivKey, err := ethereum.GenerateKey()
	if err != nil {
		return nil, nil, errors.New("creation of ephemeral private key failed")
	}
	sharedSecretX, sharedSecretY := ephemeralPrivKey.PublicKey.ScalarMult(
		ephemeralPrivKey.PublicKey.X,
		ephemeralPrivKey.PublicKey.Y,
		u.privkey.D.Bytes(),
	)

	sharedSecret := elliptic.Marshal(ethereum.S256(), sharedSecretX, sharedSecretY)
	sharedSecretHash := sha256.Sum256(sharedSecret)

	iv := encrypt.GenCBCIv()

	encryptedHiddenKey, err := encrypt.Encrypt(sharedSecretHash[:], iv, dataAESKey)
	if err != nil {
		return nil, nil, errors.New("signing failed")
	}

	hiddenSharedKey := &protods.HiddenDataKey{
		Pubkey:           ethereum.CompressPubkey(&u.privkey.PublicKey),
		EphemeralPubkey:  ethereum.CompressPubkey(&ephemeralPrivKey.PublicKey),
		EncryptedDataKey: encryptedHiddenKey,
		Iv:               iv,
	}

	return hiddenSharedKey, dataAESKey, nil
}
