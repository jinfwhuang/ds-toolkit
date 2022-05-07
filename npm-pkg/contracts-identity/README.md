

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
npx hardhat verify 0xe0822c07B9513B2bDC9bd0780f4AED4452dAc8c4
```

## References:
- Ropsten faucet: https://faucet.metamask.io/
