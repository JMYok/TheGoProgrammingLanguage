package string_test

import (
	"fmt"
	"testing"
)

func TestLowlevelString(t *testing.T) {
	str := "阿煜Go"
	t.Logf("The string: %q\n", str)
	t.Logf("  => runes(char): %q\n", []rune(str))
	t.Logf("  => runes(hex): %x\n", []rune(str))
	t.Logf("  => bytes(hex): [% x]\n", []byte(str))
}

func TestPrintString(t *testing.T) {
	str := "Hello Go"
	fmt.Printf("%c\n", str[0])
	str = "阿煜Go"
	fmt.Printf("%c\n", str[0])
	runeStr := []rune(str)
	fmt.Printf("%c\n", runeStr[1])
}

func TestStringFor(t *testing.T) {
	str := "阿煜Go"
	fmt.Printf("%q\n", str[0])
	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}
}
