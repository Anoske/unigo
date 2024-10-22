package main

import "fmt"

func main() {
    ptr := new(float64)
    *ptr = 3.14
    fmt.Println("Значение, хранящееся по указателю:", *ptr)
}