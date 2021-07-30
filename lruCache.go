package main

import (
	"fmt"
	"strings"
)

func findPos(input []string, symb string) int {
	for i, s := range input {
		if s == symb {
			return i
		}
	}

	return -1
}

func deleteSymb(input []string, symb string) []string {
	pos := findPos(input, symb)

	output := make([]string, cap(input))
	output = append(output, input[:pos]...)
	output = append(output, input[pos+1:]...)

	return output
}

func cacheLRU(input []string) string {
	lruMap := make(map[string]struct{})

	var result []string

	for _, symb := range input {
		_, exists := lruMap[symb]
		//  Если не сущетсвует, добавляем
		if !exists {
			lruMap[symb] = struct{}{}
			result = append(result, symb)
			continue
		}
		// Если существует, удаляем из слайса и добавляем в конец
		result = deleteSymb(result, symb)
		result = append(result, symb)
	}
	res := strings.Join(result, "")

	return res
}

func main() {
	//A-B-A-C-A -> B-C-A
	cacheInput := []string{"A", "B", "A", "C", "A"}
	cacheOutput := cacheLRU(cacheInput)

	fmt.Println(cacheOutput)
}
