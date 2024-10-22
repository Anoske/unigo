package main

import "fmt"

func calculateAverage(scores map[string]int) float64 {
    total := 0
    for _, score := range scores {
        total += score
    }
    return float64(total) / float64(len(scores))
}

func main() {

    studentScores := map[string]int{
        "Alice":  85,
        "Bob":    90,
        "Charlie": 78,
        "David":  92,
        "John":    88,
    }


    fmt.Println("Оригинальная карта:", studentScores)

    average := calculateAverage(studentScores)
    fmt.Printf("Средний балл: %.2f\n", average)


    delete(studentScores, "Charlie")

    fmt.Println("Карта после удаления студента:", studentScores)

    newAverage := calculateAverage(studentScores)
    fmt.Printf("Новый средний балл: %.2f\n", newAverage)
}
