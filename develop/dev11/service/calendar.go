package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/4halavandala/dev11/calendar"
)

const layout = "2006-01-02"

type CalendarService struct {
	cal *calendar.Calendar
}

type Resp struct {
	Result interface{} `json:"result,omitempty"`
	Error  error       `json:"error,omitempty"`
}

func New(cal *calendar.Calendar) *CalendarService {
	return &CalendarService{cal}
}

func (s *CalendarService) Save(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")

	date, err := time.Parse(layout, r.FormValue("date"))
	if err != nil {
		log.Println("Error date", err)
		err1 := json.NewEncoder(w).Encode(Resp{Error: err})
		if err1 != nil {
			log.Println("error send response", err)
		}
		w.WriteHeader(400)
		return
	}

	s.cal.Save(id, calendar.Event{
		ID:    id,
		Title: title,
		Date:  date,
	})

	err = json.NewEncoder(w).Encode(Resp{Result: id})
	if err != nil {
		log.Println("error send response", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func (s *CalendarService) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	s.cal.Delete(id)

	err := json.NewEncoder(w).Encode(Resp{Result: id})
	if err != nil {
		log.Println("error send response", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func (s *CalendarService) GetByDay(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	day, err := time.Parse(layout, values.Get("day"))
	if err != nil {
		log.Println("Error date", err)
		err1 := json.NewEncoder(w).Encode(Resp{Error: err})
		if err1 != nil {
			log.Println("error send response", err)
		}
		w.WriteHeader(400)
		return
	}

	events := s.cal.GetByDay(day)

	err = json.NewEncoder(w).Encode(Resp{Result: events})
	if err != nil {
		log.Println("error send response", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func (s *CalendarService) GetByWeek(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	start, err := time.Parse(layout, values.Get("day"))
	if err != nil {
		log.Println("Error date", err)
		err1 := json.NewEncoder(w).Encode(Resp{Error: err})
		if err1 != nil {
			log.Println("error send response", err)
		}
		w.WriteHeader(400)
		return
	}

	end := start.AddDate(0, 0, 7)

	events := s.cal.GetRangeDate(start, end)

	err = json.NewEncoder(w).Encode(Resp{Result: events})
	if err != nil {
		log.Println("error send response", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func (s *CalendarService) GetByMonth(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	start, err := time.Parse(layout, values.Get("day"))
	if err != nil {
		log.Println("Error date", err)
		err1 := json.NewEncoder(w).Encode(Resp{Error: err})
		if err1 != nil {
			log.Println("error send response", err)
		}
		w.WriteHeader(400)
		return
	}

	end := start.AddDate(0, 0, 30)

	events := s.cal.GetRangeDate(start, end)

	err = json.NewEncoder(w).Encode(Resp{Result: events})
	if err != nil {
		log.Println("error send response", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
