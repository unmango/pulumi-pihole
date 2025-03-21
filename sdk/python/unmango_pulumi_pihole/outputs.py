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

__all__ = [
    'GetCnameRecordsRecordResult',
    'GetDnsRecordsRecordResult',
    'GetDomainsDomainResult',
    'GetGroupsGroupResult',
]

@pulumi.output_type
class GetCnameRecordsRecordResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 target: str):
        """
        :param str domain: CNAME record domain
        :param str target: CNAME target value where traffic is routed to from the domain
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        CNAME record domain
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        CNAME target value where traffic is routed to from the domain
        """
        return pulumi.get(self, "target")


@pulumi.output_type
class GetDnsRecordsRecordResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 ip: str):
        """
        :param str domain: DNS record domain
        :param str ip: IP address where traffic is routed to from the DNS record domain
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "ip", ip)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        DNS record domain
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        IP address where traffic is routed to from the DNS record domain
        """
        return pulumi.get(self, "ip")


@pulumi.output_type
class GetDomainsDomainResult(dict):
    def __init__(__self__, *,
                 comment: str,
                 domain: str,
                 enabled: bool,
                 group_ids: Sequence[int],
                 id: int,
                 type: str,
                 wildcard: bool):
        """
        :param str comment: Comments associated with the domain
        :param str domain: Domain
        :param bool enabled: Whether the domain rule is enabled
        :param Sequence[int] group_ids: Groups to which the domain is associated
        :param int id: Domain ID
        :param str type: Whether the doamin is on the allow or deny list
        :param bool wildcard: Whether the domain should be interpreted using a wildcard parser
        """
        pulumi.set(__self__, "comment", comment)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "group_ids", group_ids)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "wildcard", wildcard)

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comments associated with the domain
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        Domain
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether the domain rule is enabled
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Sequence[int]:
        """
        Groups to which the domain is associated
        """
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Domain ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Whether the doamin is on the allow or deny list
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def wildcard(self) -> bool:
        """
        Whether the domain should be interpreted using a wildcard parser
        """
        return pulumi.get(self, "wildcard")


@pulumi.output_type
class GetGroupsGroupResult(dict):
    def __init__(__self__, *,
                 description: str,
                 enabled: bool,
                 id: int,
                 name: str):
        """
        :param str description: Group description
        :param bool enabled: Whether the group is enabled
        :param int id: Group ID
        :param str name: Name of the group
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Group description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether the group is enabled
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Group ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the group
        """
        return pulumi.get(self, "name")


