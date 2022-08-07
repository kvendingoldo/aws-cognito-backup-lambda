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
export USER_POOL_ID=<USER_POOL_ID>
export BUCKET_NAME=<S3_BUCKET_NAME>


```
2. Run lambda locally
```sh
go run main.go
```

## Environment variables

* AWS_REGION
   * Description: AWS Region. Inside of Lambda it's setting automatically by Lambda
   * Possible values: <any valid AWS region>

* MODE
   * Description: mode of application running
   * Possible values: cloud | local

* USER_POOL_ID
   * Description: ID of Cognito user pool
   * Possible values: <any valid ID>

* BUCKET_NAME
   * Description: Name of S3 bucket
   * Possible values: <any valid bucket name> 

* BACKUP_PREFIX
   * Description: Backup prefix
   * Possible values: <any string>

* ROTATION_DAYS
   * Description: Max TTL for backups
   * Possible values: any positive number, or -1 for disabling it