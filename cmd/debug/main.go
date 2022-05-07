package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/login"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
}


/*

 */
func main() {
	//ds.Put()

	login.GenPrivateKey()
	logrus.Println("finish debug main")
}


// Use secp256k1
// See https://jinsnotes.com/2020-12-30-elliptical-curve-cryptography
func RecoverPrivkey(s string)  *ecdsa.PrivateKey {
	curve := secp256k1.S256()
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	d, _ := hexutil.DecodeBig(s)
	priv.D = d

	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(priv.D.Bytes())

	// DEBUG MSG
	logrus.Println("Private Key:", hexutil.Encode(crypto.FromECDSA(priv)))
	pub := priv.Public()
	publicKeyECDSA, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		logrus.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	logrus.Println("Public Key:", hexutil.Encode(publicKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	logrus.Println("Address:", address)

	return priv
}






