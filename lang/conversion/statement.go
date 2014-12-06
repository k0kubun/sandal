package conversion

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
	"log"
)

func (x *intModConverter) convertStatements(statements []Statement, defaults map[string]string, tags []string, vars []intVar) ([]intVar, intState, []intTransition) {
	converter := newIntStatementConverter(x.env, defaults, tags, vars)

	for _, stmt := range statements {
		converter.convert(stmt)
	}

	return converter.vars, "state0", converter.trans
}

// ========================================
// Statement conversion

type intStatementConverter struct {
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

func newIntStatementConverter(upper *varEnv, defaults map[string]string, tags []string, vars []intVar) *intStatementConverter {
	x := new(intStatementConverter)
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

func (x *intStatementConverter) convert(stmt Statement) {
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
	case LabelledStatement:
		x.convertLabelled(stmt)
	case BlockStatement:
		x.convertBlock(stmt)
	case VarDeclStatement:
		x.convertVarDecl(stmt)
	case IfStatement:
		x.convertIf(stmt)
	case AssignmentStatement:
		x.convertAssignment(stmt)
	case OpAssignmentStatement:
		x.convertOpAssignment(stmt)
	case ChoiceStatement:
		x.convertChoice(stmt)
	case RecvStatement:
		x.convertRecv(stmt)
	case PeekStatement:
		x.convertPeek(stmt)
	case SendStatement:
		x.convertSend(stmt)
	case ForStatement:
		x.convertFor(stmt)
	case ForInStatement:
		x.convertForIn(stmt)
	case ForInRangeStatement:
		x.convertForInRange(stmt)
	case BreakStatement:
		x.convertBreak(stmt)
	case GotoStatement:
		x.convertGoto(stmt)
	case SkipStatement:
		x.convertSkip(stmt)
	case ExprStatement:
		x.convertExpr(stmt)
	case NullStatement:
		x.convertNull(stmt)
	}
}

func (x *intStatementConverter) hasRealName(realName string) bool {
	for _, intvar := range x.vars {
		if intvar.Name == realName {
			return true
		}
	}
	return false
}

func (x *intStatementConverter) genRealName(name string) string {
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

func (x *intStatementConverter) convertConstantDef(stmt ConstantDef) {
	panic("not implemented")
}

func (x *intStatementConverter) convertLabelled(stmt LabelledStatement) {
	x.labelToState[stmt.Label] = x.currentState
	x.convert(stmt.Statement)
}

func (x *intStatementConverter) convertBlock(stmt BlockStatement) {
	nextState := x.genNextState()
	x.pushEnv()
	for _, stmt := range stmt.Statements {
		x.convert(stmt)
	}
	x.popEnv()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}

func (x *intStatementConverter) convertVarDecl(stmt VarDeclStatement) {
	nextState := x.genNextState()

	realName := x.genRealName(stmt.Name)
	nextRealName := fmt.Sprintf("next(%s)", realName)
	if stmt.Initializer != nil {
		intExprObj := expressionToInternalObj(stmt.Initializer, x.env)
		x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, nextRealName)...)
	} else {
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	x.vars = append(x.vars, intVar{realName, convertTypeToString(stmt.Type, x.env)})
	x.env.add(stmt.Name, intInternalPrimitiveVar{realName, stmt.Type, nil})
	x.defaults[nextRealName] = realName
	x.currentState = nextState
}

func (x *intStatementConverter) convertIf(stmt IfStatement) {
	nextState := x.genNextState()
	trueBranchState := x.genNextState()
	falseBranchState := x.genNextState()

	{
		intExprObj := expressionToInternalObj(stmt.Condition, x.env)
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

func (x *intStatementConverter) convertAssignment(stmt AssignmentStatement) {
	nextState := x.genNextState()
	intExprObj := expressionToInternalObj(stmt.Expr, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, fmt.Sprintf("next(%s)", stmt.Variable))...)
	x.currentState = nextState
}

func (x *intStatementConverter) convertOpAssignment(stmt OpAssignmentStatement) {
	nextState := x.genNextState()
	intExprObj := expressionToInternalObj(BinOpExpr{
		IdentifierExpr{Name: stmt.Variable}, stmt.Operator, stmt.Expr,
	}, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, fmt.Sprintf("next(%s)", stmt.Variable))...)
	x.currentState = nextState
}

func (x *intStatementConverter) convertChoice(stmt ChoiceStatement) {
	nextState := x.genNextState()
	currentState := x.currentState
	for _, block := range stmt.Blocks {
		choicedState := x.genNextState()
		x.trans = append(x.trans, intTransition{
			FromState: currentState,
			NextState: choicedState,
		})
		x.currentState = choicedState
		x.pushEnv()
		x.convert(block)
		x.popEnv()
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})
	}
	x.currentState = nextState
}

func (x *intStatementConverter) convertRecv(stmt RecvStatement) {
	nextState := x.genNextState()

	ch, args := convertChannelExpr(stmt, x.env)
	condition := ""
	switch ch.GetType().(type) {
	case HandshakeChannelType:
		condition = fmt.Sprintf("%s.ready & !%s.received", ch, ch)
	case BufferedChannelType:
		condition = fmt.Sprintf("%s.ready", ch)
	default:
		panic("unknown channel type")
	}

	actions := []intAssign{{
		LHS: fmt.Sprintf("%s.recv_received", ch),
		RHS: "TRUE",
	}}
	for i, arg := range args {
		actions = append(actions, intAssign{
			LHS: fmt.Sprintf("next(%s)", arg),
			RHS: fmt.Sprintf("%s.value_%d", ch, i),
		})
	}
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
		Condition: condition,
		Actions:   actions,
	})
	x.currentState = nextState
}

func (x *intStatementConverter) convertPeek(stmt PeekStatement) {
	panic("not implemented")
}

func (x *intStatementConverter) convertSend(stmt SendStatement) {
	if len(stmt.Tags) == 0 {
		x.convertSendWithoutTag(stmt)
	} else {
		nextState := x.genNextState()
		currentState := x.currentState

		choicedState := x.genNextState()
		x.trans = append(x.trans, intTransition{
			FromState: currentState,
			NextState: choicedState,
		})
		x.currentState = choicedState
		x.pushEnv()
		x.convertSendWithoutTag(stmt)
		x.popEnv()
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
		})

		for _, tag := range stmt.Tags {
			choicedState := x.genNextState()
			x.trans = append(x.trans, intTransition{
				FromState: currentState,
				NextState: choicedState,
			})
			x.currentState = choicedState
			x.pushEnv()

			faultVar := x.env.lookup("send@" + tag)
			if faultVar == nil {
				log.Fatalf("Fault @%s does not exist for send", tag)
			}
			faultDef := faultVar.(intInternalFaultDef).Def
			for _, stmt := range faultDef.Statements {
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
}

func (x *intStatementConverter) convertSendWithoutTag(stmt SendStatement) {
	ch, args := convertChannelExpr(stmt, x.env)
	chType := ch.GetType()

	actions := []intAssign{}
	switch chType.(type) {
	case HandshakeChannelType:
		chVar := resolveRealObj(ch).(intInternalHandshakeChannelVar)
		firstState := x.currentState
		secondState := x.genNextState()
		lastState := x.genNextState()

		// Generate the first state transition
		actions = append(actions, intAssign{
			LHS: fmt.Sprintf("%s.send_filled", ch),
			RHS: "TRUE",
		})
		for i, arg := range args {
			actions = append(actions, intAssign{
				LHS: fmt.Sprintf("%s.send_value_%d", ch, i),
				RHS: arg.String(),
			})
		}
		x.trans = append(x.trans, intTransition{
			FromState: firstState,
			NextState: secondState,
			Condition: fmt.Sprintf("!(%s.ready)", ch),
			Actions:   actions,
		})

		// Generate the second state transition
		x.trans = append(x.trans, intTransition{
			FromState: secondState,
			NextState: lastState,
			Condition: fmt.Sprintf("(%s.ready) & (%s.received)", ch, ch),
			Actions: []intAssign{
				{LHS: fmt.Sprintf("%s.send_leaving", ch), RHS: "TRUE"},
			},
		})

		// Inject drop fault
		if hasTag(chVar.Tags, "drop") {
			x.trans = append(x.trans, intTransition{
				FromState: firstState,
				NextState: lastState,
			})
		}

		x.currentState = lastState
	case BufferedChannelType:
		chVar := resolveRealObj(ch).(intInternalBufferedChannelVar)
		nextState := x.genNextState()

		actions = append(actions, intAssign{
			LHS: fmt.Sprintf("%s.send_filled", ch),
			RHS: "TRUE",
		})
		for i, arg := range args {
			actions = append(actions, intAssign{
				LHS: fmt.Sprintf("%s.send_value_%d", ch, i),
				RHS: arg.String(),
			})
		}
		x.trans = append(x.trans, intTransition{
			FromState: x.currentState,
			NextState: nextState,
			Condition: fmt.Sprintf("!(%s.full)", ch),
			Actions:   actions,
		})

		// Inject drop fault
		if hasTag(chVar.Tags, "drop") {
			x.trans = append(x.trans, intTransition{
				FromState: x.currentState,
				NextState: nextState,
			})
		}
		x.currentState = nextState
	default:
		panic("unknown channel type")
	}
}

func (x *intStatementConverter) convertFor(stmt ForStatement) {
	savedCurrentState := x.currentState
	savedBreakState := x.breakToState
	x.breakToState = x.genNextState()
	x.pushEnv()
	for _, stmt := range stmt.Statements {
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

func (x *intStatementConverter) convertForIn(stmt ForInStatement) {
	switch container := expressionToInternalObj(stmt.Container, x.env).(type) {
	case intInternalArrayVar:
		savedBreakState := x.breakToState
		x.breakToState = x.genNextState()
		for i, elem := range container.RealLiteral.Elems {
			x.pushEnv()
			x.env.add(stmt.Variable, intInternalPrimitiveVar{
				fmt.Sprintf("__elem%d_%s", i, container.RealName),
				elem.GetType(),
				elem,
			})
			for _, stmt := range stmt.Statements {
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

func (x *intStatementConverter) convertForInRange(stmt ForInRangeStatement) {
	panic("not implemented")
}

func (x *intStatementConverter) convertBreak(stmt BreakStatement) {
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

func (x *intStatementConverter) convertGoto(stmt GotoStatement) {
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

func (x *intStatementConverter) convertSkip(stmt SkipStatement) {
	nextState := x.genNextState()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}

func (x *intStatementConverter) convertExpr(stmt ExprStatement) {
	nextState := x.genNextState()
	intExprObj := expressionToInternalObj(stmt.Expr, x.env)
	if intExprObj.Steps() > 1 {
		panic("Steps constraint violation")
	}
	x.trans = append(x.trans, intExprObj.Transition(x.currentState, nextState, "")...)
	x.currentState = nextState
}

func (x *intStatementConverter) convertNull(stmt NullStatement) {
	nextState := x.genNextState()
	x.trans = append(x.trans, intTransition{
		FromState: x.currentState,
		NextState: nextState,
	})
	x.currentState = nextState
}

// ========================================

func (x *intStatementConverter) genNextState() (state intState) {
	state = intState(fmt.Sprintf("state%d", x.nextStateNum))
	x.nextStateNum++
	return
}

func (x *intStatementConverter) pushEnv() {
	x.env = newVarEnvFromUpper(x.env)
}

func (x *intStatementConverter) popEnv() {
	x.env = x.env.upper
}
