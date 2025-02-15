package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/event-with-no-data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("HX-Trigger", "event1")
		w.Write([]byte("dispached event 1"))
	})
	http.HandleFunc("/event-with-string", func(w http.ResponseWriter, r *http.Request) {
		type Trigger struct {
			Event2 string `json:"event2"`
		}
		trigger := Trigger{
			Event2: "some string",
		}
		msg, _ := json.Marshal(trigger)
		w.Header().Set("HX-Trigger", string(msg))
		w.Write([]byte("dispatched event 2"))
	})
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
