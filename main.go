package main

import (
	"fmt"
)

// Функция для проверки строки на уникальность символов
func isUnique(input string) bool {
	word_already_exist := make(map[rune]bool)

	for _, char := range input {
		if word_already_exist[char] {
			return false
		}
		word_already_exist[char] = true
	}
	return true
}

func main() {
	//Тесты
	tests := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, test := range tests {
		result := isUnique(test)
		fmt.Printf("%s -> %t\n", test, result)
	}
}
