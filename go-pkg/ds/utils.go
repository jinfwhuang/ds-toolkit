package ds

import (
	"bytes"
	"crypto/sha256"
	"errors"

	ethereum "github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/encrypt"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

var (
	curve256 = ethereum.S256()
)

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

func findUserKey(keys []*protods.HiddenDataKey, pubKey []byte) (*protods.HiddenDataKey, error) {
	for _, k := range keys {
		if bytes.Equal(k.Pubkey, pubKey) {
			return k, nil
		}
	}
	return nil, errors.New("could not find public key")
}
