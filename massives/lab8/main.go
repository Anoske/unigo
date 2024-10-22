package main

import (
    "fmt"
    "sort"
)

func main() {
    numbers := []int{10, 3, 5, 1, 9, 7}
    fmt.Println("Изначальный срез:", numbers)
    sort.Ints(numbers)
    fmt.Println("Отсортированный срез:", numbers)
}
