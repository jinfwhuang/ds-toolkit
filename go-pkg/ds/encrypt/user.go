package encrypt

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


type UserData struct {
	User User
	Userid common.Address
	Name string
	privkey *ecdsa.PrivateKey
	pubkey *ecdsa.PublicKey
}

func (u *User) getRegistryPubkeys() []Pubkey {
	logrus.Fatal("not implemented")
	return []Pubkey{}
}

