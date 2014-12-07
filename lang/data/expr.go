package data

import (
	"strings"
)

type (
	Expr interface {
		Position() Pos
		expression()
		String() string
	}

	IdentifierExpr struct {
		Pos
		Name string
	}

	NumberExpr struct {
		Pos
		Lit string
	}

	TrueExpr struct {
		Pos
	}

	FalseExpr struct {
		Pos
	}

	NotExpr struct {
		Pos
		SubExpr Expr
	}

	UnarySubExpr struct {
		Pos
		SubExpr Expr
	}

	ParenExpr struct {
		Pos
		SubExpr Expr
	}

	BinOpExpr struct {
		LHS      Expr
		Operator string
		RHS      Expr
	}

	ArrayExpr struct {
		Pos
		Elems []Expr
	}
)

func (x BinOpExpr) Position() Pos { return x.LHS.Position() }

func (x IdentifierExpr) expression()   {}
func (x NumberExpr) expression()       {}
func (x TrueExpr) expression()         {}
func (x FalseExpr) expression()        {}
func (x NotExpr) expression()          {}
func (x UnarySubExpr) expression()     {}
func (x ParenExpr) expression()        {}
func (x BinOpExpr) expression()        {}
func (x TimeoutRecvExpr) expression()  {}
func (x TimeoutPeekExpr) expression()  {}
func (x NonblockRecvExpr) expression() {}
func (x NonblockPeekExpr) expression() {}
func (x ArrayExpr) expression()        {}

func (x IdentifierExpr) String() string { return x.Name }
func (x NumberExpr) String() string     { return x.Lit }
func (x TrueExpr) String() string       { return "true" }
func (x FalseExpr) String() string      { return "false" }
func (x NotExpr) String() string        { return "!" + x.SubExpr.String() }
func (x UnarySubExpr) String() string   { return "-" + x.SubExpr.String() }
func (x ParenExpr) String() string      { return "(" + x.SubExpr.String() + ")" }
func (x BinOpExpr) String() string      { return x.LHS.String() + x.Operator + x.RHS.String() }

func (x TimeoutRecvExpr) String() string {
	params := []string{x.Channel.String()}
	for _, arg := range x.Args {
		params = append(params, arg.String())
	}
	return "timeout_recv(" + strings.Join(params, ", ") + ")"
}

func (x TimeoutPeekExpr) String() string {
	params := []string{x.Channel.String()}
	for _, arg := range x.Args {
		params = append(params, arg.String())
	}
	return "timeout_peek(" + strings.Join(params, ", ") + ")"
}

func (x NonblockRecvExpr) String() string {
	params := []string{x.Channel.String()}
	for _, arg := range x.Args {
		params = append(params, arg.String())
	}
	return "nonblock_recv(" + strings.Join(params, ", ") + ")"
}

func (x NonblockPeekExpr) String() string {
	params := []string{x.Channel.String()}
	for _, arg := range x.Args {
		params = append(params, arg.String())
	}
	return "nonblock_peek(" + strings.Join(params, ", ") + ")"
}

func (x ArrayExpr) String() string {
	elems := []string{}
	for _, elem := range x.Elems {
		elems = append(elems, elem.String())
	}
	return "[" + strings.Join(elems, ", ") + "]"
}
