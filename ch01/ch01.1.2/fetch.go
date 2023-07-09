package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		/*
			judge := url.HasPrefix("http://")
			if judge == false {

			}
		*/
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		errCode, _ := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		fmt.Fprintf(os.Stderr, "error code is : %d\n", errCode)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
