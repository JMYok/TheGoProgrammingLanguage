package string_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
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
	str1 := "Hello Go2"
	str2 := "Hello Go2"
	sh1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	sh2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	//stringStruct地址
	fmt.Println(&str1, &str2)
	// 输出 0xc00002c420 0xc00002c430

	// 底层数组地址
	fmt.Println(sh1.Data, sh2.Data)
	fmt.Println(sh1.Data == sh2.Data)
	// 输出 10990392 10990392
	// 输出 true
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
