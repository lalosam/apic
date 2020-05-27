package glueapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Handlers for Glue API
func Handlers(subrouter *mux.Router) {
	subrouter.StrictSlash(true)
	subrouter.HandleFunc("/status/", status)
	subrouter.HandleFunc("/status1/", status1)
	subrouter.HandleFunc("/jobs/", getJobs)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func status1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "STATUS1")
}

func getJobs(w http.ResponseWriter, r *http.Request) {
	c := NewClient()
	jobs := c.GetJobs()
	for _, j := range jobs {
		fmt.Fprintln(w, *j.Name)
	}
}
