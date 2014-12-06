package typecheck

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

// ========================================
// typeCheckDef

func typeCheckDefs(defs []Def, env *typeEnv) error {
	// Put all definitions to the env first. Module and toplevel definition
	// has a scope that can see all names within the block.
	for _, def := range defs {
		switch def := def.(type) {
		case DataDef:
			namedType := NamedType{Name: def.Name}
			for _, elem := range def.Elems {
				env.add(elem, namedType)
			}
		case ModuleDef:
			params := make([]Type, len(def.Parameters))
			for i, p := range def.Parameters {
				params[i] = p.Type
			}
			env.add(def.Name, CallableType{Parameters: params})
		case ConstantDef:
			env.add(def.Name, def.Type)
		case ProcDef:
			params := make([]Type, len(def.Parameters))
			for i, p := range def.Parameters {
				params[i] = p.Type
			}
			env.add(def.Name, CallableType{Parameters: params})
		case FaultDef:
			// TODO: some type check necessary?
			// Do nothing
		case InitBlock:
			// Do nothing
		case LtlSpec:
			// Do nothing
		default:
			panic("Unknown definition type")
		}
	}

	for _, def := range defs {
		if err := typeCheckDef(def, env); err != nil {
			return err
		}
	}
	return nil
}

func typeCheckDef(x Def, env *typeEnv) error {
	switch x := x.(type) {
	case DataDef:
		return typeCheckDataDef(x, env)
	case ModuleDef:
		return typeCheckModuleDef(x, env)
	case ConstantDef:
		return typeCheckConstantDef(x, env)
	case ProcDef:
		return typeCheckProcDef(x, env)
	case FaultDef:
		return typeCheckFaultDef(x, env)
	case InitBlock:
		return typeCheckInitBlock(x, env)
	case LtlSpec:
		// TODO
		return nil
	}
	panic("Unknown Def")
}

func typeCheckDataDef(def DataDef, env *typeEnv) error {
	return nil
}

func typeCheckModuleDef(def ModuleDef, env *typeEnv) error {
	env = newTypeEnvFromUpper(env)
	for _, param := range def.Parameters {
		env.add(param.Name, param.Type)
	}
	for _, def := range def.Defs {
		if err := typeCheckDef(def, env); err != nil {
			return err
		}
	}
	return nil
}

func typeCheckConstantDef(def ConstantDef, env *typeEnv) error {
	if err := typeCheckExpr(def.Expr, env); err != nil {
		return err
	}
	actual := typeOfExpr(def.Expr, env)
	if !actual.Equal(def.Type) {
		return fmt.Errorf("Expect %+#v to have type %+#v but has %+#v (%s)",
			def.Expr, def.Type, actual, def.Position())
	}
	return nil
}

func typeCheckProcDef(def ProcDef, env *typeEnv) error {
	procEnv := newTypeEnvFromUpper(env)
	for _, param := range def.Parameters {
		env.add(param.Name, param.Type)
	}
	for _, stmt := range def.Stmts {
		if err := typeCheckStmt(stmt, procEnv); err != nil {
			return err
		}
		switch s := stmt.(type) {
		case ConstantDef:
			env.add(s.Name, s.Type)
		case VarDeclStmt:
			env.add(s.Name, s.Type)
		}
	}
	return nil
}

func typeCheckFaultDef(def FaultDef, env *typeEnv) error {
	// TODO: type check
	return nil
}

func typeCheckInitBlock(b InitBlock, env *typeEnv) error {
	env = newTypeEnvFromUpper(env)
	names := make(map[string]bool)
	for _, initVar := range b.Vars {
		if _, defined := names[initVar.VarName()]; defined {
			return fmt.Errorf("Varname %s is duplicated (%s)", initVar.VarName(), initVar.Position())
		}
		names[initVar.VarName()] = true

		switch initVar := initVar.(type) {
		case ChannelVar:
			env.add(initVar.Name, initVar.Type)
		case InstanceVar:
			calleeType := env.lookup(initVar.ProcDefName)
			if calleeType == nil {
				return fmt.Errorf("%q should be a callable type (%s)", initVar.ProcDefName, initVar.Position())
			}
			env.add(initVar.Name, calleeType)
		default:
			panic("Unknown initvar type")
		}
	}

	for _, initVar := range b.Vars {
		switch initVar := initVar.(type) {
		case ChannelVar:
			switch initVar.Type.(type) {
			case HandshakeChannelType, BufferedChannelType:
				// OK
			default:
				return fmt.Errorf("%s should be a channel (%s)", initVar.Name, initVar.Position())
			}
		case InstanceVar:
			calleeType := env.lookup(initVar.ProcDefName)
			if t, isCallableType := calleeType.(CallableType); isCallableType {
				if len(t.Parameters) != len(initVar.Args) {
					return fmt.Errorf("Argument count mismatch: %d to %d (%s)", len(initVar.Args), len(t.Parameters), initVar.Position())
				}
				for i := 0; i < len(t.Parameters); i++ {
					if err := typeCheckExpr(initVar.Args[i], env); err != nil {
						return err
					}
					argType := typeOfExpr(initVar.Args[i], env)
					if !argType.Equal(t.Parameters[i]) {
						return fmt.Errorf("Argument type mismatch: %s to %s (%s)", argType, t.Parameters[i], initVar.Position())
					}
				}
			} else {
				return fmt.Errorf("%q should be a callable type (%s)", initVar.ProcDefName, initVar.Position())
			}
		default:
			panic("Unknown initvar type")
		}
	}
	return nil
}
