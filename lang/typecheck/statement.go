package typecheck

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

// ========================================
// typeCheckStmt

func typeCheckStmt(x Stmt, env *typeEnv) error {
	switch x := x.(type) {
	case ConstantDef:
		return typeCheckConstantDef(x, env)
	case LabelledStmt:
		return typeCheckLabelledStmt(x, env)
	case BlockStmt:
		return typeCheckBlockStmt(x, env)
	case VarDeclStmt:
		return typeCheckVarDeclStmt(x, env)
	case IfStmt:
		return typeCheckIfStmt(x, env)
	case AssignmentStmt:
		return typeCheckAssignmentStmt(x, env)
	case OpAssignmentStmt:
		return typeCheckOpAssignmentStmt(x, env)
	case ChoiceStmt:
		return typeCheckChoiceStmt(x, env)
	case RecvStmt:
		return typeCheckRecvStmt(x, env)
	case PeekStmt:
		return typeCheckPeekStmt(x, env)
	case SendStmt:
		return typeCheckSendStmt(x, env)
	case ForStmt:
		return typeCheckForStmt(x, env)
	case ForInStmt:
		return typeCheckForInStmt(x, env)
	case ForInRangeStmt:
		return typeCheckForInRangeStmt(x, env)
	case BreakStmt:
		return typeCheckBreakStmt(x, env)
	case GotoStmt:
		return typeCheckGotoStmt(x, env)
	case SkipStmt:
		return typeCheckSkipStmt(x, env)
	case ExprStmt:
		return typeCheckExprStmt(x, env)
	case NullStmt:
		return typeCheckNullStmt(x, env)
	}
	panic("Unknown Stmt")
}

func typeCheckStmts(stmts []Stmt, env *typeEnv) error {
	env = newTypeEnvFromUpper(env)
	for _, stmt := range stmts {
		if err := typeCheckStmt(stmt, env); err != nil {
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

func typeCheckLabelledStmt(x LabelledStmt, env *typeEnv) error {
	return typeCheckStmt(x.Stmt, env)
}
func typeCheckBlockStmt(x BlockStmt, env *typeEnv) error {
	return typeCheckStmts(x.Stmts, env)
}
func typeCheckVarDeclStmt(x VarDeclStmt, env *typeEnv) error {
	if x.Initializer != nil {
		if err := typeCheckExpr(x.Initializer, env); err != nil {
			return err
		}
	}
	return nil
}
func typeCheckIfStmt(x IfStmt, env *typeEnv) error {
	if err := typeCheckExpr(x.Condition, env); err != nil {
		return err
	}
	if err := typeCheckStmts(x.TrueBranch, env); err != nil {
		return err
	}
	if err := typeCheckStmts(x.FalseBranch, env); err != nil {
		return err
	}
	return nil
}
func typeCheckAssignmentStmt(x AssignmentStmt, env *typeEnv) error {
	if err := typeCheckExpr(x.Expr, env); err != nil {
		return err
	}
	if ty := env.lookup(x.Variable); ty != nil {
		if !typeOfExpr(x.Expr, env).Equal(ty) {
			return fmt.Errorf("Expect %s to be a type %s (%s)", x.Expr, ty, x.Expr.Position())
		}
	} else {
		return fmt.Errorf("Undefined variable %s (%s)", x.Variable, x.Position())
	}
	return nil
}
func typeCheckOpAssignmentStmt(x OpAssignmentStmt, env *typeEnv) error {
	return typeCheckExpr(
		BinOpExpr{IdentifierExpr{Pos{}, x.Variable}, x.Operator, x.Expr},
		env,
	)
}
func typeCheckChoiceStmt(x ChoiceStmt, env *typeEnv) error {
	for _, block := range x.Blocks {
		if err := typeCheckStmt(block, env); err != nil {
			return err
		}
	}
	return nil
}
func typeCheckRecvStmt(x RecvStmt, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}
func typeCheckPeekStmt(x PeekStmt, env *typeEnv) error {
	return channelExprCheck(x, env, true)
}
func typeCheckSendStmt(x SendStmt, env *typeEnv) error {
	return channelExprCheck(x, env, false)
}
func typeCheckForStmt(x ForStmt, env *typeEnv) error {
	return typeCheckStmts(x.Stmts, env)
}
func typeCheckForInStmt(x ForInStmt, env *typeEnv) error {
	if err := typeCheckExpr(x.Container, env); err != nil {
		return err
	}
	if ty, isArrayType := typeOfExpr(x.Container, env).(ArrayType); isArrayType {
		blockEnv := newTypeEnvFromUpper(env)
		blockEnv.add(x.Variable, ty.ElemType)
		return typeCheckStmts(x.Stmts, blockEnv)
	} else {
		return fmt.Errorf("Expect %s to be an array (%s)", x.Container, x.Container.Position())
	}
}
func typeCheckForInRangeStmt(x ForInRangeStmt, env *typeEnv) error {
	if err := typeCheckExpr(x.FromExpr, env); err != nil {
		return err
	}
	if err := typeCheckExpr(x.ToExpr, env); err != nil {
		return err
	}
	if !typeOfExpr(x.FromExpr, env).Equal(NamedType{"int"}) {
		return fmt.Errorf("Expect %s to be an int (%s)", x.FromExpr, x.FromExpr.Position())
	}
	if !typeOfExpr(x.ToExpr, env).Equal(NamedType{"int"}) {
		return fmt.Errorf("Expect %s to be an int (%s)", x.ToExpr, x.ToExpr.Position())
	}
	blockEnv := newTypeEnvFromUpper(env)
	blockEnv.add(x.Variable, NamedType{"int"})
	return typeCheckStmts(x.Stmts, blockEnv)
}
func typeCheckBreakStmt(x BreakStmt, env *typeEnv) error { return nil }
func typeCheckGotoStmt(x GotoStmt, env *typeEnv) error   { return nil }
func typeCheckSkipStmt(x SkipStmt, env *typeEnv) error   { return nil }
func typeCheckExprStmt(x ExprStmt, env *typeEnv) error   { return nil }
func typeCheckNullStmt(x NullStmt, env *typeEnv) error   { return nil }
