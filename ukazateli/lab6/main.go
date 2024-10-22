package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func modifyPerson(p *Person) {
    p.Name = "Charlie"
    p.Age = 28
}

func main() {
    person := Person{
        Name: "Alice",
        Age:  30,
    }


    fmt.Println("Старые значения:")
    fmt.Println("Имя:", person.Name)
    fmt.Println("Возраст:", person.Age)

    modifyPerson(&person)


    fmt.Println("Новые значения:")
    fmt.Println("Имя:", person.Name)
    fmt.Println("Возраст:", person.Age)
}