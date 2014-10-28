package lang

import (
	"github.com/k0kubun/santalum/lang/parsing"
	"github.com/k0kubun/santalum/lang/typecheck"
	"github.com/k0kubun/santalum/lang/conversion"
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
