package model

import (
	"log"

	"github.com/aws/aws-sdk-go/service/glue"
)

//Job internal Glue Job model
type Job struct {
	Name            string  `html:"width:15%"`
	WorkerType      string  `html:"width:5%"`
	NumberOfWorkers int64   `html:"width:8%"`
	MaxCapacity     float64 `html:"width:7%"`
	GlueVersion     string  `html:"width:5%"`
	JarPaths        string  `html:"width:60%"`
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

//SetGlueVersion setter
func (j *Job) SetGlueVersion(glueVersion *string) *Job {
	if glueVersion != nil {
		j.GlueVersion = *glueVersion
	}
	return j
}

//SetJarPaths setter
func (j *Job) SetJarPaths(jarPaths *string) *Job {
	if jarPaths != nil {
		j.JarPaths = *jarPaths
	}
	return j
}

//SetGlueJob setter
func (j *Job) SetGlueJob(glueJob *glue.Job) *Job {
	j.SetName(glueJob.Name)
	j.SetWorkerType(glueJob.WorkerType)
	j.SetNumberOfWorkers(glueJob.NumberOfWorkers)
	j.SetMaxCapacity(glueJob.MaxCapacity)
	j.SetGlueVersion(glueJob.GlueVersion)
	j.SetJarPaths(glueJob.DefaultArguments["--extra-jars"])
	log.Println(glueJob.DefaultArguments)
	return j
}
