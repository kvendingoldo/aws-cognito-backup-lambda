package config

import (
	"fmt"
	"github.com/kvendingoldo/aws-cognito-backup-lambda/internal/types"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	Region       string
	UserPoolId   string
	BucketName   string
	BackupPrefix string
	RotationDays int
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func New(eventRaw interface{}) *Config {
	var config = Config{}
	var getFromEvent bool
	var event types.Event

	switch value := eventRaw.(type) {
	case types.Event:
		getFromEvent = true
		event = value
	default:
		getFromEvent = false
	}

	// Process Region
	if region := getEnv("AWS_REGION", ""); region == "" {
		log.Error("Required environment variable 'AWS_REGION' is empty. Please, specify it")
		os.Exit(1)
	} else {
		config.Region = region
	}

	// Process UserPoolId
	userPoolId := getEnv("USER_POOL_ID", "")
	if userPoolId == "" {
		log.Warnf("Environment variable 'USER_POOL_ID' is empty")
	} else {
		config.UserPoolId = userPoolId
	}

	if getFromEvent {
		if event.UserPoolId == "" {
			log.Warnf("Event contains empty user_pool_id")
			if userPoolId == "" {
				log.Error("UserPoolId is empty; Configure it via 'USER_POOL_ID' env variable OR pass in event body")
				os.Exit(1)
			}
		} else {
			config.UserPoolId = event.UserPoolId
		}
	}

	// Process BucketName
	bucketName := getEnv("BUCKET_NAME", "")
	if bucketName == "" {
		log.Warnf("Environment variable 'BUCKET_NAME' is empty")
	} else {
		config.BucketName = bucketName
	}

	if getFromEvent {
		if event.BucketName == "" {
			log.Warnf("Event contains empty bucket_name")
			if bucketName == "" {
				log.Error("BucketName is empty; Configure it via 'BUCKET_NAME' env variable OR pass in event body")
				os.Exit(1)
			}
		} else {
			config.BucketName = event.BucketName
		}
	}

	// Process backupPrefix
	backupPrefix := getEnv("BACKUP_PREFIX", "")
	if backupPrefix == "" {
		log.Warnf("Environment variable 'BACKUP_PREFIX' is empty")
	} else {
		config.BackupPrefix = backupPrefix
	}

	if getFromEvent {
		if event.BackupPrefix == "" {
			log.Warnf("Event contains empty backup_prefix")
		} else {
			config.BackupPrefix = event.BackupPrefix
		}
	}

	// Process RotationDays
	rotationDays := getEnv("ROTATION_DAYS", "")
	var rotationDaysValue int

	if rotationDays == "" {
		log.Warnf("Environment variable 'ROTATION_DAYS' is empty")
	} else {
		rotationDaysValue, err := strconv.Atoi(rotationDays)
		if err != nil {
			log.Errorf("Could not parse 'ROTATION_DAYS' variable. Error: %v", err)
			os.Exit(1)
		}

		if rotationDaysValue == 0 {
			log.Errorf("'ROTATION_DAYS' variable should be greater than 0. Error: %v", err)
			os.Exit(1)
		}

		if rotationDaysValue == -1 {
			log.Warnf("Pay attention that 'ROTATION_DAYS' = -1, it means that rotation is disabled")
		}

		config.RotationDays = rotationDaysValue
	}

	fmt.Println(rotationDaysValue)
	// TODO: please, rework this logic
	//if getFromEvent {
	//	if event.RotationDays == 0 {
	//		log.Error("Event contains RotationDays == 0; This values should be greater that 0, or -1 if you want to disable rotation")
	//		os.Exit(1)
	//	} else {
	//		if rotationDaysValue == 0 {
	//			log.Error("RotationDays is empty; Configure it via 'ROTATION_DAYS' env variable OR pass in event body")
	//			os.Exit(1)
	//		} else {
	//			config.RotationDays = event.RotationDays
	//		}
	//	}
	//}

	return &config
}
