package data

type (
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

func (x TimeoutRecvExpr) ChannelExpr() Expr  { return x.Channel }
func (x TimeoutPeekExpr) ChannelExpr() Expr  { return x.Channel }
func (x NonblockRecvExpr) ChannelExpr() Expr { return x.Channel }
func (x NonblockPeekExpr) ChannelExpr() Expr { return x.Channel }
func (x TimeoutRecvExpr) ArgExprs() []Expr   { return x.Args }
func (x TimeoutPeekExpr) ArgExprs() []Expr   { return x.Args }
func (x NonblockRecvExpr) ArgExprs() []Expr  { return x.Args }
func (x NonblockPeekExpr) ArgExprs() []Expr  { return x.Args }
