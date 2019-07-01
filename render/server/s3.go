package server

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (a *App) BuildS3(){
	region := "us-east-1"
	chainErrors := true
	sess, err := session.NewSession(&aws.Config{
		Region: &region,
		CredentialsChainVerboseErrors: &chainErrors,
	})

	if err != nil {
		panic(err)
	}
	svc := s3.New(sess)
	a.S3 = svc
}
