package conversion

import (
	. "github.com/k0kubun/santalum/lang/data"
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
