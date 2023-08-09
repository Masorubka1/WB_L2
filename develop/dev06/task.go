package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/
func main() {
	fields := flag.String("f", "", "Select fields")
	delimiter := flag.String("d", "\t", "Set delimiter")
	separated := flag.Bool("s", false, "Only separated lines")
	flag.Parse()

	selectedFields := make(map[int]bool)
	for _, f := range strings.Split(*fields, ",") {
		if f != "" {
			fieldNum := parseInt(f)
			selectedFields[fieldNum] = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}
		columns := strings.Split(line, *delimiter)
		output := make([]string, 0)
		for i, col := range columns {
			if selectedFields[i+1] {
				output = append(output, col)
			}
		}
		fmt.Println(strings.Join(output, *delimiter))
	}
}

func parseInt(s string) int {
	var result int
	for _, c := range s {
		if c < '0' || c > '9' {
			panic(fmt.Sprintf("Invalid integer: %s", s))
		}
		result = result*10 + int(c-'0')
	}
	return result
}
