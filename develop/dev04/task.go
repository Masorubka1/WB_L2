package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

func findAnagrams(words *[]string) *map[string][]string {
	anagramMap := make(map[string][]string)
	for _, word := range *words {
		// Convert word to lowercase and sort its characters
		sortedWord := sortString(strings.ToLower(word))
		// Add word to anagram map
		if _, ok := anagramMap[sortedWord]; ok {
			// Anagram already exists, add word to its array
			anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
		} else {
			// New anagram, create new array with current word
			anagramMap[sortedWord] = []string{word}
		}
	}
	// Remove single-word anagram sets
	for key, value := range anagramMap {
		if len(value) < 2 {
			delete(anagramMap, key)
		} else {
			// Sort words within each anagram set
			sort.Strings(value)
			anagramMap[key] = value
		}
	}
	return &anagramMap
}

// Helper function to sort a string's characters
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
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

	anagrams := findAnagrams(&dictionary)

	for key, value := range *anagrams {
		fmt.Printf("Key: %s\nValue: %v\n\n", key, value)
	}
}
