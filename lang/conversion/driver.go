package conversion

import (
	"fmt"
	"github.com/k0kubun/pretty"
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

func DumpIR1(defs []Definition) {
	err, intMods := convertASTToIntModule(defs)
	if err != nil {
		log.Fatal("Conversion error: ", err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(intMods))
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
	fmt.Printf("%# v\n", pretty.Formatter(tmplMods))
}
