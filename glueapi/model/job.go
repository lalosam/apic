package model

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/glue"
)

//Job internal Glue Job model
type Job struct {
	Name            string
	WorkerType      string
	NumberOfWorkers int64
	MaxCapacity     float64
}

//SetName setter
func (j *Job) SetName(name *string) *Job {
	if name != nil {
		j.Name = *name
	}
	return j
}

//SetWorkerType setter
func (j *Job) SetWorkerType(workerType *string) *Job {
	if workerType != nil {
		j.WorkerType = *workerType
	} else {
		j.WorkerType = "Std"
	}
	return j
}

//SetNumberOfWorkers setter
func (j *Job) SetNumberOfWorkers(numberOfWorkers *int64) *Job {
	if numberOfWorkers != nil {
		j.NumberOfWorkers = *numberOfWorkers
	}
	return j
}

//SetMaxCapacity setter
func (j *Job) SetMaxCapacity(maxCapacity *float64) *Job {
	if maxCapacity != nil {
		j.MaxCapacity = *maxCapacity
	}
	return j
}

//SetGlueJob setter
func (j *Job) SetGlueJob(glueJob *glue.Job) *Job {
	j.SetName(glueJob.Name)
	j.SetWorkerType(glueJob.WorkerType)
	j.SetNumberOfWorkers(glueJob.NumberOfWorkers)
	j.SetMaxCapacity(glueJob.MaxCapacity)
	return j
}

func (j *Job) String() string {
	return fmt.Sprintf("%s, %s, %d, %f", j.Name, j.WorkerType, j.NumberOfWorkers, j.MaxCapacity)
}

func (J *Job) ToHtml() string {
	return ""
}
