## How to use it within AWS

1. Lambda image should be pulled from docker hub and pushed into your personal ECR repository; AWS Lambda is not able to
   work with any other docker registry except ECR.
2. Apply TF module into your infrastructure

### How to trigger lambda manually via UI

1. Go to Lambda function that has been created via Terraform -> Tests
2. Fill "Test Event" and click "Test"

```
{
  "awsRegion": "<AWS_REGION | optional>",
  
  "cognitoUserPoolID": "<COGNITO_USER_POOL_ID>",
  "cognitoRegion": "<COGNITO_POOL_AWS_REGION | AWS_REGION will be used if this var is omitted>",
  
  "s3BucketName": "<S3_BUCKET_NAME>",
  "s3BucketRegion": "<S3_AWS_REGION | AWS_REGION will be used if this var is omitted>",
  
  "backupPrefix": "<can be escaped or any valid string>",
  
  "rotationEnabled": <true | false>,
  "rotationDaysLimit": <can be escaped or any valid number>
}
```

#### Example #1:

```json
{
  "awsRegion": "us-west-2",
  "cognitoUserPoolID": "ap-southeast-2_EPyUfpQq7",
  "s3BucketName": "mybuckettest"
}
```

#### Example #2:

```json
{
  "cognitoUserPoolID": "ap-southeast-2_EPyUfpQq7",
  "cognitoRegion": "us-west-2",
  "s3BucketName": "mybuckettest",
  "s3BucketRegion": "us-east-1",
  "backupPrefix": "testprefix",
  "rotationEnabled": true,
  "rotationDaysLimit": 7
}
```
