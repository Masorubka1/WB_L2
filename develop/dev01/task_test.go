package main

import (
	"os"
	"testing"
	"time"

	"github.com/beevik/ntp"
)

func TestTime(t *testing.T) {
	// Получаем текущее время с помощью time.Now()
	currentTime := time.Now()

	// Получаем точное время с помощью ntp.Time()
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		t.Errorf("Error getting NTP time: %v\n", err)
	}

	// Проверяем, что полученное время не пустое
	if ntpTime.IsZero() {
		t.Errorf("NTP time is zero")
	}

	// Проверяем, что точное время больше текущего времени
	if ntpTime.Before(currentTime) {
		t.Errorf("NTP time is before current time")
	}
}

func TestMain(m *testing.M) {
	// Выполняем тесты
	exitCode := m.Run()

	// Завершаем программу с кодом выхода, соответствующим результатам тестов
	os.Exit(exitCode)
}
