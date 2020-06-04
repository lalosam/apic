package model

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/glue"
)

//JobRun internal Glue Job Run model
type JobRun struct {
	Status          string
	ID              string
	StartedOn       string
	ExecutionTime   int64
	ErrorMessage    string
	LogURL          string `html:"elem:href@Log"`
	MaxCapacity     float64
	WorkerType      string
	NumberOfWorkers int64
}

//SetID setter
func (j *JobRun) SetID(id *string) *JobRun {
	if id != nil {
		j.ID = *id
	}
	return j
}

//SetStatus setter
func (j *JobRun) SetStatus(status *string) *JobRun {
	if status != nil {
		j.Status = *status
	}
	return j
}

//SetStartedOn setter
func (j *JobRun) SetStartedOn(startedOn *time.Time) *JobRun {
	if startedOn != nil {
		j.StartedOn = startedOn.String()
	}
	return j
}

//SetExecutionTime setter
func (j *JobRun) SetExecutionTime(executionTime *int64) *JobRun {
	if executionTime != nil {
		j.ExecutionTime = *executionTime
	}
	return j
}

//SetErrorMessage setter
func (j *JobRun) SetErrorMessage(errorMessage *string) *JobRun {
	if errorMessage != nil {
		j.ErrorMessage = *errorMessage
	}
	return j
}

//SetLogURL setter
func (j *JobRun) SetLogURL(id *string) *JobRun {
	if id != nil {
		j.LogURL = fmt.Sprintf("https://us-west-2.console.aws.amazon.com/cloudwatch/home?region=us-west-2#logsV2:log-groups/log-group/$252Faws-glue$252Fjobs$252Flogs-v2$3FlogStreamNameFilter$3D%s", *id)
	}
	return j
}

//SetMaxCapacity setter
func (j *JobRun) SetMaxCapacity(maxCapacity *float64) *JobRun {
	if maxCapacity != nil {
		j.MaxCapacity = *maxCapacity
	}
	return j
}

//SetWorkerType setter
func (j *JobRun) SetWorkerType(workerType *string) *JobRun {
	if workerType != nil {
		j.WorkerType = *workerType
	} else {
		j.WorkerType = "Std"
	}
	return j
}

//SetNumberOfWorkers setter
func (j *JobRun) SetNumberOfWorkers(numberOfWorkers *int64) *JobRun {
	if numberOfWorkers != nil {
		j.NumberOfWorkers = *numberOfWorkers
	}
	return j
}

//SetGlueJobRun setter
func (j *JobRun) SetGlueJobRun(jobRun *glue.JobRun) *JobRun {
	j.SetID(jobRun.Id)
	j.SetStatus(jobRun.JobRunState)
	j.SetStartedOn(jobRun.StartedOn)
	j.SetExecutionTime(jobRun.ExecutionTime)
	j.SetErrorMessage(jobRun.ErrorMessage)
	j.SetLogURL(jobRun.Id)
	j.SetMaxCapacity(jobRun.MaxCapacity)
	j.SetWorkerType(jobRun.WorkerType)
	j.SetNumberOfWorkers(jobRun.NumberOfWorkers)
	return j
}
