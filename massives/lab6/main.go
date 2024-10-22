package main

import "fmt"

func main() {
    grades := map[string][]int{
        "Math":     {90, 85, 88},
        "Physics":  {78, 82, 91},
        "Chemistry": {89, 95, 80},
    }
    for subject, scores := range grades {
        total := 0
        for _, score := range scores {
            total += score
        }
        avg := float64(total) / float64(len(scores))
        fmt.Printf("%s: Средний балл: %.2f\n", subject, avg)
    }
}