package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strings"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type Color struct {
	Id int
	R  string
	G  string
	B  string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func ColorHandle(w http.ResponseWriter, r *http.Request) {
	var colors []Color

	data, err := ioutil.ReadFile("LEDL.txt")
	if err != nil {
		panic("asdf")
	}

	if err := json.Unmarshal(data, &colors); err != nil {
		panic(err)
	}

	WriteJSON(w, colors)
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

func ChangeHandler(w http.ResponseWriter, r *http.Request){
	 r.ParseForm();
	 var colors []Color

	 data, err := ioutil.ReadFile("LEDL.txt")
	 if err != nil {
		 panic("asdf")
	 }

	 if err := json.Unmarshal(data, &colors); err != nil {
		 panic(err)
	 }

	 id, err:= strconv.Atoi(strings.Split(r.Form["word"][0],",")[3]);
	 colors[id] = Color{
		 id,
		 strings.Split(r.Form["word"][0],",")[0],
		 strings.Split(r.Form["word"][0],",")[1],
		 strings.Split(r.Form["word"][0],",")[2],
	 }
	 b, err := json.MarshalIndent(colors, "", " ")
	 error := ioutil.WriteFile("LEDL.txt", b, 0644)
	 _ = error
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/color", ColorHandle).Methods("POST")
	r.HandleFunc("/color", ColorHandle).Methods("GET")
	r.HandleFunc("/submit", ChangeHandler).Methods("POST")

	//r.HandleFunc("/import", App.ImportHandler).Methods("GET")
	// Serve requests
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Unable to ListenAndServe: %v", err)
	}
}
