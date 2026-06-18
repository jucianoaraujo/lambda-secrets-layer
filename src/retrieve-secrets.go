package main

import (
	"context"
	"fmt"
	"os"

	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/config"
)

// The main function will pull command line arg and retrieve the secret.  The resulting
// secret will be dumped as JSON to the output
func main() {

	// Expect the Secret ID/ARN to be passed via a bootstrap env var
	secretsArn := os.Getenv("SECRETS_ARN")
	if secretsArn == "" {
		panic("SECRETS_ARN environment variable is not set")
	}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("failed to load SDK config: " + err.Error())
	}

	client := secretsmanager.NewFromConfig(cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretsArn),
	}

	result, err := client.GetSecretValue(ctx, input)
	if err != nil {
		panic("failed to retrieve secret: " + err.Error())
	}

	if result.SecretString == nil {
		panic("secret string is empty")
	}

	// Parse JSON secrets
	var secrets map[string]string
	if err := json.Unmarshal([]byte(*result.SecretString), &secrets); err != nil {
		panic("failed to unmarshal secret JSON: " + err.Error())
	}

	// Output standard shell export commands
	for key, value := range secrets {
		fmt.Printf("export %s='%s'\n", key, value)
	}
}

