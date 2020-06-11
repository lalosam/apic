package s3api

import (
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
)

//Bucket internal S3 bucket model
type Bucket struct {
	Name         string
	CreationDate string
}

//SetName setter
func (b *Bucket) SetName(name *string) *Bucket {
	if name != nil {
		b.Name = *name
	}
	return b
}

//SetCreationDate setter
func (b *Bucket) SetCreationDate(creationDate *time.Time) *Bucket {
	if creationDate != nil {
		b.CreationDate = creationDate.String()
	}
	return b
}

//SetS3Bucket setter
func (b *Bucket) SetS3Bucket(s3Bucket *s3.Bucket) *Bucket {
	b.SetName(s3Bucket.Name)
	b.SetCreationDate(s3Bucket.CreationDate)
	return b
}
