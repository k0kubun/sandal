package typecheck

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

// ========================================
// typeOfExpr

func typeOfExpr(x Expr, env *typeEnv) Type {
	switch x := x.(type) {
	case IdentifierExpr:
		return typeOfIdentifierExpr(x, env)
	case NumberExpr:
		return typeOfNumberExpr(x, env)
	case TrueExpr, FalseExpr:
		return NamedType{"bool"}
	case NotExpr:
		return typeOfNotExpr(x, env)
	case UnarySubExpr:
		return typeOfUnarySubExpr(x, env)
	case ParenExpr:
		return typeOfParenExpr(x, env)
	case BinOpExpr:
		return typeOfBinOpExpr(x, env)
	case TimeoutRecvExpr:
		return typeOfTimeoutRecvExpr(x, env)
	case TimeoutPeekExpr:
		return typeOfTimeoutPeekExpr(x, env)
	case NonblockRecvExpr:
		return typeOfNonblockRecvExpr(x, env)
	case NonblockPeekExpr:
		return typeOfNonblockPeekExpr(x, env)
	case ArrayExpr:
		return typeOfArrayExpr(x, env)
	default:
		panic("Unknown Expr")
	}
}

func typeOfIdentifierExpr(x IdentifierExpr, env *typeEnv) Type {
	if x.Name == "true" || x.Name == "false" {
		return NamedType{Name: "bool"}
	}
	return env.lookup(x.Name)
}

func typeOfNumberExpr(x NumberExpr, env *typeEnv) Type {
	return NamedType{Name: "int"}
}

func typeOfNotExpr(x NotExpr, env *typeEnv) Type {
	return typeOfExpr(x.SubExpr, env)
}

func typeOfUnarySubExpr(x UnarySubExpr, env *typeEnv) Type {
	return typeOfExpr(x.SubExpr, env)
}

func typeOfParenExpr(x ParenExpr, env *typeEnv) Type {
	return typeOfExpr(x.SubExpr, env)
}

var operatorResultType = map[string]Type{
	"+":  NamedType{"int"},
	"-":  NamedType{"int"},
	"*":  NamedType{"int"},
	"/":  NamedType{"int"},
	"%":  NamedType{"int"},
	"&":  NamedType{"int"},
	"|":  NamedType{"int"},
	"^":  NamedType{"int"},
	"<<": NamedType{"int"},
	">>": NamedType{"int"},
	"&&": NamedType{"bool"},
	"||": NamedType{"bool"},
	"==": NamedType{"bool"},
	"<":  NamedType{"bool"},
	">":  NamedType{"bool"},
	"!=": NamedType{"bool"},
	"<=": NamedType{"bool"},
	">=": NamedType{"bool"},
}

func typeOfBinOpExpr(x BinOpExpr, env *typeEnv) Type {
	if ty, exist := operatorResultType[x.Operator]; exist {
		return ty
	} else {
		panic("Unknown operator: " + x.Operator)
	}
}

func typeOfTimeoutRecvExpr(x TimeoutRecvExpr, env *typeEnv) Type {
	return NamedType{Name: "bool"}
}

func typeOfTimeoutPeekExpr(x TimeoutPeekExpr, env *typeEnv) Type {
	return NamedType{Name: "bool"}
}

func typeOfNonblockRecvExpr(x NonblockRecvExpr, env *typeEnv) Type {
	return NamedType{Name: "bool"}
}

func typeOfNonblockPeekExpr(x NonblockPeekExpr, env *typeEnv) Type {
	return NamedType{Name: "bool"}
}

func typeOfArrayExpr(x ArrayExpr, env *typeEnv) Type {
	if len(x.Elems) == 0 {
		panic("An array should have at least one element")
	}
	// Every element of an array has the same type.
	return ArrayType{ElemType: typeOfExpr(x.Elems[0], env)}
}

// ========================================
// typeCheckExpr

func typeCheckExpr(x Expr, env *typeEnv) error {
	switch x := x.(type) {
	case IdentifierExpr:
		return typeCheckIdentifierExpr(x, env)
	case NumberExpr, TrueExpr, FalseExpr:
		return nil
	case NotExpr:
		return typeCheckNotExpr(x, env)
	case UnarySubExpr:
		return typeCheckUnarySubExpr(x, env)
	case ParenExpr:
		return typeCheckParenExpr(x, env)
	case BinOpExpr:
		return typeCheckBinOpExpr(x, env)
	case TimeoutRecvExpr:
		return typeCheckTimeoutRecvExpr(x, env)
	case TimeoutPeekExpr:
		return typeCheckTimeoutPeekExpr(x, env)
	case NonblockRecvExpr:
		return typeCheckNonblockRecvExpr(x, env)
	case NonblockPeekExpr:
		return typeCheckNonblockPeekExpr(x, env)
	case ArrayExpr:
		return typeCheckArrayExpr(x, env)
	default:
		panic("Unknown Expr")
	}
}

func typeCheckIdentifierExpr(x IdentifierExpr, env *typeEnv) error {
	if x.Name == "true" || x.Name == "false" {
		return nil
	}
	if env.lookup(x.Name) == nil {
		return fmt.Errorf("Undefined variable %s (%s)", x.Name, x.Position())
	}
	return nil
}

func typeCheckNotExpr(x NotExpr, env *typeEnv) error {
	if err := typeCheckExpr(x.SubExpr, env); err != nil {
		return err
	}
	if !typeOfExpr(x.SubExpr, env).Equal(NamedType{"bool"}) {
		return fmt.Errorf("Expect %s to have type bool, but got %s (%s)",
			x.SubExpr, typeOfExpr(x.SubExpr, env), x.Position())
	}
	return nil
}

func typeCheckUnarySubExpr(x UnarySubExpr, env *typeEnv) error {
	if err := typeCheckExpr(x.SubExpr, env); err != nil {
		return err
	}
	if !typeOfExpr(x.SubExpr, env).Equal(NamedType{"int"}) {
		return fmt.Errorf("Expect %s to have type int, but got %s (%s)",
			x.SubExpr, typeOfExpr(x.SubExpr, env), x.Position())
	}
	return nil
}

func typeCheckParenExpr(x ParenExpr, env *typeEnv) error {
	if err := typeCheckExpr(x.SubExpr, env); err != nil {
		return err
	}
	return nil
}

var operatorOperandType = map[string]Type{
	"+":  NamedType{"int"},
	"-":  NamedType{"int"},
	"*":  NamedType{"int"},
	"/":  NamedType{"int"},
	"%":  NamedType{"int"},
	"&":  NamedType{"int"},
	"|":  NamedType{"int"},
	"^":  NamedType{"int"},
	"<<": NamedType{"int"},
	">>": NamedType{"int"},
	"&&": NamedType{"bool"},
	"||": NamedType{"bool"},
	"==": nil,
	"<":  NamedType{"int"},
	">":  NamedType{"int"},
	"!=": nil,
	"<=": NamedType{"int"},
	">=": NamedType{"int"},
}

func typeCheckBinOpExpr(x BinOpExpr, env *typeEnv) error {
	if err := typeCheckExpr(x.LHS, env); err != nil {
		return err
	}
	if err := typeCheckExpr(x.RHS, env); err != nil {
		return err
	}
	if ty, exist := operatorOperandType[x.Operator]; exist {
		if ty != nil {
			lhsType := typeOfExpr(x.LHS, env)
			if !lhsType.Equal(ty) {
				return fmt.Errorf("Expect %s to have type %s, but got %s (%s)",
					x.LHS, ty, lhsType, x.Position())
			}
			rhsType := typeOfExpr(x.RHS, env)
			if !rhsType.Equal(ty) {
				return fmt.Errorf("Expect %s to have type %s, but got %s (%s)",
					x.RHS, ty, rhsType, x.Position())
			}
		} else {
			lhsType := typeOfExpr(x.LHS, env)
			rhsType := typeOfExpr(x.RHS, env)
			if !lhsType.Equal(rhsType) {
				return fmt.Errorf("Expect %s and %s to have the same type but got %s and %s (%s)",
					x.LHS, x.RHS, lhsType, rhsType, x.Position())
			}
		}
	} else {
		panic("Unknown operator: " + x.Operator)
	}
	return nil
}

func typeCheckTimeoutRecvExpr(x TimeoutRecvExpr, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}

func typeCheckTimeoutPeekExpr(x TimeoutPeekExpr, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}

func typeCheckNonblockRecvExpr(x NonblockRecvExpr, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}

func typeCheckNonblockPeekExpr(x NonblockPeekExpr, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}

func typeCheckArrayExpr(x ArrayExpr, env *typeEnv) error {
	ty := typeOfExpr(x.Elems[0], env)
	for _, elem := range x.Elems {
		if err := typeCheckExpr(elem, env); err != nil {
			return err
		}
		if !typeOfExpr(elem, env).Equal(ty) {
			return fmt.Errorf("Expect %s to be a %s (%s)", elem, ty, elem.Position())
		}
	}
	return nil
}
