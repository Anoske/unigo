package main

import "fmt"

func modifyArray(arr *[5]int) {
    for i := range arr {
        arr[i] += 10 
    }
}

func main() {
    numbers := [5]int{1, 2, 3, 4, 5}

    fmt.Println("Старые значения массива:", numbers)

    modifyArray(&numbers)

    fmt.Println("Новые значения массива:", numbers)
}