package typecheck

import (
	. "github.com/k0kubun/sandal/lang/data"
	"testing"
)

func TestIdentifierExprTypecheck(t *testing.T) {
	expr := IdentifierExpr{Pos{}, "a"}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestNotExprTypecheck(t *testing.T) {
	expr := NotExpr{Pos{}, IdentifierExpr{Pos{}, "a"}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"bool"})
		expectValid(t, expr, env)
	}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectInvalid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestUnarySubExprTypecheck(t *testing.T) {
	expr := UnarySubExpr{Pos{}, IdentifierExpr{Pos{}, "a"}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"bool"})
		expectInvalid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestParenExprTypecheck(t *testing.T) {
	expr := ParenExpr{Pos{}, IdentifierExpr{Pos{}, "a"}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	expectInvalid(t, expr, newTypeEnv())
}

func TestBinOpExprTypecheck(t *testing.T) {
	{
		expr := BinOpExpr{IdentifierExpr{Pos{}, "a"}, "+", IdentifierExpr{Pos{}, "b"}}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			env.add("b", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			env.add("b", NamedType{"bool"})
			expectInvalid(t, expr, env)
		}
		expectInvalid(t, expr, newTypeEnv())
	}
	{
		expr := BinOpExpr{IdentifierExpr{Pos{}, "a"}, "==", IdentifierExpr{Pos{}, "b"}}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			env.add("b", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			env.add("b", NamedType{"bool"})
			expectInvalid(t, expr, env)
		}
	}
}

func TestTimeoutRecvExprTypecheck(t *testing.T) {
	chExp := IdentifierExpr{Pos{}, "ch"}
	{
		expr := TimeoutRecvExpr{Pos{}, chExp, []Expr{IdentifierExpr{Pos{}, "a"}}}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{[]Type{NamedType{"int"}}})
			env.add("a", NamedType{"int"})
			expectValid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{[]Type{NamedType{"int"}}})
			expectInvalid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
		{
			env := newTypeEnv()
			env.add("ch", NamedType{"int"})
			env.add("a", NamedType{"int"})
			expectInvalid(t, expr, env)
		}
	}
	{
		expr := TimeoutRecvExpr{Pos{}, chExp, []Expr{NumberExpr{Pos{}, "1"}}}
		{
			env := newTypeEnv()
			env.add("ch", HandshakeChannelType{[]Type{NamedType{"int"}}})
			expectInvalid(t, expr, env)
		}
	}
}

func TestArrayExprTypecheck(t *testing.T) {
	expr := ArrayExpr{Pos{}, []Expr{IdentifierExpr{Pos{}, "a"}, NumberExpr{Pos{}, "1"}}}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"int"})
		expectValid(t, expr, env)
	}
	{
		env := newTypeEnv()
		env.add("a", NamedType{"bool"})
		expectInvalid(t, expr, env)
	}
}
