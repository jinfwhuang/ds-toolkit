# Decenetralized Storage Networks research

## Arweave

Arweave tries to achieve permanent storage in the web, however, this can only be the objective and the ecosystem can only take technological decisions to be as close as possible to “permaweb”. The storage persistence highly depends on the incentives the miners have and of course how highly the token is valued. This does not differ than any other protocol/service/DLT. Arweave uses for consensus algorithm Succinct Proof of Random Access (SPoRA). SPoRA is similar to the traditional PoW blockchain, but the newly mined block links not only to the previous block, but also to a random historical block in the chain, called “recall” block. This adds another unique mechanism to traverse the chain, instead only by the linked list in the blockchain. Because of that Arweave is a blockweave, rather than a blockchain. Given that the Arweave stores huge amounts of data (compared to Bitcoin for example), miners are not required to store the full weave in order to be eligible to create a new block. Moreover, in order a miner to mine a new block, the miner needs to prove it has stored the chosen recall block. Because of this, miners are incentivised to store the most rare blocks, in order to have the highest chance of being chosen to mine the next block. This is the main mechanic ensuring that no matter how rarely a data chunk is used or how big it is, miners are still incentivised to store it, so that they can mine the next block.

Further reads:

- [Yellow Paper [3.2 Token Economy; page 15]](https://www.arweave.org/yellow-paper.pdf) (overall - well written, highly recommend)
- [Overview of Arweave](https://www.texasblockchain.org/blog/arweave-paying-for-permanence)
- [Arweave use cases](https://cryptowallet.com/academy/arweave-use-case/)
- [SPoRA consensus algorithm](https://arweave.medium.com/the-arweave-network-is-now-running-succinct-random-proofs-of-access-spora-e2732cbcbb46)

*Developer's notes: Easy to implement in golang, had well maintained SDK. Everything worked as expected and has good documentation.*

## Ceramic

- A modified IPFS
- On Ceramic, each piece of information is represented as an append-only log of commits, called a Stream. Each stream is a DAG stored in IPLD.
- A stream has immutable streamid

*Developer's notes: Quite bad documentation. Managed to run using the CLI, but didn't manage using HTTP. Ceramic have only JavaScript SDK, which is untested.*

## Celestia

- formerly lazyledger
- a data availability-focus blockchain

## Estuary

- https://github.com/application-research/estuary
- An experimental ipfs node
- https://estuary.tech/home
- A reliable way to upload public data onto Filecoin and pin it to IPFS.

*Developer's notes: The example code they provide does not work. Golang SDK and curl request via HTTP client were tested. Contacted support, no response for 4 days already.*

## Storj

*Developer's notes: Worked quite well with the golang SDK, no issues.*

## Sia

*Developer's notes: Replica of the full node is required in order to proceed. No running client on either Sia's premises or community. Quite slow download of the whole chain, most likely will get worse with time.*
