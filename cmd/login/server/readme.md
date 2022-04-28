


## An example identityy service
An example service that demonstrates how to implement a login flow that uses key signing.

## Login Flow
- The client requests a sign-in message from the server.
- The sign-in message is generated and returned to the client.
- Singing mechanism:
  - The message is hashed by Keccak256Hash
  - The message hash is signed and a ECDSA signature is generated. The elliptical curve used is secp256k1.
- The signature is sent back to server for validation 

## Quick commands
```bash
go run ./cmd/login/server --log-caller=false
```

```bash
grpcurl -plaintext localhost:4000 list
grpcurl -plaintext localhost:4000 list jinfwhuang.dstoolkit.identity.Identity

# Debug
grpcurl -plaintext localhost:4000 jinfwhuang.dstoolkit.identity.Identity.Debug

# request login
grpcurl -plaintext -d '{
"pubKey": "4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12"
}' \
localhost:4000 jinfwhuang.dstoolkit.identity.Identity.RequestLogin
```

## Notes
- grpcurl automatically convert byte into base64 string
