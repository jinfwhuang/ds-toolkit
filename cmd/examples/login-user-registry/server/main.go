package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	cmd_utils "github.com/jinfwhuang/ds-toolkit/go-pkg/cmd-utils"
	eth_ds "github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	protoId "github.com/jinfwhuang/ds-toolkit/proto/identity"
)

const (
	OwnerAddr    = "0xaABcEa31ac2c76B5d11ad579d26A671D4F20171B"
	OwnerPrivkey = "0x8b88ac23aa273fb6c8b4c6783f883d157ad01d425abee25df014d7528c5c6452"
	ChainID      = 1337
)

var (
	ethAddr = &cli.StringFlag{
		Name:  "eth-address",
		Usage: "TODO: xxx",
		Value: "http://127.0.0.1:8545",
	}

	contractAddr = &cli.StringFlag{
		Name:  "contract-address",
		Usage: "TODO: xxx",
		Value: "0xC26f4289DFB6138C5f12e7D52D283F0Ce15FF985",
	}

	appFlags = []cli.Flag{
		cmd_utils.GrpcPort,
		cmd_utils.LogLevel,
		cmd_utils.LogCaller,
		ethAddr,
		contractAddr,
	}

	callOpt = &bind.CallOpts{
		Pending: true,
	}
)

func main() {
	app := cli.App{}
	app.Name = "UserRegistryLogin"
	app.Flags = appFlags
	app.Action = start
	if err := app.Run(os.Args); err != nil {
		logrus.Info(err)
	}
}

func start(cliCtx *cli.Context) error {
	cmd_utils.SetupLogrus(cliCtx)

	stop := make(chan struct{})

	// Tcp Listener
	grpcPort := cliCtx.Int(cmd_utils.GrpcPort.Name)
	address := fmt.Sprintf("%s:%d", "0.0.0.0", grpcPort)
	logrus.Infof("grpc address: %s", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatalf("could not listen to port in Start() %s: %v", address, err)
	}
	// GRPC
	grpcServer := grpc.NewServer()
	userRegistryServer := NewUserRegistryServer(cliCtx.String(ethAddr.Name), contractAddr.GetValue())
	logrus.Infof("ethaddr: ", ethAddr.GetValue())
	protoId.RegisterUserRegistryLoginServer(grpcServer, userRegistryServer)
	reflection.Register(grpcServer) // Enable reflection
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	logrus.Info("wait for stop channel to be closed.")
	<-stop
	return nil
}

type UserRegistryServer struct {
	protoId.UnimplementedUserRegistryLoginServer

	userRegistry *eth_ds.UserRegistry
	ethConn      *ethclient.Client
	userList     []*protoId.User
}

func NewUserRegistryServer(ethAddr, contractAddr string) *UserRegistryServer {
	ethconn, err := ethclient.Dial(ethAddr)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	userRegistryAcc := common.HexToAddress(contractAddr)
	userRegistry, err := eth_ds.NewUserRegistry(userRegistryAcc, ethconn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	return &UserRegistryServer{
		userRegistry: userRegistry,
		ethConn:      ethconn,
	}
}

func (s *UserRegistryServer) ListAllUsers(context.Context, *emptypb.Empty) (*protoId.UserList, error) {
	// Gathers all users by userID
	addrs, err := s.userRegistry.GetAllUsers(callOpt)
	if err != nil {
		log.Fatalf("Failed to retrieve all users: %v", err)
	}

	userList := make([]*protoId.User, len(addrs))
	for i, addr := range addrs {
		// Retrieving username with userId
		userName, err := s.userRegistry.GetName(callOpt, addr)
		if err != nil {
			log.Fatalf("Failed to find user with userId: %v", err)
		}
		// Retrieving first key with userId
		pubKey, err := s.userRegistry.GetKey(callOpt, addr, 0)
		if err != nil {
			log.Fatalf("Failed to find keys with userId %v", err)
		}

		userList[i] = &protoId.User{
			UserName: userName,
			PubKey:   pubKey.Key,
		}
	}
	// Assigning userList field of UserRegistryServer
	s.userList = userList

	return &protoId.UserList{
		Users: userList,
	}, nil
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

func (s *UserRegistryServer) AddUser(ctx context.Context, user *protoId.User) (*emptypb.Empty, error) {
	privkey, err := ecdsa_util.RecoverPrivkey(user.PrivKey)
	if err != nil {
		panic(err)
	}
	pubkey := crypto.FromECDSAPub(&privkey.PublicKey)

	userAddr := crypto.PubkeyToAddress(privkey.PublicKey)

	// Instantiate the contract and display its name
	nonce := getNonce(ctx, s.ethConn, common.HexToAddress(OwnerAddr))
	log.Println("nonce=", nonce)

	gasPrice, err := s.ethConn.SuggestGasPrice(context.Background())
	log.Println("gas price=", gasPrice)

	callOpt := bind.CallOpts{
		Pending: true,
	}
	txOpt := bind.TransactOpts{
		From:     common.HexToAddress(OwnerAddr),
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   signer,
		GasPrice: gasPrice.Mul(gasPrice, big.NewInt(2)),
	}

	sess := eth_ds.UserRegistrySession{
		Contract:     s.userRegistry,
		CallOpts:     callOpt,
		TransactOpts: txOpt,
	}
	tx, err := sess.NewUser(userAddr, user.UserName, eth_ds.Secp25661, eth_ds.Admin, pubkey)
	if err != nil {
		return nil, fmt.Errorf("Failed to add new user: %v", err)

	}
	log.Println("tx.nonce=", tx.Nonce())

	return &emptypb.Empty{}, nil
}

// GetUserByPubKey retrieves user using pubkey and returns error if not found.
func (s *UserRegistryServer) GetUserByPubKey(ctx context.Context, pubKey *protoId.PubKey) (*protoId.User, error) {
	if len(pubKey.PubKey) == 0 {
		return nil, errors.New("Public Key not specified")
	}

	for _, u := range s.userList {
		if bytes.Compare(u.PubKey, pubKey.PubKey) == 0 {
			return u, nil
		}
	}
	return nil, errors.New("User with public key not found")
}

// GetUserByUserName retrieves user using username and returns error if not found.
func (s *UserRegistryServer) GetUserByUserName(ctx context.Context, userName *protoId.UserName) (*protoId.User, error) {
	if len(userName.UserName) == 0 {
		return nil, errors.New("Username not specified")
	}

	for _, u := range s.userList {
		if u.UserName == userName.UserName {
			return u, nil
		}
	}
	return nil, errors.New("User with username not found")
}
