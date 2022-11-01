package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/jinfwhuang/ds-toolkit/go-pkg/ds"
	"github.com/jinfwhuang/ds-toolkit/go-pkg/dsn"
	ecdsa_util "github.com/jinfwhuang/ds-toolkit/go-pkg/ecdsa-util"
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

func main() {
	app := &cli.App{
		Name:     "pm",
		Version:  "0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Jin Huang",
				Email: "huang.jin.f@gmail.com",
			},
			{
				Name:  "Kaloyan Tanev",
				Email: "tanevdev@gmail.com",
			},
		},
		HelpName:  "pm",
		Usage:     "pm [pm options] command [command options] params",
		UsageText: "Create, update and manage passwords, stored on a decentralised storage network",
		Commands: []*cli.Command{
			{
				Name:        "encryptPassword",
				Usage:       "Encrypt password using ECDSA public key in Hex format",
				Description: "Encrypt password using ECDSA public key in Hex format",
				ArgsUsage:   "",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "password",
						Aliases: []string{"p", "pass"},
						Value:   "",
						Usage:   "Password to encrypt in string format",
					},
					&cli.StringFlag{
						Name:    "publicKey",
						Aliases: []string{"k", "pk", "key"},
						Value:   "",
						Usage:   "Public key to encrypt the password in hex format",
					},
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				HelpName:        "",
				Action: func(c *cli.Context) error {
					password := c.String("password")
					keyHex := c.String("key")
					if password == "" {
						return errors.New("no password provided")
					}
					if keyHex == "" {
						return errors.New("no public key provided")
					}

					key, err := ecdsa_util.RecoverPubkey(keyHex)
					if err != nil {
						return errors.New("could not recover ecdsa public key from hex")
					}

					dataBlob, err := ds.CreateDataBlob([]byte(password), key)
					if err != nil {
						return errors.New("could not encrypt password")
					}

					res, err := json.Marshal(dataBlob)
					if err != nil {
						return errors.New("could not marshal encrypted password to JSON")
					}

					println(string(res))

					return nil
				},
			},
			{
				Name:        "decryptPassword",
				Usage:       "Decrypt password using ECDSA private key in Hex format",
				Description: "Decrypt password using ECDSA private key in Hex format",
				ArgsUsage:   "",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "passwordBlob",
						Aliases: []string{"p", "pass", "blob"},
						Value:   "",
						Usage:   "Password blob to dencrypt in string format",
					},
					&cli.StringFlag{
						Name:    "key",
						Aliases: []string{"k", "pk"},
						Value:   "",
						Usage:   "Private key to decrypt the password blob in hex format",
					},
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				HelpName:        "",
				Action: func(c *cli.Context) error {
					stringBlob := c.String("passwordBlob")
					keyHex := c.String("key")
					if stringBlob == "" {
						return errors.New("no password provided")
					}
					if keyHex == "" {
						return errors.New("no private key provided")
					}

					key, err := ecdsa_util.RecoverPrivkey(keyHex)
					if err != nil {
						return errors.New("could not recover ecdsa private key from HEX")
					}

					var blob protods.DataBlob
					err = json.Unmarshal([]byte(stringBlob), &blob)
					if err != nil {
						return errors.New("could not unmarshal password blob to JSON")
					}

					decryptedPassword, err := ds.ExtractData(&blob, key)
					if err != nil {
						println(err.Error())
						return errors.New("could not decrypt password blob")
					}

					println(string(decryptedPassword))

					return nil
				},
			},
			{
				Name:        "uploadPassword",
				Usage:       "Upload password to Arweave DSN using encrypted password blob and JWK",
				Description: "Upload password to Arweave DSN using encrypted password blob and JWK",
				ArgsUsage:   "",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "passwordBlob",
						Aliases: []string{"p", "pass", "blob"},
						Value:   "",
						Usage:   "Password blob to dencrypt in string format",
					},
					&cli.StringFlag{
						Name:    "wallet",
						Aliases: []string{"w"},
						Value:   "",
						Usage:   "Arweave wallet in JWK format",
					},
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				HelpName:        "",
				Action: func(c *cli.Context) error {
					stringBlob := c.String("passwordBlob")
					walletString := c.String("wallet")
					if stringBlob == "" {
						return errors.New("no password provided")
					}
					if walletString == "" {
						return errors.New("no JWK provided")
					}

					wallet, err := dsn.GenerateWalletFromJWK([]byte(walletString))
					if err != nil {
						return errors.New("could not recover wallet from JWK")
					}

					var blob protods.DataBlob
					err = json.Unmarshal([]byte(stringBlob), &blob)
					if err != nil {
						return errors.New("could not unmarshal password blob to JSON")
					}

					blobBytes, err := json.Marshal(&blob)
					if err != nil {
						return errors.New("could not marshal password blob from JSON")
					}

					txId, err := dsn.Write(blobBytes, wallet)
					if err != nil {
						return errors.New("could not upload to Arweave")
					}

					fmt.Printf("Transaction ID: %v", txId)

					return nil
				},
			},
			{
				Name:        "retrievePassword",
				Usage:       "Retrieve password from Arweave DSN using transaction ID",
				Description: "Retrieve password from Arweave DSN using transaction ID",
				ArgsUsage:   "",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "transactionID",
						Aliases: []string{"tx", "id", "txid", "txId", "txID"},
						Value:   "",
						Usage:   "Transaction ID on Arweave",
					},
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				HelpName:        "",
				Action: func(c *cli.Context) error {
					id := c.String("txID")
					if id == "" {
						return errors.New("no transaction ID provided")
					}

					tx, err := dsn.Read(id)
					if err != nil {
						return errors.New("could not read Arweave transaction")
					}

					println(string(tx))

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		panic(err)
	}
}
