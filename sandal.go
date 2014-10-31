package main

import (
	"fmt"
	"github.com/k0kubun/sandal/lang"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: sandal [programfile]")
		os.Exit(1)
	}

	filePath := os.Args[1]
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(filePath, err)
	}

	compiled, err := lang.CompileFile(string(body))
	if err != nil {
		log.Fatal(filePath, err)
	}
	fmt.Print(compiled)
}
