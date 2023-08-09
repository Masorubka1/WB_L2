package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS

*/

func main() {
	// Получаем текущее время с помощью time.Now()
	currentTime := time.Now()

	// Получаем точное время с помощью ntp.Time()
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting NTP time: %v\n", err)
		os.Exit(1)
	}

	// Выводим текущее и точное время на экран
	fmt.Printf("Current time: %s\n", currentTime)
	fmt.Printf("NTP time: %s\n", ntpTime)
}
