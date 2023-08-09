package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Паттерн «Цепочка вызовов» (Chain of Responsibility) является поведенческим паттерном проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков, пока один из них не сможет обработать запрос. Он позволяет избежать явного связывания отправителя запроса с его получателем и дает возможность динамически устанавливать порядок обработки запроса.

Применимость паттерна «Цепочка вызовов»:

Когда есть несколько объектов, способных обработать один и тот же запрос, и нужно выбрать один из них динамически.
Когда набор объектов и порядок их обработки могут изменяться динамически.
Когда нужно обеспечить гибкую возможность добавления новых обработчиков без изменения клиентского кода.

Плюсы использования паттерна «Цепочка вызовов»:

Разделение отправителя запроса и его получателей.
Увеличение гибкости и расширяемости кода.
Уменьшение зависимости между отправителем и получателем запроса.
Упрощение добавления новых обработчиков без изменения существующего кода.

Минусы использования паттерна «Цепочка вызовов»:

Возможность неполной обработки запроса, если не задан соответствующий обработчик.
Возможное усложнение процесса отладки и понимания потока выполнения запроса из-за его динамической природы.

Реальные примеры использования паттерна «Цепочка вызовов»:

Обработка запросов веб-приложения, где различные обработчики могут выполнять различные действия, например, аутентификацию, авторизацию, валидацию и т.д.
Обработка событий в графическом интерфейсе, где различные элементы интерфейса могут быть обработчиками событий, и события могут передаваться по цепочке для выполнения соответствующих действий.
Логирование и аудит в системах, где различные обработчики могут выполнять логирование событий в различных форматах и сохранять информацию в различных источниках.
*/

import "fmt"

// Handler - Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// ConcreteHandlerA - Конкретный обработчик A
type ConcreteHandlerA struct {
	nextHandler Handler
}

func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *ConcreteHandlerA) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("Обработчик A обрабатывает запрос")
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	} else {
		fmt.Println("Ни один из обработчиков не может обработать запрос")
	}
}

// ConcreteHandlerB - Конкретный обработчик B
type ConcreteHandlerB struct {
	nextHandler Handler
}

func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *ConcreteHandlerB) HandleRequest(request string) {
	if request == "B" {
		fmt.Println("Обработчик B обрабатывает запрос")
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	} else {
		fmt.Println("Ни один из обработчиков не может обработать запрос")
	}
}

// ConcreteHandlerC - Конкретный обработчик C
type ConcreteHandlerC struct {
	nextHandler Handler
}

func (h *ConcreteHandlerC) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *ConcreteHandlerC) HandleRequest(request string) {
	if request == "C" {
		fmt.Println("Обработчик C обрабатывает запрос")
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	} else {
		fmt.Println("Ни один из обработчиков не может обработать запрос")
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	handlerA.SetNext(handlerB)
	handlerB.SetNext(handlerC)

	handlerA.HandleRequest("B")
	handlerA.HandleRequest("C")
	handlerA.HandleRequest("D")
}
