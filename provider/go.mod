module github.com/pulumi/pulumi-kind/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
	github.com/kyma-incubator/terraform-provider-kind v0.0.11 => github.com/w9n/terraform-provider-kind v0.0.12-0.20220219020836-8e1e3a1971aa
)

require (
	github.com/kyma-incubator/terraform-provider-kind v0.0.11
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.13.0
	github.com/pulumi/pulumi/sdk/v3 v3.19.0
)

replace k8s.io/client-go => k8s.io/client-go v0.20.2
