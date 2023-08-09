package calendar

import (
	"sync"
	"time"
)

type Calendar struct {
	data sync.Map
}

type Event struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

func New() *Calendar {
	return &Calendar{}
}

func (c *Calendar) Save(id string, event Event) {
	c.data.Store(id, event)
}

func (c *Calendar) Delete(id string) {
	c.data.Delete(id)
}

func (c *Calendar) GetByDay(day time.Time) map[string][]Event {
	events := make(map[string][]Event)

	c.data.Range(func(k, v interface{}) bool {
		event, _ := v.(Event)
		if event.Date.Equal(day) {
			events[event.Date.String()] = append(events[event.Date.String()], event)
		}

		return true
	})

	return events
}

func (c *Calendar) GetRangeDate(start, end time.Time) map[string][]Event {
	events := make(map[string][]Event)

	c.data.Range(func(k, v interface{}) bool {
		event, _ := v.(Event)
		if event.Date.After(start) && event.Date.Before(end) {
			events[event.Date.String()] = append(events[event.Date.String()], event)
		}

		return true
	})

	return events
}
