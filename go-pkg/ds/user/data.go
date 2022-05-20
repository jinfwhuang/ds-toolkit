package user

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/**



 */
type DsData struct {
	// Header
	EncryptionKey []byte  // AES128, the key itself is encrypted
	userKeys []UserKey // Each entry could be used to recover AES_key

	// Data
	EncryptedData []byte  // Encrypted by AES_key
}

type UserKey struct {
	counterPartyPubkey []byte  // This should be ephemeral; H_b, p_b
	pubkey []byte  // H_a, private key is p_a

	// Encrypted with shared secret: s = p_a * p_b * G = p_a * H_b
	// Anyone with knowledge of p_a is able to calculate s
	// encryptedKey = AES_encrypt(ase_key, a)
	encryptedAesKey []byte // Be able to recover aes_key

}


func (u *User) Encrypt(data []byte, sig []byte) bool {
	// Keccak hash
	dataDigest :=  crypto.Keccak256(data)

	// Verify
	crypto.VerifySignature(pubkey, digestHash, sigWithoutID)

	// Keccak hash
	// Sign
	//crypto.Keccak256()
	//crypto.Sign(digestHash, privateKey)

}
