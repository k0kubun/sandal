package lang

import (
	"github.com/k0kubun/pp"
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

// -- debug functions --

func DumpAST(body string) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)

	pp.Println(defs)
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

func DumpGraph(body string) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)

	if err := typecheck.TypeCheck(defs); err != nil {
		log.Fatal("TypeCheck error:", err)
	}

	conversion.DumpGraph(defs)
}
