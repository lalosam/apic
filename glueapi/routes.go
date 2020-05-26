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
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func status1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "STATUS1")
}
