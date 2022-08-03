package config

import (
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/types"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Region string
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
		log.Errorf("Required environment variable 'AWS_REGION' is empty. Please, specify", region)
		os.Exit(1)
	} else {
		config.Region = region
	}

	return &config
}
