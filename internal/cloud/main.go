package cloud

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	CognitoClient *cognitoidentityprovider.Client
	S3Client      *s3.Client
	Region        string
}

func New(ctx context.Context, region string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &Client{
		CognitoClient: cognitoidentityprovider.NewFromConfig(cfg),
		S3Client:      s3.NewFromConfig(cfg),
		Region:        region,
	}, nil
}
