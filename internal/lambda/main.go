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
	"net/http"
	"os"
	"path/filepath"
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

func uploadToS3(ctx context.Context, client *cloud.Client, bucketName, filePath, filePrefix, timestamp string) error {
	uploadFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open local filepath [%v]: %+v", filePath, err)
	}

	defer uploadFile.Close()

	uploadFileInfo, _ := uploadFile.Stat()
	var fileSize int64 = uploadFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	uploadFile.Read(fileBuffer)

	var keyName string
	if filePrefix == "" {
		keyName = fmt.Sprintf("%v/%v", timestamp, filepath.Base(filePath))
	} else {
		keyName = fmt.Sprintf("%v/%v/%v", filePrefix, timestamp, filepath.Base(filePath))
	}

	_, err = client.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(keyName),
		ACL:                  types.ObjectCannedACLPrivate,
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        fileSize,
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: types.ServerSideEncryptionAes256,
	})
	return err

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
	err = writeBackupToFile("users.json", usersData)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to write cognito users backup to disk"), "error", err)
		os.Exit(1)
	}
	err = uploadToS3(ctx, client, config.BucketName, "./users.json", config.BackupPrefix, timestamp)
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
	err = writeBackupToFile("groups.json", groupsData)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to write cognito groups backup to disk"), "error", err)
		os.Exit(1)
	}
	err = uploadToS3(ctx, client, config.BucketName, "./groups.json", config.BackupPrefix, timestamp)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to upload cognito groups backup to S3"), "error", err)
		os.Exit(1)
	}
}
