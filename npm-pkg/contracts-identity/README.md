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
# export abi
npx hardhat export-abi

# deploy
NETWORK=ropsten npx hardhat run scripts/deploy/user-registry.ts

# verify: publish the source code and make the contracts verifiable.
npx hardhat verify 0xe0822c07B9513B2bDC9bd0780f4AED4452dAc8c4
```

## Local dev

```bash
# Start a local chain
npx hardhat node

# Setup the accounts
NETWORK=localhost npx hardhat run scripts/hardhat-setup.ts

# Test
npx hardhat test
NETWORK=localhost npx hardhat test
```

## References:

- Ropsten faucet: https://faucet.metamask.io/
- https://docs.near.org/docs/concepts/account

## TODOs:
- verified accounts
- Implicit accounts
