package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func updatePerson(p *Person) {
    p.Name = "Alice"
    p.Age = 30
}

func main() {
    person := Person{
        Name: "Bob",
        Age:  25,
    }

    fmt.Println("Старые значения:")
    fmt.Println("Имя:", person.Name)
    fmt.Println("Возраст:", person.Age)

    personPtr := &person

    updatePerson(personPtr)

    fmt.Println("Новые значения:")
    fmt.Println("Имя:", person.Name)
    fmt.Println("Возраст:", person.Age)
}