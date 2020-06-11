package glueapi

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"rojosam.com/apic/apiccore"
)

//Handlers for Glue API
func Handlers(subrouter *mux.Router) {
	subrouter.StrictSlash(true)
	subrouter.HandleFunc("/job/{name}", getJob)
	subrouter.HandleFunc("/jobs/", getJobs)
	subrouter.HandleFunc("/jobdetail/{name}", getJobDetail)
	subrouter.HandleFunc("/jobsdetail/", getJobsDetail)
}

func getJob(w http.ResponseWriter, r *http.Request) {
	c := NewClient()
	vars := mux.Vars(r)
	job := c.GetJob(vars["name"])
	apiccore.Encode(w, r.Header.Get("Accept"), *job, "TABLE")
}

func getJobs(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")
	c := NewClient()
	jobs := c.GetJobs(true)
	apiccore.Encode(w, accept, *jobs, "TABLE")
	//html := apiccore.ToHTML(*jobs, "TABLE")
	//fmt.Fprintln(w, html)
	//enc := json.NewEncoder(w)
	//enc.Encode(*jobs)
}

func getJobsDetail(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")
	c := NewClient()
	jobsDetail := c.GetJobsDetail(true)
	apiccore.Encode(w, accept, *jobsDetail, "DIV")
	//enc := json.NewEncoder(w)
	//enc.Encode(*jobsDetail)
}

func getJobDetail(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")
	log.Println("Get Job Detail")
	c := NewClient()
	vars := mux.Vars(r)
	jobDetail := c.GetJobDetail(vars["name"])
	apiccore.Encode(w, accept, *jobDetail, "DIV")
	//enc := json.NewEncoder(w)
	//enc.Encode(*jobDetail)
}
