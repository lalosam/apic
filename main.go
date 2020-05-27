package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"rojosam.com/apic/glueapi"
)

var (
	port = flag.String("port", "9090", "port")
)

/*
curl -I 'localhost:8080/api/'
curl -I 'localhost:8080/api/v1/'
curl -I 'localhost:8080/api/v1/status'
curl -I 'localhost:8080/api/v2/'
curl -I 'localhost:8080/api/v2/status'
curl -I 'localhost:8080/api/v1/status' -H "x-auth-token: admin"
curl -I 'localhost:8080/api/v1/status' -H "x-auth-token: notadmin"
*/
func main() {
	log.SetOutput(os.Stdout)
	flag.Parse()
	var router = mux.NewRouter()
	var api = router.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	api.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("x-auth-token") != "admin" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})
	var api1 = api.PathPrefix("/v1").Subrouter()
	api1.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	api1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	glueapi.Handlers(api1.PathPrefix("/glue/").Subrouter())

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + *port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	//http.ListenAndServe(":"+*port, router)
}
