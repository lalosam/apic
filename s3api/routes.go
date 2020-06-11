package s3api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Handlers for S3 API
func Handlers(subrouter *mux.Router) {
	subrouter.StrictSlash(true)
	subrouter.HandleFunc("/stack/{prefix}", getStackBuckets)
	subrouter.HandleFunc("/stacks/", getBuckets)
}

func getStackBuckets(w http.ResponseWriter, r *http.Request) {

}

func getBuckets(w http.ResponseWriter, r *http.Request) {

}
