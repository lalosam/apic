package model

//JobDetail internal Glue Job Detail model
type JobDetail struct {
	Job     *Job      `html:"layout:TABLE"`
	JobRuns *[]JobRun `html:"layout:TABLE"`
}
