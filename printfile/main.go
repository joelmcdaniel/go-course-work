package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(b))

	/* f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f) */
}
