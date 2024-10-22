package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Введите строку:")
    var text string
    fmt.Scanln(&text)
    words := strings.Fields(text)

    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }

    fmt.Println("Частота слов:", wordCount)
}