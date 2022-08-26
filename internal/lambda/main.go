package lambda

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/kvendingoldo/aws-cognito-backup-lambda/internal/cloud"
	"github.com/kvendingoldo/aws-cognito-backup-lambda/internal/config"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func uploadToS3(ctx context.Context, client *cloud.Client, bucketName, keyName string, data []byte) error {
	_, err := client.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(keyName),
		ACL:                  types.ObjectCannedACLPrivate,
		Body:                 bytes.NewReader(data),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: types.ServerSideEncryptionAes256,
		Tagging:              aws.String(fmt.Sprintf("producer=%v", lambdaAwsTag)),
	})
	return err
}

func getKeyName(prefix, timestamp, name string) string {
	var output string
	if prefix == "" {
		output = fmt.Sprintf("%v/%v", timestamp, name)
	} else {
		output = fmt.Sprintf("%v/%v/%v", prefix, timestamp, name)
	}

	return output
}

func rotateBackups(ctx context.Context, client *cloud.Client, bucketName string, rotationDaysLimit int64) error {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(bucketName)
	objects, err := client.S3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println(err)
	}

	for _, obj := range objects.Contents {
		diffDays := int64(now.Sub(*obj.LastModified).Hours() / 24)
		if diffDays >= rotationDaysLimit {
			log.Infof("Object %v is %v days old. It's greater (or equal) than rotation days limit (%v days). Due to that it will be deleted", *obj.Key, diffDays, rotationDaysLimit)
			_, err = client.S3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
				Bucket: aws.String(bucketName),
				Key:    obj.Key,
			})

			if err != nil {
				log.Warnf("Failed to delete %v from bucket %v. Error: %v", *obj.Key, bucketName, err)
			} else {
				log.Debugf("Object %v has been successfully deleted from bucket %v", *obj.Key, bucketName)
			}
		} else {
			log.Debugf("Object %v is %v days old. It's less than rotation days limit (%v days). Due to that it will be skipped", *obj.Key, diffDays, rotationDaysLimit)
		}
	}

	return nil
}

func Execute(config config.Config) {
	client, err := cloud.New(context.TODO(), config.CognitoRegion, config.S3BucketRegion)
	if err != nil {
		log.Errorf("Could not create AWS client. Error: %v", err)
		os.Exit(1)
	}

	ctx := context.TODO()
	timestamp := time.Now().Format(time.RFC3339)

	// Backup Cognito users
	users, err := client.CognitoClient.ListUsers(ctx, &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(config.CognitoUserPoolId),
	})
	if err != nil {
		log.Errorf("Failed to get list of cognito users. Error: %v", err)
		os.Exit(1)
	}

	usersData, err := json.Marshal(users)
	if err != nil {
		log.Errorf("Failed to marshal cognito users structure. Error: %v", err)
		os.Exit(1)
	}

	err = uploadToS3(ctx, client, config.S3BucketName, getKeyName(config.BackupPrefix, timestamp, "users.json"), usersData)
	if err != nil {
		log.Errorf("Failed to upload cognito users backup to S3. Error: %v", err)
		os.Exit(1)
	}

	// Backup Cognito groups
	groups, err := client.CognitoClient.ListGroups(ctx, &cognitoidentityprovider.ListGroupsInput{
		UserPoolId: aws.String(config.CognitoUserPoolId),
	})
	if err != nil {
		log.Errorf("Failed to get list of cognito groups. Error: %v", err)
		os.Exit(1)
	}

	groupsData, err := json.Marshal(groups)
	if err != nil {
		log.Errorf("Failed to marshal cognito groups structure. Error: %v", err)
		os.Exit(1)
	}

	err = uploadToS3(ctx, client, config.S3BucketName, getKeyName(config.BackupPrefix, timestamp, "groups.json"), groupsData)
	if err != nil {
		log.Errorf("Failed to upload cognito groups backup to S3. Error: %v", err)
		os.Exit(1)
	}

	if config.RotationDaysLimit != -1 {
		err = rotateBackups(ctx, client, config.S3BucketName, config.RotationDaysLimit)
		if err != nil {
			log.Errorf("Rotation has been failed. Error: %v", err)
			os.Exit(1)
		}
	} else {
		log.Warnf("Pay attention that rotation is disabled; If you want to disable it, activate rotation via env variables, or event body")
	}
}
