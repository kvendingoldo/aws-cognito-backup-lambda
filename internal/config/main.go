package config

import (
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

	RotationEnabled   bool
	RotationDaysLimit int
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
		if event.AWSRegion == "" {
			if config.AWSRegion == "" {
				log.Error("awsRegion is empty; Configure it via 'AWS_REGION' env variable OR pass in event body")
				os.Exit(1)
			}
		} else {
			config.AWSRegion = event.AWSRegion
		}
	}

	// Process CognitoRegion
	if cognitoRegion := getEnv("COGNITO_REGION", ""); cognitoRegion != "" {
		config.CognitoRegion = cognitoRegion
	} else {
		log.Warn("Environment variable 'COGNITO_REGION' is empty")
	}
	if getFromEvent {
		if event.CognitoRegion == "" {
			log.Warn("Event contains empty cognitoRegion variable")
			if config.CognitoRegion == "" {
				log.Warnf("cognitoRegion is empty; Default region %s will be used", config.AWSRegion)
				config.CognitoRegion = config.AWSRegion
			}
		} else {
			config.CognitoRegion = event.CognitoRegion
		}
	}

	// Process S3BucketRegion
	if bucketRegion := getEnv("BUCKET_REGION", ""); bucketRegion != "" {
		config.S3BucketRegion = bucketRegion
	} else {
		log.Warn("Environment variable 'BUCKET_REGION' is empty")
	}
	if getFromEvent {
		if event.S3BucketRegion == "" {
			log.Warn("Event contains empty s3BucketRegion variable")
			if config.S3BucketRegion == "" {
				log.Warnf("bucketRegion is empty; Default region %s will be used", config.AWSRegion)
				config.S3BucketRegion = config.AWSRegion
			}
		} else {
			config.S3BucketRegion = event.CognitoRegion
		}
	}

	// Process CognitoUserPoolId
	if cognitoUserPoolId := getEnv("COGNITO_USER_POOL_ID", ""); cognitoUserPoolId != "" {
		config.CognitoUserPoolId = cognitoUserPoolId
	} else {
		log.Warn("Environment variable 'COGNITO_USER_POOL_ID' is empty")
	}
	if getFromEvent {
		if event.CognitoUserPoolId == "" {
			log.Warn("Event contains empty cognitoUserPoolID")
			if config.CognitoUserPoolId == "" {
				log.Error("cognitoUserPoolID is empty; Configure it via 'COGNITO_USER_POOL_ID' env variable OR pass in event body")
				os.Exit(1)
			}
		} else {
			config.CognitoUserPoolId = event.CognitoUserPoolId
		}
	}

	// Process S3BucketName
	if s3BucketName := getEnv("S3_BUCKET_NAME", ""); s3BucketName != "" {
		config.S3BucketName = s3BucketName
	} else {
		log.Warn("Environment variable 'S3_BUCKET_NAME' is empty")
	}
	if getFromEvent {
		if event.S3BucketName == "" {
			log.Warn("Event contains empty s3BucketName")
			if config.S3BucketName == "" {
				log.Error("BucketName is empty; Configure it via 'S3_BUCKET_NAME' env variable OR pass in event body")
				os.Exit(1)
			}
		} else {
			config.S3BucketName = event.S3BucketName
		}
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

	// Process RotationDays
	rotationDays := getEnv("ROTATION_DAYS_LIMIT", "")
	var rotationDaysValue int

	if rotationDays == "" {
		log.Warnf("Environment variable 'ROTATION_DAYS_LIMIT' is empty")
		config.RotationDaysLimit = -1
	} else {
		rotationDaysValue, err := strconv.Atoi(rotationDays)
		if err != nil {
			log.Errorf("Could not parse 'ROTATION_DAYS_LIMIT' variable. Error: %v", err)
			os.Exit(1)
		}

		if rotationDaysValue == 0 {
			log.Errorf("'ROTATION_DAYS_LIMIT' variable should be greater than 0. Error: %v", err)
			os.Exit(1)
		}

		if rotationDaysValue == -1 {
			log.Warnf("Pay attention that 'ROTATION_DAYS_LIMIT' = -1, it means that rotation is disabled")
		}

		config.RotationDaysLimit = rotationDaysValue
	}

	if getFromEvent {
		if event.RotationDays == "" {
			log.Error("Event contains empty RotationDays")
			if config.RotationDaysLimit == 0 {
				config.RotationDaysLimit = -1
			}
		} else {
			if rotationDaysValue != 0 {
				rotationDaysEventValue, err := strconv.Atoi(event.RotationDays)
				if err != nil {
					log.Errorf("Could not parse RotationDays variable from event. Error: %v", err)
					os.Exit(1)
				}
				config.RotationDaysLimit = rotationDaysEventValue
			} else {
				log.Error("RotationDays is empty; Configure it via 'ROTATION_DAYS_LIMIT' env variable OR pass in event body")
				os.Exit(1)
			}
		}
	}

	return config
}
