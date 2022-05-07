```bash
# Copy over ABI
cp -r modules/contracts/artifacts/abi experimental/jin/go/cmd/web3/abi

# Generate code
go run github.com/ethereum/go-ethereum/cmd/abigen --abi experimental/eth-client/abi/Token.json --pkg main --type Token --out experimental/eth-client/token.go

# Run main
go run ./experimental/jin/go/cmd/web3
```
