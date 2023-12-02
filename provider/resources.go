// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate go run generate.go
package pihole

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ettle/strcase"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	shimprovider "github.com/ryanwholey/terraform-provider-pihole/shim"
	"github.com/unmango/pulumi-pihole/provider/pkg/version"
)

const (
	// This variable controls the default name of the package in the package
	mainMod = "index" // the pihole module
)

func convertName(name string) string {
	idx := strings.Index(name, "_")
	contract.Assertf(idx > 0 && idx < len(name)-1, "Invalid snake case name %s", name)
	name = name[idx+1:]
	contract.Assertf(len(name) > 0, "Invalid snake case name %s", name)
	return strcase.ToPascal(name)
}

func makeDataSource(mod string, name string) tokens.ModuleMember {
	name = convertName(name)
	return tfbridge.MakeDataSource("pihole", mod, "get"+name)
}

func makeResource(mod string, res string) tokens.Type {
	return tfbridge.MakeResource("pihole", mod, convertName(res))
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(shimprovider.NewProvider())

	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "pihole",
		DisplayName:       "Pi-hole",
		Publisher:         "UnstoppableMango",
		LogoURL:           "https://raw.githubusercontent.com/unmango/pulumi-pihole/main/docs/pihole.png",
		PluginDownloadURL: "github://api.github.com/unmango/pulumi-pihole",
		Description:       "A Pulumi package for creating and managing Pi-hole resources",
		Keywords: []string{
			"pulumi",
			"pihole",
			"category/network",
		},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/unmango/pulumi-pihole",
		Repository: "https://github.com/unmango/pulumi-pihole",
		Version:    version.Version,
		GitHubOrg:  "ryanwholey",
		Config: map[string]*tfbridge.SchemaInfo{
			"api_token": {
				Name: "apiToken",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PIHOLE_API_TOKEN"},
				},
			},
			"ca_file": {
				Name: "caFile",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PIHOLE_CA_FILE"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PIHOLE_PASSWORD"},
				},
			},
			"url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PIHOLE_URL"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"pihole_ad_blocker_status": {Tok: makeResource(mainMod, "pihole_ad_blocker_status")},
			"pihole_cname_record":      {Tok: makeResource(mainMod, "pihole_cname_record")},
			"pihole_dns_record":        {Tok: makeResource(mainMod, "pihole_dns_record")},
			"pihole_group":             {Tok: makeResource(mainMod, "pihole_group")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"pihole_cname_records": {Tok: makeDataSource(mainMod, "pihole_cname_records")},
			"pihole_dns_records":   {Tok: makeDataSource(mainMod, "pihole_dns_records")},
			"pihole_domains":       {Tok: makeDataSource(mainMod, "pihole_domains")},
			"pihole_groups":        {Tok: makeDataSource(mainMod, "pihole_groups")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@unmango/pulumi-pihole",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "unmango_pulumi_pihole",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/unmango/pulumi-%[1]s/sdk/", "pihole"),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"pihole",
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "UnMango.Pulumi",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.unmango.pulumi",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
