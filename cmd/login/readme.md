


## An example login service

An example service that demonstrates how to implement a key signing login flow.

```bash
go run ./cmd/login


```

- grpcurl automatically convert byte into base64 string


```bash
grpcurl -plaintext localhost:4000 list
grpcurl -plaintext localhost:4000 list jinfwhuang.ds.identity.Identity

# get
grpcurl -plaintext localhost:4000 jinfwhuang.ds.identity.Identity.Debug

# request login
grpcurl -plaintext -d '{
"pubKey": "4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12"
}' \
localhost:4000 jinfwhuang.ds.identity.Identity.RequestLogin

# login
grpcurl -plaintext -d '{
"pubKey": "4E6B0228A5bc0Ca7f2a8bfaC93B13aA9cc506F12"
}' \
localhost:4000 jinfwhuang.ds.identity.Identity.Login


# post request

-d '{"destination": "172.31.255.0", "count": 2, "do_not_resolve":true}'
```