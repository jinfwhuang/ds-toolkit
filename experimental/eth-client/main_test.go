package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/bytesutil"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"testing"
)

func init() {
	log.SetFlags(log.Llongfile)
	//logrus.SetReportCaller(true)
}

func TestGenPrivateKey(t *testing.T) {
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

var users = []User{
	{"0x89df41bce61452e4bed3e6325873022ba877a1683cfca9d03c0a5bf542944321", "jin001", Secp25661, Admin},
	{"0xa329248cb59ccd5429a67e11529a831147bf1193bafcc45255129288362095cb", "jin001", Secp25661, Admin},
	{"0x9eb5b98580d351a7607732cf85a03fac5e2a86e9f3eaf833470c80c316d17c6b", "jin001", Secp25661, Admin},
	{"0x045294d477c6c43c7b564a2df54f281d7ad4a0a19b533cf793fc6a576a1a28be", "jin001", Secp25661, Admin},
	{"0xc93e2f86fc1457a956ee366ecd43d2bfdcda2837f48eeca746dbd4a8ef6e2fc8", "jin001", Secp25661, Admin},
}

func addUser(_user User) {
	ctx := context.Background()
	ethconn, err := ethclient.Dial(RpcAddr)
	if err != nil {
		logrus.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privkey, err := ecdsa_util.RecoverPrivkey(_user.privkey)
	if err != nil {
		panic(err)
	}
	pubkey := crypto.FromECDSAPub(&privkey.PublicKey)
	userAddr := crypto.PubkeyToAddress(privkey.PublicKey)

	// Instantiate the contract and display its name
	userRegistry, err := NewUserRegistry(common.HexToAddress(UserRegistryContractAddr), ethconn)
	if err != nil {
		logrus.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	name := _user.name
	nonce := getNonce(ctx, ethconn, common.HexToAddress(OwnerAddr))
	log.Println("nonce=", nonce)

	gasPrice, err := ethconn.SuggestGasPrice(context.Background())
	log.Println("gas price=", gasPrice)

	txOpt := &bind.TransactOpts{
		From: common.HexToAddress(OwnerAddr),
		Nonce: big.NewInt(int64(nonce)),
		Signer: signer,
		GasPrice: gasPrice.Mul(gasPrice, big.NewInt(2)),
	}
	log.Println(txOpt)
	tx, err := userRegistry.NewUser(txOpt, userAddr, name, Secp25661, Admin, pubkey)
	if err != nil {
		logrus.Fatal(err)
	}
	log.Println("nonce=", tx.Nonce())
	log.Println("type=", tx.Type())
	log.Println("to=", tx.To())
	log.Println("hash=", tx.Hash())
}

func getRandPrivkey() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	return hexutil.Encode(privateKeyBytes)
}

func TestAddOneUser(t *testing.T) {
	user := User {
		privkey: getRandPrivkey(),
		name: hexutil.Encode(bytesutil.RandBytes(13)),
		keyType: Secp25661,
		keyStatus: Admin,
	}

	addUser(user)
}


func Test_GetAllUsers(t *testing.T) {
	// Create an IPC based RPC connection to a remote node
	//ctx := context.Background()
	ethconn, err := ethclient.Dial(RpcAddr)
	if err != nil {
		logrus.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	userRegistryAcc := common.HexToAddress(UserRegistryContractAddr)
	userRegistry, err := NewUserRegistry(userRegistryAcc, ethconn)
	if err != nil {
		logrus.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	callOpt := &bind.CallOpts{
		Pending: true,
	}

	// How many users
	log.Println("-------")
	addrs, err := userRegistry.GetAllUsers(callOpt)
	for i, addr := range addrs {
		log.Println(i, addr.Hex())
	}
	log.Println("-------")
	// https://ropsten.etherscan.io/address/0xd52C607197467c4200d3cd02BFe667A42e05aa1d
}