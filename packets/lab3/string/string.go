package mypackage

import "strings"

type StringOperations struct{}

func (s StringOperations) Concatenate(a, b string) string {
    return a + b
}

func (s StringOperations) Length(str string) int {
    return len(str)
}

func (s StringOperations) ToUpper(str string) string {
    return strings.ToUpper(str)
}

func (s StringOperations) ToLower(str string) string {
    return strings.ToLower(str)
}
