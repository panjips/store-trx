package r2

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)


var R2Client *s3.Client
var once sync.Once

type R2Config struct {
	BucketName string
	AccountID string
	AccessKeyID string
	SecretKey string
}

func InitR2Client() *s3.Client {
	once.Do( func ()  {
		configKey := R2Config{
			BucketName: os.Getenv("R2_BUCKET_NAME"),
			AccountID: os.Getenv("R2_ACCOUNT_ID"),
			AccessKeyID: os.Getenv("R2_ACCESS_KEY"),
			SecretKey: os.Getenv("R2_SECRET_KEY"),
		}

		cfg, err := config.LoadDefaultConfig(context.TODO(), 
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				configKey.AccessKeyID, configKey.SecretKey, "")), 
		config.WithRegion("auto"))

		if err != nil {
			log.Fatal(err)
		}

		R2Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", configKey.AccountID))
		})

		fmt.Println("Successfully initialized R2")
	})

	return R2Client
}

func GetClient() *s3.Client {
	if R2Client != nil {
		return R2Client
	}

	return InitR2Client()
}
