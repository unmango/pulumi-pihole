---
title: Pi-hole Installation & Configuration
meta_desc: Information on how to install the Pi-hole provider.
layout: installation
---

## Installation

The Pulumi Pihole provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@unmango/pulumi-pihole`](https://www.npmjs.com/package/@unmango/pulumi-pihole)
* Python: [`unmango_pulumi_pihole`](https://pypi.org/project/unmango_pulumi_pihole/)
* Go: [`github.com/unmango/pulumi-pihole/sdk/go/pihole`](https://pkg.go.dev/github.com/unmango/pulumi-pihole/sdk/go/pihole)
* .NET: [`UnMango.Pulumi.Pihole`](https://www.nuget.org/packages/UnMango.Pulumi.Pihole)

## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `pihole` provider:

* `pihole:apiToken` (environment: `PIHOLE_API_TOKEN`) - Experimental: Pi-hole API token. Conflicts with `password`.
* `pihole:caFile` (environment: `PIHOLE_CA_FILE`) - CA file to connect to Pi-hole with TLS.
* `pihole:password` (environment: `PIHOLE_PASSWORD`) - The admin password used to login to the admin dashboard. Conflicts with `api_token`.
* `pihole:url` (environment: `PIHOLE_URL`) - URL where Pi-hole is deployed.

### Provider Binary

The Pihole provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource pihole <version>
```

Replace the version string `<version>` with your desired version.
