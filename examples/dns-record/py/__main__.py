"""A Python Pulumi program"""

import unmango_pulumi_pihole as pihole

dns_record = pihole.DnsRecord('example-record-py',
    domain='py.unmango.net',
    ip='192.168.1.69'
)
