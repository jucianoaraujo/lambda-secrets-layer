# lambda-secrets-layer

First, create a lambda layer with the zip files on layers folder.

Second, associate the layer to the lambda you want, and set the envs for your lambda:
```
SECRETS_ARN={your_secrets_arn}
AWS_LAMBDA_EXEC_WRAPPER=/opt/secrets-wrapper.sh
```