package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/bytesutil"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	//"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"testing"
)

func init() {
	log.SetFlags(log.Llongfile)
	//log.SetReportCaller(true)
}

var (
	ethconn, _ = ethclient.Dial(RpcAddr)
	privkey, _ = ecdsa_util.RecoverPrivkey(OwnerPrivkey)

	// Instantiate the contract and display its name
	userRegistryAcc = common.HexToAddress(UserRegistryContractAddr)
	userRegistry, _ = NewUserRegistry(userRegistryAcc, ethconn)

	callOpt = bind.CallOpts{
		Pending: true,
	}

	ctx = context.Background()

)

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
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
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
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	name := _user.name
	nonce := getNonce(ctx, ethconn, common.HexToAddress(OwnerAddr))
	log.Println("nonce=", nonce)

	gasPrice, err := ethconn.SuggestGasPrice(context.Background())
	log.Println("gas price=", gasPrice)

	callOpt := bind.CallOpts{
		Pending: true,
	}
	txOpt := bind.TransactOpts{
		From: common.HexToAddress(OwnerAddr),
		Nonce: big.NewInt(int64(nonce)),
		Signer: signer,
		GasPrice: gasPrice.Mul(gasPrice, big.NewInt(2)),
	}

	sess := UserRegistrySession {
		Contract: userRegistry,
		CallOpts: callOpt,
		TransactOpts: txOpt,
	}
	log.Println(txOpt)
	//tx, err := userRegistry.NewUser(txOpt, userAddr, name, Secp25661, Admin, pubkey)
	tx, err := sess.NewUser(userAddr, name, Secp25661, Admin, pubkey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("nonce=", tx.Nonce())
	log.Println("type=", tx.Type())
	log.Println("to=", tx.To())
	log.Println("hash=", tx.Hash())
	log.Println("dfdf", string(tx.Data()))
	log.Println("pubkey bytes", hexutil.Encode(pubkey))
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
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	userRegistryAcc := common.HexToAddress(UserRegistryContractAddr)
	userRegistry, err := NewUserRegistry(userRegistryAcc, ethconn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
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

func Test_SolidityComputeAddr(t *testing.T) {
	privkey, err := ecdsa_util.RecoverPrivkey(OwnerPrivkey)
	if err != nil {
		panic(err)
	}
	pubkey := &privkey.PublicKey
	pubkeyBytes := crypto.FromECDSAPub(pubkey)
	addr := crypto.PubkeyToAddress(privkey.PublicKey)

	callOpt := &bind.CallOpts{
		Pending: true,
	}
	userRegistryAcc := common.HexToAddress(UserRegistryContractAddr)
	userRegistry, err := NewUserRegistry(userRegistryAcc, ethconn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	solidityAddr, err := userRegistry.ComputeAddr(callOpt, pubkeyBytes)

	log.Println("Address      :", addr)
	log.Println("Solidity addr:", solidityAddr)
}

func Test_AddressFromKey(t *testing.T) {
	privkey, err := ecdsa_util.RecoverPrivkey(OwnerPrivkey)
	if err != nil {
		panic(err)
	}
	pubkey := &privkey.PublicKey
	pubkeyBytes := crypto.FromECDSAPub(&privkey.PublicKey)
	addr := crypto.PubkeyToAddress(privkey.PublicKey)

	log.Println("Address:", addr)
	log.Println("Public Key:", hexutil.Encode(pubkeyBytes))
	log.Println("Public Key:", pubkeyBytes)
	log.Println("Public Key:", len(pubkeyBytes))
	log.Println("compressed form:", hexutil.Encode(secp256k1.CompressPubkey(pubkey.X, pubkey.Y)))

	_fromFull := crypto.Keccak256(pubkeyBytes)
	addrFull := common.BytesToAddress(_fromFull[12:])

	_fromComp := crypto.Keccak256(secp256k1.CompressPubkey(pubkey.X, pubkey.Y))
	addrComp := common.BytesToAddress(_fromComp[12:])

	log.Println("From Full", addrFull.Hex())
	log.Println("From comp", addrComp.Hex())
	log.Println("From func", PubkeyToAddress(*pubkey).Hex())

}

func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
	pubBytes := crypto.FromECDSAPub(&p)
	return common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
}

func packByte() {
	// https://gist.github.com/miguelmota/bc4304bb21a8f4cc0a37a0f9347b8bbb
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func Test_Sign_Recover(t *testing.T) {
	privkey, err := ecdsa_util.RecoverPrivkey(OwnerPrivkey)
	if err != nil {
		panic(err)
	}
	pubkey := &privkey.PublicKey
	pubkeyBytes := crypto.FromECDSAPub(&privkey.PublicKey)
	addr := crypto.PubkeyToAddress(privkey.PublicKey)


	log.Println("Address:", addr)
	log.Println("Public Key:", hexutil.Encode(pubkeyBytes))
	log.Println("compressed form:", hexutil.Encode(secp256k1.CompressPubkey(pubkey.X, pubkey.Y)))

	msg := []byte("random message")
	msgHash := crypto.Keccak256Hash(msg)
	//userRegistry.VerifyUser(callOpt, addr, )

	var nonce uint16 = 0
	nonceByte := make([]byte, 2)
	binary.LittleEndian.PutUint16(nonceByte, nonce)
	concatMsg := append(nonceByte, msgHash.Bytes()...) // nonce, message

	hashToSign := crypto.Keccak256Hash(concatMsg).Bytes()

	// 961562c9f1d50aba6a58ad744ac2b4efc7aa00df2538f78132c5dbf8e3796112
	log.Println("concat message", hexutil.Encode(concatMsg))
	log.Println("hashToSign", hexutil.Encode(hashToSign))

	sig, err := crypto.Sign(hashToSign, privkey)
	if err != nil {
		panic(err)
	}

	//SignatureLength := 65
	//sigWithoutID := sig[:len(sig)-1] // remove recovery id

	log.Println("signature length", len(sig))

	pubeyBytesRecovered, err := crypto.Ecrecover(hashToSign, sig)
	if err != nil {
		panic(err)
	}
	log.Println("Public Key:", hexutil.Encode(pubkeyBytes))
	log.Println("Public Key:", hexutil.Encode(pubeyBytesRecovered))
}


func Test_VerifyUser(t *testing.T) {
	// Create an IPC based RPC connection to a remote node
	//ctx := context.Background()
	ethconn, err := ethclient.Dial(RpcAddr)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	privkey, err := ecdsa_util.RecoverPrivkey(OwnerPrivkey)
	if err != nil {
		panic(err)
	}
	pubkey := &privkey.PublicKey
	pubkeyBytes := crypto.FromECDSAPub(&privkey.PublicKey)
	addr := crypto.PubkeyToAddress(privkey.PublicKey)


	log.Println("Address:", addr)
	log.Println("Public Key:", hexutil.Encode(pubkeyBytes))
	log.Println("compressed form:", hexutil.Encode(secp256k1.CompressPubkey(pubkey.X, pubkey.Y)))

	msg := []byte("random message")
	msgHash := crypto.Keccak256Hash(msg)

	var nonce uint16 = 0
	nonceByte := make([]byte, 2)
	binary.LittleEndian.PutUint16(nonceByte, nonce)
	concatMsg := append(nonceByte, msgHash.Bytes()...) // nonce, message

	log.Println(concatMsg)
	log.Println(hexutil.Encode(concatMsg))

	hashToSign := crypto.Keccak256Hash(concatMsg).Bytes()

	sig, err := crypto.Sign(hashToSign, privkey)
	if err != nil {
		panic(err)
	}



	// Instantiate the contract and display its name
	userRegistryAcc := common.HexToAddress(UserRegistryContractAddr)
	userRegistry, err := NewUserRegistry(userRegistryAcc, ethconn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	callOpt := &bind.CallOpts{
		Pending: true,
	}

	// How many users
	log.Println("-------")
	_addrs, err := userRegistry.GetAllUsers(callOpt)
	for i, _addr := range _addrs {
		log.Println(i, _addr.Hex())
	}
	log.Println("-------")

	// Lookup user nonce
	//addr := common.HexToAddress(OwnerAddr)
	nonce, err = userRegistry.GetUserNonce(callOpt, addr)
	if err != nil {
		panic(err)
	}
	log.Println("nonce from the server", nonce)

	// Call Verify function
	verified, err := userRegistry.VerifyUser(callOpt, addr, msgHash, sig[:len(sig)-1])
	log.Println(verified)



	// https://ropsten.etherscan.io/address/0xd52C607197467c4200d3cd02BFe667A42e05aa1d
}

func Test_AddPubkey(t *testing.T) {
	//pubkey := &privkey.PublicKey
	pubkeyBytes := crypto.FromECDSAPub(&privkey.PublicKey)
	addr := crypto.PubkeyToAddress(privkey.PublicKey)

	//bytes memory msgToKeccak = abi.encodePacked(user, keytype, keystatus, pubkey);
	//bytes32 msgToSign = keccak256(msgToKeccak);
	//address signAddr = msgToSign.recover(sig);

	//const msgToKeccak = ethers.utils.solidityPack(
	//["address", "uint8", "uint8", "bytes"],
	//[owner.address, 1, 1, randomPubkeyHex]);
	//let msgToSign = ethers.utils.keccak256(msgToKeccak);


	msgToKeccak := encodePacked(addr.Bytes(), []byte{1}, []byte{1}, pubkeyBytes);
	msgToSign := crypto.Keccak256(msgToKeccak)
	sig, err := crypto.Sign(msgToSign, privkey)

	log.Println("addr", addr.Hex())
	log.Println("msgToKeccak", hexutil.Encode(msgToKeccak))
	log.Println("msgToSign", hexutil.Encode(msgToSign))

	nonce := getNonce(ctx, ethconn, common.HexToAddress(OwnerAddr))
	gasPrice, err := ethconn.SuggestGasPrice(context.Background())
	txOpt := bind.TransactOpts{
		From: common.HexToAddress(OwnerAddr),
		Nonce: big.NewInt(int64(nonce)),
		Signer: signer,
		GasPrice: gasPrice.Mul(gasPrice, big.NewInt(2)),
	}

	//_, err = userRegistry.AddPubkey(&txOpt, addr, 1, 1, pubkeyBytes, sig[:len(sig)-1])
	_, err = userRegistry.AddPubkey(&txOpt, addr, 1, 1, pubkeyBytes, sig)
	if err != nil {
		panic(err)
	}


	// https://ropsten.etherscan.io/address/0xd52C607197467c4200d3cd02BFe667A42e05aa1d
}

func Test_AddMainUser(t *testing.T) {
	user := User {
		privkey: OwnerPrivkey,
		name: hexutil.Encode(bytesutil.RandBytes(13)),
		keyType: Secp25661,
		keyStatus: Admin,
	}

	addUser(user)
}