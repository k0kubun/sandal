package data

type (
	Stmt interface {
		Position() Pos
		statement()
		String() string
	}

	LabelledStmt struct {
		Pos   Pos
		Label string
		Stmt  Stmt
	}

	BlockStmt struct {
		Pos   Pos
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
		Pos   Pos
		Stmts []Stmt
	}

	ForInStmt struct {
		Pos       Pos
		Variable  string
		Container Expr
		Stmts     []Stmt
	}

	ForInRangeStmt struct {
		Pos      Pos
		Variable string
		FromExpr Expr
		ToExpr   Expr
		Stmts    []Stmt
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
