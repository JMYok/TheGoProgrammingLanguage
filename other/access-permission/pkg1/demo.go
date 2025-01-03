package pkg1

import "fmt"

type A struct {
	private int
	Public  int
}

func NewA() *A {
	return &A{private: 11}
}

func (a *A) OutPrivate() {
	fmt.Println(a.private)
}
