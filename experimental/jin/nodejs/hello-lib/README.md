```bash
bazel run //experimental/jinhuan/nodejs/hello-lib:binary

bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64  //experimental/jinhuan/nodejs/hello-deps:image
```
