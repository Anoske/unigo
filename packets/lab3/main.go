package main

import (
    "fmt"
    "lab3/string"
    "lab3/math" 
)

func main() {

    mathOps := mypackage.MathOperations{}
    stringOps := mypackage.StringOperations{}

    sum := mathOps.Add(10, 5)
    fmt.Println("Сумма:", sum)

    difference := mathOps.Subtract(10, 5)
    fmt.Println("Разность:", difference)

    product := mathOps.Multiply(10, 5)
    fmt.Println("Произведение:", product)

    quotient := mathOps.Divide(10, 5)
    fmt.Println("Частное:", quotient)


    concatenated := stringOps.Concatenate("Hello, ", "World!")
    fmt.Println("Объединенная строка:", concatenated)

    length := stringOps.Length(concatenated)
    fmt.Println("Длина объединенной строки:", length)

    upper := stringOps.ToUpper(concatenated)
    fmt.Println("Верхний регистр:", upper)

    lower := stringOps.ToLower(concatenated)
    fmt.Println("Нижний регистр:", lower)
}
