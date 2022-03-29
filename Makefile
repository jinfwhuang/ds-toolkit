clean-proto:
	rm -rf /Users/jin/code/repos/ds-sdk/build/github.com

proto:
	rm -rf /Users/jin/code/repos/ds-sdk/build/github.com/jinfwhuang/ds-sdk/proto
	protoc -I=./third_party/googleapis -I=./proto/identity --go_out=build --go-grpc_out=build proto/identity/login.proto
	cp build/github.com/jinfwhuang/ds-sdk/proto/identity/*  proto/identity/

# Targets that are not associated with explicit filename or filedir
.PHONY: clean proto
