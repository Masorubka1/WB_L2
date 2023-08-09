package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Masorubka1/dev11/calendar"
	"github.com/Masorubka1/dev11/service"
)

/*
Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
Реализовать middleware для логирования запросов


Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
Реализовать все методы.
Бизнес логика НЕ должна зависеть от кода HTTP сервера.
В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
*/

const (
	portNumber = ":8060"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Page"))
}

func handleRequest(service *services.CalendarService) {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/create", service.Save).Methods("POST")
	rtr.HandleFunc("/update", service.Save).Methods("POST")
	rtr.HandleFunc("/delete", service.DeleteById).Methods("POST")
	rtr.HandleFunc("/get_day", service.GetByDay).Methods("GET")
	rtr.HandleFunc("/get_week", service.GetByWeek).Methods("GET")
	rtr.HandleFunc("/get_month", service.GetByMonth).Methods("GET")

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func main() {
	cal := calendar.New()
	service := services.New(cal)

	handleRequest(service)
}
