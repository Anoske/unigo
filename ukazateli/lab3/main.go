package main

import "fmt"

func swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x := 5
    y := 10

    fmt.Println("Старые значения:")
    fmt.Println("x:", x)
    fmt.Println("y:", y)

    swap(&x, &y)


    fmt.Println("Новые значения:")
    fmt.Println("x:", x)
    fmt.Println("y:", y)
}
