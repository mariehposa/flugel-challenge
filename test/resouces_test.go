package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestResources(t *testing.T) {
	t.Parallel()

	awsRegion := "us-east-1"

	expectedName := fmt.Sprintf("Flugel")

	expectedOwner := fmt.Sprintf("InfraTeam")

	expectedBucket := "buc955815154140"

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform/",
	}

	// Destroy the infrastructure after testing.
	defer terraform.Destroy(t, terraformOptions)

	// Deploy the resources with the terraform options declared above.
	terraform.InitAndApply(t, terraformOptions)

	// Get EC2 instance ID to query AWS.
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	// Get EC2 instance tags to perform the testing.
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)

	// Test EC2 instance if properly tagged.
	nameTag, containsNameTag := instanceTags["Name"]
	assert.True(t, containsNameTag)
	assert.Equal(t, expectedName, nameTag)

	OwnerTag, containsOwnerTag := instanceTags["Owner"]
	assert.True(t, containsOwnerTag)
	assert.Equal(t, expectedOwner, OwnerTag)

	// Test S3 bucket if properly tagged.
	bucketWithTag := aws.FindS3BucketWithTag(t, awsRegion, "Name", expectedName)
	assert.Equal(t, expectedBucket, bucketWithTag)

	bucketWithOwner := aws.FindS3BucketWithTag(t, awsRegion, "Owner", expectedOwner)
	assert.Equal(t, expectedBucket, bucketWithOwner)
}