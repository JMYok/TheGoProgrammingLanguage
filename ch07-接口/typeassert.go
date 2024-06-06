package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)

	//var ErrNotExist = errors.New("file does not exist")

}

func IsNotExist(err error) bool {
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == os.ErrNotExist
}
