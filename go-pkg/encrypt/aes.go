package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func CbcPadBytes(in []byte) []byte {
	paddingSize := aes.BlockSize - len(in) % aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(in, padding...)
}

func Encrypt(key, iv, data []byte) ([]byte, error) {
	in := CbcPadBytes(data)
	out := make([]byte, len(in))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(out, in)

	//return out[:len(data)], nil
	return out, nil
}


func Decrypt(key, iv, enc []byte) ([]byte, error) {
	paddedEnc := CbcPadBytes(enc)
	dec := make([]byte, len(paddedEnc))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dec, paddedEnc)
	return dec[:len(enc)], nil
}

func SecureRandBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func GenAes128Key() []byte {
	return SecureRandBytes(128 / 8)
}

func GenCBCIv() []byte {
	return SecureRandBytes(aes.BlockSize)
}
