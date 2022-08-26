## Environment variables

* FORMATTER_TYPE
    * Description: formatter type for lambda's logs
    * Possible values: JSON | TEXT

* MODE
    * Description: mode of application running
    * Possible values: cloud | local

* LOG_LEVEL
    * Description: log level for lambda
    * Possible values: panic|fatal|error|warn|info|debug|trace

* AWS_REGION
    * Description: AWS Region. Inside of Lambda it's setting automatically by Lambda
    * Possible values: <any valid AWS region>

* COGNITO_USER_POOL_ID
    * Description: ID of Cognito user pool
    * Possible values: <any valid ID>

* S3_BUCKET_NAME
    * Description: Name of S3 bucket
    * Possible values: <any valid bucket name>

* BACKUP_PREFIX
    * Description: Backup prefix
    * Possible values: <any string>

* ROTATION_DAYS_LIMIT
    * Description: Max TTL for file in S3 bucket
    * Possible values: any positive number, or -1 for disabling it