#### Terraform code example

1. Add module execution to your TF code

```terraform
module "cognito_backup_lambda" {
  source = "git@github.com:kvendingoldo/aws-cognito-backup-lambda.git//files/terraform/module?ref=rc/0.1.0"

  blank_name = "${module.naming.common_name}-cognito-backup"
  tags       = local.tags

  cron_schedule = var.cognito_backup_lambda_cron_schedule
  image_uri     = var.cognito_backup_lambda_image_uri
  events        = var.cognito_backup_lambda_events
}
```

2. Specify variables

```terraform
variable "cognito_backup_lambda_cron_schedule" {
  default = "rate(168 hours)"
}
variable "cognito_backup_lambda_image_uri" {
  default = "<YOUR_ACCOUNT_ID>.dkr.ecr.us-east-2.amazonaws.com/aws-cognito_backup_lambda:rc-0.1.0"
}
variable "cognito_backup_lambda_events" {
  default = [
    {
      CognitoUserPoolID : "<YOUR_POOL_ID>",
      S3BucketName : "<YOUR_BUCKET_NAME>",
    },
  ]
}
```