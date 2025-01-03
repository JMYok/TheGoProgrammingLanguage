package pkg2

import (
	"fmt"
	"jmy/goprogramminglanguage/other/access-permission/pkg1"
)

type B struct {
	pkg1.A
	private int
	Public  int
}

func NewB() *B {
	return &B{private: 22}
}

func (b *B) OutPrivate() {
	fmt.Println(b.private)
}
