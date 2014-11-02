package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/k0kubun/sandal/lang"
	"io/ioutil"
	"log"
	"os"
)

type Options struct {
	Ast bool `short:"a" long:"ast" default:"false" description:"dump parsed ast"`
}

func run(filePath string, options *Options) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(filePath, err)
	}

	if options.Ast {
		lang.DumpAST(string(body))
		return
	}

	compiled, err := lang.CompileFile(string(body))
	if err != nil {
		log.Fatal(filePath, err)
	}
	fmt.Print(compiled)
}

func main() {
	options := new(Options)
	args, err := flags.Parse(options)
	if err != nil {
		return
	}

	if len(args) != 1 {
		fmt.Println("Usage: sandal [programfile]")
		os.Exit(1)
	}

	run(args[0], options)
}
