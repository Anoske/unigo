package main

import "fmt"

func main() {
    // Создаем срез строк
    originalSlice := []string{"apple", "banana", "cherry", "date", "elderberry"}

    // Создаем новый срез для копирования элементов
    copiedSlice := make([]string, len(originalSlice))

    // Копируем элементы с помощью функции copy
    copy(copiedSlice, originalSlice)

    // Выводим оригинальный и скопированный срез
    fmt.Println("Оригинальный срез:", originalSlice)
    fmt.Println("Скопированный срез:", copiedSlice)

    // Удаляем элемент "cherry" из оригинального среза
    elementToRemove := 2 // индекс элемента "cherry"
    originalSlice = append(originalSlice[:elementToRemove], originalSlice[elementToRemove+1:]...)

    // Выводим срез после удаления элемента
    fmt.Println("Срез после удаления элемента:", originalSlice)
}
