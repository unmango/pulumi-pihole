import * as pihole from '@unmango/pulumi-pihole';

const record = new pihole.DnsRecord('example-record-ts', {
  domain: 'ts.unmango.net',
  ip: '192.168.1.69',
});
