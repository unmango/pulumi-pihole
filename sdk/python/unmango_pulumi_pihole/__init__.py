# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .ad_blocker_status import *
from .cname_record import *
from .dns_record import *
from .get_cname_records import *
from .get_dns_records import *
from .get_domains import *
from .get_groups import *
from .group import *
from .provider import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import unmango_pulumi_pihole.config as __config
    config = __config
else:
    config = _utilities.lazy_import('unmango_pulumi_pihole.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "pihole",
  "mod": "index/adBlockerStatus",
  "fqn": "unmango_pulumi_pihole",
  "classes": {
   "pihole:index/adBlockerStatus:AdBlockerStatus": "AdBlockerStatus"
  }
 },
 {
  "pkg": "pihole",
  "mod": "index/cnameRecord",
  "fqn": "unmango_pulumi_pihole",
  "classes": {
   "pihole:index/cnameRecord:CnameRecord": "CnameRecord"
  }
 },
 {
  "pkg": "pihole",
  "mod": "index/dnsRecord",
  "fqn": "unmango_pulumi_pihole",
  "classes": {
   "pihole:index/dnsRecord:DnsRecord": "DnsRecord"
  }
 },
 {
  "pkg": "pihole",
  "mod": "index/group",
  "fqn": "unmango_pulumi_pihole",
  "classes": {
   "pihole:index/group:Group": "Group"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "pihole",
  "token": "pulumi:providers:pihole",
  "fqn": "unmango_pulumi_pihole",
  "class": "Provider"
 }
]
"""
)
