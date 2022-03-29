load("@build_bazel_rules_nodejs//:index.bzl", "js_library")

js_library(
    name = "lib",
    package_name = "@exphz/hello-lib",
    srcs = [
        "index.js",
        "package.json",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@npm_experiment_jinhuan_hello_lib//smallest",
    ],
)
