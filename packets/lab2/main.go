package main

import (
    "fmt"
    "lab2/mathopp" 
    "lab2/stringopp"
)

func main() {

    sum := mathoperations.Add(10, 5)
    fmt.Println("Сумма:", sum)

    difference := mathoperations.Subtract(10, 5)
    fmt.Println("Разность:", difference)

    product := mathoperations.Multiply(10, 5)
    fmt.Println("Произведение:", product)

    quotient := mathoperations.Divide(10, 5)
    fmt.Println("Частное:", quotient)


    concatenated := stringoperations.Concatenate("Hello, ", "World!")
    fmt.Println("Объединенная строка:", concatenated)

    length := stringoperations.Length(concatenated)
    fmt.Println("Длина объединенной строки:", length)

    upper := stringoperations.ToUpper(concatenated)
    fmt.Println("Верхний регистр:", upper)

    lower := stringoperations.ToLower(concatenated)
    fmt.Println("Нижний регистр:", lower)
}