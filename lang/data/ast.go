package data

import (
	"fmt"
)

type (
	Pos struct {
		Line   int
		Column int
	}

	Definition interface {
		Position() Pos
		definition()
	}

	Statement interface {
		Position() Pos
		statement()
		String() string
	}

	Expr interface {
		Position() Pos
		expression()
		String() string
	}

	// For type-checking
	ChanExpr interface {
		Position() Pos
		ChannelExpr() Expr
		ArgExprs() []Expr
		String() string
	}

	LtlExpr interface {
		ltlexpression()
	}
)

func (x Pos) String() string {
	return fmt.Sprintf("Line: %d, Column %d", x.Line, x.Column)
}

// ========================================
// Definitions

type (
	DataDefinition struct {
		Pos   Pos
		Name  string
		Elems []string
	}

	ModuleDefinition struct {
		Pos         Pos
		Name        string
		Parameters  []Parameter
		Definitions []Definition
	}

	// ConstantDefinition is a definition but also is a statement.
	ConstantDefinition struct {
		Pos  Pos
		Name string
		Type Type
		Expr Expr
	}

	ProcDefinition struct {
		Pos        Pos
		Name       string
		Parameters []Parameter
		Statements []Statement
	}

	FaultDefinition struct {
		Pos        Pos
		Name       string
		Tag        string
		Parameters []Parameter
		Statements []Statement
	}

	InitBlock struct {
		Pos  Pos
		Vars []InitVar
	}

	LtlSpec struct {
		Expr LtlExpr
	}
)

func (x DataDefinition) definition()     {}
func (x ModuleDefinition) definition()   {}
func (x ConstantDefinition) definition() {}
func (x ConstantDefinition) statement()  {}
func (x ProcDefinition) definition()     {}
func (x FaultDefinition) definition()    {}
func (x InitBlock) definition()          {}
func (x LtlSpec) definition()            {}

func (x DataDefinition) Position() Pos     { return x.Pos }
func (x ModuleDefinition) Position() Pos   { return x.Pos }
func (x ConstantDefinition) Position() Pos { return x.Pos }
func (x ProcDefinition) Position() Pos     { return x.Pos }
func (x FaultDefinition) Position() Pos    { return x.Pos }
func (x InitBlock) Position() Pos          { return x.Pos }
func (x LtlSpec) Position() Pos            { panic("not implemented") }

// ========================================
// Statements

type (
	LabelledStatement struct {
		Pos       Pos
		Label     string
		Statement Statement
	}

	BlockStatement struct {
		Pos        Pos
		Statements []Statement
	}

	VarDeclStatement struct {
		Pos         Pos
		Name        string
		Type        Type
		Initializer Expr
	}

	IfStatement struct {
		Pos         Pos
		Condition   Expr
		TrueBranch  []Statement
		FalseBranch []Statement
	}

	AssignmentStatement struct {
		Pos      Pos
		Variable string
		Expr     Expr
	}

	OpAssignmentStatement struct {
		Pos      Pos
		Variable string
		Operator string
		Expr     Expr
	}

	ChoiceStatement struct {
		Pos    Pos
		Blocks []BlockStatement
	}

	RecvStatement struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
		Tags    []string
	}

	PeekStatement struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
	}

	SendStatement struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
		Tags    []string
	}

	ForStatement struct {
		Pos        Pos
		Statements []Statement
	}

	ForInStatement struct {
		Pos        Pos
		Variable   string
		Container  Expr
		Statements []Statement
	}

	ForInRangeStatement struct {
		Pos        Pos
		Variable   string
		FromExpr   Expr
		ToExpr     Expr
		Statements []Statement
	}

	BreakStatement struct {
		Pos Pos
	}

	GotoStatement struct {
		Pos   Pos
		Label string
	}

	SkipStatement struct {
		Pos Pos
	}

	ExprStatement struct {
		Expr Expr
	}

	NullStatement struct {
		Pos Pos
	}
)

func (x LabelledStatement) statement()     {}
func (x BlockStatement) statement()        {}
func (x VarDeclStatement) statement()      {}
func (x IfStatement) statement()           {}
func (x AssignmentStatement) statement()   {}
func (x OpAssignmentStatement) statement() {}
func (x ChoiceStatement) statement()       {}
func (x RecvStatement) statement()         {}
func (x PeekStatement) statement()         {}
func (x SendStatement) statement()         {}
func (x ForStatement) statement()          {}
func (x ForInStatement) statement()        {}
func (x ForInRangeStatement) statement()   {}
func (x BreakStatement) statement()        {}
func (x GotoStatement) statement()         {}
func (x SkipStatement) statement()         {}
func (x ExprStatement) statement()         {}
func (x NullStatement) statement()         {}

func (x LabelledStatement) Position() Pos     { return x.Pos }
func (x BlockStatement) Position() Pos        { return x.Pos }
func (x VarDeclStatement) Position() Pos      { return x.Pos }
func (x IfStatement) Position() Pos           { return x.Pos }
func (x AssignmentStatement) Position() Pos   { return x.Pos }
func (x OpAssignmentStatement) Position() Pos { return x.Pos }
func (x ChoiceStatement) Position() Pos       { return x.Pos }
func (x RecvStatement) Position() Pos         { return x.Pos }
func (x PeekStatement) Position() Pos         { return x.Pos }
func (x SendStatement) Position() Pos         { return x.Pos }
func (x ForStatement) Position() Pos          { return x.Pos }
func (x ForInStatement) Position() Pos        { return x.Pos }
func (x ForInRangeStatement) Position() Pos   { return x.Pos }
func (x BreakStatement) Position() Pos        { return x.Pos }
func (x GotoStatement) Position() Pos         { return x.Pos }
func (x SkipStatement) Position() Pos         { return x.Pos }
func (x ExprStatement) Position() Pos         { return x.Expr.Position() }
func (x NullStatement) Position() Pos         { return x.Pos }

func (x RecvStatement) ChannelExpr() Expr { return x.Channel }
func (x PeekStatement) ChannelExpr() Expr { return x.Channel }
func (x SendStatement) ChannelExpr() Expr { return x.Channel }
func (x RecvStatement) ArgExprs() []Expr  { return x.Args }
func (x PeekStatement) ArgExprs() []Expr  { return x.Args }
func (x SendStatement) ArgExprs() []Expr  { return x.Args }

// ========================================
// Exprs

type (
	IdentifierExpr struct {
		Pos  Pos
		Name string
	}

	NumberExpr struct {
		Pos Pos
		Lit string
	}

	TrueExpr struct {
		Pos Pos
	}

	FalseExpr struct {
		Pos Pos
	}

	NotExpr struct {
		Pos     Pos
		SubExpr Expr
	}

	UnarySubExpr struct {
		Pos     Pos
		SubExpr Expr
	}

	ParenExpr struct {
		Pos     Pos
		SubExpr Expr
	}

	BinOpExpr struct {
		LHS      Expr
		Operator string
		RHS      Expr
	}

	TimeoutRecvExpr struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
	}

	TimeoutPeekExpr struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
	}

	NonblockRecvExpr struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
	}

	NonblockPeekExpr struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
	}

	ArrayExpr struct {
		Pos   Pos
		Elems []Expr
	}
)

func (x TimeoutRecvExpr) ChannelExpr() Expr  { return x.Channel }
func (x TimeoutPeekExpr) ChannelExpr() Expr  { return x.Channel }
func (x NonblockRecvExpr) ChannelExpr() Expr { return x.Channel }
func (x NonblockPeekExpr) ChannelExpr() Expr { return x.Channel }
func (x TimeoutRecvExpr) ArgExprs() []Expr   { return x.Args }
func (x TimeoutPeekExpr) ArgExprs() []Expr   { return x.Args }
func (x NonblockRecvExpr) ArgExprs() []Expr  { return x.Args }
func (x NonblockPeekExpr) ArgExprs() []Expr  { return x.Args }

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

func (x IdentifierExpr) Position() Pos   { return x.Pos }
func (x NumberExpr) Position() Pos       { return x.Pos }
func (x TrueExpr) Position() Pos         { return x.Pos }
func (x FalseExpr) Position() Pos        { return x.Pos }
func (x NotExpr) Position() Pos          { return x.Pos }
func (x UnarySubExpr) Position() Pos     { return x.Pos }
func (x ParenExpr) Position() Pos        { return x.Pos }
func (x BinOpExpr) Position() Pos        { return x.LHS.Position() }
func (x TimeoutRecvExpr) Position() Pos  { return x.Pos }
func (x TimeoutPeekExpr) Position() Pos  { return x.Pos }
func (x NonblockRecvExpr) Position() Pos { return x.Pos }
func (x NonblockPeekExpr) Position() Pos { return x.Pos }
func (x ArrayExpr) Position() Pos        { return x.Pos }

// ========================================
// Misc

type (
	Parameter struct {
		Name string
		Type Type
	}

	InitVar interface {
		Position() Pos
		initvar()
		VarName() string
	}

	ChannelVar struct {
		Pos  Pos
		Name string
		Type Type
		Tags []string
	}

	InstanceVar struct {
		Pos         Pos
		Name        string
		ProcDefName string
		Args        []Expr
		Tags        []string
	}

	Type interface {
		typetype()
		Equal(Type) bool
		String() string
	}

	NamedType struct {
		Name string
	}

	CallableType struct {
		Parameters []Type
	}

	ArrayType struct {
		ElemType Type
	}

	HandshakeChannelType struct {
		Elems []Type
	}

	BufferedChannelType struct {
		BufferSize Expr
		Elems      []Type
	}
)

func (x ChannelVar) initvar()         {}
func (x InstanceVar) initvar()        {}
func (x ChannelVar) Position() Pos    { return x.Pos }
func (x InstanceVar) Position() Pos   { return x.Pos }
func (x ChannelVar) VarName() string  { return x.Name }
func (x InstanceVar) VarName() string { return x.Name }

func (x NamedType) typetype()            {}
func (x CallableType) typetype()         {}
func (x ArrayType) typetype()            {}
func (x HandshakeChannelType) typetype() {}
func (x BufferedChannelType) typetype()  {}

func (x NamedType) Equal(ty Type) bool {
	if ty, b := ty.(NamedType); b {
		return (ty.Name == x.Name)
	} else {
		return false
	}
}

func (x CallableType) Equal(ty Type) bool {
	if ty, b := ty.(CallableType); b {
		if len(ty.Parameters) != len(x.Parameters) {
			return false
		}
		for i := 0; i < len(x.Parameters); i++ {
			if !ty.Parameters[i].Equal(x.Parameters[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (x ArrayType) Equal(ty Type) bool {
	if ty, b := ty.(ArrayType); b {
		return ty.ElemType.Equal(x.ElemType)
	} else {
		return false
	}
}

func (x HandshakeChannelType) Equal(ty Type) bool {
	if ty, b := ty.(HandshakeChannelType); b {
		if len(ty.Elems) != len(x.Elems) {
			return false
		}
		for i := 0; i < len(x.Elems); i++ {
			if !ty.Elems[i].Equal(x.Elems[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (x BufferedChannelType) Equal(ty Type) bool {
	if ty, b := ty.(BufferedChannelType); b {
		if len(ty.Elems) != len(x.Elems) {
			return false
		}
		for i := 0; i < len(x.Elems); i++ {
			if !ty.Elems[i].Equal(x.Elems[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

// ========================================
// LtlExpr

type (
	LtlAtomExpr struct {
		Names []string
	}

	ParenLtlExpr struct {
		SubExpr LtlExpr
	}

	UnOpLtlExpr struct {
		Operator string
		SubExpr  LtlExpr
	}

	BinOpLtlExpr struct {
		LHS      LtlExpr
		Operator string
		RHS      LtlExpr
	}
)

func (x LtlAtomExpr) ltlexpression()  {}
func (x ParenLtlExpr) ltlexpression() {}
func (x UnOpLtlExpr) ltlexpression()  {}
func (x BinOpLtlExpr) ltlexpression() {}
