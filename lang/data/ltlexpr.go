package data

type (
	LtlExpr interface {
		ltlexpression()
	}

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
