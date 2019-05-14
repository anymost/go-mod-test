package mod

import "fmt"

const (
	A = 1
	b = 2
)

func sayA() {
	fmt.Println(A)
}

func SayB() {
	fmt.Println(b)
}

type User struct {
	Name string
	age  int
}
