package ecdsa_util

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const (
	privkeyHexStr = "0x289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	privkeyStr = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	pubkeyHexStr = "0x04e32df42865e97135acfb65f3bae71bdc86f4d49150ad6a440b6f15878109880a0a2b2667f7e725ceea70c673093bf67663e0312623c8e091b13cf2c0f11ef652"
	xxxxxxxxxxxx = "0x047db227d7094ce215c3a0f57e1bcc732551fe351f94249471934567e0f5dc1bf795962b8cccb87a2eb56b29fbe37d614e2f4c3c45b789ae4f1f51f4cb21972ffd"
)

func init() {
	log.SetFlags(log.Llongfile)  // Make debugging easier
}

func TestRecoverPrivkey(t *testing.T) {
	key1, _ := RecoverPrivkey(privkeyStr)
	assert.Empty(t, key1)
	key2, _ := RecoverPrivkey(privkeyHexStr)
	assert.NotEmpty(t, key2)
	log.Println(key1)
	log.Println(key2)

	pubkeyByte := crypto.FromECDSAPub(&key2.PublicKey)
	log.Println(hex.EncodeToString(pubkeyByte))
}

func TestRecoverPubkey(t *testing.T) {
	key, err := RecoverPubkey(pubkeyHexStr)
	if err != nil {
		log.Fatal(err)
	}
	b := crypto.FromECDSAPub(key)
	keyHexStr := "0x" + hex.EncodeToString(b)

	assert.Equal(t, pubkeyHexStr, keyHexStr)
}




func TestSign(t *testing.T) {

	testPrivHex := "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
//var testpubkey  = hexutil.MustDecode("0x04e32df42865e97135acfb65f3bae71bdc86f4d49150ad6a440b6f15878109880a0a2b2667f7e725ceea70c673093bf67663e0312623c8e091b13cf2c0f11ef652")


	key, _ := crypto.HexToECDSA(testPrivHex)
	pubkey := crypto.FromECDSAPub(&key.PublicKey)
	//addr := common.HexToAddress(testAddrHex)

	msg := crypto.Keccak256([]byte("foo"))
	sig, err := crypto.Sign(msg, key)
	if err != nil {
		logrus.Errorf("Sign error: %s", err)
	}
	recoveredPub, err := crypto.Ecrecover(msg, sig)
	if err != nil {
		logrus.Fatalf("ECRecover error: %s", err)
	}


	// sign
	//digest := crypto.Keccak256([]byte("jin huang"))
	//sig, err := crypto.Sign(digest, privateKey)
	//if err != nil {
	//	logrus.Fatal(err)
	//}

	// Verify
	sig2 := sig[:len(sig)-1] // remove recovery id

	validated := crypto.VerifySignature(pubkey, msg, sig2)

	logrus.Info("pubkey           ", pubkey)
	logrus.Info("recovered pubkey ", recoveredPub)
	logrus.Info("pubkey len       ", len(pubkey))
	logrus.Info("msg len          ", len(msg))
	logrus.Info("sig len          ", len(sig))
	logrus.Info("result           ", validated)

}

func recoverPrivkey(s string)  *ecdsa.PrivateKey {
	curve := secp256k1.S256()
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	d, _ := hexutil.DecodeBig(s)
	priv.D = d

	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(priv.D.Bytes())
	return priv
}

