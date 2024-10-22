package main

import "fmt"

func increaseByTen(num *int) {
    *num += 10
}

func main() {
    value := 20
    fmt.Println("Старое значение:", value)

    increaseByTen(&value)

    fmt.Println("Новое значение:", value)
}