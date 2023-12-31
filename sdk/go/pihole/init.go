// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pihole

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unmango/pulumi-pihole/sdk/go/pihole/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "pihole:index/adBlockerStatus:AdBlockerStatus":
		r = &AdBlockerStatus{}
	case "pihole:index/cnameRecord:CnameRecord":
		r = &CnameRecord{}
	case "pihole:index/dnsRecord:DnsRecord":
		r = &DnsRecord{}
	case "pihole:index/group:Group":
		r = &Group{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:pihole" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"pihole",
		"index/adBlockerStatus",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"pihole",
		"index/cnameRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"pihole",
		"index/dnsRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"pihole",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"pihole",
		&pkg{version},
	)
}
