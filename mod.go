package go_mod_test

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
