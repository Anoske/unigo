package main

import "fmt"

func main() {
    a := 10
    b := 20

    ptrA := &a
    ptrB := &b

    fmt.Println("Указатель ptrA указывает на:", *ptrA)
    fmt.Println("Указатель ptrB указывает на:", *ptrB)
    
    if ptrA == ptrB {
        fmt.Println("Указатели равны.")
    } else {
        fmt.Println("Указатели не равны.")
    }

    fmt.Printf("Адрес ptrA: %p\n", ptrA)
    fmt.Printf("Адрес ptrB: %p\n", ptrB)
}