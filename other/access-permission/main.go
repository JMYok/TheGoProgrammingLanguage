package main

import (
	"jmy/goprogramminglanguage/other/access-permission/interf"
	"jmy/goprogramminglanguage/other/access-permission/pkg1"
	"jmy/goprogramminglanguage/other/access-permission/pkg2"
)

func main() {
	A := pkg1.NewA()
	A.Public = 111
	B := pkg2.NewB()
	B.Public = 222

	printBInfo(A)
	printBInfo(B)
}

func printBInfo(out interf.Out) {
	out.OutPrivate()
}
