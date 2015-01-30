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

func TestStmtStringify(t *testing.T) {
	expectString(t, LabelledStmt{Pos{}, "label", ExprStmt{NumberExpr{Pos{}, "1"}}}, "label: 1;")
	expectString(t, BlockStmt{Pos{}, []Stmt{SkipStmt{}, NullStmt{}}, []string{}}, "{ skip; ; };")
	expectString(t, VarDeclStmt{Pos{}, "a", NamedType{"int"}, NumberExpr{Pos{}, "1"}}, "var a int = 1;")
	expectString(t, VarDeclStmt{Pos{}, "a", NamedType{"int"}, nil}, "var a int;")
	expectString(t, IfStmt{Pos{}, IdentifierExpr{Pos{}, "a"}, []Stmt{SkipStmt{}}, nil}, "if a { skip; };")
	expectString(t, IfStmt{Pos{}, IdentifierExpr{Pos{}, "a"}, []Stmt{SkipStmt{}}, []Stmt{}}, "if a { skip; } else {  };")
	expectString(t, IfStmt{Pos{}, IdentifierExpr{Pos{}, "a"}, []Stmt{SkipStmt{}}, []Stmt{SkipStmt{}}}, "if a { skip; } else { skip; };")
	expectString(t, AssignmentStmt{Pos{}, "a", NumberExpr{Pos{}, "1"}}, "a = 1;")
	expectString(t, OpAssignmentStmt{Pos{}, "a", "+", NumberExpr{Pos{}, "1"}}, "a += 1;")
	expectString(t,
		ChoiceStmt{Pos{}, []BlockStmt{
			BlockStmt{Pos{}, []Stmt{SkipStmt{}}, []string{}},
			BlockStmt{Pos{}, []Stmt{ExprStmt{NumberExpr{Pos{}, "1"}}}, []string{}}}},
		"choice { skip; }, { 1; };")
	expectString(t, RecvStmt{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}, []string{}}, "recv(ch, a);")
	expectString(t, PeekStmt{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}}, "peek(ch, a);")
	expectString(t, SendStmt{Pos{}, IdentifierExpr{Pos{}, "ch"}, []Expr{IdentifierExpr{Pos{}, "a"}}, []string{}}, "send(ch, a);")
	expectString(t, ForStmt{Pos{}, []Stmt{SkipStmt{}}}, "for { skip; };")
	expectString(t, ForInStmt{Pos{}, "ch", IdentifierExpr{Pos{}, "chs"}, []Stmt{SkipStmt{}}}, "for ch in chs { skip; };")
	expectString(t, ForInRangeStmt{Pos{}, "i", NumberExpr{Pos{}, "1"}, NumberExpr{Pos{}, "5"}, []Stmt{SkipStmt{}}}, "for i in range 1 to 5 { skip; };")
	expectString(t, BreakStmt{}, "break;")
	expectString(t, GotoStmt{Pos{}, "label"}, "goto label;")
	expectString(t, SkipStmt{}, "skip;")
	expectString(t, ExprStmt{NumberExpr{Pos{}, "1"}}, "1;")
	expectString(t, NullStmt{}, ";")
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
