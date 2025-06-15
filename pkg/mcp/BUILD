load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "mcp",
    srcs = [
        "server.go",
        "tools.go",
    ],
    importpath = "k8s.io/kubernetes/cmd/kubectl-mcp/pkg/mcp",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_mark3labs_mcp_go//mcp:go_default_library",
        "@com_github_mark3labs_mcp_go//server:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
        "@com_github_spf13_pflag//:go_default_library",
        "@io_k8s_kubectl//pkg/cmd:go_default_library",
        "@io_k8s_kubectl//pkg/cmd/util:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
