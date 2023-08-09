package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Паттерн «Посетитель» (Visitor) является поведенческим паттерном проектирования, который позволяет добавлять новые операции к объектам без изменения их классов. Он предоставляет возможность разделить алгоритмы от объектной структуры, в которой эти алгоритмы выполняются.

Применимость паттерна «Посетитель»:

Когда необходимо выполнить операции над набором объектов с разными классами без изменения их структуры.
Когда новые операции могут появляться в будущем, и необходимо обеспечить легкое добавление этих операций без модификации существующих классов.
Когда структура объектов стабильна, но операции над ними часто меняются.

Плюсы использования паттерна «Посетитель»:

Позволяет добавлять новые операции без изменения классов объектов.
Упрощает добавление и модификацию операций.
Разделяет алгоритмы и структуры объектов, повышая гибкость и удобство поддержки кода.
Упрощает процесс добавления операций для разных классов объектов.

Минусы использования паттерна «Посетитель»:

Может привести к нарушению инкапсуляции объектов, так как операции могут иметь доступ к приватным членам объектов.
Введение новых операций требует модификации интерфейса посетителя и всех классов, которые его реализуют.
Паттерн может усложнить код и увеличить его объемность.

Реальные примеры использования паттерна «Посетитель»:

Обработка деревьев разбора (парсеров) в компиляторах или интерпретаторах. Посетитель может быть использован для обхода дерева разбора и выполнения различных операций над его узлами.
Обработка элементов графического интерфейса. Посетитель может применяться для обхода и взаимодействия с различными элементами интерфейса, например, для реализации функциональности рисования, обновления или сохранения состояния.
Генерация отчетов или обработка структур данных. Посетитель может быть использован для обхода и выполнения различных операций над структурами данных, например, для генерации отчетов на основе данных из базы данных или для преобразования структуры данных.
*/

import "fmt"

// Element - Интерфейс элемента, которые могут принимать посетителя
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA - Конкретный элемент A
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// ConcreteElementB - Конкретный элемент B
type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Visitor - Интерфейс посетителя
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor - Конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Посетитель посещает элемент A")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Посетитель посещает элемент B")
}

// ObjectStructure - Объектная структура, содержащая элементы
type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) Attach(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Detach(element Element) {
	for i, el := range os.elements {
		if el == element {
			os.elements = append(os.elements[:i], os.elements[i+1:]...)
			break
		}
	}
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range os.elements {
		element.Accept(visitor)
	}
}

func main() {
	objectStructure := &ObjectStructure{}

	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}

	objectStructure.Attach(elementA)
	objectStructure.Attach(elementB)

	visitor := &ConcreteVisitor{}
	objectStructure.Accept(visitor)
}
