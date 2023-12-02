package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		RunUpdateTest: false,
	}
}

func assertHasRepoDigest(t *testing.T, stack integration.RuntimeValidationStackInfo) {
	repoDigest, ok := stack.Outputs["repoDigest"].(string)
	assert.True(t, ok, "expected repoDigest output")
	assert.NotEmpty(t, repoDigest)
}
