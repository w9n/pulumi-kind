import * as pulumi from "@pulumi/pulumi";
import * as kind from "@pulumi/kind";

export const kind_cluster = new kind.Kind_cluster("kind-cluster", {
  name: "test-cluster",
  kindConfig: {
    kind: "Cluster",
    apiVersion: "kind.x-k8s.io/v1alpha4",
    nodes: [
      {
        role: "control-plane",
      },
      {
        role: "worker",
      },
    ],
  },
});
