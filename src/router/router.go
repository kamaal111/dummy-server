package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamaal111/dummy-server/src/utils"
)

func HandleRequests(port string) {
	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(rootHandler)))
	mux.Handle("/post", loggerMiddleware(http.HandlerFunc(postHandler)))

	log.Printf("Listening on %s\n", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mappedBody := make(map[string]string)
	err = json.Unmarshal(body, &mappedBody)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(mappedBody)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
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
