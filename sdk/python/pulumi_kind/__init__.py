# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .kind_cluster import *
from .provider import *
from ._inputs import *
from . import outputs
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "kind",
  "mod": "index/kind_cluster",
  "fqn": "pulumi_kind",
  "classes": {
   "kind:index/kind_cluster:kind_cluster": "Kind_cluster"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "kind",
  "token": "pulumi:providers:kind",
  "fqn": "pulumi_kind",
  "class": "Provider"
 }
]
"""
)