package parsing

import (
	"github.com/cookieo9/go-misc/pp"
	. "github.com/k0kubun/sandal/lang/data"
	"reflect"
	"testing"
)

func parse(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune(src), 0)
	definitions := Parse(s)
	expectPP := pp.PP(expect)
	actualPP := pp.PP(definitions)
	if expectPP != actualPP {
		t.Errorf("\nExpected %s\nGot      %s", expectPP, actualPP)
	}
}

func TestDataDef(t *testing.T) {
	parse(t, "data Maybe { Just, Nothing };",
		[]Def{DataDef{Pos{1, 1}, "Maybe", []string{"Just", "Nothing"}}})
}

func TestParseModuleDef(t *testing.T) {
	parse(t, "module A(ch channel { bool }, chs []channel { bit }) { init { }; };",
		[]Def{ModuleDef{Pos{1, 1}, "A",
			[]Parameter{Parameter{"ch", HandshakeChannelType{[]Type{NamedType{"bool"}}}},
				Parameter{"chs", ArrayType{HandshakeChannelType{[]Type{NamedType{"bit"}}}}}},
			[]Def{InitBlock{Pos: Pos{1, 56}}}}})
}

func TestParseConstantDef(t *testing.T) {
	parse(t, "const a int = 1;", []Def{ConstantDef{Pos{1, 1}, "a", NamedType{"int"}, NumberExpr{Pos{1, 15}, "1"}}})
}

func TestParseProcDef(t *testing.T) {
	parse(t, "proc A(ch channel { bool }, chs []channel { bit }) { ; };",
		[]Def{ProcDef{Pos{1, 1}, "A",
			[]Parameter{Parameter{"ch", HandshakeChannelType{[]Type{NamedType{"bool"}}}},
				Parameter{"chs", ArrayType{HandshakeChannelType{[]Type{NamedType{"bit"}}}}}},
			[]Stmt{NullStmt{Pos{1, 54}}}}})
}

func TestParseInitBlock(t *testing.T) {
	parse(
		t,
		"init { };",
		[]Def{
			InitBlock{
				Pos: Pos{1, 1},
			},
		},
	)

	parse(
		t,
		"init { a : M(b) };",
		[]Def{
			InitBlock{
				Pos{1, 1},
				[]InitVar{
					InstanceVar{
						Pos{1, 8},
						"a",
						"M",
						[]Expr{IdentifierExpr{Pos{1, 14}, "b"}},
						[]string{},
					},
				},
			},
		},
	)

	parse(
		t,
		"init { a : M(b) @unstable };",
		[]Def{
			InitBlock{
				Pos{1, 1},
				[]InitVar{
					InstanceVar{
						Pos{1, 8},
						"a",
						"M",
						[]Expr{IdentifierExpr{Pos{1, 14}, "b"}},
						[]string{"unstable"},
					},
				},
			},
		},
	)

	parse(
		t,
		"init { a : channel { bool } };",
		[]Def{
			InitBlock{
				Pos{1, 1},
				[]InitVar{
					ChannelVar{
						Pos{1, 8},
						"a",
						HandshakeChannelType{[]Type{NamedType{"bool"}}},
						[]string{},
					},
				},
			},
		},
	)

	parse(
		t,
		"init { a : channel { bool } @unstable };",
		[]Def{
			InitBlock{
				Pos{1, 1},
				[]InitVar{
					ChannelVar{
						Pos{1, 8},
						"a",
						HandshakeChannelType{[]Type{NamedType{"bool"}}},
						[]string{"unstable"},
					},
				},
			},
		},
	)
}

const parseBlockOffset = 11

func parseInBlock(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune("proc A() { "+src+" }"), 0)
	definitions := Parse(s)
	if len(definitions) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if procDef, isProcDef := definitions[0].(ProcDef); isProcDef {
		if len(procDef.Stmts) != 1 {
			t.Errorf("Expect %q to be parsed in ProcDef", src)
			return
		}
		if !reflect.DeepEqual(procDef.Stmts[0], expect) {
			t.Errorf("\nExpected %s\nGot      %s", pp.PP(expect), pp.PP(procDef.Stmts[0]))
			return
		}
	} else {
		t.Errorf("Expect %q to be parsed in ProcDef", src)
		return
	}
}

func TestParseStmt(t *testing.T) {
	parseInBlock(t, "test: ;", LabelledStmt{Pos{1, 1 + parseBlockOffset}, "test", NullStmt{Pos{1, 7 + parseBlockOffset}}})
	parseInBlock(t, "{ ; };", BlockStmt{Pos{1, 1 + parseBlockOffset}, []Stmt{NullStmt{Pos{1, 3 + parseBlockOffset}}}})
	parseInBlock(t, "var abc bool;", VarDeclStmt{Pos{1, 1 + parseBlockOffset}, "abc", NamedType{"bool"}, nil})
	parseInBlock(t, "var abc bool = false;", VarDeclStmt{Pos{1, 1 + parseBlockOffset}, "abc", NamedType{"bool"}, FalseExpr{Pos{1, 16 + parseBlockOffset}}})
	parseInBlock(t, "if false { ; };", IfStmt{Pos{1, 1 + parseBlockOffset}, FalseExpr{Pos{1, 4 + parseBlockOffset}}, []Stmt{NullStmt{Pos{1, 12 + parseBlockOffset}}}, nil})
	parseInBlock(t, "if false { ; } else { skip; };", IfStmt{Pos{1, 1 + parseBlockOffset}, FalseExpr{Pos{1, 4 + parseBlockOffset}}, []Stmt{NullStmt{Pos{1, 12 + parseBlockOffset}}}, []Stmt{SkipStmt{Pos{1, 23 + parseBlockOffset}}}})

	parseInBlock(t, "a=b;", AssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", IdentifierExpr{Pos{1, 3 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a+=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "+", IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a-=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "-", IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a*=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "*", IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a/=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "/", IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a%=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "%", IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a&=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "&", IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a|=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "|", IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a^=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "^", IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a<<=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", "<<", IdentifierExpr{Pos{1, 5 + parseBlockOffset}, "b"}})
	parseInBlock(t, "a>>=b;", OpAssignmentStmt{Pos{1, 1 + parseBlockOffset}, "a", ">>", IdentifierExpr{Pos{1, 5 + parseBlockOffset}, "b"}})

	parseInBlock(t, "choice { ; }, { skip; };", ChoiceStmt{Pos{1, 1 + parseBlockOffset}, []BlockStmt{BlockStmt{Pos{1, 8 + parseBlockOffset}, []Stmt{NullStmt{Pos{1, 10 + parseBlockOffset}}}}, BlockStmt{Pos{1, 15 + parseBlockOffset}, []Stmt{SkipStmt{Pos{1, 17 + parseBlockOffset}}}}}})
	parseInBlock(t, "recv(ch, 1, 2);", RecvStmt{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 6 + parseBlockOffset}, "ch"}, []Expr{NumberExpr{Pos{1, 10 + parseBlockOffset}, "1"}, NumberExpr{Pos{1, 13 + parseBlockOffset}, "2"}}, nil})
	parseInBlock(t, "recv(ch, 1, 2) @omission;", RecvStmt{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 6 + parseBlockOffset}, "ch"}, []Expr{NumberExpr{Pos{1, 10 + parseBlockOffset}, "1"}, NumberExpr{Pos{1, 13 + parseBlockOffset}, "2"}}, []string{"omission"}})
	parseInBlock(t, "peek(ch);", PeekStmt{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 6 + parseBlockOffset}, "ch"}, []Expr{}})
	parseInBlock(t, "send(ch, 1, 2);", SendStmt{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 6 + parseBlockOffset}, "ch"}, []Expr{NumberExpr{Pos{1, 10 + parseBlockOffset}, "1"}, NumberExpr{Pos{1, 13 + parseBlockOffset}, "2"}}, nil})
	parseInBlock(t, "send(ch, 1, 2) @commission;", SendStmt{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 6 + parseBlockOffset}, "ch"}, []Expr{NumberExpr{Pos{1, 10 + parseBlockOffset}, "1"}, NumberExpr{Pos{1, 13 + parseBlockOffset}, "2"}}, []string{"commission"}})
	parseInBlock(t, "for { ; };", ForStmt{Pos{1, 1 + parseBlockOffset}, []Stmt{NullStmt{Pos{1, 7 + parseBlockOffset}}}})
	parseInBlock(t, "for ch in chs { ; };", ForInStmt{Pos{1, 1 + parseBlockOffset}, "ch", IdentifierExpr{Pos{1, 11 + parseBlockOffset}, "chs"}, []Stmt{NullStmt{Pos{1, 17 + parseBlockOffset}}}})
	parseInBlock(t, "for i in range 1 to 5 { ; };", ForInRangeStmt{Pos{1, 1 + parseBlockOffset}, "i", NumberExpr{Pos{1, 16 + parseBlockOffset}, "1"}, NumberExpr{Pos{1, 21 + parseBlockOffset}, "5"}, []Stmt{NullStmt{Pos{1, 25 + parseBlockOffset}}}})
	parseInBlock(t, "break;", BreakStmt{Pos{1, 1 + parseBlockOffset}})
	parseInBlock(t, "goto here;", GotoStmt{Pos{1, 1 + parseBlockOffset}, "here"})
	parseInBlock(t, "skip;", SkipStmt{Pos{1, 1 + parseBlockOffset}})
	parseInBlock(t, ";", NullStmt{Pos{1, 1 + parseBlockOffset}})
	parseInBlock(t, "1;", ExprStmt{NumberExpr{Pos{1, 1 + parseBlockOffset}, "1"}})
	parseInBlock(t, "const a int = 1;", ConstantDef{Pos{1, 1 + parseBlockOffset}, "a", NamedType{"int"}, NumberExpr{Pos{1, 15 + parseBlockOffset}, "1"}})
}

func TestParseExpr(t *testing.T) {
	parseInBlock(t, "abc;", ExprStmt{IdentifierExpr{Pos{1, 1 + parseBlockOffset}, "abc"}})
	parseInBlock(t, "123;", ExprStmt{NumberExpr{Pos{1, 1 + parseBlockOffset}, "123"}})
	parseInBlock(t, "true;", ExprStmt{TrueExpr{Pos{1, 1 + parseBlockOffset}}})
	parseInBlock(t, "false;", ExprStmt{FalseExpr{Pos{1, 1 + parseBlockOffset}}})
	parseInBlock(t, "!abc;", ExprStmt{NotExpr{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 2 + parseBlockOffset}, "abc"}}})
	parseInBlock(t, "-abc;", ExprStmt{UnarySubExpr{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 2 + parseBlockOffset}, "abc"}}})
	parseInBlock(t, "(abc);", ExprStmt{ParenExpr{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 2 + parseBlockOffset}, "abc"}}})

	aExp := IdentifierExpr{Pos{1, 1 + parseBlockOffset}, "a"}
	bExp := IdentifierExpr{Pos{1, 4 + parseBlockOffset}, "b"}
	parseInBlock(t, "a+ b;", ExprStmt{BinOpExpr{aExp, "+", bExp}})
	parseInBlock(t, "a- b;", ExprStmt{BinOpExpr{aExp, "-", bExp}})
	parseInBlock(t, "a* b;", ExprStmt{BinOpExpr{aExp, "*", bExp}})
	parseInBlock(t, "a/ b;", ExprStmt{BinOpExpr{aExp, "/", bExp}})
	parseInBlock(t, "a% b;", ExprStmt{BinOpExpr{aExp, "%", bExp}})
	parseInBlock(t, "a& b;", ExprStmt{BinOpExpr{aExp, "&", bExp}})
	parseInBlock(t, "a| b;", ExprStmt{BinOpExpr{aExp, "|", bExp}})
	parseInBlock(t, "a^ b;", ExprStmt{BinOpExpr{aExp, "^", bExp}})
	parseInBlock(t, "a<<b;", ExprStmt{BinOpExpr{aExp, "<<", bExp}})
	parseInBlock(t, "a>>b;", ExprStmt{BinOpExpr{aExp, ">>", bExp}})
	parseInBlock(t, "a&&b;", ExprStmt{BinOpExpr{aExp, "&&", bExp}})
	parseInBlock(t, "a||b;", ExprStmt{BinOpExpr{aExp, "||", bExp}})
	parseInBlock(t, "a==b;", ExprStmt{BinOpExpr{aExp, "==", bExp}})
	parseInBlock(t, "a< b;", ExprStmt{BinOpExpr{aExp, "<", bExp}})
	parseInBlock(t, "a> b;", ExprStmt{BinOpExpr{aExp, ">", bExp}})
	parseInBlock(t, "a!=b;", ExprStmt{BinOpExpr{aExp, "!=", bExp}})
	parseInBlock(t, "a<=b;", ExprStmt{BinOpExpr{aExp, "<=", bExp}})
	parseInBlock(t, "a>=b;", ExprStmt{BinOpExpr{aExp, ">=", bExp}})

	parseInBlock(t, "timeout_recv(ch);", ExprStmt{TimeoutRecvExpr{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 14 + parseBlockOffset}, "ch"}, []Expr{}}})
	parseInBlock(t, "timeout_peek(ch);", ExprStmt{TimeoutPeekExpr{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 14 + parseBlockOffset}, "ch"}, []Expr{}}})
	parseInBlock(t, "nonblock_recv(ch);", ExprStmt{NonblockRecvExpr{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 15 + parseBlockOffset}, "ch"}, []Expr{}}})
	parseInBlock(t, "nonblock_peek(ch);", ExprStmt{NonblockPeekExpr{Pos{1, 1 + parseBlockOffset}, IdentifierExpr{Pos{1, 15 + parseBlockOffset}, "ch"}, []Expr{}}})
	parseInBlock(t, "[a, b];", ExprStmt{ArrayExpr{Pos{1, 1 + parseBlockOffset}, []Expr{
		IdentifierExpr{Pos{1, 2 + parseBlockOffset}, "a"},
		IdentifierExpr{Pos{1, 5 + parseBlockOffset}, "b"},
	}}})
}

const parseTypeOffset = 17

func parseType(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init([]rune("proc A() { var a "+src+"; }"), 0)
	definitions := Parse(s)
	if len(definitions) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if procDef, isInitBlock := definitions[0].(ProcDef); isInitBlock {
		if len(procDef.Stmts) != 1 {
			t.Errorf("Expect %q to be parsed in ProcDef", src)
			return
		}
		if stmt, isVarDecl := procDef.Stmts[0].(VarDeclStmt); isVarDecl {
			if !reflect.DeepEqual(stmt.Type, expect) {
				t.Errorf("\nExpected %s\nGot      %s", pp.PP(expect), pp.PP(stmt.Type))
				return
			}
		} else {
			t.Errorf("Expect %q to be parsed in ProcDef", src)
			return
		}
	} else {
		t.Errorf("Expect %q to be parsed in ProcDef", src)
		return
	}
}

func TestParseType(t *testing.T) {
	parseType(t, "bool", NamedType{"bool"})
	parseType(t, "[]bool", ArrayType{NamedType{"bool"}})
	parseType(t, "channel { bool }", HandshakeChannelType{[]Type{NamedType{"bool"}}})
	parseType(t, "channel [] { bool }",
		BufferedChannelType{nil, []Type{NamedType{"bool"}}})
	parseType(t, "channel [1+2] { bool }",
		BufferedChannelType{BinOpExpr{NumberExpr{Pos{1, 10 + parseTypeOffset}, "1"}, "+", NumberExpr{Pos{1, 12 + parseTypeOffset}, "2"}}, []Type{NamedType{"bool"}}})
}
