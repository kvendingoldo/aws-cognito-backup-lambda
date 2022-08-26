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

