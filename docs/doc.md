## How to use it within AWS

1. Lambda image should be pulled from docker hub and pushed into your personal ECR repository; AWS Lambda is not able to
   work with any other docker registry except ECR.
2. Apply TF module into your infrastructure

#### Example of input variables for Terraform module

```terraform
TODO
```

### How to trigger lambda manually via UI

1. Go to Lambda function that has been created via Terraform -> Tests
2. Fill "Test Event" and click "Test"

```json
TODO
```

## How to use it locally
1. Set the following environment variables (do not forget to change placeholders)
```shell
export AWS_REGION=<REGION>
export MODE=local

```
2. Run lambda locally
```sh
go run main.go
```

## Environment variables

TODO