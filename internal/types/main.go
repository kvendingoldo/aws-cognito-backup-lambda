package types

type Response struct {
	Message string `json:"Answer:"`
}

type BackupEvent struct {
	ID          string `json:"id"`
	UserPoolArn string `json:"user_pool_arn"`
	BucketArn   string `json:"bucket_arn"`
	BackupName  string `json:"backup_name"`
	CleanUp     bool   `json:"cleanup"`
}

type RestoreEvent struct {
	ID          string `json:"id"`
	UserPoolArn string `json:"user_pool_arn"`
	BucketArn   string `json:"bucket_arn"`
	BackupName  string `json:"backup_name"`
	CleanUp     bool   `json:"cleanup"`
}
