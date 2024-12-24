package task12

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

const (
	datePattern      = "2006-01-02"
	monthDatePattern = "2006-01"
)

var service = Service{Events: make([]Event, 0, 10)}

type Event struct {
	UserID      int       `json:"user_id"`
	EventID     int       `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Service struct {
	mu     *sync.Mutex
	Events []Event
}

func (s *Service) save(event Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Events = append(s.Events, event)
}

func (s *Service) update(event Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, e := range s.Events {
		if e.EventID == event.EventID {
			s.Events[i] = event
		}
	}
}

func (s *Service) delete(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, e := range s.Events {
		if e.EventID == id {
			s.Events[i] = s.Events[len(s.Events)-1]
			s.Events = s.Events[:len(s.Events)-1]
		}
	}
}

func (s *Service) getEventsByDay(userID int, date time.Time) []Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []Event

	for _, e := range s.Events {
		if e.UserID == userID && e.Date == date {
			result = append(result, e)
		}
	}

	return result
}

func (s *Service) getEventsByWeek(userID int, date time.Time) []Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	dateYear, dateWeek := date.ISOWeek()
	var result []Event

	for _, e := range s.Events {
		year, week := e.Date.ISOWeek()
		if e.UserID == userID && year == dateYear && dateWeek == week {
			result = append(result, e)
		}
	}

	return result
}

func (s *Service) getEventsByMonth(userID int, date time.Time) []Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []Event

	for _, e := range s.Events {
		if e.UserID == userID && e.Date.Year() == date.Year() && e.Date.Month() == date.Month() {
			result = append(result, e)
		}
	}

	return result
}

func parseBody(r *http.Request) (*Event, error) {
	event := Event{}

	eventID, err := strconv.Atoi(r.FormValue("event_id"))
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return nil, err
	}

	date, err := time.Parse(datePattern, r.FormValue("date"))
	if err != nil {
		return nil, err
	}

	event.EventID = eventID
	event.UserID = userID
	event.Title = r.FormValue("title")
	event.Description = r.FormValue("description")
	event.Date = date
	return &event, nil
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "request method should be POST", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseBody(r)
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request body: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	service.save(*event)
	fmt.Println(service.Events)
	successResponse(w, http.StatusCreated, []Event{*event})
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "request method should be POST", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseBody(r)
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request body: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	service.update(*event)
	fmt.Println(service.Events)
	successResponse(w, http.StatusOK, []Event{*event})
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "request method should be POST", http.StatusMethodNotAllowed)
		return
	}

	eventID, err := strconv.Atoi(r.FormValue("event_id"))
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request body: %s\n", err.Error()), http.StatusBadRequest)
	}

	service.delete(eventID)
	fmt.Println(service.Events)
	successResponse(w, http.StatusNoContent, []Event{})
}

func getEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, "request method should be GET", http.StatusMethodNotAllowed)
		return
	}

	date, err := time.Parse(datePattern, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request params: %s\n", err.Error()), http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request params: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	events := service.getEventsByDay(userID, date)
	successResponse(w, http.StatusOK, events)
}

func getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, "request method should be GET", http.StatusMethodNotAllowed)
		return
	}

	date, err := time.Parse(datePattern, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request params: %s\n", err.Error()), http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request params: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	events := service.getEventsByWeek(userID, date)
	successResponse(w, http.StatusOK, events)
}

func getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, "request method should be GET", http.StatusMethodNotAllowed)
		return
	}

	date, err := time.Parse(monthDatePattern, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request params: %s\n", err.Error()), http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		errorResponse(w, fmt.Sprintf("error parsing request params: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	events := service.getEventsByMonth(userID, date)
	successResponse(w, http.StatusOK, events)
}

func logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func successResponse(w http.ResponseWriter, status int, body []Event) {
	jsonResponse, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
}

func errorResponse(w http.ResponseWriter, e string, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{Error: e}
	jsonResponse, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	router := http.NewServeMux()
	router.HandleFunc("/create_event", logger(http.HandlerFunc(createEvent)))
	router.HandleFunc("/update_event", logger(http.HandlerFunc(updateEvent)))
	router.HandleFunc("/delete_event", logger(http.HandlerFunc(deleteEvent)))
	router.HandleFunc("/events_for_day", logger(http.HandlerFunc(getEventsForDay)))
	router.HandleFunc("/events_for_week", logger(http.HandlerFunc(getEventsForWeek)))
	router.HandleFunc("/events_for_month", logger(http.HandlerFunc(getEventsForMonth)))

	log.Printf("starting server\n")
	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatalf("error starting server: %s\n", err.Error())
		}
	}()

	<-done
	log.Printf("server stopped\n")
}
