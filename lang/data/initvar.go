package data

type (
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
)

func (x ChannelVar) initvar()         {}
func (x InstanceVar) initvar()        {}
func (x ChannelVar) Position() Pos    { return x.Pos }
func (x InstanceVar) Position() Pos   { return x.Pos }
func (x ChannelVar) VarName() string  { return x.Name }
func (x InstanceVar) VarName() string { return x.Name }
