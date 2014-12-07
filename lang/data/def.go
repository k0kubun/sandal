package data

type (
	Def interface {
		Position() Pos
		definition()
	}

	DataDef struct {
		Pos   Pos
		Name  string
		Elems []string
	}

	ModuleDef struct {
		Pos        Pos
		Name       string
		Parameters []Parameter
		Defs       []Def
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
		Stmts      []Stmt
	}

	FaultDef struct {
		Pos        Pos
		Name       string
		Tag        string
		Parameters []Parameter
		Stmts      []Stmt
	}

	InitBlock struct {
		Pos  Pos
		Vars []InitVar
	}

	LtlSpec struct {
		Expr LtlExpr
	}

	Parameter struct {
		Name string
		Type Type
	}
)

func (x DataDef) definition()     {}
func (x ModuleDef) definition()   {}
func (x ConstantDef) definition() {}
func (x ConstantDef) statement()  {}
func (x ProcDef) definition()     {}
func (x FaultDef) definition()    {}
func (x InitBlock) definition()   {}
func (x LtlSpec) definition()     {}

func (x DataDef) Position() Pos     { return x.Pos }
func (x ModuleDef) Position() Pos   { return x.Pos }
func (x ConstantDef) Position() Pos { return x.Pos }
func (x ProcDef) Position() Pos     { return x.Pos }
func (x FaultDef) Position() Pos    { return x.Pos }
func (x InitBlock) Position() Pos   { return x.Pos }
func (x LtlSpec) Position() Pos     { panic("not implemented") }
