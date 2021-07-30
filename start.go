package main

import "fmt"

func main() {
	//A-B-A-C-A -> B-C-A
	cacheInput := []string{"A", "B", "A", "C", "A"}
	cacheOutput := cacheLRU(cacheInput)
	fmt.Println(cacheOutput)

	//35132 -> 3-5+1+3-2
	inp := 35132
	res := PlusMinus(inp)
	fmt.Printf(res)
}
