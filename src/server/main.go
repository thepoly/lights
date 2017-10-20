package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type Color struct {
	Id int
	R  string
	G  string
	B  string
}

type Stats struct{
	Changes int
}

var statistics Stats

var color Color

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func StyleHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "styles.css")
}
func JSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "app.js")
}
func ImgHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/logo_m.png")
}
func ImgHandlerShirl(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/important.png")
}


func BowerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getInArduinoFormat(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/http")

	s := color.R + "\n" + color.G + "\n" + color.B
	w.Write([]byte(s));
}

func getStats(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/http")

	WriteJSON(w, statistics)
}

func ColorHandle(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, color)
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.MarshalIndent(	data, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(b)
	return nil
}

func ChangeHandler(w http.ResponseWriter, r *http.Request){
	 statistics.Changes += 1;
	 r.ParseForm();
		 color = Color{
			 0,
			 strings.Split(r.Form["color"][0],",")[0],
			 strings.Split(r.Form["color"][0],",")[1],
			 strings.Split(r.Form["color"][0],",")[2],
		 }
		 b, _ := json.MarshalIndent(color, "", " ")
		 error := ioutil.WriteFile("led/LEDL.txt", b, 0644)
		 _ = error

}

func main() {
	color.Id = 0;
	color.R = "255";
	color.G = "255";
	color.B = "255";
	statistics.Changes = 0;

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/images/logo_m.png",ImgHandler).Methods("GET");
	r.HandleFunc("/images/important.png",ImgHandlerShirl).Methods("GET");

	r.HandleFunc("/styles.css", StyleHandler).Methods("GET")
	r.HandleFunc("/app.js", JSHandler).Methods("GET")
	r.HandleFunc("/stats", getStats).Methods("GET")
	r.HandleFunc("/color", ColorHandle).Methods("POST")
	r.HandleFunc("/color", ColorHandle).Methods("GET")
	r.HandleFunc("/color/", ColorHandle).Methods("POST")
	r.HandleFunc("/color/", ColorHandle).Methods("GET")
	r.HandleFunc("/submit", ChangeHandler).Methods("POST")
	r.HandleFunc("/submit/", ChangeHandler).Methods("POST")
	r.HandleFunc("/LEDP.txt", getInArduinoFormat).Methods("GET")
	r.HandleFunc("/LEDL.txt", getInArduinoFormat).Methods("GET")

	//r.HandleFunc("/import", App.ImportHandler).Methods("GET")
	// Serve requests
	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Unable to ListenAndServe: %v", err)
	}
}
