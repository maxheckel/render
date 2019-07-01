package server

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jinzhu/gorm"
)

type App struct {
	Database *gorm.DB
	S3 *s3.S3
	Config Config
}
