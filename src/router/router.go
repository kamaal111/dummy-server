package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kamaal111/dummy-server/src/utils"
)

func HandleRequests(port string) {
	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(rootHandler)))

	log.Printf("Listening on %s\n", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Hello string `json:"hello"`
	}{
		Hello: "welcome",
	}
	output, err := json.Marshal(response)
	if err != nil {
		utils.MLogger("something went wrong while marshaling response", http.StatusInternalServerError, err)
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
