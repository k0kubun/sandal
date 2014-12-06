package conversion

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

func (x *stmtConverter) convertSend(stmt SendStmt) {
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
			x.convertTags(stmt, tag, currentState, nextState)
		}

		x.currentState = nextState
	}
}

func (x *stmtConverter) convertSendWithoutTag(stmt SendStmt) {
	ch, args := convertChannelExpr(stmt, x.env)
	chType := ch.GetType()

	actions := []intAssign{}
	switch chType.(type) {
	case HandshakeChannelType:
		chVar := resolveRealObj(ch).(ir1HandshakeChannelVar)
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
		chVar := resolveRealObj(ch).(ir1BufferedChannelVar)
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
