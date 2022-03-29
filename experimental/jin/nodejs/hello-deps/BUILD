load("@build_bazel_rules_nodejs//:index.bzl", "nodejs_binary")
load("@io_bazel_rules_docker//nodejs:image.bzl", "nodejs_image")

nodejs_binary(
    name = "binary",
    data = [
        "//experimental/jinhuan/nodejs/hello-lib:lib",
        "@npm_experiment_jinhuan_hello_dep//chalk",
        "@npm_experiment_jinhuan_hello_dep//color",
        "@npm_experiment_jinhuan_hello_dep//colors",
    ],
    entry_point = ":main.js",
)

nodejs_image(
    name = "image",
    # base = "@alpine_nodejs//image",
    base = "@nodejs_image_base//image",
    binary = ":binary",
)
