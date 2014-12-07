package conversion

import (
	. "github.com/k0kubun/sandal/lang/data"
	"log"
	"reflect"
)

var (
	tagPrefixByStmtType = map[string]string{
		"data.SendStmt": "send",
		"data.RecvStmt": "recv",
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

func (x *stmtConverter) convertTags(stmt Stmt, tag string) {
	prefix := tagPrefixByStmt(stmt)
	faultVar := x.env.lookup(prefix + "@" + tag)
	if faultVar == nil {
		log.Fatalf("Fault @%s does not exist for %s", tag, prefix)
	}

	faultDef := faultVar.(ir1FaultDef).Def
	for _, stmt := range faultDef.Stmts {
		x.convertStmt(stmt)
	}
}
