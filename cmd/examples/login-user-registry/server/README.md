## Running this example

- Make sure you are running an Ethereum node and have deployed the contract [(More information here)](/npm-pkg/contracts-identity/)
- Run both the client and server files to see the list of users and test the GetUser functions
- Another way to run ListAllUsers: `grpcurl -plaintext localhost:4000 jinfwhuang.dstoolkit.identity.UserRegistryLogin.ListAllUsers`
- To add users, use flags when running the client file: `go run . --user-name="name" --priv-key="Ox..."`
