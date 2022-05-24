package encrypt

import (
	"bytes"
	"crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/bytesutil"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Llongfile)
}



func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}



func Test_AES001(t *testing.T) {
	aesKey := GenAes128Key()
	aesIv := GenCBCIv()

	plaintext := bytesutil.RandBytes(20)
	ciphertext, _ := Encrypt(aesKey, aesIv, plaintext)

	recoveredPlain, _ := Decrypt(aesKey, aesIv, ciphertext)
	assert.Equal(t, plaintext, recoveredPlain[:len(plaintext)])
}

func Test_AES002(t *testing.T) {
	aesKey := GenAes128Key()
	aesIv := GenCBCIv()

	plaintext := bytesutil.RandBytes(32)
	ciphertext, _ := Encrypt(aesKey, aesIv, plaintext)

	recoveredPlain, _ := Decrypt(aesKey, aesIv, ciphertext)
	assert.Equal(t, plaintext, recoveredPlain[:len(plaintext)])

	crypto.SHA1.HashFunc()
}

