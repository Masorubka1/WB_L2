package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

*/

func UnpackString(s string) (string, error) {
	var result strings.Builder
	var prev rune
	escape := false

	for i, char := range s {
		if escape {
			result.WriteRune(char)
			escape = false
			continue
		}
		if char == '\\' {
			escape = true
			continue
		}
		if i == 0 && char >= '0' && char <= '9' {
			return "", errors.New("invalid string")
		}
		if i == 0 || (char < '0' || char > '9') {
			result.WriteRune(char)
			prev = char
			continue
		}

		repeat, err := strconv.Atoi(string(char))
		if err != nil {
			return "", errors.New("invalid string")
		}
		if prev == '\\' {
			result.WriteString(strings.Repeat(string(prev), repeat-1))
		} else {
			result.WriteString(strings.Repeat(string(prev), repeat))
		}
	}

	return result.String(), nil
}

func main() {
	s1, _ := UnpackString("a4bc2d5e")
	s2, _ := UnpackString("abcd")
	_, err := UnpackString("45")
	s4, _ := UnpackString("")
	fmt.Printf(s1 + "\n")
	fmt.Printf(s2 + "\n")
	fmt.Printf(error.Error(err) + "\n")
	fmt.Printf(s4 + "\n")
}
