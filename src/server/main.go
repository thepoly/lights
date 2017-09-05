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

func getP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/http")
	var colors []Color

	data, err := ioutil.ReadFile("led/LEDL.txt")
	if err != nil {
		panic("asdf")
	}

	if err := json.Unmarshal(data, &colors); err != nil {
		panic(err)
	}
	s := colors[0].R + "\n" + colors[0].G + "\n" + colors[0].B
	w.Write([]byte(s));
}

func getO(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/http")
	var colors []Color

	data, err := ioutil.ReadFile("led/LEDL.txt")
	if err != nil {
		panic("asdf")
	}

	if err := json.Unmarshal(data, &colors); err != nil {
		panic(err)
	}
	s := colors[1].R + "\n" + colors[1].G + "\n" + colors[1].B
	w.Write([]byte(s));
}

func getL(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/http")
	var colors []Color

	data, err := ioutil.ReadFile("led/LEDL.txt")
	if err != nil {
		panic("asdf")
	}

	if err := json.Unmarshal(data, &colors); err != nil {
		panic(err)
	}
	s := colors[2].R + "\n" + colors[2].G + "\n" + colors[2].B
	w.Write([]byte(s));
}

func getY(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/http")
	var colors []Color

	data, err := ioutil.ReadFile("led/LEDL.txt")
	if err != nil {
		panic("asdf")
	}

	if err := json.Unmarshal(data, &colors); err != nil {
		panic(err)
	}
	s := colors[3].R + "\n" + colors[3].G + "\n" + colors[3].B
	w.Write([]byte(s));
}

func ColorHandle(w http.ResponseWriter, r *http.Request) {
	var colors []Color

	data, err := ioutil.ReadFile("led/LEDL.txt")
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

	 data, err := ioutil.ReadFile("led/LEDL.txt")
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
	 error := ioutil.WriteFile("led/LEDL.txt", b, 0644)
	 _ = error
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/color", ColorHandle).Methods("POST")
	r.HandleFunc("/color", ColorHandle).Methods("GET")
	r.HandleFunc("/color/", ColorHandle).Methods("POST")
	r.HandleFunc("/color/", ColorHandle).Methods("GET")
	r.HandleFunc("/submit", ChangeHandler).Methods("POST")
	r.HandleFunc("/submit/", ChangeHandler).Methods("POST")
	r.HandleFunc("/LEDL.txt", getL).Methods("GET")
	r.HandleFunc("/LEDO.txt", getO).Methods("GET")
	r.HandleFunc("/LEDP.txt", getP).Methods("GET")
	r.HandleFunc("/LEDY.txt", getY).Methods("GET")
	//r.HandleFunc("/import", App.ImportHandler).Methods("GET")
	// Serve requests
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Unable to ListenAndServe: %v", err)
	}
}
