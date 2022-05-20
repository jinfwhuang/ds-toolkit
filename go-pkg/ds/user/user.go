package user

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/sirupsen/logrus"
	log "log"
)

func GenPrivateKey() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	log.Println("SAVE BUT DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKeyBytes))

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyBytesCompressed := secp256k1.CompressPubkey(publicKeyECDSA.X, publicKeyECDSA.Y)

	log.Println("Public Key:", hexutil.Encode(publicKeyBytes))
	log.Println("Public Key, compressed form:", hexutil.Encode(publicKeyBytesCompressed))

	x, y := secp256k1.DecompressPubkey(publicKeyBytesCompressed)
	_key := &ecdsa.PublicKey {
		Curve: elliptic.P256(),
		X: x,
		Y: y,
	}
	log.Println(x, y)
	log.Println(publicKeyECDSA.X, publicKeyECDSA.Y)
	log.Println(hexutil.Encode(crypto.FromECDSAPub(_key)))
	log.Println(hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA)))


	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	log.Println("Address:", address)
}

type User struct {
	Userid common.Address
	Name string
	privkey *ecdsa.PrivateKey
	pubkey *ecdsa.PublicKey
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
	crypto.Sign(digestHash, privateKey)

}

func (u *User) VerifySignature(data []byte, sig []byte) bool {
	// Keccak hash
	dataDigest :=  crypto.Keccak256(data)

	// Verify
	crypto.VerifySignature(pubkey, digestHash, sigWithoutID)

	// Keccak hash
	// Sign
	//crypto.Keccak256()
	//crypto.Sign(digestHash, privateKey)

}

