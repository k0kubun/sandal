package lang

import (
	"github.com/k0kubun/santalum/lang/parsing"
	"github.com/k0kubun/santalum/lang/typecheck"
	"github.com/k0kubun/santalum/lang/conversion_deprecated"
)

func CompileFile(body string) (error, string) {
	scanner := new(parsing.Scanner)
	scanner.Init([]rune(body), 0)
	defs := parsing.Parse(scanner)
	if err := typecheck.TypeCheck(defs); err != nil {
		return err, ""
	}
	return conversion_deprecated.ConvertASTToNuSMV(defs)
}
