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
	Ir1 bool `short:"1" long:"ir1" default:"false" description:"dump IR1"`
	Ir2 bool `short:"2" long:"ir2" default:"false" description:"dump IR2"`
	Graph bool `short:"g" long:"graph" default:"false" description:"dump state transition to dot lang"`
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

	if options.Ir1 {
		lang.DumpIR1(string(body))
		return
	}

	if options.Ir2 {
		lang.DumpIR2(string(body))
		return
	}

	if options.Graph {
		lang.DumpGraph(string(body))
		return
	}

	compiled, err := lang.CompileFile(string(body))
	if err != nil {
		log.Fatalf("%s: %s", filePath, err.Error())
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
