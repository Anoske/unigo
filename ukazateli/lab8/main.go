package main

import "fmt"

func modifySlice(s *[]int) {
    for i := range *s {
        (*s)[i] *= 2 
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Старые значения среза:", numbers)
    modifySlice(&numbers)
    fmt.Println("Новые значения среза:", numbers)
}