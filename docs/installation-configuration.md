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

* `pihole:apiKey` (environment: `pihole_API_KEY`) - the API key for `pihole`
* `pihole:region` (environment: `pihole_REGION`) - the region in which to deploy resources

### Provider Binary

The Pihole provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource pihole <version>
```

Replace the version string `<version>` with your desired version.
