load("@build_bazel_rules_nodejs//:index.bzl", "nodejs_binary", "pkg_web")
load("@io_bazel_rules_docker//nodejs:image.bzl", "nodejs_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")
load("@com_adobe_rules_gitops//gitops:defs.bzl", "k8s_deploy")

package(default_visibility = ["//visibility:public"])

container_image(
    name = "image_base_exposed_ports",
    base = "@nodejs_image_base//image",
    docker_run_flags = "-i --publish=8080:8080",
)

nodejs_binary(
    name = "binary",
    data = [
        "app.js",
    ],
    entry_point = "app.js",
)

nodejs_image(
    name = "image",
    base = ":image_base_exposed_ports",
    binary = ":binary",
)

SERVICE_NAME = "hello-nodejs"

[k8s_deploy(
    name = "{}{}".format(
        SERVICE_NAME + "-" if ENV != "mynamespace" else "",
        ENV,
    ),
    cluster = CLUSTER,
    gitops = True if ENV != "mynamespace" else False,
    image_registry = "docker-horizon-snapshot.dr-uw2.adobeitc.com",

    # image
    images = {
        "application_image": ":image",
    },
    # manifest
    manifests = glob([
        "manifests/*.yaml",
    ]),
    namespace = NAMESPACE,
    substitutions = {
        "DNS_NAME": DNS_NAME,
    },

    # This needs to match the "user" in the context in the kube credential file ~/.kube/config
    user = "adobe",
    visibility = ["//visibility:public"],
) for ENV, NAMESPACE, CLUSTER, DNS_NAME in [
    ("mynamespace", "ns-personal-{BUILD_USER}", "ethos01devva6", "ethos01-dev-va6"),

    # This namespace has not been setup yet
    ("ccweb-dev", "ns-team-ccweb-dev", "ethos11prodva6", "ethos11-prod-va6"),
]]
