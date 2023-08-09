package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	dictionary := []string{
		"Пятак",
		"Пятак",
		"пятка",
		"Тяпка",
		"слиток",
		"слиток",
		"столик",
		"листок",
		"Топот",
		"Потоп",
	}

	expected := map[string][]string{
		"акптя":  {"Пятак", "Пятак", "пятка", "Тяпка"},
		"илкост": {"листок", "слиток", "слиток", "столик"},
		"опотт":  {"Потоп", "Топот"},
	}

	result := findAnagrams(&dictionary)

	if !reflect.DeepEqual(*result, expected) {
		t.Errorf("Expected %v, but got %v", expected, *result)
	}
}

func TestSortString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "ehllo"},
		{"racecar", "aaccerr"},
		{"abbcccddddeeeee", "abbcccddddeeeee"},
	}

	for _, testCase := range testCases {
		result := sortString(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %s, but got %s", testCase.expected, result)
		}
	}
}
