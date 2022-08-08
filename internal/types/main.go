package types

type Response struct {
	Message string `json:"Answer:"`
}

type Event struct {
	ID           string `json:"id"`
	UserPoolId   string `json:"user_pool_id"`
	BucketName   string `json:"bucket_name"`
	BackupPrefix string `json:"backup_prefix"`
	RotationDays string `json:"rotation_days_limit"`
}
