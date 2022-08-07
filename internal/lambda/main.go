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

func writeBackupToFile(fileName string, data []byte) error {
	outputFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("Failed to open output file (%v)", fileName)
	}
	defer outputFile.Close()

	n, err := outputFile.Write(data)
	if err != nil {
		return fmt.Errorf("Failed to write to output (%v) file; Written bytes %v", fileName, n)
	}

	return nil
}

func uploadToS3(ctx context.Context, client *cloud.Client, bucketName, keyName string, data []byte) error {
	_, err := client.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(keyName),
		ACL:                  types.ObjectCannedACLPrivate,
		Body:                 bytes.NewReader(data),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: types.ServerSideEncryptionAes256,
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

func Execute(config config.Config) {
	client, err := cloud.New(context.TODO(), config.Region)
	if err != nil {
		log.Error(fmt.Sprintf("Could not create AWS client"), "error", err)
		os.Exit(1)
	}

	ctx := context.TODO()
	timestamp := time.Now().Format(time.RFC3339)

	// Backup Cognito users
	users, err := client.CognitoClient.ListUsers(ctx, &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(config.UserPoolId),
	})
	if err != nil {
		log.Error(fmt.Sprintf("Failed to get list of cognito users"), "error", err)
		os.Exit(1)
	}

	usersData, err := json.Marshal(users)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to marshal cognito users structure"), "error", err)
		os.Exit(1)
	}

	err = uploadToS3(ctx, client, config.BucketName, getKeyName(config.BackupPrefix, timestamp, "users.json"), usersData)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to upload cognito users backup to S3"), "error", err)
		os.Exit(1)
	}

	// Backup Cognito groups
	groups, err := client.CognitoClient.ListGroups(ctx, &cognitoidentityprovider.ListGroupsInput{
		UserPoolId: aws.String(config.UserPoolId),
	})
	if err != nil {
		log.Error(fmt.Sprintf("Failed to get list of cognito groups"), "error", err)
		os.Exit(1)
	}

	groupsData, err := json.Marshal(groups)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to marshal cognito groups structure"), "error", err)
		os.Exit(1)
	}

	err = uploadToS3(ctx, client, config.BucketName, getKeyName(config.BackupPrefix, timestamp, "groups.json"), groupsData)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to upload cognito groups backup to S3"), "error", err)
		os.Exit(1)
	}
}
