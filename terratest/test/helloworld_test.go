package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformWorkspaces(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../module",
		Vars: map[string]any{
			"var1": "name",
		},
	})

	for _, wp := range []string{"dev", "qa"} {
		// Clean up resources with "terraform destroy" at the end of the test.
		{
			defer terraform.Destroy(t, terraformOptions)

			terraform.WorkspaceSelectOrNew(t, terraformOptions, wp)
			// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
			terraform.InitAndApply(t, terraformOptions)
			// Run `terraform output` to get the values of output variables and check they have the expected values.
			{
				expected := wp
				actual := terraform.Output(t, terraformOptions, "env")
				assert.Equal(t, expected, actual)

			}
			{
				expected := "name_test"
				actual := terraform.Output(t, terraformOptions, "name")
				assert.Equal(t, expected, actual)
			}
		}
	}

}
