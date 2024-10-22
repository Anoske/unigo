package main

import "fmt"

func main() {

    originalSlice := []string{"apple", "banana", "cherry", "date", "elderberry"}

    copiedSlice := make([]string, len(originalSlice))

    copy(copiedSlice, originalSlice)

    fmt.Println("Оригинальный срез:", originalSlice)
    fmt.Println("Скопированный срез:", copiedSlice)

    elementToRemove := 2 // индекс элемента "cherry"
    originalSlice = append(originalSlice[:elementToRemove], originalSlice[elementToRemove+1:]...)

    fmt.Println("Срез после удаления элемента:", originalSlice)
}
