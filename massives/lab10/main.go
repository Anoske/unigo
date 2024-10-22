package main

import "fmt"


func removeDuplicates(elements []string) []string {

    encountered := map[string]bool{}
    result := []string{} 

    for _, v := range elements {
        if !encountered[v] {
            encountered[v] = true 
            result = append(result, v) 
        }
    }
    return result 
}

func main() {
    data := []string{"apple", "banana", "apple", "orange", "banana"}
    fmt.Println("Срез с дубликатами:", data)

    unique := removeDuplicates(data)
    fmt.Println("Срез без дубликатов:", unique)
}