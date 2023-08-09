package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов

*/

type Line struct {
	Text string
	Key  string
}

type ByKey []Line

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func main() {
	filePath := flag.String("f", "", "path to input file")
	column := flag.Int("k", 0, "column to sort by (1-based)")
	numeric := flag.Bool("n", false, "sort numerically")
	reverse := flag.Bool("r", false, "reverse sort order")
	unique := flag.Bool("u", false, "output only unique lines")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a path to input file with -f flag")
		return
	}

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []Line
	for scanner.Scan() {
		text := scanner.Text()
		key := text
		if *column > 0 {
			fields := strings.Fields(text)
			if *column > len(fields) {
				fmt.Printf("Column %d is out of range for line: %s\n", *column, text)
				return
			}
			key = fields[*column-1]
		}
		if *numeric {
			num, err := strconv.Atoi(key)
			if err != nil {
				fmt.Printf("Line is not numeric: %s\n", text)
				return
			}
			key = fmt.Sprintf("%d", num)
		}
		lines = append(lines, Line{Text: text, Key: key})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if *unique {
		uniqueLines := make(map[string]bool)
		for _, line := range lines {
			if !uniqueLines[line.Text] {
				uniqueLines[line.Text] = true
			}
		}
		lines = nil
		for line := range uniqueLines {
			lines = append(lines, Line{Text: line})
		}
	}

	if *reverse {
		sort.Sort(sort.Reverse(ByKey(lines)))
	} else {
		sort.Sort(ByKey(lines))
	}

	output := ""
	for _, line := range lines {
		output += line.Text + "\n"
	}

	err = ioutil.WriteFile(*filePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File sorted successfully")
}
