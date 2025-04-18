# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetDomainsResult',
    'AwaitableGetDomainsResult',
    'get_domains',
    'get_domains_output',
]

@pulumi.output_type
class GetDomainsResult:
    """
    A collection of values returned by getDomains.
    """
    def __init__(__self__, domains=None, id=None, type=None):
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetDomainsDomainResult']:
        """
        Domains
        """
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
        """
        return pulumi.get(self, "type")


class AwaitableGetDomainsResult(GetDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainsResult(
            domains=self.domains,
            id=self.id,
            type=self.type)


def get_domains(type: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainsResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_pihole as pihole

    all = pihole.get_domains()
    denied = pihole.get_domains(type="deny")
    ```


    :param str type: Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
    """
    __args__ = dict()
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('pihole:index/getDomains:getDomains', __args__, opts=opts, typ=GetDomainsResult).value

    return AwaitableGetDomainsResult(
        domains=pulumi.get(__ret__, 'domains'),
        id=pulumi.get(__ret__, 'id'),
        type=pulumi.get(__ret__, 'type'))
def get_domains_output(type: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDomainsResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_pihole as pihole

    all = pihole.get_domains()
    denied = pihole.get_domains(type="deny")
    ```


    :param str type: Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
    """
    __args__ = dict()
    __args__['type'] = type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('pihole:index/getDomains:getDomains', __args__, opts=opts, typ=GetDomainsResult)
    return __ret__.apply(lambda __response__: GetDomainsResult(
        domains=pulumi.get(__response__, 'domains'),
        id=pulumi.get(__response__, 'id'),
        type=pulumi.get(__response__, 'type')))
