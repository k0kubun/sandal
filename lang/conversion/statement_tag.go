package conversion

import (
	. "github.com/k0kubun/sandal/lang/data"
	"log"
	"reflect"
)

var (
	tagPrefixByStmtType = map[string]string{
		"data.SendStmt": "send",
	}
)

func tagPrefixByStmt(stmt Stmt) string {
	t := reflect.TypeOf(stmt).String()
	if prefix, ok := tagPrefixByStmtType[t]; ok {
		return prefix
	}
	log.Fatalf("Fault can't be defined for %s", t)
	return ""
}

func (x *stmtConverter) convertTags(stmt Stmt, tag string, startState, endState intState) {
	choicedState := x.genNextState()
	x.trans = append(x.trans, intTransition{
		FromState: startState,
		NextState: choicedState,
	})
	x.currentState = choicedState
	x.pushEnv()

	prefix := tagPrefixByStmt(stmt)
	faultVar := x.env.lookup(prefix + "@" + tag)
	if faultVar == nil {
		log.Fatalf("Fault @%s does not exist for %s", tag, prefix)
	}
	faultDef := faultVar.(ir1FaultDef).Def
	for _, stmt := range faultDef.Stmts {
		x.convert(stmt)
	}

	x.popEnv()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: endState,
	})
}
