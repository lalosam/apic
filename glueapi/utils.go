package glueapi

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
	"rojosam.com/apic/apiccore"
	"rojosam.com/apic/glueapi/model"
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

//GetJob get the Glue Job
func (c *GlueClient) GetJob(name string) *model.Job {
	log.Printf("Getting Glue Job [%s]\n", name)
	request := glue.GetJobInput{JobName: &name}
	res := new(model.Job)
	response, e := c.client.GetJob(&request)
	if e != nil {
		panic(e)
	}
	job := response.Job
	res.SetGlueJob(job)
	return res
}

//GetJobs get the list of Glue Jobs
func (c *GlueClient) GetJobs(onlyOfficial bool) *[]model.Job {
	log.Println("Getting Glue Jobs")
	request := glue.GetJobsInput{}
	more := true
	var res []model.Job
	for more {
		response, e := c.client.GetJobs(&request)
		if e != nil {
			panic(e)
		}
		jobs := response.Jobs
		for i := range response.Jobs {
			job := jobs[i]
			name := job.Name
			if onlyOfficial && apiccore.IsOfficial(*name) {
				jj := new(model.Job)
				jj.SetGlueJob(job)
				res = append(res, *jj)
			}
		}
		nextToken := response.NextToken
		if nextToken == nil {
			log.Println("No more data")
			more = false
		} else {
			request.SetNextToken(*nextToken)
			log.Println(request)
		}
	}
	log.Printf("Retrieved jobs: %d\n", len(res))
	return &res
}

//GetJobRuns get the list of Glue Jobs
func (c *GlueClient) GetJobRuns(j *model.Job) *[]model.JobRun {
	log.Printf("Getting Glue Job Runs [%s]", j.Name)
	request := &glue.GetJobRunsInput{}
	request.SetJobName(j.Name).SetMaxResults(int64(3))
	log.Println(request)
	var res []model.JobRun
	response, e := c.client.GetJobRuns(request)
	if e != nil {
		panic(e)
	}
	jobRuns := response.JobRuns
	log.Printf("Job Runs: %d\n", len(jobRuns))
	for _, jr := range response.JobRuns {
		o := model.JobRun{}
		o.SetGlueJobRun(jr)
		res = append(res, o)
	}
	log.Printf("Retrieved Job Runs: %d\n", len(res))
	return &res
}

//GetJobsDetail get the list of Glue Jobs
func (c *GlueClient) GetJobsDetail(onlyOfficial bool) *[]model.JobDetail {
	jobs := *(c.GetJobs(onlyOfficial))
	var res []model.JobDetail
	for i := range jobs {
		job := jobs[i]
		log.Println(job.Name)
		jobRuns := c.GetJobRuns(&job)
		o := model.JobDetail{
			Job:     &job,
			JobRuns: jobRuns,
		}
		res = append(res, o)
	}
	log.Printf("JOBS: [%d]", len(jobs))
	log.Printf("JOBS_DETAIL: [%d]", len(res))
	return &res
}

//GetJobDetail get the list of Glue Jobs
func (c *GlueClient) GetJobDetail(name string) *[]model.JobDetail {
	job := *(c.GetJob(name))
	var res []model.JobDetail
	log.Println(job.Name)
	jobRuns := c.GetJobRuns(&job)
	o := model.JobDetail{
		Job:     &job,
		JobRuns: jobRuns,
	}
	res = append(res, o)
	log.Printf("JOBS_DETAIL: [%d]", len(res))
	return &res
}
