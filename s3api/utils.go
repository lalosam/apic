package s3api

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//S3Client return a S3 client
type S3Client struct {
	client *s3.S3
}

//NewClient return a S3 client
func NewClient() *S3Client {
	log.Println("Getting a S3 Client")
	mySession := session.Must(session.NewSession())
	svc := s3.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
	return &S3Client{client: svc}
}

/*func(c *S3Client) GetBuckets(onlyOfficial bool) *[]model.Bucket {
	log.Println("Getting S3 Buckets")
	request := s3.ListObjectsV2Input{}
	///request.
	log.Println(request)
}*/
