# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('pihole')


class _ExportableConfig(types.ModuleType):
    @property
    def api_token(self) -> Optional[str]:
        """
        Experimental: Pi-hole API token. Conflicts with `password`.
        """
        return __config__.get('apiToken')

    @property
    def ca_file(self) -> Optional[str]:
        """
        CA file to connect to Pi-hole with TLS
        """
        return __config__.get('caFile')

    @property
    def password(self) -> Optional[str]:
        """
        The admin password used to login to the admin dashboard. Conflicts with `api_token`.
        """
        return __config__.get('password')

    @property
    def url(self) -> Optional[str]:
        """
        URL where Pi-hole is deployed
        """
        return __config__.get('url')
