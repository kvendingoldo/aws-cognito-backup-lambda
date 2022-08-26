## Environment variables

* FORMATTER_TYPE
    * Description: formatter type for Lambda's logs
    * Possible values: JSON | TEXT
    * Required: no

* MODE
    * Description: application mode
    * Possible values: cloud | local
    * Required: yes

* LOG_LEVEL
    * Description: Lambda's log level
    * Possible values: panic|fatal|error|warn|info|debug|trace
    * Required: no

* AWS_REGION
    * Description: Default AWS Region. Inside of Lambda it's setting automatically by AWS
    * Possible values: <any valid AWS region>
    * Required: yes

* COGNITO_USER_POOL_ID
    * Description: Cognito user pool ID
    * Possible values: <any valid ID>
    * Required: yes

* COGNITO_REGION
    * Description: AWS Region for Cognito client; If this variable is omitted value from AWS REGION will be used
    * Possible values: <any valid AWS region>
    * Required: no

* S3_BUCKET_NAME
    * Description: S3 bucket name
    * Possible values: <any valid bucket name>
    * Required: no

* S3_BUCKET_REGION
    * Description: AWS Region for S3 client; If this variable is omitted value from AWS REGION will be used
    * Possible values: <any valid AWS region>
    * Required: no

* BACKUP_PREFIX
    * Description: Backup prefix
    * Possible values: <any string>
    * Required: no

* ROTATION_ENABLED
    * Description: Is rotation enabled? By default, S3 bucket rotation will be disabled
    * Possible values: true|false
    * Required: no

* ROTATION_DAYS_LIMIT
    * Description: Max TTL for files inside S3 bucket
    * Possible values: any integer number greater than zero
    * Required: no