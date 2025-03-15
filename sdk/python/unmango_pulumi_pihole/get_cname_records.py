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
    'GetCnameRecordsResult',
    'AwaitableGetCnameRecordsResult',
    'get_cname_records',
    'get_cname_records_output',
]

@pulumi.output_type
class GetCnameRecordsResult:
    """
    A collection of values returned by getCnameRecords.
    """
    def __init__(__self__, id=None, records=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if records and not isinstance(records, list):
            raise TypeError("Expected argument 'records' to be a list")
        pulumi.set(__self__, "records", records)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def records(self) -> Sequence['outputs.GetCnameRecordsRecordResult']:
        """
        List of CNAME Pi-hole records
        """
        return pulumi.get(self, "records")


class AwaitableGetCnameRecordsResult(GetCnameRecordsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCnameRecordsResult(
            id=self.id,
            records=self.records)


def get_cname_records(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCnameRecordsResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_pihole as pihole

    records = pihole.get_cname_records()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('pihole:index/getCnameRecords:getCnameRecords', __args__, opts=opts, typ=GetCnameRecordsResult).value

    return AwaitableGetCnameRecordsResult(
        id=pulumi.get(__ret__, 'id'),
        records=pulumi.get(__ret__, 'records'))
def get_cname_records_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCnameRecordsResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_pihole as pihole

    records = pihole.get_cname_records()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('pihole:index/getCnameRecords:getCnameRecords', __args__, opts=opts, typ=GetCnameRecordsResult)
    return __ret__.apply(lambda __response__: GetCnameRecordsResult(
        id=pulumi.get(__response__, 'id'),
        records=pulumi.get(__response__, 'records')))
