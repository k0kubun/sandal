package lang

import (
	"fmt"
	"github.com/k0kubun/pretty"
	"github.com/k0kubun/sandal/lang/conversion"
	"github.com/k0kubun/sandal/lang/parsing"
	"github.com/k0kubun/sandal/lang/typecheck"
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
