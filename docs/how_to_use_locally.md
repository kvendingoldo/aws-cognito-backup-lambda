## How to use it locally

1. Set the following environment variables (do not forget to change placeholders)

```shell
export AWS_REGION=<REGION>
export MODE=local
export COGNITO_USER_POOL_ID=<COGNITO_USER_POOL_ID>
export S3_BUCKET_NAME=<S3_BUCKET_NAME>
```

2. Run lambda locally

```sh
go run main.go
```

