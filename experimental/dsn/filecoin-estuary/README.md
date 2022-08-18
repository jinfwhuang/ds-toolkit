# Filecoin - Estuary

## Summary

Filecoin is a DSN provider, on which different implementations are built. One of the most used solutions is Estuary. Filecoin itself does not have implementation. Estuary is an invite-only platform, but getting approved was fast - next day. Filecoin is built with IPFS in mind, being complementary products with each other. Another pillar of Filecoin (and Estuary) is integration with Ethereum. Clients to access Estuary are - Javascript SDK, HTTP and CLI.

There are Swagger docs for the HTTP client and curl examples. However, I seemed to get the same error over and over. The error was present when testing with the CLI and HTTP client as well . I have contacted people from the Estuary team, but got no response. There is an attempt to run Estuary, using the HTTP client in `filecoin_estuary.go`.

The Javascript SDK was not tested.

Useful links:

- [Filecoin and IPFS](https://docs.filecoin.io/about-filecoin/ipfs-and-filecoin/)
- [Filecoin and Ethereum](https://filecoin.io/blog/posts/building-web3-filecoin-ethereum-better-together/)
- [Estuary](https://estuary.tech/)
- [API docs](https://docs.estuary.tech/)
