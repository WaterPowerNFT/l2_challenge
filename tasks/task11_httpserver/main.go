package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Event struct {
	ID     string `json:"id"`
	UID    string `json:"uid"`
	Date   Date   `json:"date"`
	Action string `json:"action"`
}

type Date struct {
	Day   uint64 `json:"day"`
	Month uint64 `json:"month"`
	Year  uint64 `json:"year"`
}
type MassiveMutex struct {
	events []Event
	*sync.RWMutex
	*log.Logger
}

func sendError(w http.ResponseWriter, errorStr string, status int) {
	response := struct {
		Error string `json:"error"`
	}{errorStr}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
}

func (mm *MassiveMutex) deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	mm.Println(r.URL)
	mm.RWMutex.Lock()
	defer mm.RWMutex.Unlock()
	for index, item := range mm.events {
		if item.ID == params["id"] {
			mm.events = append(mm.events[:index], mm.events[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(mm.events)
}

func (mm *MassiveMutex) createEvent(w http.ResponseWriter, r *http.Request) {
	mm.Println(r.URL)
	mm.RWMutex.Lock()
	defer mm.RWMutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	var med Event
	_ = json.NewDecoder(r.Body).Decode(&med)
	med.ID = strconv.Itoa(rand.Intn(100000000))
	mm.events = append(mm.events, med)
	json.NewEncoder(w).Encode(mm.events)
}

func (mm *MassiveMutex) updateEvent(w http.ResponseWriter, r *http.Request) {
	mm.Println(r.URL)
	mm.RWMutex.Lock()
	defer mm.RWMutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range mm.events {
		if item.ID == params["ID"] {
			mm.events = append(mm.events[:index], mm.events[index+1:]...)
			var med Event
			_ = json.NewDecoder(r.Body).Decode(&med)
			med.ID = params["ID"]
			json.NewEncoder(w).Encode(mm.events)
			return
		}
	}
}

func (mm *MassiveMutex) eventsForDay(w http.ResponseWriter, r *http.Request) {
	mm.Println(r.URL)
	mm.RWMutex.RLock()
	defer mm.RWMutex.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var curEvents []Event
	u64, err := strconv.ParseUint(params["Day"], 10, 32)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, item := range mm.events {
		if (item.ID == params["ID"]) && (item.Date.Day == u64) {
			curEvents = append(curEvents, item)
		}
	}
	json.NewEncoder(w).Encode(curEvents)
}

func (mm *MassiveMutex) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	mm.Println(r.URL)
	mm.RWMutex.RLock()
	defer mm.RWMutex.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var curEvents []Event
	u64, err := strconv.ParseUint(params["Day"], 10, 32)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u64_month, err := strconv.ParseUint(params["Month"], 10, 32)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, item := range mm.events {
		if (item.ID == params["ID"]) && (item.Date.Day >= u64 && item.Date.Day <= u64 && u64_month == item.Date.Month) {
			curEvents = append(curEvents, item)
		}
	}
	json.NewEncoder(w).Encode(curEvents)
}

func (mm *MassiveMutex) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	mm.Println(r.URL)
	mm.RWMutex.RLock()
	defer mm.RWMutex.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var curEvents []Event
	u64, err := strconv.ParseUint(params["Month"], 10, 32)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, item := range mm.events {
		if (item.ID == params["ID"]) && (item.Date.Month == u64) {
			curEvents = append(curEvents, item)
		}
	}
	json.NewEncoder(w).Encode(curEvents)
}

func initMassiveMutex() *MassiveMutex {
	return &MassiveMutex{Logger: log.New(os.Stdout, "logger: ", log.Lshortfile),
		RWMutex: new(sync.RWMutex),
		events:  make([]Event, 10),
	}
}

func main() {
	r := mux.NewRouter()
	massiveMutexUnit := initMassiveMutex()
	massiveMutexUnit.events = append(massiveMutexUnit.events,
		Event{ID: "1", UID: "1", Date: Date{Day: 2, Month: 2, Year: 2010}, Action: "Go to medical center"})

	r.HandleFunc("/create_event", massiveMutexUnit.createEvent).Methods("POST")

	r.HandleFunc("/update_event", massiveMutexUnit.updateEvent).Methods("POST")

	r.HandleFunc("/delete_event", massiveMutexUnit.deleteEvent).Methods("POST")

	r.HandleFunc("/events_for_day", massiveMutexUnit.eventsForDay).Methods("GET")
	r.HandleFunc("/events_for_week", massiveMutexUnit.eventsForWeek).Methods("GET")
	r.HandleFunc("/events_for_month", massiveMutexUnit.eventsForMonth).Methods("GET")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
