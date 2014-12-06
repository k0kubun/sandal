package conversion

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

func (x *stmtConverter) convertRecv(stmt RecvStmt) {
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
