```bash
# Copy over ABI
cp -r npm-pkg/contracts-identity/artifacts/abi experimental/eth-client/

# Generate code
go run github.com/ethereum/go-ethereum/cmd/abigen --abi experimental/eth-client/abi/Token.json --pkg main --type Token --out experimental/eth-client/token.go

export name=UserRegistry
go run github.com/ethereum/go-ethereum/cmd/abigen --abi "experimental/eth-client/abi/${name}.json" --pkg main --type "${name}" --out "experimental/eth-client/${name}.go"

# Run main
go run ./experimental/jin/go/cmd/web3
```
