package conversion

import (
	. "github.com/k0kubun/sandal/lang/data"
)

func expressionToInternObj(expr Expr, env *varEnv) intInternExprObj {
	// This function does not return nil.
	switch expr := expr.(type) {
	case IdentifierExpr:
		intObj := env.lookup(expr.Name)
		if intExprObj, isExprObj := intObj.(intInternExprObj); isExprObj {
			return intExprObj
		} else {
			panic("Referenced name is not expression")
		}
	case NumberExpr:
		return intInternLiteral{Lit: expr.Lit, Type: NamedType{"number"}}
	case TrueExpr:
		return intInternLiteral{Lit: "TRUE", Type: NamedType{"bool"}}
	case FalseExpr:
		return intInternLiteral{Lit: "FALSE", Type: NamedType{"bool"}}
	case NotExpr:
		return intInternNot{Sub: expressionToInternObj(expr.SubExpr, env)}
	case UnarySubExpr:
		return intInternUnarySub{Sub: expressionToInternObj(expr.SubExpr, env)}
	case ParenExpr:
		return intInternParen{Sub: expressionToInternObj(expr.SubExpr, env)}
	case BinOpExpr:
		intObjLHS := expressionToInternObj(expr.LHS, env)
		intObjRHS := expressionToInternObj(expr.RHS, env)
		return intInternBinOp{LHS: intObjLHS, Op: expr.Operator, RHS: intObjRHS}
	case TimeoutRecvExpr:
		ch, args := convertChannelExpr(expr, env)
		return intInternTimeoutRecv{Channel: ch, Args: args}
	case TimeoutPeekExpr:
		ch, args := convertChannelExpr(expr, env)
		return intInternTimeoutPeek{Channel: ch, Args: args}
	case NonblockRecvExpr:
		ch, args := convertChannelExpr(expr, env)
		return intInternNonblockRecv{Channel: ch, Args: args}
	case NonblockPeekExpr:
		ch, args := convertChannelExpr(expr, env)
		return intInternNonblockPeek{Channel: ch, Args: args}
	case ArrayExpr:
		elems := []intInternExprObj{}
		for _, subExpr := range expr.Elems {
			elems = append(elems, expressionToInternObj(subExpr, env))
		}
		return intInternArrayLiteral{Elems: elems}
	default:
		panic("Unknown Expr")
	}
}

func convertChannelExpr(expr ChanExpr, env *varEnv) (ch intInternExprObj, args []intInternExprObj) {
	ch = expressionToInternObj(expr.ChannelExpr(), env)
	if ch.Steps() != 0 {
		panic("Steps constraint violation")
	}
	for _, arg := range expr.ArgExprs() {
		argObj := expressionToInternObj(arg, env)
		if argObj.Steps() != 0 {
			panic("Steps constraint violation")
		}
		args = append(args, argObj)
	}
	return
}
