package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3Bucket(t *testing.T) {
	awsRegion := "eu-west-2"
	terraformOpts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../aws/",

		Vars: map[string]interface{}{
			"bucket_name": fmt.Sprintf("-%v", strings.ToLower(random.UniqueId())),
		},

		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// We want to destroy the infrastructure after testing.
	defer terraform.Destroy(t, terraformOpts)

	// Deploy the infrastructure with the options defined above
	terraform.InitAndApply(t, terraformOpts)

	// Get the bucket ID so we can query AWS
	bucketID := terraform.Output(t, terraformOpts, "bucket_id")

	// Get the versioning status to test that versioning is enabled
	actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)

	// Test that the status we get back from AWS is "Enabled" for versioning
	assert.Equal(t, "Enabled", actualStatus)
}
