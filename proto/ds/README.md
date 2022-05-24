





```bash
rm -rf build/github.com/jinfwhuang/ds-toolkit/proto/

protoc \
-I=./third_party/googleapis -I=./proto \
--go_out=build --go-grpc_out=build \
proto/**/*.proto

cp -r build/github.com/jinfwhuang/ds-toolkit/proto/  proto

```