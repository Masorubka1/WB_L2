package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Паттерн "Команда" (Command) является поведенческим паттерном проектирования, который позволяет инкапсулировать запросы в отдельные объекты, делая их полноценными объектами первого класса. Он позволяет параметризовать клиентов с использованием различных запросов, ставить их в очередь или регистрировать их для отмены и повторного выполнения.

Применимость паттерна "Команда":

Когда необходимо параметризовать объекты с использованием операций и запросов.
Когда нужно обеспечить выполнение операций в разное время, ставить их в очередь, а также поддерживать отмену и повторное выполнение.
Когда требуется создавать системы, способные логировать и восстанавливать операции.
Когда необходимо унифицировать интерфейс для выполнения различных операций, действий или команд.

Плюсы использования паттерна "Команда":

Разделение клиента, который инициирует выполнение операций, от получателя, который фактически выполняет операции.
Упрощение добавления новых операций без изменения существующего кода.
Возможность создания очереди операций и выполнения их в разное время.
Поддержка отмены и повторного выполнения операций.
Унифицированный интерфейс для выполнения различных команд.

Минусы использования паттерна "Команда":

Увеличение сложности кода из-за создания дополнительных классов и интерфейсов.
Возможное усложнение архитектуры системы из-за введения множества команд и их взаимодействия.
Дополнительные затраты на память для хранения объектов команд.

Реальные примеры использования паттерна "Команда":

Реализация командной строки или интерфейса команд в системе операционной системы.
Реализация системы "Отмена" и "Повтор" в текстовых редакторах или графических приложениях.
Реализация очереди задач или обработки сообщений в системах планирования и управления.
Реализация системы управления устройствами, где команды могут быть представлены в виде объектов и выполняться по запросу.
*/
import "fmt"

// Command - Интерфейс команды
type Command interface {
	Execute()
}

// Receiver - Получатель, который выполняет операции
type Receiver struct {
	name string
}

func NewReceiver(name string) *Receiver {
	return &Receiver{name: name}
}

func (r *Receiver) Action() {
	fmt.Printf("Выполнение операции для получателя %s\n", r.name)
}

// ConcreteCommand - Конкретная команда
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Invoker - Инициатор, который выполняет команды
type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	receiver := NewReceiver("Получатель")

	command1 := NewConcreteCommand(receiver)
	command2 := NewConcreteCommand(receiver)
	command3 := NewConcreteCommand(receiver)

	invoker := Invoker{}
	invoker.AddCommand(command1)
	invoker.AddCommand(command2)
	invoker.AddCommand(command3)

	invoker.ExecuteCommands()
}
