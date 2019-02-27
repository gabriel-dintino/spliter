package routers_v1

import (
	"fmt"
	"net/http"
	handlers "handlers"
)

func PingV1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	handlers.PingV1(w, r)
}

func IngestFileV1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	//fmt.Fprintf(w, "Ingest")
	handlers.IngestFileV1(w, r)
}
