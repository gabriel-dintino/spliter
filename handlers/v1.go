package handlers

import (
	"fmt"
	"net/http"

	configurations "configurations"
	file "utilities"
)

func PingV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

func IngestFileV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ingest!")
	if exists, err := file.Exists(configurations.LineFileFullpath); err == nil {
		if exists {
			Run()
		}
	}

}
