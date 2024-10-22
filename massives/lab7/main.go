package main

import "fmt"

type Student struct {
    Name string
    Age  int
}

func main() {
    classes := map[string][]Student{
        "Class A": {
            {"Alice", 15},
            {"Bob", 16},
        },
        "Class B": {
            {"Carol", 17},
            {"Dave", 16},
        },
    }

    for class, students := range classes {
        fmt.Printf("Класс %s:\n", class)
        for _, student := range students {
            fmt.Printf("- Имя: %s, Возраст: %d\n", student.Name, student.Age)
        }
    }
}
