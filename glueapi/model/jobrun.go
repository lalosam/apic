package model

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/glue"
)

//JobRun internal Glue Job Run model
type JobRun struct {
	Status          string  `html:"style:SUCCEEDED@color$green$$font-weight$bold|RUNNING@color$blue$$font-weight$bold|FAILED@color$red$$font-weight$bold|STOPPED@text-decoration$line-through$$font-weight$bold;width:5%"`
	ID              string  `html:"width:30%"`
	StartedOn       string  `html:"width:15%"`
	ExecutionTime   int64   `html:"width:7%"`
	ErrorMessage    string  `html:"width:22%"`
	Log             string  `html:"elem:href@Log;width:3%"`
	MaxCapacity     float64 `html:"width:6%"`
	WorkerType      string  `html:"width:6%"`
	NumberOfWorkers int64   `html:"width:6%"`
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

//SetLog setter
func (j *JobRun) SetLog(id *string) *JobRun {
	if id != nil {
		j.Log = fmt.Sprintf("https://us-west-2.console.aws.amazon.com/cloudwatch/home?region=us-west-2#logsV2:log-groups/log-group/$252Faws-glue$252Fjobs$252Flogs-v2/log-events/%s-driver?filterPattern$3D$253F$2522DataLakeGlueJob$253A$2522", *id)
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
	j.SetLog(jobRun.Id)
	j.SetMaxCapacity(jobRun.MaxCapacity)
	j.SetWorkerType(jobRun.WorkerType)
	j.SetNumberOfWorkers(jobRun.NumberOfWorkers)
	return j
}
