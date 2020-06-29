package test

import (
	"testing"

	"fmt"
	"log"
	"net/http"
	"time"

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
	fmt.Println("Sleeping for 10 seconds...")
	time.Sleep(10 * time.Second)

	resp, err := http.Get("http://localhost")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)

	assert.Equal(t, 200, resp.StatusCode)
}
