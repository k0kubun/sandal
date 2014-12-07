package data

type (
	// For type-checking
	ChanExpr interface {
		Position() Pos
		ChannelExpr() Expr
		ArgExprs() []Expr
		String() string
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
)

func (x TimeoutRecvExpr) ChannelExpr() Expr  { return x.Channel }
func (x TimeoutPeekExpr) ChannelExpr() Expr  { return x.Channel }
func (x NonblockRecvExpr) ChannelExpr() Expr { return x.Channel }
func (x NonblockPeekExpr) ChannelExpr() Expr { return x.Channel }
func (x TimeoutRecvExpr) ArgExprs() []Expr   { return x.Args }
func (x TimeoutPeekExpr) ArgExprs() []Expr   { return x.Args }
func (x NonblockRecvExpr) ArgExprs() []Expr  { return x.Args }
func (x NonblockPeekExpr) ArgExprs() []Expr  { return x.Args }
