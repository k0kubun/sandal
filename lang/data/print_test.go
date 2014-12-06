package data

import (
	"testing"
)

type Printable interface {
	String() string
}

func expectString(t *testing.T, printable Printable, expected string) {
	if s := printable.String(); s != expected {
		t.Errorf("Expect %q to be %q", s, expected)
	}
}

func TestDefStringify(t *testing.T) {
	expectString(t, ConstantDef{Pos{}, "a", NamedType{"int"}, NumberExpr{Pos{}, "1"}}, "const a int = 1;")
}

func TestStatementStringify(t *testing.T) {
	expectString(t, LabelledStatement{Pos{}, "label", ExprStatement{NumberExpr{Pos{}, "1"}}}, "label: 1;")
	expectString(t, BlockStatement{Pos{}, []Statement{SkipStatement{}, NullStatement{}}}, "{ skip; ; };")
	expectString(t, VarDeclStatement{Pos{}, "a", NamedType{"int"}, NumberExpr{Pos{}, "1"}}, "var a int = 1;")
	expectString(t, VarDeclStatement{Pos{}, "a", NamedType{"int"}, nil}, "var a int;")
	expectString(t, IfStatement{Pos{}, IdentifierExpr{Pos{}, "a"}, []Statement{SkipStatement{}}, nil}, "if a { skip; };")
	expectString(t, IfStatement{Pos{}, IdentifierExpr{Pos{}, "a"}, []Statement{SkipStatement{}}, []Statement{}}, "if a { skip; } else {  };")
	expectString(t, IfStatement{Pos{}, IdentifierExpr{Pos{}, "a"}, []Statement{SkipStatement{}}, []Statement{SkipStatement{}}}, "if a { skip; } else { skip; };")
	expectString(t, AssignmentStatement{Pos{}, "a", NumberExpr{Pos{}, "1"}}, "a = 1;")
	expectString(t, OpAssignmentStatement{Pos{}, "a", "+", NumberExpr{Pos{}, "1"}}, "a += 1;")
	expectString(t,
		ChoiceStatement{Pos{}, []BlockStatement{
			BlockStatement{Pos{}, []Statement{SkipStatement{}}},
			BlockStatement{Pos{}, []Statement{ExprStatement{NumberExpr{Pos{}, "1"}}}}}},
		"choice { skip; }, { 1; };")
	expectString(t, RecvStatement{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}, []string{}}, "recv(ch, a);")
	expectString(t, PeekStatement{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}}, "peek(ch, a);")
	expectString(t, SendStatement{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}, []string{}}, "send(ch, a);")
	expectString(t, ForStatement{Pos{}, []Statement{SkipStatement{}}}, "for { skip; };")
	expectString(t, ForInStatement{Pos{}, "ch", IdentifierExpr{Pos{}, "chs"}, []Statement{SkipStatement{}}}, "for ch in chs { skip; };")
	expectString(t, ForInRangeStatement{Pos{}, "i", NumberExpr{Pos{}, "1"}, NumberExpr{Pos{}, "5"}, []Statement{SkipStatement{}}}, "for i in range 1 to 5 { skip; };")
	expectString(t, BreakStatement{}, "break;")
	expectString(t, GotoStatement{Pos{}, "label"}, "goto label;")
	expectString(t, SkipStatement{}, "skip;")
	expectString(t, ExprStatement{NumberExpr{Pos{}, "1"}}, "1;")
	expectString(t, NullStatement{}, ";")
}

func TestExprStringify(t *testing.T) {
	expectString(t, IdentifierExpr{Pos{}, "a"}, "a")
	expectString(t, NumberExpr{Pos{}, "1"}, "1")
	expectString(t, TrueExpr{Pos{}}, "true")
	expectString(t, FalseExpr{Pos{}}, "false")
	expectString(t, NotExpr{Pos{}, IdentifierExpr{Pos{}, "a"}}, "!a")
	expectString(t, UnarySubExpr{Pos{}, IdentifierExpr{Pos{}, "a"}}, "-a")
	expectString(t, ParenExpr{Pos{}, IdentifierExpr{Pos{}, "a"}}, "(a)")
	expectString(t, BinOpExpr{IdentifierExpr{Pos{}, "a"}, "+", IdentifierExpr{Pos{}, "b"}}, "a+b")
	expectString(t, TimeoutRecvExpr{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}}, "timeout_recv(ch, a)")
	expectString(t, TimeoutPeekExpr{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}}, "timeout_peek(ch, a)")
	expectString(t, NonblockRecvExpr{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}}, "nonblock_recv(ch, a)")
	expectString(t, NonblockPeekExpr{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}}, "nonblock_peek(ch, a)")
	expectString(t, ArrayExpr{Pos{}, []Expr{IdentifierExpr{Pos{}, "a"}, IdentifierExpr{Pos{}, "b"}}}, "[a, b]")
}

func TestTypeStringify(t *testing.T) {
	expectString(t, NamedType{"int"}, "int")
	expectString(t, CallableType{[]Type{NamedType{"int"}, NamedType{"bool"}}}, "callable(int, bool)")
	expectString(t, ArrayType{NamedType{"int"}}, "[]int")
	expectString(t, HandshakeChannelType{[]Type{NamedType{"int"}}}, "channel {int}")
	expectString(t, BufferedChannelType{IdentifierExpr{Pos{}, "a"}, []Type{NamedType{"int"}}}, "channel [a] {int}")
	expectString(t, BufferedChannelType{nil, []Type{NamedType{"int"}}}, "channel [] {int}")
}
