// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kind

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Kind_cluster struct {
	pulumi.CustomResourceState

	// Client certificate for authenticating to cluster.
	ClientCertificate pulumi.StringOutput `pulumi:"clientCertificate"`
	// Client key for authenticating to cluster.
	ClientKey pulumi.StringOutput `pulumi:"clientKey"`
	// Client verifies the server certificate with this CA cert.
	ClusterCaCertificate pulumi.StringOutput `pulumi:"clusterCaCertificate"`
	// Kubernetes APIServer endpoint.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The kind_config that kind will use to bootstrap the cluster.
	KindConfig Kind_clusterKindConfigPtrOutput `pulumi:"kindConfig"`
	// Kubeconfig set after the the cluster is created.
	Kubeconfig pulumi.StringOutput `pulumi:"kubeconfig"`
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath pulumi.StringOutput `pulumi:"kubeconfigPath"`
	// The kind name that is given to the created cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The node_image that kind will use (ex: kindest/node:v1.15.3).
	NodeImage pulumi.StringOutput `pulumi:"nodeImage"`
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false
	WaitForReady pulumi.BoolPtrOutput `pulumi:"waitForReady"`
}

// NewKind_cluster registers a new resource with the given unique name, arguments, and options.
func NewKind_cluster(ctx *pulumi.Context,
	name string, args *Kind_clusterArgs, opts ...pulumi.ResourceOption) (*Kind_cluster, error) {
	if args == nil {
		args = &Kind_clusterArgs{}
	}

	var resource Kind_cluster
	err := ctx.RegisterResource("kind:index/kind_cluster:kind_cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKind_cluster gets an existing Kind_cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKind_cluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Kind_clusterState, opts ...pulumi.ResourceOption) (*Kind_cluster, error) {
	var resource Kind_cluster
	err := ctx.ReadResource("kind:index/kind_cluster:kind_cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Kind_cluster resources.
type kind_clusterState struct {
	// Client certificate for authenticating to cluster.
	ClientCertificate *string `pulumi:"clientCertificate"`
	// Client key for authenticating to cluster.
	ClientKey *string `pulumi:"clientKey"`
	// Client verifies the server certificate with this CA cert.
	ClusterCaCertificate *string `pulumi:"clusterCaCertificate"`
	// Kubernetes APIServer endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The kind_config that kind will use to bootstrap the cluster.
	KindConfig *Kind_clusterKindConfig `pulumi:"kindConfig"`
	// Kubeconfig set after the the cluster is created.
	Kubeconfig *string `pulumi:"kubeconfig"`
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath *string `pulumi:"kubeconfigPath"`
	// The kind name that is given to the created cluster.
	Name *string `pulumi:"name"`
	// The node_image that kind will use (ex: kindest/node:v1.15.3).
	NodeImage *string `pulumi:"nodeImage"`
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false
	WaitForReady *bool `pulumi:"waitForReady"`
}

type Kind_clusterState struct {
	// Client certificate for authenticating to cluster.
	ClientCertificate pulumi.StringPtrInput
	// Client key for authenticating to cluster.
	ClientKey pulumi.StringPtrInput
	// Client verifies the server certificate with this CA cert.
	ClusterCaCertificate pulumi.StringPtrInput
	// Kubernetes APIServer endpoint.
	Endpoint pulumi.StringPtrInput
	// The kind_config that kind will use to bootstrap the cluster.
	KindConfig Kind_clusterKindConfigPtrInput
	// Kubeconfig set after the the cluster is created.
	Kubeconfig pulumi.StringPtrInput
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath pulumi.StringPtrInput
	// The kind name that is given to the created cluster.
	Name pulumi.StringPtrInput
	// The node_image that kind will use (ex: kindest/node:v1.15.3).
	NodeImage pulumi.StringPtrInput
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false
	WaitForReady pulumi.BoolPtrInput
}

func (Kind_clusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kind_clusterState)(nil)).Elem()
}

type kind_clusterArgs struct {
	// The kind_config that kind will use to bootstrap the cluster.
	KindConfig *Kind_clusterKindConfig `pulumi:"kindConfig"`
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath *string `pulumi:"kubeconfigPath"`
	// The kind name that is given to the created cluster.
	Name *string `pulumi:"name"`
	// The node_image that kind will use (ex: kindest/node:v1.15.3).
	NodeImage *string `pulumi:"nodeImage"`
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false
	WaitForReady *bool `pulumi:"waitForReady"`
}

// The set of arguments for constructing a Kind_cluster resource.
type Kind_clusterArgs struct {
	// The kind_config that kind will use to bootstrap the cluster.
	KindConfig Kind_clusterKindConfigPtrInput
	// Kubeconfig path set after the the cluster is created or by the user to override defaults.
	KubeconfigPath pulumi.StringPtrInput
	// The kind name that is given to the created cluster.
	Name pulumi.StringPtrInput
	// The node_image that kind will use (ex: kindest/node:v1.15.3).
	NodeImage pulumi.StringPtrInput
	// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false
	WaitForReady pulumi.BoolPtrInput
}

func (Kind_clusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kind_clusterArgs)(nil)).Elem()
}

type Kind_clusterInput interface {
	pulumi.Input

	ToKind_clusterOutput() Kind_clusterOutput
	ToKind_clusterOutputWithContext(ctx context.Context) Kind_clusterOutput
}

func (*Kind_cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind_cluster)(nil)).Elem()
}

func (i *Kind_cluster) ToKind_clusterOutput() Kind_clusterOutput {
	return i.ToKind_clusterOutputWithContext(context.Background())
}

func (i *Kind_cluster) ToKind_clusterOutputWithContext(ctx context.Context) Kind_clusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Kind_clusterOutput)
}

// Kind_clusterArrayInput is an input type that accepts Kind_clusterArray and Kind_clusterArrayOutput values.
// You can construct a concrete instance of `Kind_clusterArrayInput` via:
//
//          Kind_clusterArray{ Kind_clusterArgs{...} }
type Kind_clusterArrayInput interface {
	pulumi.Input

	ToKind_clusterArrayOutput() Kind_clusterArrayOutput
	ToKind_clusterArrayOutputWithContext(context.Context) Kind_clusterArrayOutput
}

type Kind_clusterArray []Kind_clusterInput

func (Kind_clusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kind_cluster)(nil)).Elem()
}

func (i Kind_clusterArray) ToKind_clusterArrayOutput() Kind_clusterArrayOutput {
	return i.ToKind_clusterArrayOutputWithContext(context.Background())
}

func (i Kind_clusterArray) ToKind_clusterArrayOutputWithContext(ctx context.Context) Kind_clusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Kind_clusterArrayOutput)
}

// Kind_clusterMapInput is an input type that accepts Kind_clusterMap and Kind_clusterMapOutput values.
// You can construct a concrete instance of `Kind_clusterMapInput` via:
//
//          Kind_clusterMap{ "key": Kind_clusterArgs{...} }
type Kind_clusterMapInput interface {
	pulumi.Input

	ToKind_clusterMapOutput() Kind_clusterMapOutput
	ToKind_clusterMapOutputWithContext(context.Context) Kind_clusterMapOutput
}

type Kind_clusterMap map[string]Kind_clusterInput

func (Kind_clusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kind_cluster)(nil)).Elem()
}

func (i Kind_clusterMap) ToKind_clusterMapOutput() Kind_clusterMapOutput {
	return i.ToKind_clusterMapOutputWithContext(context.Background())
}

func (i Kind_clusterMap) ToKind_clusterMapOutputWithContext(ctx context.Context) Kind_clusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Kind_clusterMapOutput)
}

type Kind_clusterOutput struct{ *pulumi.OutputState }

func (Kind_clusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind_cluster)(nil)).Elem()
}

func (o Kind_clusterOutput) ToKind_clusterOutput() Kind_clusterOutput {
	return o
}

func (o Kind_clusterOutput) ToKind_clusterOutputWithContext(ctx context.Context) Kind_clusterOutput {
	return o
}

type Kind_clusterArrayOutput struct{ *pulumi.OutputState }

func (Kind_clusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kind_cluster)(nil)).Elem()
}

func (o Kind_clusterArrayOutput) ToKind_clusterArrayOutput() Kind_clusterArrayOutput {
	return o
}

func (o Kind_clusterArrayOutput) ToKind_clusterArrayOutputWithContext(ctx context.Context) Kind_clusterArrayOutput {
	return o
}

func (o Kind_clusterArrayOutput) Index(i pulumi.IntInput) Kind_clusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Kind_cluster {
		return vs[0].([]*Kind_cluster)[vs[1].(int)]
	}).(Kind_clusterOutput)
}

type Kind_clusterMapOutput struct{ *pulumi.OutputState }

func (Kind_clusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kind_cluster)(nil)).Elem()
}

func (o Kind_clusterMapOutput) ToKind_clusterMapOutput() Kind_clusterMapOutput {
	return o
}

func (o Kind_clusterMapOutput) ToKind_clusterMapOutputWithContext(ctx context.Context) Kind_clusterMapOutput {
	return o
}

func (o Kind_clusterMapOutput) MapIndex(k pulumi.StringInput) Kind_clusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Kind_cluster {
		return vs[0].(map[string]*Kind_cluster)[vs[1].(string)]
	}).(Kind_clusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Kind_clusterInput)(nil)).Elem(), &Kind_cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*Kind_clusterArrayInput)(nil)).Elem(), Kind_clusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Kind_clusterMapInput)(nil)).Elem(), Kind_clusterMap{})
	pulumi.RegisterOutputType(Kind_clusterOutput{})
	pulumi.RegisterOutputType(Kind_clusterArrayOutput{})
	pulumi.RegisterOutputType(Kind_clusterMapOutput{})
}