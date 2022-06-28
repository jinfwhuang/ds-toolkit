package ds

import (
	"crypto/sha256"
	"errors"
	"fmt"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/encrypt"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

var (
	curve256 = ethereum.S256()
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

func generateHiddenDataKey(dataAESKey []byte, pubKey []byte) (*protods.HiddenDataKey, error) {
	ephemeralPrivKey, err := ethereum.GenerateKey()
	if err != nil {
		return nil, errors.New("creation of ephemeral private key failed")
	}

	decompressedPubKey, err := ethereum.DecompressPubkey(pubKey)
	if err != nil {
		return nil, errors.New("failed to decompress public key")
	}

	secretPointX, _ := curve256.ScalarMult(
		decompressedPubKey.X,
		decompressedPubKey.Y,
		ephemeralPrivKey.D.Bytes(),
	)

	sharedSecret := secretPointX.Bytes()
	sharedSecretHash := sha256.Sum256(sharedSecret)
	ephemeralAesKey := sharedSecretHash[:16]

	iv := encrypt.GenCBCIv()

	encryptedHiddenKey, err := encrypt.Encrypt(ephemeralAesKey, iv, dataAESKey)
	if err != nil {
		return nil, errors.New("signing failed")
	}

	hiddenSharedKey := &protods.HiddenDataKey{
		Pubkey:           pubKey,
		EphemeralPubkey:  ethereum.CompressPubkey(&ephemeralPrivKey.PublicKey),
		EncryptedDataKey: encryptedHiddenKey,
		Iv:               iv,
	}

	return hiddenSharedKey, nil
}

func recoverHiddenDataKey(hiddenSharedKey *protods.HiddenDataKey, privKey []byte) ([]byte, error) {
	decompressedEphemeralPubkey, err := ethereum.DecompressPubkey(hiddenSharedKey.EphemeralPubkey)
	if err != nil {
		return nil, errors.New("failed to decompress public key")
	}
	secretPointX, _ := curve256.ScalarMult(
		decompressedEphemeralPubkey.X,
		decompressedEphemeralPubkey.Y,
		privKey,
	)

	sharedSecret := secretPointX.Bytes()
	sharedSecretHash := sha256.Sum256(sharedSecret)
	ephemeralAesKey := sharedSecretHash[:16]

	hiddenKey, err := encrypt.Decrypt(ephemeralAesKey, hiddenSharedKey.Iv, hiddenSharedKey.EncryptedDataKey)
	if err != nil {
		return nil, errors.New("failed to decrypt hidden key")
	}

	return hiddenKey[:16], nil
}
