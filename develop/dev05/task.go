package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/
func main() {
	var after, before, context, count, ignoreCase, invert, fixed, lineNum bool
	var numLines int
	var pattern string

	flag.BoolVar(&after, "A", false, "Print N lines after the match")
	flag.BoolVar(&before, "B", false, "Print N lines before the match")
	flag.BoolVar(&context, "C", false, "Print N lines before and after the match")
	flag.BoolVar(&count, "c", false, "Print only the count of matching lines")
	flag.BoolVar(&ignoreCase, "i", false, "Ignore case when matching")
	flag.BoolVar(&invert, "v", false, "Invert the match (print non-matching lines)")
	flag.BoolVar(&fixed, "F", false, "Match fixed string instead of pattern")
	flag.BoolVar(&lineNum, "n", false, "Print line numbers")

	flag.IntVar(&numLines, "num", 0, "Number of lines to print")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Usage: grep [OPTIONS] PATTERN FILE")
		os.Exit(1)
	}

	pattern = args[0]
	filename := args[1]

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(fileData), "\n")

	var matches []string

	for i, line := range lines {
		if ignoreCase {
			pattern = "(?i)" + pattern
		}
		if fixed {
			match := strings.Contains(line, pattern)
			if (invert && !match) || (!invert && match) {
				matches = append(matches, line)
			}
		} else {
			re := regexp.MustCompile(pattern)
			match := re.MatchString(line)
			if (invert && !match) || (!invert && match) {
				matches = append(matches, line)
			}
		}
		if count && len(matches) > 0 {
			break
		}
		if numLines > 0 && len(matches) >= numLines {
			break
		}
		if before && len(matches) == 0 {
			if i < numLines {
				matches = append(matches, lines[:i]...)
			} else {
				matches = append(matches, lines[i-numLines:i]...)
			}
		}
		if after && len(matches) > 0 {
			if len(lines)-i-1 < numLines {
				matches = append(matches, lines[i+1:]...)
			} else {
				matches = append(matches, lines[i+1:i+numLines+1]...)
			}
		}
		if context && len(matches) > 0 {
			start := i - numLines
			if start < 0 {
				start = 0
			}
			end := i + numLines
			if end > len(lines)-1 {
				end = len(lines) - 1
			}
			matches = append(matches, lines[start:i]...)
			matches = append(matches, lines[i:i+1]...)
			matches = append(matches, lines[i+1:end+1]...)
		}
	}

	if count {
		fmt.Println(len(matches))
	} else if lineNum {
		for i, line := range matches {
			fmt.Printf("%d:%s\n", i+1, line)
		}
	} else {
		for _, line := range matches {
			fmt.Println(line)
		}
	}
}
