package main

import (
	"net/http"
  "encoding/json"


	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

)

type Color struct {
	R             string
	G             string
	B             string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func ColorHandle(w http.ResponseWriter, r *http.Request){
  c := Color{
    "255",
    "255",
    "255",
  }
  WriteJSON(w,c);
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(b)
	return nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
  r.HandleFunc("/color", ColorHandle).Methods("GET")

  //r.HandleFunc("/import", App.ImportHandler).Methods("GET")
	// Serve requests
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Unable to ListenAndServe: %v", err)
	}
}
