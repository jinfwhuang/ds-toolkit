clean-proto:
	rm -rf build/github.com

proto:
	rm -rf build/github.com/jinfwhuang/ds-toolkit/proto
	protoc -I=./third_party/googleapis -I=./proto/identity --go_out=build --go-grpc_out=build proto/identity/login.proto
	cp build/github.com/jinfwhuang/ds-toolkit/proto/identity/*  proto/identity/

# Targets that are not associated with explicit filename or filedir
.PHONY: clean proto
