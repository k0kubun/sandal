package data

type (
	Def interface {
		Position() Pos
		definition()
	}

	DataDef struct {
		Pos
		Name  string
		Elems []string
	}

	ModuleDef struct {
		Pos
		Name       string
		Parameters []Parameter
		Defs       []Def
	}

	// ConstantDef is a definition but also is a statement.
	ConstantDef struct {
		Pos
		Name string
		Type Type
		Expr Expr
	}

	ProcDef struct {
		Pos
		Name       string
		Parameters []Parameter
		Stmts      []Stmt
	}

	FaultDef struct {
		Pos
		Name       string
		Tag        string
		Parameters []Parameter
		Stmts      []Stmt
	}

	InitBlock struct {
		Pos
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

func (x LtlSpec) Position() Pos { panic("not implemented") }
