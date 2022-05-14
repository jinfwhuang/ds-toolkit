package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// TokenContractAddr = "0x2Be679b6cee5F2811144274c65ADAed1C1d96e24"
	TokenContractAddr = "0xC26f4289DFB6138C5f12e7D52D283F0Ce15FF985"
)

func fastTokenContract() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(RpcAddr)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	token, err := NewToken(common.HexToAddress(TokenContractAddr), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	// Query the contract name
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	log.Println("Token name:", name)

	// Query balance
	ownerAddr := common.HexToAddress(OwnerAddr)
	balance, err := token.BalanceOf(&bind.CallOpts{}, ownerAddr)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	log.Printf("address=%s balance=%d", ownerAddr, balance)

	ownerAddr = common.HexToAddress(RandAddr) // This address has balance zero
	balance, err = token.BalanceOf(&bind.CallOpts{}, ownerAddr)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	log.Printf("address=%s balance=%d", ownerAddr, balance)

}



