package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	"github.com/sirupsen/logrus"
	"math/big"
)

func init() {
	//log.SetFlags(log.Llongfile)
	logrus.SetReportCaller(true)
}

const (
	OwnerAddr         = "0xaABcEa31ac2c76B5d11ad579d26A671D4F20171B"
	OwnerUsername = "jin.huang003"
	OwnerPrivkey = "0x8b88ac23aa273fb6c8b4c6783f883d157ad01d425abee25df014d7528c5c6452"
	RandAddr          = "0x4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12"
)

const (
	//UserRegistryContractAddr = "0xd52C607197467c4200d3cd02BFe667A42e05aa1d" // v1
	//UserRegistryContractAddr = "0x75B1BaB85bF717FbDAaeBfCaba6a2AeBd7f292Aa" // v2
	//UserRegistryContractAddr = "0xc9091eA0d76C760Aa080Af1E6aAfC127C373a563" // v3

	// Ropsten
	//RpcAddr           = "https://ropsten.infura.io/v3/f05e591322b64bdba3ee32673817254f"
	//UserRegistryContractAddr = "0x2925076a3f606BEaE300905632C2C4d92D32C610" // ropsten v4
	//ChainID = 3 // ropsten

	// hardhast
	RpcAddr           = "http://127.0.0.1:8545"
	UserRegistryContractAddr = "0x66e513FA7b40baCB7C1B0391FDE0af3232b1579E"
	ChainID = 1337 // hardhat


)

//KeyType keytype; // 1=secp256k1, 2=bls-12-381
//KeyStatus status; // 1=admin, 2=active, 3=canceled

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



func getChainId(ctx context.Context, client *ethclient.Client) {
	logrus.Info(client.ChainID(ctx))
}



func getPubkey(s string) []byte {
	privkey, err := ecdsa_util.RecoverPrivkey(OwnerPrivkey)
	if err != nil {
		logrus.Fatal(err)
	}
	pubkey := crypto.FromECDSAPub(&privkey.PublicKey)

	return pubkey
}

func getNonce(ctx context.Context, client *ethclient.Client, acc common.Address) uint64 {
	nonce, err := client.PendingNonceAt(ctx, acc)
	if err != nil {
		logrus.Fatal(err)
	}
	return nonce
}

func signer(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
	privkey, err := ecdsa_util.RecoverPrivkey(OwnerPrivkey)
	if err != nil {
		logrus.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(ChainID)), privkey)
	return signedTx, err
}

func getSigner(privkey *ecdsa.PrivateKey) bind.SignerFn {
	return func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		addrFromPrivkey := crypto.PubkeyToAddress(privkey.PublicKey)
		if addrFromPrivkey != addr {
			logrus.Fatal("address and privkey does not match")
		}
		signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(ChainID)), privkey)
		return signedTx, err
	}
}



type User struct {
	privkey string
	name string
	keyType KeyType
	keyStatus KeyStatus
}

// How to use go eth client: https://goethereumbook.org/transfer-eth/
// https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/transfer_eth.go
func newUser() {



	// Create an IPC based RPC connection to a remote node
	ctx := context.Background()
	ethconn, err := ethclient.Dial(RpcAddr)
	if err != nil {
		logrus.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	userRegistry, err := NewUserRegistry(common.HexToAddress(UserRegistryContractAddr), ethconn)
	if err != nil {
		logrus.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	userId := common.HexToAddress(OwnerAddr)
	name := OwnerUsername
	pubkey := getPubkey(OwnerPrivkey)
	nonce := getNonce(ctx, ethconn, common.HexToAddress(OwnerAddr))
	logrus.Info(nonce)
	txOpt := &bind.TransactOpts{
		From: userId,
		Nonce: big.NewInt(int64(nonce)),
		Signer: signer,
	}
	logrus.Info(txOpt)
	tx, err := userRegistry.NewUser(txOpt, userId, name, Secp25661, Admin, pubkey)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("nonce=", tx.Nonce())
	logrus.Info("type=", tx.Type())
	logrus.Info("to=", tx.To())
	logrus.Info("hash=", tx.Hash())
}


func updateUserStatus() {
	ethconn, err := ethclient.Dial(RpcAddr)
	if err != nil {
		logrus.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	userRegistry, err := NewUserRegistry(common.HexToAddress(UserRegistryContractAddr), ethconn)
	if err != nil {
		logrus.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	// Lookup name
	callOpt := &bind.CallOpts{
		Pending: true,
	}
	name, _ := userRegistry.GetName(callOpt, common.HexToAddress(OwnerAddr))
	logrus.Info(name)

	ctx := context.Background()
	userId := common.HexToAddress(OwnerAddr)
	nonce := getNonce(ctx, ethconn, common.HexToAddress(OwnerAddr))
	logrus.Info(nonce)

	gasPrice, err := ethconn.SuggestGasPrice(context.Background())
	logrus.Info(gasPrice)

	txOpt := &bind.TransactOpts{
		From: userId,
		Nonce: big.NewInt(int64(nonce + 1)),
		Signer: signer,
		GasPrice: gasPrice.Mul(gasPrice, big.NewInt(25)),
	}
	logrus.Info(txOpt)
	tx, err := userRegistry.UpdateKeyStatus(txOpt, userId, 0, 3) // Update to "Cancel"
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("nonce=", tx.Nonce())
	logrus.Info("type=", tx.Type())
	logrus.Info("to=", tx.To())
	logrus.Info("hash=", tx.Hash())
}


func inspectUserRegistry() {
	// Create an IPC based RPC connection to a remote node
	ctx := context.Background()
	ethconn, err := ethclient.Dial(RpcAddr)
	if err != nil {
		logrus.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	getChainId(ctx, ethconn)

	// Instantiate the contract and display its name
	userRegistryAcc := common.HexToAddress(UserRegistryContractAddr)
	userRegistry, err := NewUserRegistry(userRegistryAcc, ethconn)
	if err != nil {
		logrus.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	logrus.Info(userRegistry)

	nonce, err := ethconn.PendingNonceAt(ctx, userRegistryAcc)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(nonce)
	nonce, err = ethconn.NonceAt(ctx, userRegistryAcc, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(nonce)

	balance, _ := ethconn.BalanceAt(ctx, userRegistryAcc, nil)
	logrus.Info(balance)
	balance, _ = ethconn.BalanceAt(ctx, common.HexToAddress(OwnerAddr), nil)
	logrus.Info(balance)

	callOpt := &bind.CallOpts{
		Pending: true,
	}

	// Lookup name
	name, _ := userRegistry.GetName(callOpt, common.HexToAddress(OwnerAddr))
	logrus.Info(name)

	// Lookup id
	id, _ := userRegistry.GetUser(callOpt, name)
	logrus.Info(id)

	howmanyKeys, err := userRegistry.GetKeyLen(callOpt, common.HexToAddress(OwnerAddr))
	logrus.Info("how many keys:", howmanyKeys)

	// Lookup key
	tx, err := userRegistry.GetKey(callOpt, common.HexToAddress(OwnerAddr), 0)
	logrus.Println("admin=", Admin)
	logrus.Println("Cancel=", Cancel)
	logrus.Println("status=", KeyStatus(tx.Status))
	logrus.Println("type=", KeyType(tx.Keytype))

	pubkey, err := crypto.UnmarshalPubkey(tx.Key)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("public-key=", common.Bytes2Hex(tx.Key))
	logrus.Println("address=", crypto.PubkeyToAddress(*pubkey))
	logrus.Println("owner-address=", OwnerAddr)
	logrus.Println("user registry add=", UserRegistryContractAddr)
	// https://ropsten.etherscan.io/address/0xd52C607197467c4200d3cd02BFe667A42e05aa1d

}

func main() {
	//fastTokenContract()

	//newUser()

	//updateUserStatus()
	inspectUserRegistry()
}
