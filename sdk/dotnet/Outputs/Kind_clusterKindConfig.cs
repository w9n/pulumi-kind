// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kind.Outputs
{

    [OutputType]
    public sealed class Kind_clusterKindConfig
    {
        public readonly string ApiVersion;
        public readonly ImmutableArray<string> ContainerdConfigPatches;
        public readonly string Kind;
        public readonly Outputs.Kind_clusterKindConfigNetworking? Networking;
        public readonly ImmutableArray<Outputs.Kind_clusterKindConfigNode> Nodes;

        [OutputConstructor]
        private Kind_clusterKindConfig(
            string apiVersion,

            ImmutableArray<string> containerdConfigPatches,

            string kind,

            Outputs.Kind_clusterKindConfigNetworking? networking,

            ImmutableArray<Outputs.Kind_clusterKindConfigNode> nodes)
        {
            ApiVersion = apiVersion;
            ContainerdConfigPatches = containerdConfigPatches;
            Kind = kind;
            Networking = networking;
            Nodes = nodes;
        }
    }
}
