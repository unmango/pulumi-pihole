//go:build python || all
// +build python all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestDnsRecordPy(t *testing.T) {
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "dns-record", "py"),
		})

	integration.ProgramTest(t, &test)
}

func getPyOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	basePy := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			path.Join("..", "sdk", "python", "bin"),
		},
		Config: map[string]string{
			"pihole:url":      "http://localhost:8080",
			"pihole:password": "totally-secure-password",
		},
	})

	return basePy
}
