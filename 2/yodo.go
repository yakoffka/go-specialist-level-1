package main

import "fmt"

func main() {
	word1, word2, word3, word4 := "имя", "твоё", "мне", "знакомо"
	fmt.Println(word4, word3, word2, word1)
	fmt.Println(word3, word1, word4, word2)
	fmt.Println(word2, word4, word1, word3)
}
