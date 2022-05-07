package ecdsa_util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/bytesutil"
)

func ZeroXToByte(s string) []byte {
	b, err := hexutil.Decode(s)
	if err != nil {
		panic(err)
	}
	return b
}


// Use secp256k1
// See https://jinsnotes.com/2020-12-30-elliptical-curve-cryptography
func RecoverPrivkey(s string) (*ecdsa.PrivateKey, error) {
	s, err := bytesutil.CheckNumber(s)
	if err != nil {
		return nil, err
	}
	key, err := crypto.HexToECDSA(s)
	if err != nil {
		return nil, err
	}
	return key, err
}

func RecoverPubkey(s string) (*ecdsa.PublicKey, error) {
	b, err := hexutil.Decode(s)
	if err != nil {
		return nil, err
	}
	pubkey, err := crypto.UnmarshalPubkey(b)
	if err != nil {
		return nil, err
	}
	return pubkey, nil
}


func ToAddress(pub []byte) (common.Address, error) {
	pubkey, err := crypto.UnmarshalPubkey(pub)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pubkey), nil
}







