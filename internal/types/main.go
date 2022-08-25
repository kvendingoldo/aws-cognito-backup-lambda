package types

import (
	"github.com/guregu/null"
)

type Response struct {
	Message string `json:"Answer:"`
}

type Event struct {
	ID string `json:"id"`

	AWSRegion string `json:"awsRegion"`

	CognitoUserPoolId string `json:"cognitoUserPoolID"`
	CognitoRegion     string `json:"cognitoRegion"`

	S3BucketName   string `json:"s3BucketName"`
	S3BucketRegion string `json:"s3BucketRegion"`

	BackupPrefix string `json:"backupPrefix"`

	RotationEnabled null.Bool `json:"rotation_enabled"`
	RotationDays    string    `json:"rotation_days_limit"`
}
