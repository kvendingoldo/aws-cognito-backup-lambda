package config

import (
	"github.com/guregu/null"
	"github.com/kvendingoldo/aws-cognito-backup-lambda/internal/types"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	AWSRegion string

	CognitoUserPoolId string
	CognitoRegion     string

	S3BucketName   string
	S3BucketRegion string

	BackupPrefix string

	RotationEnabled   null.Bool
	RotationDaysLimit int64
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func New(eventRaw interface{}) *Config {
	var config = &Config{}

	var getFromEvent bool
	var event types.Event

	switch value := eventRaw.(type) {
	case types.Event:
		getFromEvent = true
		event = value
	default:
		getFromEvent = false
	}

	// Process AWSRegion
	if awsRegion := getEnv("AWS_REGION", ""); awsRegion != "" {
		config.AWSRegion = awsRegion
	} else {
		log.Warn("Environment variable AWS_REGION is empty")
	}
	if getFromEvent {
		if event.AWSRegion != "" {
			config.AWSRegion = event.AWSRegion
		} else {
			log.Warn("Event contains empty awsRegion variable")
		}
	}
	if config.AWSRegion == "" {
		log.Error("awsRegion is empty; Configure it via 'AWS_REGION' env variable OR pass in event body")
		os.Exit(1)
	}

	// Process CognitoRegion
	if cognitoRegion := getEnv("COGNITO_REGION", ""); cognitoRegion != "" {
		config.CognitoRegion = cognitoRegion
	} else {
		log.Warn("Environment variable 'COGNITO_REGION' is empty")
	}
	if getFromEvent {
		if event.CognitoRegion != "" {
			config.CognitoRegion = event.CognitoRegion
		} else {
			log.Warn("Event contains empty cognitoRegion variable")
		}
	}
	if config.CognitoRegion == "" {
		log.Warnf("cognitoRegion is empty; Default region %s will be used", config.AWSRegion)
		config.CognitoRegion = config.AWSRegion
	}

	// Process S3BucketRegion
	if bucketRegion := getEnv("S3_BUCKET_REGION", ""); bucketRegion != "" {
		config.S3BucketRegion = bucketRegion
	} else {
		log.Warn("Environment variable 'S3_BUCKET_REGION' is empty")
	}
	if getFromEvent {
		if event.S3BucketRegion != "" {
			config.S3BucketRegion = event.CognitoRegion
		} else {
			log.Warn("Event contains empty s3BucketRegion variable")
		}
	}
	if config.S3BucketRegion == "" {
		log.Warnf("bucketRegion is empty; Default region %s will be used", config.AWSRegion)
		config.S3BucketRegion = config.AWSRegion
	}

	// Process CognitoUserPoolId
	if cognitoUserPoolId := getEnv("COGNITO_USER_POOL_ID", ""); cognitoUserPoolId != "" {
		config.CognitoUserPoolId = cognitoUserPoolId
	} else {
		log.Warn("Environment variable 'COGNITO_USER_POOL_ID' is empty")
	}
	if getFromEvent {
		if event.CognitoUserPoolId != "" {
			config.CognitoUserPoolId = event.CognitoUserPoolId
		} else {
			log.Warn("Event contains empty cognitoUserPoolID")
		}
	}
	if config.CognitoUserPoolId == "" {
		log.Error("cognitoUserPoolID is empty; Configure it via 'COGNITO_USER_POOL_ID' env variable OR pass in event body")
		os.Exit(1)
	}

	// Process S3BucketName
	if s3BucketName := getEnv("S3_BUCKET_NAME", ""); s3BucketName != "" {
		config.S3BucketName = s3BucketName
	} else {
		log.Warn("Environment variable 'S3_BUCKET_NAME' is empty")
	}
	if getFromEvent {
		if event.S3BucketName != "" {
			config.S3BucketName = event.S3BucketName
		} else {
			log.Warn("Event contains empty s3BucketName")
		}
	}
	if config.S3BucketName == "" {
		log.Error("BucketName is empty; Configure it via 'S3_BUCKET_NAME' env variable OR pass in event body")
		os.Exit(1)
	}

	// Process BackupPrefix
	if backupPrefix := getEnv("BACKUP_PREFIX", ""); backupPrefix != "" {
		config.BackupPrefix = backupPrefix
	} else {
		log.Warn("Environment variable 'BACKUP_PREFIX' is empty")
	}
	if getFromEvent {
		if event.BackupPrefix == "" {
			log.Warnf("Event contains empty backupPrefix")
		} else {
			config.BackupPrefix = event.BackupPrefix
		}
	}

	// Process RotationEnabled
	if rotationEnabled := getEnv("ROTATION_ENABLED", ""); rotationEnabled != "" {
		rotationEnabledValue, err := strconv.ParseBool(rotationEnabled)
		if err != nil {
			log.Errorf("Could not parse 'ROTATION_ENABLED' variable. Error: %v", err)
			os.Exit(1)
		}

		config.RotationEnabled = null.NewBool(rotationEnabledValue, true)
	} else {
		log.Warn("Environment variable 'ROTATION_ENABLED' is empty")
	}
	if getFromEvent {
		if event.RotationEnabled.Valid {
			config.RotationEnabled = event.RotationEnabled
		}
	}
	if !config.RotationEnabled.Valid {
		log.Warn("rotationEnabled is not specified; Rotation will be disabled")
		config.RotationEnabled = null.NewBool(false, true)
	}

	if config.RotationEnabled.Bool {
		// Process RotationDaysLimit

		if rotationDaysLimit := getEnv("ROTATION_DAYS_LIMIT", ""); rotationDaysLimit != "" {
			rotationDaysValue, err := strconv.ParseInt(rotationDaysLimit, 10, 64)
			if err != nil {
				log.Errorf("Could not parse 'ROTATION_DAYS_LIMIT' variable. Error: %v", err)
				os.Exit(1)
			}

			config.RotationDaysLimit = rotationDaysValue
		} else {
			log.Warnf("Environment variable 'ROTATION_DAYS_LIMIT' is empty")
		}

		if getFromEvent {
			if event.RotationDaysLimit.Valid {
				config.RotationDaysLimit = event.RotationDaysLimit.Int64
			}
		}

		if config.RotationDaysLimit == 0 {
			log.Error("RotationDaysLimit variable should be greater than 0")
			os.Exit(1)
		}
	}

	return config
}
