package mypackage

type MathOperations struct{}

func (m MathOperations) Add(a, b int) int {
    return a + b
}

func (m MathOperations) Subtract(a, b int) int {
    return a - b
}

func (m MathOperations) Multiply(a, b int) int {
    return a * b
}


func (m MathOperations) Divide(a, b int) float64 {
    if b == 0 {
        return 0 // Возвращаем 0, если деление на 0
    }
    return float64(a) / float64(b)
}