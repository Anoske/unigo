
package mypackage

import "fmt"


func Greet(name string) string {
    return fmt.Sprintf("Привет, %s!", name)
}


func Add(a, b int) int {
    return a + b
}

type Number struct {
    Value int
}

func (n *Number) Double() int {
    return n.Value * 2
}

func (n *Number) Add(value int) int {
    n.Value += value
    return n.Value
}