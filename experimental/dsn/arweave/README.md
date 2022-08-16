# Arweave

## Summary

Arweave worked seamlessly, using the GoLang SDK.

Other ways to connect with the Arweave protocol are: Javascript SDK, php SDK, Scala SDK, HTTP client.
Those have not been tested and cannot guarantee completeness or ease of use.

Useful links:

- [Developer docs](https://docs.arweave.org/developers/)
- [GoLang SDK repository](https://github.com/everFinance/goar)
- [Web wallet guide](https://docs.arweave.org/info/wallets/arweave-wallet)

## How to run - Go SDK

1. [Create new wallet](https://arweave.app/)
2. Fund the wallet with Arweave Coin
3. [Export the wallet](https://arweave.app/settings)
4. Run `arweave.Write`, providing data bytes and wallet path.
5. Transaction is sent to the Arweave mainet.
6. In 1-3 minutes your transaction will be minted. Run `arweave.Read`, providing transaction ID.
7. The byte slice returned should be the data you have uploaded initially.
