import * as pulumi from '@pulumi/pulumi';
import * as pihole from '@unmango/pulumi-pihole';

const record = new pihole.DnsRecord('example-record', {
  domain: 'unmango.net',
  ip: '192.168.1.69',
});
