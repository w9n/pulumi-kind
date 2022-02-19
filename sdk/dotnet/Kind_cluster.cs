// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kind
{
    [KindResourceType("kind:index/kind_cluster:kind_cluster")]
    public partial class Kind_cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Client certificate for authenticating to cluster.
        /// </summary>
        [Output("clientCertificate")]
        public Output<string> ClientCertificate { get; private set; } = null!;

        /// <summary>
        /// Client key for authenticating to cluster.
        /// </summary>
        [Output("clientKey")]
        public Output<string> ClientKey { get; private set; } = null!;

        /// <summary>
        /// Client verifies the server certificate with this CA cert.
        /// </summary>
        [Output("clusterCaCertificate")]
        public Output<string> ClusterCaCertificate { get; private set; } = null!;

        /// <summary>
        /// Kubernetes APIServer endpoint.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The kind_config that kind will use to bootstrap the cluster.
        /// </summary>
        [Output("kindConfig")]
        public Output<Outputs.Kind_clusterKindConfig?> KindConfig { get; private set; } = null!;

        /// <summary>
        /// Kubeconfig set after the the cluster is created.
        /// </summary>
        [Output("kubeconfig")]
        public Output<string> Kubeconfig { get; private set; } = null!;

        /// <summary>
        /// Kubeconfig path set after the the cluster is created or by the user to override defaults.
        /// </summary>
        [Output("kubeconfigPath")]
        public Output<string> KubeconfigPath { get; private set; } = null!;

        /// <summary>
        /// The kind name that is given to the created cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The node_image that kind will use (ex: kindest/node:v1.15.3).
        /// </summary>
        [Output("nodeImage")]
        public Output<string> NodeImage { get; private set; } = null!;

        /// <summary>
        /// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false
        /// </summary>
        [Output("waitForReady")]
        public Output<bool?> WaitForReady { get; private set; } = null!;


        /// <summary>
        /// Create a Kind_cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Kind_cluster(string name, Kind_clusterArgs? args = null, CustomResourceOptions? options = null)
            : base("kind:index/kind_cluster:kind_cluster", name, args ?? new Kind_clusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Kind_cluster(string name, Input<string> id, Kind_clusterState? state = null, CustomResourceOptions? options = null)
            : base("kind:index/kind_cluster:kind_cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Kind_cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Kind_cluster Get(string name, Input<string> id, Kind_clusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Kind_cluster(name, id, state, options);
        }
    }

    public sealed class Kind_clusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The kind_config that kind will use to bootstrap the cluster.
        /// </summary>
        [Input("kindConfig")]
        public Input<Inputs.Kind_clusterKindConfigArgs>? KindConfig { get; set; }

        /// <summary>
        /// Kubeconfig path set after the the cluster is created or by the user to override defaults.
        /// </summary>
        [Input("kubeconfigPath")]
        public Input<string>? KubeconfigPath { get; set; }

        /// <summary>
        /// The kind name that is given to the created cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The node_image that kind will use (ex: kindest/node:v1.15.3).
        /// </summary>
        [Input("nodeImage")]
        public Input<string>? NodeImage { get; set; }

        /// <summary>
        /// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false
        /// </summary>
        [Input("waitForReady")]
        public Input<bool>? WaitForReady { get; set; }

        public Kind_clusterArgs()
        {
        }
    }

    public sealed class Kind_clusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client certificate for authenticating to cluster.
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// Client key for authenticating to cluster.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// Client verifies the server certificate with this CA cert.
        /// </summary>
        [Input("clusterCaCertificate")]
        public Input<string>? ClusterCaCertificate { get; set; }

        /// <summary>
        /// Kubernetes APIServer endpoint.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The kind_config that kind will use to bootstrap the cluster.
        /// </summary>
        [Input("kindConfig")]
        public Input<Inputs.Kind_clusterKindConfigGetArgs>? KindConfig { get; set; }

        /// <summary>
        /// Kubeconfig set after the the cluster is created.
        /// </summary>
        [Input("kubeconfig")]
        public Input<string>? Kubeconfig { get; set; }

        /// <summary>
        /// Kubeconfig path set after the the cluster is created or by the user to override defaults.
        /// </summary>
        [Input("kubeconfigPath")]
        public Input<string>? KubeconfigPath { get; set; }

        /// <summary>
        /// The kind name that is given to the created cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The node_image that kind will use (ex: kindest/node:v1.15.3).
        /// </summary>
        [Input("nodeImage")]
        public Input<string>? NodeImage { get; set; }

        /// <summary>
        /// Defines wether or not the provider will wait for the control plane to be ready. Defaults to false
        /// </summary>
        [Input("waitForReady")]
        public Input<bool>? WaitForReady { get; set; }

        public Kind_clusterState()
        {
        }
    }
}
