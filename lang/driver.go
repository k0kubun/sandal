package lang

import (
	"fmt"
	"github.com/k0kubun/pretty"
	"github.com/k0kubun/sandal/lang/conversion"
	"github.com/k0kubun/sandal/lang/parsing"
	"github.com/k0kubun/sandal/lang/typecheck"
	"log"
)

func CompileFile(body string) (string, error) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)

	if err := typecheck.TypeCheck(defs); err != nil {
		return "", err
	}

	return conversion.ConvertASTToNuSMV(defs)
}

func DumpAST(body string) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)

	fmt.Printf("%# v\n", pretty.Formatter(defs))
}

func DumpIR1(body string) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)

	if err := typecheck.TypeCheck(defs); err != nil {
		log.Fatal("TypeCheck error:", err)
	}

	conversion.DumpIR1(defs)
}

func DumpIR2(body string) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)

	if err := typecheck.TypeCheck(defs); err != nil {
		log.Fatal("TypeCheck error:", err)
	}

	conversion.DumpIR2(defs)
}
