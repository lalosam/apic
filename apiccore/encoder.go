package apiccore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//Encode generate the required Output
func Encode(w http.ResponseWriter, accept string, i interface{}, htmlLayout string) {
	log.Println("Encoding...")
	log.Println(accept)
	processed := false
	for _, a := range strings.Split(accept, ",") {
		switch a {
		case "*/*", "application/json", "":
			w.Header().Set("Content-Type", "application/json")
			enc := json.NewEncoder(w)
			enc.Encode(i)
		case "text/html":
			log.Println("HTML")
			w.Header().Set("Content-Type", "text/html")
			html := ToHTML(i, htmlLayout)
			fmt.Fprintln(w, html)
			processed = true
		}
		if processed {
			break
		}
	}

}
