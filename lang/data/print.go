package data

import (
	"fmt"
	"strings"
)

// ========================================
// Defs

func (x ConstantDef) String() string {
	return fmt.Sprintf("const %s %s = %s;", x.Name, x.Type, x.Expr)
}

// ========================================
// Statements

func (x LabelledStatement) String() string {
	return x.Label + ": " + x.Statement.String()
}
func (x BlockStatement) String() string {
	stmts := []string{}
	for _, stmt := range x.Statements {
		stmts = append(stmts, stmt.String())
	}
	return "{ " + strings.Join(stmts, " ") + " };"
}
func (x VarDeclStatement) String() string {
	if x.Initializer != nil {
		return fmt.Sprintf("var %s %s = %s;", x.Name, x.Type, x.Initializer)
	} else {
		return fmt.Sprintf("var %s %s;", x.Name, x.Type)
	}
}
func (x IfStatement) String() string {
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
func (x AssignmentStatement) String() string {
	return fmt.Sprintf("%s = %s;", x.Variable, x.Expr)
}
func (x OpAssignmentStatement) String() string {
	return fmt.Sprintf("%s %s= %s;", x.Variable, x.Operator, x.Expr)
}
func (x ChoiceStatement) String() string {
	blocks := []string{}
	for _, block := range x.Blocks {
		stmts := []string{}
		for _, stmt := range block.Statements {
			stmts = append(stmts, stmt.String())
		}
		blocks = append(blocks, "{ "+strings.Join(stmts, " ")+" }")
	}
	return "choice " + strings.Join(blocks, ", ") + ";"
}
func (x RecvStatement) String() string {
	args := []string{x.Channel.String()}
	for _, arg := range x.Args {
		args = append(args, arg.String())
	}
	return "recv(" + strings.Join(args, ", ") + ");"
}
func (x PeekStatement) String() string {
	args := []string{x.Channel.String()}
	for _, arg := range x.Args {
		args = append(args, arg.String())
	}
	return "peek(" + strings.Join(args, ", ") + ");"
}
func (x SendStatement) String() string {
	args := []string{x.Channel.String()}
	for _, arg := range x.Args {
		args = append(args, arg.String())
	}
	return "send(" + strings.Join(args, ", ") + ");"
}
func (x ForStatement) String() string {
	stmts := []string{}
	for _, stmt := range x.Statements {
		stmts = append(stmts, stmt.String())
	}
	return fmt.Sprintf("for { %s };", strings.Join(stmts, " "))
}
func (x ForInStatement) String() string {
	stmts := []string{}
	for _, stmt := range x.Statements {
		stmts = append(stmts, stmt.String())
	}
	return fmt.Sprintf("for %s in %s { %s };", x.Variable, x.Container,
		strings.Join(stmts, " "))
}
func (x ForInRangeStatement) String() string {
	stmts := []string{}
	for _, stmt := range x.Statements {
		stmts = append(stmts, stmt.String())
	}
	return fmt.Sprintf("for %s in range %s to %s { %s };", x.Variable,
		x.FromExpr, x.ToExpr, strings.Join(stmts, " "))
}
func (x BreakStatement) String() string {
	return "break;"
}
func (x GotoStatement) String() string {
	return "goto " + x.Label + ";"
}
func (x SkipStatement) String() string {
	return "skip;"
}
func (x ExprStatement) String() string {
	return x.Expr.String() + ";"
}
func (x NullStatement) String() string {
	return ";"
}

// ========================================
// Exprs

func (x IdentifierExpr) String() string {
	return x.Name
}
func (x NumberExpr) String() string {
	return x.Lit
}
func (x TrueExpr) String() string {
	return "true"
}
func (x FalseExpr) String() string {
	return "false"
}
func (x NotExpr) String() string {
	return "!" + x.SubExpr.String()
}
func (x UnarySubExpr) String() string {
	return "-" + x.SubExpr.String()
}
func (x ParenExpr) String() string {
	return "(" + x.SubExpr.String() + ")"
}
func (x BinOpExpr) String() string {
	return x.LHS.String() + x.Operator + x.RHS.String()
}
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

// ========================================
// Misc

func (x NamedType) String() string {
	return x.Name
}

func (x CallableType) String() string {
	params := []string{}
	for _, param := range x.Parameters {
		params = append(params, param.String())
	}
	return "callable(" + strings.Join(params, ", ") + ")"
}

func (x ArrayType) String() string {
	return "[]" + x.ElemType.String()
}

func (x HandshakeChannelType) String() string {
	elems := []string{}
	for _, elem := range x.Elems {
		elems = append(elems, elem.String())
	}
	return "channel {" + strings.Join(elems, ", ") + "}"
}

func (x BufferedChannelType) String() string {
	bufsize := ""
	if x.BufferSize != nil {
		bufsize = x.BufferSize.String()
	}

	elems := []string{}
	for _, elem := range x.Elems {
		elems = append(elems, elem.String())
	}

	return "channel [" + bufsize + "] {" + strings.Join(elems, ", ") + "}"
}
