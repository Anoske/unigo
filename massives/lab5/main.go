package main

import "fmt"

func main() {

    scores := map[string]int{
        "Alice": 85,
        "Bob":   92,
    }

    name := "Alice"

    if score, found := scores[name]; found {

        fmt.Println(name, "найдена с баллом", score)
        scores[name] = 90 
    } else {

        fmt.Println(name, "не найдена, добавляю запись")
        scores[name] = 75
    }

    fmt.Println("Обновлённые результаты:", scores)
}
