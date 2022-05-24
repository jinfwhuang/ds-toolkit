
## Identity Management

The identity system does not require a centralized database. Users are registered in a decentralized network. The registry is implemented as an EVM smart contract. It could have been any smart contract system, or it could be implemented its own purpose built state machine. EVM smart contract is chosen purely for convenience. Once the contract is deployed, everyone has equal access to the identity system. Anyone can register an entry in the UserRegistry contract. Each entry is equivalent to a user. No one other than the user owner could make changes to their entry. The identity system is the foundation that many other applications could build on top.



#### Implementation
- EVM [smart contract](../npm-pkg/contracts-identity) implementation of the UserRegistry
- Here is a [yaml example](yaml-examples/sample-id.yaml) that 
describes a user
- Key features
  - Each user is uniquely identified by a 20 bytes ID
  - User could add more keys to their entry
  - User could rotate keys out of service
  - Each user has a human-readable name
- See [examples](../experimental/eth-client) on how to register and modify user entries

## Examples of What the Identity System Could be Used For

#### Login System
A backend service could leverage Elliptical Curve Digital Signature Algorithm (ECDSA) to implement a login system

- See [simple-login](cmd/examples/login-pubkey)
- See [pubkey-jwt-login](cmd/examples/login-user-registry); this is a work in progress
- Here is a [short introduction](https://jinsnotes.com/2020-12-30-elliptical-curve-cryptography#signing) to ECDSA

#### Data Wallet
Users could use the identity system to manage and share data without intermediaries. 

- See documentation on the [data permission module](./data-permission.md)
- See how [data blob](../proto/ds/dsdata.proto) is represented 