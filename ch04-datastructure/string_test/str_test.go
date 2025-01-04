package string_test

import (
	"bytes"
	"fmt"
	"strings"
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

func TestStringPool(t *testing.T) {
	str1 := "Hello Go"
	str2 := "Hello Go"
	fmt.Println(&str1, &str2)
}

func TestStringModify(t *testing.T) {
	str := "Hello Go"
	bytes := []byte(str)
	bytes[0] = 'X'
	str = string(bytes)
	fmt.Println(str)
}

func TestStringAdd(t *testing.T) {
	str1 := "Hello"
	str2 := "Go"
	str := str1 + str2
	fmt.Println(str)

	str = fmt.Sprintf("%s%s", str1, str2)
	fmt.Println(str)

	var sb strings.Builder
	sb.WriteString(str1)
	sb.WriteString(str2)
	str = sb.String()
	fmt.Println(str)

	strs := []string{str1, str2}
	str = strings.Join(strs, "")
	fmt.Println(str)

	var buffer bytes.Buffer
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	str = buffer.String()
	fmt.Println(str)
}
