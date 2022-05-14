package login

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	log "log"
)

func init() {
	log.SetFlags(log.Llongfile)
}

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


