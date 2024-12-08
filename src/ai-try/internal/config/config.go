package config

import (
	"fmt"
	"log"
	"os"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

type (
	Env struct {
		// Environment Current environment
		Environment string
		// DefaultAWSRegion default AWS region
		DefaultAWSRegion string
		// TableName the name of table
		TableName string
		// LogLevel defines the log level
		LogLevel string
		// AWSConfig aws config
		AWSConfig awsconfig.Config
		// S3Bucket the name of the s3 bucket
		S3Bucket string
		// S3Directory the directory in the s3 bucket
		S3Directory string
		// VisionAPIKey the api key for OpenAI vision API
		VisionAPIKey string
		// GeminiAPIKey the api key for OpenAI vision API
		GeminiAPIKey string
		// CFSignKey the key to sign the cloudfront url
		CFSignKey string
		// CDNUrls the urls of the cloudfront, separated by comma
		CDNUrls string
		// SecretName the name of the secret in secrets manager
		SecretName string
	}
)

const (
	// EnvironmentDevelopment development environment
	EnvironmentDevelopment = "development"
)

// NewEnv creates a new instance of Env
// tries to load the env variables from .env
func NewEnv() *Env {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting path")
	}
	dotenvError := godotenv.Load(fmt.Sprintf("%s/.env", path))
	if dotenvError != nil {
		log.Printf("error loading .env file, ignoring dotenv")
	}

	return &Env{
		Environment:      os.Getenv("ENVIRONMENT"),
		DefaultAWSRegion: os.Getenv("AWS_REGION"),
		TableName:        os.Getenv("DYNAMODB_TABLE_NAME_SERVICE"),
		LogLevel:         os.Getenv("LOG_LEVEL"),
		S3Bucket:         os.Getenv("S3_BUCKET_NAME"),
		S3Directory:      os.Getenv("S3_DIRECTORY"),
		VisionAPIKey:     os.Getenv("OPENAI_VISION_API_KEY"),
		GeminiAPIKey:     os.Getenv("GEMINI_API_KEY"),
		CFSignKey:        os.Getenv("CLOUDFRONT_PUBLIC_KEY_ID"),
		CDNUrls:          os.Getenv("CDN_URLS"),
		SecretName:       os.Getenv("SECRET_NAME"),
	}
}
