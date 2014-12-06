package conversion

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

func (x *intModConverter) convertStmts(statements []Stmt, defaults map[string]string, tags []string, vars []intVar) ([]intVar, intState, []intTransition) {
	converter := newStmtConverter(x.env, defaults, tags, vars)

	for _, stmt := range statements {
		converter.convert(stmt)
	}

	return converter.vars, "state0", converter.trans
}

// ========================================
// Stmt conversion

type stmtConverter struct {
	env           *varEnv
	vars          []intVar
	defaults      map[string]string
	trans         []intTransition
	currentState  intState
	nextStateNum  int
	labelToState  map[string]intState
	breakToState  intState
	tags          []string
	unstable      bool
	unstableState intState
}

func newStmtConverter(upper *varEnv, defaults map[string]string, tags []string, vars []intVar) *stmtConverter {
	x := new(stmtConverter)
	x.env = newVarEnvFromUpper(upper)
	x.vars = vars
	x.defaults = defaults
	x.currentState = "state0"
	x.nextStateNum = 1
	x.labelToState = make(map[string]intState)
	x.tags = tags

	if hasTag(x.tags, "unstable") || hasTag(x.tags, "reboot") {
		x.unstable = true
		x.unstableState = x.genNextState()

		if hasTag(x.tags, "reboot") {
			x.trans = append(x.trans, intTransition{
				FromState: x.unstableState,
				NextState: "state0",
				Condition: "",
			})
		}
	}
	return x
}

func hasTag(tags []string, tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func (x *stmtConverter) convert(stmt Stmt) {
	if x.unstable {
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: x.unstableState,
			Condition: "",
		})
	}

	switch stmt := stmt.(type) {
	case ConstantDef:
		x.convertConstantDef(stmt)
	case LabelledStmt:
		x.convertLabelled(stmt)
	case BlockStmt:
		x.convertBlock(stmt)
	case VarDeclStmt:
		x.convertVarDecl(stmt)
	case IfStmt:
		x.convertIf(stmt)
	case AssignmentStmt:
		x.convertAssignment(stmt)
	case OpAssignmentStmt:
		x.convertOpAssignment(stmt)
	case ChoiceStmt:
		x.convertChoice(stmt)
	case RecvStmt:
		x.convertRecv(stmt)
	case PeekStmt:
		x.convertPeek(stmt)
	case SendStmt:
		x.convertSend(stmt)
	case ForStmt:
		x.convertFor(stmt)
	case ForInStmt:
		x.convertForIn(stmt)
	case ForInRangeStmt:
		x.convertForInRange(stmt)
	case BreakStmt:
		x.convertBreak(stmt)
	case GotoStmt:
		x.convertGoto(stmt)
	case SkipStmt:
		x.convertSkip(stmt)
	case ExprStmt:
		x.convertExpr(stmt)
	case NullStmt:
		x.convertNull(stmt)
	}
}

func (x *stmtConverter) hasRealName(realName string) bool {
	for _, intvar := range x.vars {
		if intvar.Name == realName {
			return true
		}
	}
	return false
}

func (x *stmtConverter) genRealName(name string) string {
	realName := name
	if x.hasRealName(realName) {
		i := 2
		realName = fmt.Sprintf("%s_%d", name, i)
		for x.hasRealName(realName) {
			i += 1
			realName = fmt.Sprintf("%s_%d", name, i)
		}
	}
	return realName
}

// ========================================

func (x *stmtConverter) convertConstantDef(stmt ConstantDef) {
	panic("not implemented")
}

func (x *stmtConverter) convertLabelled(stmt LabelledStmt) {
	x.labelToState[stmt.Label] = x.currentState
	x.convert(stmt.Stmt)
}

func (x *stmtConverter) convertBlock(stmt BlockStmt) {
	nextState := x.genNextState()
	x.pushEnv()
	for _, stmt := range stmt.Stmts {
		x.convert(stmt)
	}
	x.popEnv()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}

func (x *stmtConverter) convertVarDecl(stmt VarDeclStmt) {
	nextState := x.genNextState()

	realName := x.genRealName(stmt.Name)
	nextRealName := fmt.Sprintf("next(%s)", realName)
	if stmt.Initializer != nil {
		intExprObj := exprToIr1Obj(stmt.Initializer, x.env)
		x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, nextRealName)...)
	} else {
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	x.vars = append(x.vars, intVar{realName, convertTypeToString(stmt.Type, x.env)})
	x.env.add(stmt.Name, ir1PrimitiveVar{realName, stmt.Type, nil})
	x.defaults[nextRealName] = realName
	x.currentState = nextState
}

func (x *stmtConverter) convertIf(stmt IfStmt) {
	nextState := x.genNextState()
	trueBranchState := x.genNextState()
	falseBranchState := x.genNextState()

	{
		intExprObj := exprToIr1Obj(stmt.Condition, x.env)
		if intExprObj.Steps() != 0 {
			panic("Steps constraint violation")
		}
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: trueBranchState,
			Condition: intExprObj.String(),
		})
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: falseBranchState,
			Condition: "!(" + intExprObj.String() + ")",
		})
	}
	{
		x.currentState = trueBranchState
		x.pushEnv()
		for _, stmt := range stmt.TrueBranch {
			x.convert(stmt)
		}
		x.popEnv()
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	{
		x.currentState = falseBranchState
		x.pushEnv()
		for _, stmt := range stmt.FalseBranch {
			x.convert(stmt)
		}
		x.popEnv()
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	x.currentState = nextState
}

func (x *stmtConverter) convertAssignment(stmt AssignmentStmt) {
	nextState := x.genNextState()
	intExprObj := exprToIr1Obj(stmt.Expr, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, fmt.Sprintf("next(%s)", stmt.Variable))...)
	x.currentState = nextState
}

func (x *stmtConverter) convertOpAssignment(stmt OpAssignmentStmt) {
	nextState := x.genNextState()
	intExprObj := exprToIr1Obj(BinOpExpr{
		IdentifierExpr{Name: stmt.Variable}, stmt.Operator, stmt.Expr,
	}, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, fmt.Sprintf("next(%s)", stmt.Variable))...)
	x.currentState = nextState
}

func (x *stmtConverter) convertChoice(stmt ChoiceStmt) {
	nextState := x.genNextState()
	for _, block := range stmt.Blocks {
		x.branched(nextState, func(x *stmtConverter) {
			x.convert(block)
		})
	}
	x.currentState = nextState
}

func (x *stmtConverter) convertPeek(stmt PeekStmt) {
	panic("not implemented")
}

func (x *stmtConverter) convertFor(stmt ForStmt) {
	savedCurrentState := x.currentState
	savedBreakState := x.breakToState
	x.breakToState = x.genNextState()
	x.pushEnv()
	for _, stmt := range stmt.Stmts {
		x.convert(stmt)
	}
	x.popEnv()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: savedCurrentState,
	})
	x.currentState = x.breakToState
	x.breakToState = savedBreakState
}

func (x *stmtConverter) convertForIn(stmt ForInStmt) {
	switch container := exprToIr1Obj(stmt.Container, x.env).(type) {
	case ir1ArrayVar:
		savedBreakState := x.breakToState
		x.breakToState = x.genNextState()
		for i, elem := range container.RealLiteral.Elems {
			x.pushEnv()
			x.env.add(stmt.Variable, ir1PrimitiveVar{
				fmt.Sprintf("__elem%d_%s", i, container.RealName),
				elem.GetType(),
				elem,
			})
			for _, stmt := range stmt.Stmts {
				x.convert(stmt)
			}
			x.popEnv()
		}
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: x.breakToState,
		})
		x.currentState = x.breakToState
		x.breakToState = savedBreakState
	default:
		// TODO
		panic("not implemented")
	}
}

func (x *stmtConverter) convertForInRange(stmt ForInRangeStmt) {
	panic("not implemented")
}

func (x *stmtConverter) convertBreak(stmt BreakStmt) {
	nextState := x.genNextState()
	if x.breakToState == "" {
		panic("Invalid break statement")
	}
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: x.breakToState,
	})
	x.currentState = nextState
}

func (x *stmtConverter) convertGoto(stmt GotoStmt) {
	nextState := x.genNextState()
	jumpState := x.labelToState[stmt.Label]
	if jumpState == "" {
		panic("Invalid jump label")
	}
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: jumpState,
	})
	x.currentState = nextState
}

func (x *stmtConverter) convertSkip(stmt SkipStmt) {
	nextState := x.genNextState()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}

func (x *stmtConverter) convertExpr(stmt ExprStmt) {
	nextState := x.genNextState()
	intExprObj := exprToIr1Obj(stmt.Expr, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, "")...)
	x.currentState = nextState
}

func (x *stmtConverter) convertNull(stmt NullStmt) {
	nextState := x.genNextState()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}

// ========================================

func (x *stmtConverter) genNextState() (state intState) {
	state = intState(fmt.Sprintf("state%d", x.nextStateNum))
	x.nextStateNum++
	return
}

func (x *stmtConverter) pushEnv() {
	x.env = newVarEnvFromUpper(x.env)
}

func (x *stmtConverter) popEnv() {
	x.env = x.env.upper
}

func (x *stmtConverter) branched(nextState intState, f func(*stmtConverter)) {
	currentState := x.currentState
	choicedState := x.genNextState()
	x.trans = append(x.trans, intTransition{
		FromState: currentState,
		NextState: choicedState,
	})
	x.currentState = choicedState
	x.pushEnv()

	f(x)

	x.popEnv()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = currentState
}
