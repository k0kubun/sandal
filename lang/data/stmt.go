package data

import (
	"fmt"
	"strings"
)

type (
	Stmt interface {
		Position() Pos
		statement()
		String() string
	}

	LabelledStmt struct {
		Pos
		Label string
		Stmt  Stmt
	}

	BlockStmt struct {
		Pos
		Stmts []Stmt
	}

	VarDeclStmt struct {
		Pos
		Name        string
		Type        Type
		Initializer Expr
	}

	IfStmt struct {
		Pos
		Condition   Expr
		TrueBranch  []Stmt
		FalseBranch []Stmt
	}

	AssignmentStmt struct {
		Pos
		Variable string
		Expr     Expr
	}

	OpAssignmentStmt struct {
		Pos
		Variable string
		Operator string
		Expr     Expr
	}

	ChoiceStmt struct {
		Pos
		Blocks []BlockStmt
	}

	RecvStmt struct {
		Pos
		Channel Expr
		Args    []Expr
		Tags    []string
	}

	PeekStmt struct {
		Pos
		Channel Expr
		Args    []Expr
	}

	SendStmt struct {
		Pos
		Channel Expr
		Args    []Expr
		Tags    []string
	}

	ForStmt struct {
		Pos
		Stmts []Stmt
	}

	ForInStmt struct {
		Pos
		Variable  string
		Container Expr
		Stmts     []Stmt
	}

	ForInRangeStmt struct {
		Pos
		Variable string
		FromExpr Expr
		ToExpr   Expr
		Stmts    []Stmt
	}

	BreakStmt struct {
		Pos
	}

	GotoStmt struct {
		Pos
		Label string
	}

	SkipStmt struct {
		Pos
	}

	ExprStmt struct {
		Expr
	}

	NullStmt struct {
		Pos
	}
)

func (x LabelledStmt) statement()     {}
func (x BlockStmt) statement()        {}
func (x VarDeclStmt) statement()      {}
func (x IfStmt) statement()           {}
func (x AssignmentStmt) statement()   {}
func (x OpAssignmentStmt) statement() {}
func (x ChoiceStmt) statement()       {}
func (x RecvStmt) statement()         {}
func (x PeekStmt) statement()         {}
func (x SendStmt) statement()         {}
func (x ForStmt) statement()          {}
func (x ForInStmt) statement()        {}
func (x ForInRangeStmt) statement()   {}
func (x BreakStmt) statement()        {}
func (x GotoStmt) statement()         {}
func (x SkipStmt) statement()         {}
func (x ExprStmt) statement()         {}
func (x NullStmt) statement()         {}

func (x RecvStmt) ChannelExpr() Expr { return x.Channel }
func (x PeekStmt) ChannelExpr() Expr { return x.Channel }
func (x SendStmt) ChannelExpr() Expr { return x.Channel }
func (x RecvStmt) ArgExprs() []Expr  { return x.Args }
func (x PeekStmt) ArgExprs() []Expr  { return x.Args }
func (x SendStmt) ArgExprs() []Expr  { return x.Args }

func (x LabelledStmt) String() string   { return x.Label + ": " + x.Stmt.String() }
func (x AssignmentStmt) String() string { return fmt.Sprintf("%s = %s;", x.Variable, x.Expr) }
func (x BreakStmt) String() string      { return "break;" }
func (x GotoStmt) String() string       { return "goto " + x.Label + ";" }
func (x SkipStmt) String() string       { return "skip;" }
func (x ExprStmt) String() string       { return x.Expr.String() + ";" }
func (x NullStmt) String() string       { return ";" }

func (x OpAssignmentStmt) String() string {
	return fmt.Sprintf("%s %s= %s;", x.Variable, x.Operator, x.Expr)
}

func (x BlockStmt) String() string {
	stmts := []string{}
	for _, stmt := range x.Stmts {
		stmts = append(stmts, stmt.String())
	}
	return "{ " + strings.Join(stmts, " ") + " };"
}

func (x VarDeclStmt) String() string {
	if x.Initializer != nil {
		return fmt.Sprintf("var %s %s = %s;", x.Name, x.Type, x.Initializer)
	} else {
		return fmt.Sprintf("var %s %s;", x.Name, x.Type)
	}
}

func (x IfStmt) String() string {
	cond := x.Condition.String()
	tBranch := []string{}
	for _, stmt := range x.TrueBranch {
		tBranch = append(tBranch, stmt.String())
	}
	if x.FalseBranch == nil {
		return fmt.Sprintf("if %s { %s };", cond, strings.Join(tBranch, " "))
	} else {
		fBranch := []string{}
		for _, stmt := range x.FalseBranch {
			fBranch = append(fBranch, stmt.String())
		}
		return fmt.Sprintf("if %s { %s } else { %s };", cond,
			strings.Join(tBranch, " "), strings.Join(fBranch, " "))
	}
}

func (x ChoiceStmt) String() string {
	blocks := []string{}
	for _, block := range x.Blocks {
		stmts := []string{}
		for _, stmt := range block.Stmts {
			stmts = append(stmts, stmt.String())
		}
		blocks = append(blocks, "{ "+strings.Join(stmts, " ")+" }")
	}
	return "choice " + strings.Join(blocks, ", ") + ";"
}

func (x RecvStmt) String() string {
	args := []string{x.Channel.String()}
	for _, arg := range x.Args {
		args = append(args, arg.String())
	}
	return "recv(" + strings.Join(args, ", ") + ");"
}

func (x PeekStmt) String() string {
	args := []string{x.Channel.String()}
	for _, arg := range x.Args {
		args = append(args, arg.String())
	}
	return "peek(" + strings.Join(args, ", ") + ");"
}

func (x SendStmt) String() string {
	args := []string{x.Channel.String()}
	for _, arg := range x.Args {
		args = append(args, arg.String())
	}
	return "send(" + strings.Join(args, ", ") + ");"
}

func (x ForStmt) String() string {
	stmts := []string{}
	for _, stmt := range x.Stmts {
		stmts = append(stmts, stmt.String())
	}
	return fmt.Sprintf("for { %s };", strings.Join(stmts, " "))
}

func (x ForInStmt) String() string {
	stmts := []string{}
	for _, stmt := range x.Stmts {
		stmts = append(stmts, stmt.String())
	}
	return fmt.Sprintf("for %s in %s { %s };", x.Variable, x.Container,
		strings.Join(stmts, " "))
}

func (x ForInRangeStmt) String() string {
	stmts := []string{}
	for _, stmt := range x.Stmts {
		stmts = append(stmts, stmt.String())
	}
	return fmt.Sprintf("for %s in range %s to %s { %s };", x.Variable,
		x.FromExpr, x.ToExpr, strings.Join(stmts, " "))
}
