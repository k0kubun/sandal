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

	TimeoutRecvExpr struct {
		Pos
		Channel Expr
		Args    []Expr
	}

	TimeoutPeekExpr struct {
		Pos
		Channel Expr
		Args    []Expr
	}

	NonblockRecvExpr struct {
		Pos
		Channel Expr
		Args    []Expr
	}

	NonblockPeekExpr struct {
		Pos
		Channel Expr
		Args    []Expr
	}

	ArrayExpr struct {
		Pos
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

func (x BinOpExpr) Position() Pos { return x.LHS.Position() }

func (x TimeoutRecvExpr) ChannelExpr() Expr  { return x.Channel }
func (x TimeoutPeekExpr) ChannelExpr() Expr  { return x.Channel }
func (x NonblockRecvExpr) ChannelExpr() Expr { return x.Channel }
func (x NonblockPeekExpr) ChannelExpr() Expr { return x.Channel }
func (x TimeoutRecvExpr) ArgExprs() []Expr   { return x.Args }
func (x TimeoutPeekExpr) ArgExprs() []Expr   { return x.Args }
func (x NonblockRecvExpr) ArgExprs() []Expr  { return x.Args }
func (x NonblockPeekExpr) ArgExprs() []Expr  { return x.Args }
