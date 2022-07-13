package ds

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

const (
	NotImplemented = "not implemented"
)

type KeyType uint8

const (
	Secp25661 = iota + 1
	Bls12381
)

func (e KeyType) String() string {
	return [...]string{
		"Secp25661",
		"Bls12381",
	}[e-1]
}

type KeyStatus uint8

const (
	Admin = iota + 1
	Active
	Cancel
)

func (e KeyStatus) String() string {
	return [...]string{
		"Admin",
		"Active",
		"Cancel",
	}[e-1]
}

type User struct {
	Userid  common.Address
	Name    string
	privkey *ecdsa.PrivateKey
	pubkey  *ecdsa.PublicKey
}

func (u *User) getRegistryPubkeys() []Pubkey {
	logrus.Fatal("not implemented")
	return []Pubkey{}
}

func (u *User) getPubkey() *ecdsa.PublicKey {
	if u.pubkey == nil {
		publicKey := u.privkey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			logrus.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		u.pubkey = publicKeyECDSA
	}

	return u.pubkey
}

func (u *User) getPubkeyBytes() []byte {
	return crypto.FromECDSAPub(&u.privkey.PublicKey)
}

// The produced signature is in the [R || S || V] format where V is 0 or 1.
func (u *User) Sign(data []byte) (sig []byte, err error) {
	// Keccak hash
	// Sign
	//crypto.Keccak256()
	//crypto.Sign(digestHash, privateKey)

	panic(NotImplemented)
}

func (u *User) VerifySignature(data []byte, sig []byte) bool {
	// Keccak hash
	//dataDigest :=  crypto.Keccak256(data)

	//// Verify
	//crypto.VerifySignature(pubkey, digestHash, sigWithoutID)

	// Keccak hash
	// Sign
	//crypto.Keccak256()
	//crypto.Sign(digestHash, privateKey)

	panic(NotImplemented)

}
