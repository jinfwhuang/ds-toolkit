

## Quick Commands
```bash
# Use different networks
NETWORK=mainnet npx hardhat fast
NETWORK=ropsten npx hardhat fast

# compile
npx hardhat compile 

# export abi
npx hardhat export-abi

# Test
npx hardhat test

# Solidity Linting
yarn lint:sol
```

## Deploy
```bash
# deploy
NETWORK=ropsten npx hardhat run scripts/deploy.ts

# verify: publish the source code and make the contracts verifiable.
npx hardhat verify 0x6B66f47D14ED039D6C0492b85a18c1cd85762Ac5
```
