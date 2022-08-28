# aws-cognito-backup-lambda

## Overview

Users of AWS Cognito periodically face issues with backups. The Cognito user pool backup mechanism is not embedded into AWS. 
Therefore, this Lambda function can assist with the problem; using it, 
you can backup all users and groups from the Cognito user pool into an S3 bucket in JSON format. 
The [aws-cognito-restore-lambda](https://github.com/kvendingoldo/aws-cognito-restore-lambda) project can be used to recover users.

## Documentation
You can review the following documents on the Lambda to learn more:
* [How to use the Lambda inside of AWS](docs/how_to_use_aws.md)
* [How to use the Lambda locally](docs/how_to_use_locally.md)
* [How to use Terraform automation](docs/how_to_use_terraform.md)
* [Labmda's environment variables](docs/environment_variables.md)