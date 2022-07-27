## Running this example on command prompt

- Make sure you are running an Ethereum node and have deployed the contract [(More information here)](/npm-pkg/contracts-identity/)
- Run both the client and server files to see the list of users and test the GetUser functions
- Another way to run ListAllUsers: `grpcurl -plaintext localhost:4000 jinfwhuang.dstoolkit.identity.UserRegistryLogin.ListAllUsers`
- To add users, use flags when running the client file: `go run . --user-name="name" --priv-key="Ox..."`

## Using a simple web application

### Set Up

- You may need to run `npm install google-protobuf --save` and `npm install googleapis` to download necessary packages
- To generate the protobuf messages and client service stub class from your .proto definitions, run `bash protoc \ -I=./third_party/googleapis -I=./proto proto/identity/login.proto \ --js_out=import_style=commonjs:build \ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:build `
- Make sure you have downloaded `protoc-gen-grpc-web plugin` and installed `protoc` ([Reference here for more information](https://github.com/grpc/grpc-web/tree/master/net/grpc/gateway/examples/helloworld))

### To run the example

- Make sure you are running an Ethereum node and have deployed the contract ([More information here](/npm-pkg/contracts-identity/))
- Run envoy using `envoy -c envoy.yaml`
- Run the React web application using `npm start` in the correct directory
- Test by inputting the username and corresponding private key (should already be registered in the contract) to login
