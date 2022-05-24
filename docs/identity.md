
## Identity Management
- UserRegistry is a publicly deployed smart contract
- Anyone can register an entry in UserRegistry; each entry is equivalent to a user
- Each user is uniquely identified by a 20 bytes ID
- Key rotation
- Human-readable name

- A [yaml example](yaml-examples/sample-id.yaml) of information contained in entry that describes a user 
- Implemented in [smart contract](../npm-pkg/contracts-identity)
- See [examples](../experimental/eth-client) on how to interact with the smart contract in golang

## Examples of What the Identity System Could be Used For

#### Login System
A backend service could leverage Elliptical Curve Digital Signature Algorithm (ECDSA) to implement a login system

- See [simple-login](cmd/examples/login-pubkey)
- See [pubkey-jwt-login](cmd/examples/login-user-registry); this is a work in progress
- Here is a [short introduction](https://jinsnotes.com/2020-12-30-elliptical-curve-cryptography#signing) to ECDSA

#### Data Encryption and
Users could use the identity system to manage and share data without intermediaries. 

- See documentation on the [data permission module](./data-permission.md)
- See how [data blob](../proto/ds/dsdata.proto) is represented 