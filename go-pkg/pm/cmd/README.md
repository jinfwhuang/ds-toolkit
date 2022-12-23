# Manage passwords on Arweave DLT

## Before run

  1. Export your Arweave wallet in JSON format and save it as `wallet.json` in `go-pkg/pm/cmd`.
  2. Create secp256k1 keypair and save the private and public keys in hex format as `priv_key` and `public_key` respectively in `go-pkg/pm/cmd`. (example: `0x289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032`, `0x047db227d7094ce215c3a0f57e1bcc732551fe351f94249471934567e0f5dc1bf795962b8cccb87a2eb56b29fbe37d614e2f4c3c45b789ae4f1f51f4cb21972ffd`)

## Run

## Create

Encrypt password using the public key and upload to Arweave, using your wallet.

1. Run `go run pm.go create -p "{PASSWORD}"` (i.e.: `go run pm.go create -p "my_secret_password123"`)
2. The output of the CLI command is the Arweave transaction, which you can use for fetching your password later.

## Get

Fetch password from Arweave using the tx id and decrypt it using the private key.
N.B.: it takes 2-5 minutes before a tx is available on Arweave chain.

1. Run `go run pm.go get -tx {ARWEAVE_TX_ID}` (i.e.: `go run pm.go get -tx jm72tlhv-13r7ytZj_eWfJvOvrRZ33_bFu5ncSLuWtY`)
2. The output of the CLI command is the password.
