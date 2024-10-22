package main

import (
    "fmt"
    "C:\Users\puk1\Documents\Уник Хуюник\packets\lab1\mypackage"
)

func main() {
    greeting := mypackage.Greet("Иван")
    fmt.Println(greeting)

    sum := mypackage.Add(3, 5)
    fmt.Println("Сумма:", sum)


    num := mypackage.Number{Value: 10}
    fmt.Println("Исходное значение:", num.Value)

    doubled := num.Double()
    fmt.Println("Удвоенное значение:", doubled)

    newValue := num.Add(5)
    fmt.Println("Новое значение после добавления:", newValue)
}
