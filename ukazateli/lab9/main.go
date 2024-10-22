package main

import "fmt"

func addPeople(ages *map[string]int) {
    (*ages)["Alice"] = 30
    (*ages)["Bob"] = 25
    (*ages)["Charlie"] = 35
}

func main() {
    ages := make(map[string]int)

    fmt.Println("Старая карта:", ages)

    addPeople(&ages)

    fmt.Println("Обновленная карта:", ages)
}