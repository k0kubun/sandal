package conversion

import (
	"fmt"
	"github.com/k0kubun/pp"
	. "github.com/k0kubun/sandal/lang/data"
	"log"
	"strings"
)

func ConvertASTToNuSMV(defs []Definition) (string, error) {
	err, intMods := convertASTToIntModule(defs)
	if err != nil {
		return "", err
	}

	err, tmplMods := convertIntermediateModuleToTemplate(intMods)
	if err != nil {
		return "", err
	}

	mods := []string{}
	for _, tmplMod := range tmplMods {
		mods = append(mods, instantiateTemplate(tmplMod))
	}

	return strings.Join(mods, ""), nil
}

// -- debug functions --

func DumpIR1(defs []Definition) {
	err, intMods := convertASTToIntModule(defs)
	if err != nil {
		log.Fatal("Conversion error: ", err)
	}
	pp.Println(intMods)
}

func DumpIR2(defs []Definition) {
	err, intMods := convertASTToIntModule(defs)
	if err != nil {
		log.Fatal("Conversion error: ", err)
	}

	err, tmplMods := convertIntermediateModuleToTemplate(intMods)
	if err != nil {
		log.Fatal("Conversion error: ", err)
	}
	pp.Println(tmplMods)
}

func DumpGraph(defs []Definition) {
	err, intMods := convertASTToIntModule(defs)
	if err != nil {
		log.Fatal("Conversion error: ", err)
	}

	fmt.Println("digraph MyGraph {")
	for _, intMod := range intMods {
		switch intMod.(type) {
		case intProcModule:
			mod := intMod.(intProcModule)
			for _, trans := range mod.Trans {
				cond := trans.Condition
				if len(cond) > 0 {
					fmt.Printf("	%s -> %s [label=\"%s\"]\n", trans.FromState, trans.NextState, trans.Condition)
				} else {
					fmt.Printf("	%s -> %s\n", trans.FromState, trans.NextState)
				}
			}
		}
	}
	fmt.Println("}")
}
