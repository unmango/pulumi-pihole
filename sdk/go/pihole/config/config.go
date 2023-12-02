// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/unmango/pulumi-pihole/sdk/go/pihole/internal"
)

var _ = internal.GetEnvOrDefault

// Experimental: Pi-hole API token. Conflicts with `password`.
func GetApiToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "pihole:apiToken")
}

// CA file to connect to Pi-hole with TLS
func GetCaFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "pihole:caFile")
}

// The admin password used to login to the admin dashboard. Conflicts with `api_token`.
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "pihole:password")
}

// URL where Pi-hole is deployed
func GetUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "pihole:url")
}