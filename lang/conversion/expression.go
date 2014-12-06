package conversion

import (
	. "github.com/k0kubun/sandal/lang/data"
)

func exprToIr1Obj(expr Expr, env *varEnv) ir1ExprObj {
	// This function does not return nil.
	switch expr := expr.(type) {
	case IdentifierExpr:
		intObj := env.lookup(expr.Name)
		if intExprObj, isExprObj := intObj.(ir1ExprObj); isExprObj {
			return intExprObj
		} else {
			panic("Referenced name is not expression")
		}
	case NumberExpr:
		return ir1Literal{Lit: expr.Lit, Type: NamedType{"number"}}
	case TrueExpr:
		return ir1Literal{Lit: "TRUE", Type: NamedType{"bool"}}
	case FalseExpr:
		return ir1Literal{Lit: "FALSE", Type: NamedType{"bool"}}
	case NotExpr:
		return ir1Not{Sub: exprToIr1Obj(expr.SubExpr, env)}
	case UnarySubExpr:
		return ir1UnarySub{Sub: exprToIr1Obj(expr.SubExpr, env)}
	case ParenExpr:
		return ir1Paren{Sub: exprToIr1Obj(expr.SubExpr, env)}
	case BinOpExpr:
		intObjLHS := exprToIr1Obj(expr.LHS, env)
		intObjRHS := exprToIr1Obj(expr.RHS, env)
		return ir1BinOp{LHS: intObjLHS, Op: expr.Operator, RHS: intObjRHS}
	case TimeoutRecvExpr:
		ch, args := convertChannelExpr(expr, env)
		return ir1TimeoutRecv{Channel: ch, Args: args}
	case TimeoutPeekExpr:
		ch, args := convertChannelExpr(expr, env)
		return ir1TimeoutPeek{Channel: ch, Args: args}
	case NonblockRecvExpr:
		ch, args := convertChannelExpr(expr, env)
		return ir1NonblockRecv{Channel: ch, Args: args}
	case NonblockPeekExpr:
		ch, args := convertChannelExpr(expr, env)
		return ir1NonblockPeek{Channel: ch, Args: args}
	case ArrayExpr:
		elems := []ir1ExprObj{}
		for _, subExpr := range expr.Elems {
			elems = append(elems, exprToIr1Obj(subExpr, env))
		}
		return ir1ArrayLiteral{Elems: elems}
	default:
		panic("Unknown Expr")
	}
}

func convertChannelExpr(expr ChanExpr, env *varEnv) (ch ir1ExprObj, args []ir1ExprObj) {
	ch = exprToIr1Obj(expr.ChannelExpr(), env)
	if ch.Steps() != 0 {
		panic("Steps constraint violation")
	}
	for _, arg := range expr.ArgExprs() {
		argObj := exprToIr1Obj(arg, env)
		if argObj.Steps() != 0 {
			panic("Steps constraint violation")
		}
		args = append(args, argObj)
	}
	return
}
