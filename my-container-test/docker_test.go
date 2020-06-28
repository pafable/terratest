package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformDocker(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../my-container-test",
	}

	// executes terraform destroy after creating
	defer terraform.Destroy(t, terraformOptions)

	// executes terraform init and apply to check containers
	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "container_name")
	assert.Equal(t, "ghost-cont-01", output)
}
