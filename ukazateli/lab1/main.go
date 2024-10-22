package main

import "fmt"

func main() {
    var num int = 42

    ptr := &num

    fmt.Println("Старое значение:", num)

    *ptr = 100

    fmt.Println("Новое значение:", num)
}
