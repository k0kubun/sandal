package data

import (
	"fmt"
)

type (
	Pos struct {
		Line   int
		Column int
	}

	Def interface {
		Position() Pos
		definition()
	}

	Stmt interface {
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
// Defs

type (
	DataDef struct {
		Pos   Pos
		Name  string
		Elems []string
	}

	ModuleDef struct {
		Pos         Pos
		Name        string
		Parameters  []Parameter
		Defs []Def
	}

	// ConstantDef is a definition but also is a statement.
	ConstantDef struct {
		Pos  Pos
		Name string
		Type Type
		Expr Expr
	}

	ProcDef struct {
		Pos        Pos
		Name       string
		Parameters []Parameter
		Stmts []Stmt
	}

	FaultDef struct {
		Pos        Pos
		Name       string
		Tag        string
		Parameters []Parameter
		Stmts []Stmt
	}

	InitBlock struct {
		Pos  Pos
		Vars []InitVar
	}

	LtlSpec struct {
		Expr LtlExpr
	}
)

func (x DataDef) definition()     {}
func (x ModuleDef) definition()   {}
func (x ConstantDef) definition() {}
func (x ConstantDef) statement()  {}
func (x ProcDef) definition()     {}
func (x FaultDef) definition()    {}
func (x InitBlock) definition()          {}
func (x LtlSpec) definition()            {}

func (x DataDef) Position() Pos     { return x.Pos }
func (x ModuleDef) Position() Pos   { return x.Pos }
func (x ConstantDef) Position() Pos { return x.Pos }
func (x ProcDef) Position() Pos     { return x.Pos }
func (x FaultDef) Position() Pos    { return x.Pos }
func (x InitBlock) Position() Pos          { return x.Pos }
func (x LtlSpec) Position() Pos            { panic("not implemented") }

// ========================================
// Stmts

type (
	LabelledStmt struct {
		Pos       Pos
		Label     string
		Stmt Stmt
	}

	BlockStmt struct {
		Pos        Pos
		Stmts []Stmt
	}

	VarDeclStmt struct {
		Pos         Pos
		Name        string
		Type        Type
		Initializer Expr
	}

	IfStmt struct {
		Pos         Pos
		Condition   Expr
		TrueBranch  []Stmt
		FalseBranch []Stmt
	}

	AssignmentStmt struct {
		Pos      Pos
		Variable string
		Expr     Expr
	}

	OpAssignmentStmt struct {
		Pos      Pos
		Variable string
		Operator string
		Expr     Expr
	}

	ChoiceStmt struct {
		Pos    Pos
		Blocks []BlockStmt
	}

	RecvStmt struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
		Tags    []string
	}

	PeekStmt struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
	}

	SendStmt struct {
		Pos     Pos
		Channel Expr
		Args    []Expr
		Tags    []string
	}

	ForStmt struct {
		Pos        Pos
		Stmts []Stmt
	}

	ForInStmt struct {
		Pos        Pos
		Variable   string
		Container  Expr
		Stmts []Stmt
	}

	ForInRangeStmt struct {
		Pos        Pos
		Variable   string
		FromExpr   Expr
		ToExpr     Expr
		Stmts []Stmt
	}

	BreakStmt struct {
		Pos Pos
	}

	GotoStmt struct {
		Pos   Pos
		Label string
	}

	SkipStmt struct {
		Pos Pos
	}

	ExprStmt struct {
		Expr Expr
	}

	NullStmt struct {
		Pos Pos
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

func (x LabelledStmt) Position() Pos     { return x.Pos }
func (x BlockStmt) Position() Pos        { return x.Pos }
func (x VarDeclStmt) Position() Pos      { return x.Pos }
func (x IfStmt) Position() Pos           { return x.Pos }
func (x AssignmentStmt) Position() Pos   { return x.Pos }
func (x OpAssignmentStmt) Position() Pos { return x.Pos }
func (x ChoiceStmt) Position() Pos       { return x.Pos }
func (x RecvStmt) Position() Pos         { return x.Pos }
func (x PeekStmt) Position() Pos         { return x.Pos }
func (x SendStmt) Position() Pos         { return x.Pos }
func (x ForStmt) Position() Pos          { return x.Pos }
func (x ForInStmt) Position() Pos        { return x.Pos }
func (x ForInRangeStmt) Position() Pos   { return x.Pos }
func (x BreakStmt) Position() Pos        { return x.Pos }
func (x GotoStmt) Position() Pos         { return x.Pos }
func (x SkipStmt) Position() Pos         { return x.Pos }
func (x ExprStmt) Position() Pos         { return x.Expr.Position() }
func (x NullStmt) Position() Pos         { return x.Pos }

func (x RecvStmt) ChannelExpr() Expr { return x.Channel }
func (x PeekStmt) ChannelExpr() Expr { return x.Channel }
func (x SendStmt) ChannelExpr() Expr { return x.Channel }
func (x RecvStmt) ArgExprs() []Expr  { return x.Args }
func (x PeekStmt) ArgExprs() []Expr  { return x.Args }
func (x SendStmt) ArgExprs() []Expr  { return x.Args }

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
