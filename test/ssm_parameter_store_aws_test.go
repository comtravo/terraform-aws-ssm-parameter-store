package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestFetchParameterStoreArn_basic(t *testing.T) {
	t.Parallel()

	workspaceName := strings.ToUpper(fmt.Sprintf("WS_%s", random.UniqueId()))
	secretName := fmt.Sprintf("FOO_%s", random.UniqueId())

	expectedParameterName := fmt.Sprintf("%s_%s", workspaceName, secretName)
	secretValue := random.UniqueId()

	SSMPutParameter(t, expectedParameterName, secretValue)

	exampleDir := "../examples/basic/"

	terraformOptions := &terraform.Options{
		TerraformDir: exampleDir,
		Vars: map[string]interface{}{
			"secret_name": secretName,
		},
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, workspaceName)
	TerraformApplyAndValidateOutputs(t, terraformOptions, expectedParameterName)

	t.Logf("Terraform module inputs: %+v", *terraformOptions)
}

func TestFetchParameterStoreArn_customPrefix(t *testing.T) {
	t.Parallel()

	customPrefixName := strings.ToUpper(fmt.Sprintf("PF_%s", random.UniqueId()))
	secretName := fmt.Sprintf("FOO_%s_", random.UniqueId())

	expectedParameterName := fmt.Sprintf("%s%s", customPrefixName, secretName)
	secretValue := random.UniqueId()

	SSMPutParameter(t, expectedParameterName, secretValue)

	exampleDir := "../examples/custom_prefix/"

	terraformOptions := &terraform.Options{
		TerraformDir: exampleDir,
		Vars: map[string]interface{}{
			"secret_name": secretName,
			"prefix":      customPrefixName,
		},
	}
	defer terraform.Destroy(t, terraformOptions)

	TerraformApplyAndValidateOutputs(t, terraformOptions, expectedParameterName)

	t.Logf("Terraform module inputs: %+v", *terraformOptions)
}

func TestFetchParameterStoreArn_disable(t *testing.T) {
	t.Parallel()

	customPrefixName := strings.ToUpper(fmt.Sprintf("PF_%s", random.UniqueId()))
	secretName := fmt.Sprintf("FOO_%s_", random.UniqueId())

	expectedParameterName := fmt.Sprintf("%s%s", customPrefixName, secretName)
	secretValue := random.UniqueId()

	SSMPutParameter(t, expectedParameterName, secretValue)

	exampleDir := "../examples/disable/"

	terraformOptions := &terraform.Options{
		TerraformDir: exampleDir,
		Vars: map[string]interface{}{
			"secret_name": secretName,
			"prefix":      customPrefixName,
		},
	}
	defer terraform.Destroy(t, terraformOptions)

	TerraformApplyAndValidateOutputs(t, terraformOptions, "")

	t.Logf("Terraform module inputs: %+v", *terraformOptions)
}

func SSMPutParameter(t *testing.T, secretName string, secretValue string) *ssm.PutParameterOutput {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	ssmClient := ssm.New(sess, &aws.Config{})

	params := ssm.PutParameterInput{
		Name:  aws.String(secretName),
		Value: aws.String(secretValue),
		Type:  aws.String("SecureString"),
	}

	t.Logf("Creating secret with parameters: %+v", params)
	res, err := ssmClient.PutParameter(&params)

	if err != nil {
		panic(err)
	}

	t.Logf("Creating secret response: %+v", res)
	return res
}

func TerraformApplyAndValidateOutputs(t *testing.T, terraformOptions *terraform.Options, expectedParameterName string) {
	terraformApplyOutput := terraform.InitAndApply(t, terraformOptions)
	resourceCount := terraform.GetResourceCount(t, terraformApplyOutput)

	require.Equal(t, resourceCount.Add, 0)
	require.Equal(t, resourceCount.Change, 0)
	require.Equal(t, resourceCount.Destroy, 0)

	if expectedParameterName == "" {
		require.Equal(t, "", terraform.Output(t, terraformOptions, "arn"))
	} else {
		require.Regexp(t,
			regexp.MustCompile(fmt.Sprintf("arn:aws:ssm:us-east-1:\\d{12}:parameter/%s", expectedParameterName)),
			terraform.Output(t, terraformOptions, "arn"))
	}
}
