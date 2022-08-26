## Environment variables

* FORMATTER_TYPE
    * Description: formatter type for lambda's logs
    * Possible values: JSON | TEXT
    * Required: no

* MODE
    * Description: mode of application running
    * Possible values: cloud | local
    * Required: yes

* LOG_LEVEL
    * Description: log level for lambda
    * Possible values: panic|fatal|error|warn|info|debug|trace
    * Required: no

* AWS_REGION
    * Description: Default AWS Region for S3 & Cognito clients. Inside of Lambda it's setting automatically by AWS;
    * Possible values: <any valid AWS region>
    * Required: yes

* COGNITO_USER_POOL_ID
    * Description: ID of Cognito user pool
    * Possible values: <any valid ID>
    * Required: yes

* COGNITO_REGION
  * Description: AWS Region for Cognito client; If this variable is omitted value from AWS REGION will be used
  * Possible values: <any valid AWS region>
  * Required: no

* S3_BUCKET_NAME
    * Description: Name of S3 bucket
    * Possible values: <any valid bucket name>
    * Required: yes

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
    * Description: Max TTL for file in S3 bucket
    * Possible values: any integer number greater than zero
    * Required: no