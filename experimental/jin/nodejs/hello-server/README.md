```
# binary
bazel run //experimental/jinhuan/nodejs/hello-server:binary

# image
# Note that the binary target is linux
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //experimental/jinhuan/nodejs/hello-server:image

# deploy
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //experimental/jinhuan/nodejs/hello-server:mynamespace.apply
```
