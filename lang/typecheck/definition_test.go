package typecheck

import (
	. "github.com/k0kubun/sandal/lang/data"
	"testing"
)

func TestConstantDefinitionTypeCheck(t *testing.T) {
	intType := NamedType{"int"}
	boolType := NamedType{"bool"}
	numberExpr := NumberExpr{Pos{}, "1"}

	expectValid(t, ConstantDefinition{Pos{}, "a", intType, numberExpr}, newTypeEnv())
	expectInvalid(t, ConstantDefinition{Pos{}, "a", boolType, numberExpr}, newTypeEnv())
}

func TestInitBlockTypeCheck(t *testing.T) {
	{
		expectValid(t, InitBlock{Pos{}, []InitVar{
			ChannelVar{Pos{}, "ch", HandshakeChannelType{[]Type{NamedType{"int"}}}, nil},
		}}, newTypeEnv())

		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			ChannelVar{Pos{}, "ch", HandshakeChannelType{[]Type{NamedType{"int"}}}, nil},
			ChannelVar{Pos{}, "ch", HandshakeChannelType{[]Type{NamedType{"int"}}}, nil},
		}}, newTypeEnv())

		expectInvalid(t, InitBlock{Pos{}, []InitVar{
			ChannelVar{Pos{}, "a", NamedType{"int"}, nil},
		}}, newTypeEnv())
	}

	{
		typeEnv := newTypeEnv()
		typeEnv.add("A", CallableType{[]Type{NamedType{"int"}}})
		typeEnv.add("a", NamedType{"int"})
		typeEnv.add("b", NamedType{"bool"})

		expectValid(
			t,
			InitBlock{
				Pos{},
				[]InitVar{
					InstanceVar{
						Pos{},
						"proc1",
						"A",
						[]Expr{IdentifierExpr{Pos{}, "a"}},
						nil,
					},
				},
			},
			typeEnv,
		)

		expectInvalid(
			t,
			InitBlock{
				Pos{},
				[]InitVar{
					InstanceVar{
						Pos{},
						"proc1",
						"a",
						[]Expr{IdentifierExpr{Pos{}, "a"}},
						nil,
					},
				},
			},
			typeEnv,
		)

		expectInvalid(
			t,
			InitBlock{
				Pos{},
				[]InitVar{
					InstanceVar{
						Pos{},
						"proc1",
						"A",
						[]Expr{IdentifierExpr{Pos{}, "a"}, IdentifierExpr{Pos{}, "a"}},
						nil,
					},
				},
			},
			typeEnv,
		)

		expectInvalid(
			t,
			InitBlock{
				Pos{},
				[]InitVar{
					InstanceVar{
						Pos{},
						"proc1",
						"A",
						[]Expr{IdentifierExpr{Pos{}, "b"}},
						nil,
					},
				},
			},
			typeEnv,
		)

		expectInvalid(
			t,
			InitBlock{
				Pos{},
				[]InitVar{
					InstanceVar{
						Pos{},
						"proc1",
						"A",
						[]Expr{IdentifierExpr{Pos{}, "c"}},
						nil,
					},
				},
			},
			typeEnv,
		)

		expectInvalid(
			t,
			InitBlock{
				Pos{},
				[]InitVar{
					InstanceVar{
						Pos{},
						"proc1",
						"A",
						[]Expr{IdentifierExpr{Pos{}, "a"}},
						nil,
					},
				},
			},
			newTypeEnv(),
		)
	}

	{
		typeEnv := newTypeEnv()
		typeEnv.add("A", CallableType{[]Type{HandshakeChannelType{[]Type{NamedType{"int"}}}}})
		expectValid(t, InitBlock{Pos{}, []InitVar{
			InstanceVar{Pos{}, "proc1", "A", []Expr{IdentifierExpr{Pos{}, "ch"}}, nil},
			ChannelVar{Pos{}, "ch", HandshakeChannelType{[]Type{NamedType{"int"}}}, nil},
		}}, typeEnv)
	}
}
