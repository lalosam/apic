package glueapi

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

//GlueClient return a glue client
type GlueClient struct {
	client *glue.Glue
}

//NewClient return a Glue client
func NewClient() *GlueClient {
	log.Println("Getting a Glue Client")
	mySession := session.Must(session.NewSession())
	svc := glue.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
	return &GlueClient{client: svc}
}

//GetJobs get the list of Glue Jobs
func (c *GlueClient) GetJobs() []*glue.Job {
	log.Println("Getting Glue Jobs")
	request := glue.GetJobsInput{}
	response, e := c.client.GetJobs(&request)
	log.Println(e)
	if e != nil {
		panic(e)
	}
	jobs := response.Jobs
	for i, j := range jobs {
		log.Println(i, *j.Name)
	}
	log.Println(len(jobs))
	log.Println(response.NextToken)
	return jobs
}
