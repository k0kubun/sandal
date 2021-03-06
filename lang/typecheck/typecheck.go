package typecheck

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

func TypeCheck(defs []Def) error {
	return typeCheckDefs(defs, newTypeEnv())
}

// ========================================

type typeEnv struct {
	upper *typeEnv
	scope map[string]Type
}

func newTypeEnv() (ret *typeEnv) {
	ret = new(typeEnv)
	ret.scope = make(map[string]Type)
	return
}

func newTypeEnvFromUpper(upper *typeEnv) (ret *typeEnv) {
	ret = newTypeEnv()
	ret.upper = upper
	return
}

func (env *typeEnv) add(name string, ty Type) {
	env.scope[name] = ty
}

func (env *typeEnv) lookup(name string) Type {
	if ty, found := env.scope[name]; found {
		return ty
	}
	if env.upper != nil {
		return env.upper.lookup(name)
	} else {
		return nil
	}
}

func channelExprCheck(ch ChanExpr, env *typeEnv, recvOrPeek bool) error {
	chExpr := ch.ChannelExpr()
	args := ch.ArgExprs()
	if err := typeCheckExpr(chExpr, env); err != nil {
		return err
	}
	for _, arg := range args {
		if err := typeCheckExpr(arg, env); err != nil {
			return err
		}
	}

	var elemTypes []Type
	switch ty := typeOfExpr(chExpr, env).(type) {
	case HandshakeChannelType:
		elemTypes = ty.Elems
	case BufferedChannelType:
		elemTypes = ty.Elems
	default:
		return fmt.Errorf("Expect the first argument of %s to be a channel but got %s (%s)",
			ch, typeOfExpr(chExpr, env), ch.Position())
	}

	if len(elemTypes) != len(args) {
		return fmt.Errorf("Expect the arugments of %s to have %d elements (%s)",
			ch, len(elemTypes), ch.Position())
	}
	for i := 0; i < len(elemTypes); i++ {
		if !typeOfExpr(args[i], env).Equal(elemTypes[i]) {
			return fmt.Errorf("Expect the argument %s to be a %s (%s)", args[i], elemTypes[i], args[i].Position())
		}
		if recvOrPeek {
			if _, isIdentExpr := args[i].(IdentifierExpr); !isIdentExpr {
				return fmt.Errorf("Expect the argument %s to be an identifier (%s)", args[i], args[i].Position())
			}
		}
	}
	return nil
}
