package data

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
