package conversion

import (
	. "github.com/k0kubun/sandal/lang/data"
)

func expressionToInternalObj(expr Expr, env *varEnv) intInternalExprObj {
	// This function does not return nil.
	switch expr := expr.(type) {
	case IdentifierExpr:
		intObj := env.lookup(expr.Name)
		if intExprObj, isExprObj := intObj.(intInternalExprObj); isExprObj {
			return intExprObj
		} else {
			panic("Referenced name is not expression")
		}
	case NumberExpr:
		return intInternalLiteral{Lit: expr.Lit, Type: NamedType{"number"}}
	case TrueExpr:
		return intInternalLiteral{Lit: "TRUE", Type: NamedType{"bool"}}
	case FalseExpr:
		return intInternalLiteral{Lit: "FALSE", Type: NamedType{"bool"}}
	case NotExpr:
		return intInternalNot{Sub: expressionToInternalObj(expr.SubExpr, env)}
	case UnarySubExpr:
		return intInternalUnarySub{Sub: expressionToInternalObj(expr.SubExpr, env)}
	case ParenExpr:
		return intInternalParen{Sub: expressionToInternalObj(expr.SubExpr, env)}
	case BinOpExpr:
		intObjLHS := expressionToInternalObj(expr.LHS, env)
		intObjRHS := expressionToInternalObj(expr.RHS, env)
		return intInternalBinOp{LHS: intObjLHS, Op: expr.Operator, RHS: intObjRHS}
	case TimeoutRecvExpr:
		ch, args := convertChannelExpr(expr, env)
		return intInternalTimeoutRecv{Channel: ch, Args: args}
	case TimeoutPeekExpr:
		ch, args := convertChannelExpr(expr, env)
		return intInternalTimeoutPeek{Channel: ch, Args: args}
	case NonblockRecvExpr:
		ch, args := convertChannelExpr(expr, env)
		return intInternalNonblockRecv{Channel: ch, Args: args}
	case NonblockPeekExpr:
		ch, args := convertChannelExpr(expr, env)
		return intInternalNonblockPeek{Channel: ch, Args: args}
	case ArrayExpr:
		elems := []intInternalExprObj{}
		for _, subExpr := range expr.Elems {
			elems = append(elems, expressionToInternalObj(subExpr, env))
		}
		return intInternalArrayLiteral{Elems: elems}
	default:
		panic("Unknown Expr")
	}
}

func convertChannelExpr(expr ChanExpr, env *varEnv) (ch intInternalExprObj, args []intInternalExprObj) {
	ch = expressionToInternalObj(expr.ChannelExpr(), env)
	if ch.Steps() != 0 {
		panic("Steps constraint violation")
	}
	for _, arg := range expr.ArgExprs() {
		argObj := expressionToInternalObj(arg, env)
		if argObj.Steps() != 0 {
			panic("Steps constraint violation")
		}
		args = append(args, argObj)
	}
	return
}
