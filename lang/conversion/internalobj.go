package conversion

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

type (
	intInternObj interface {
		intinternalobj()
	}
)

// ========================================
// intInternObj

type (
	intInternConstantDef struct {
		Type Type
		Expr Expr
	}

	intInternDataTypeDef struct {
		Elems []string
	}

	intInternProcDef struct {
		Def ProcDef
	}

	intInternProcVar struct {
		Name       string
		ModuleName string
		Def        intInternProcDef
		Args       []intInternExprObj
		Pid        int
	}

	intInternFaultDef struct {
		Def FaultDef
	}
)

func (x intInternConstantDef) intinternalobj() {}
func (x intInternDataTypeDef) intinternalobj() {}
func (x intInternProcDef) intinternalobj()     {}
func (x intInternProcVar) intinternalobj()     {}
func (x intInternFaultDef) intinternalobj()    {}

// ========================================
// intInternExprObj

type (
	intInternExprObj interface {
		intInternObj
		Steps() int
		Transition(fromState, nextState intState, varName string) []intTransition
		String() string
		GetType() Type
	}

	intInternPrimitiveVar struct {
		RealName string
		Type     Type
		RealObj  intInternExprObj
	}

	intInternArrayVar struct {
		RealName    string
		RealLiteral intInternArrayLiteral
	}

	intInternLiteral struct {
		Lit  string
		Type Type
	}

	intInternNot struct {
		Sub intInternExprObj
	}

	intInternUnarySub struct {
		Sub intInternExprObj
	}

	intInternParen struct {
		Sub intInternExprObj
	}

	intInternBinOp struct {
		LHS intInternExprObj
		Op  string
		RHS intInternExprObj
	}

	intInternTimeoutRecv struct {
		Channel intInternExprObj
		Args    []intInternExprObj
	}

	intInternTimeoutPeek struct {
		Channel intInternExprObj
		Args    []intInternExprObj
	}

	intInternNonblockRecv struct {
		Channel intInternExprObj
		Args    []intInternExprObj
	}

	intInternNonblockPeek struct {
		Channel intInternExprObj
		Args    []intInternExprObj
	}

	intInternArrayLiteral struct {
		Elems []intInternExprObj
	}

	intInternHandshakeChannelVar struct {
		ModuleName string
		RealName   string
		Type       HandshakeChannelType
		Tags       []string
		Pids       map[int]bool
	}

	intInternBufferedChannelVar struct {
		ModuleName string
		RealName   string
		Type       BufferedChannelType
		Tags       []string
		Pids       map[int]bool
	}
)

func (x intInternPrimitiveVar) intinternalobj()        {}
func (x intInternArrayVar) intinternalobj()            {}
func (x intInternLiteral) intinternalobj()             {}
func (x intInternNot) intinternalobj()                 {}
func (x intInternUnarySub) intinternalobj()            {}
func (x intInternParen) intinternalobj()               {}
func (x intInternBinOp) intinternalobj()               {}
func (x intInternTimeoutRecv) intinternalobj()         {}
func (x intInternTimeoutPeek) intinternalobj()         {}
func (x intInternNonblockRecv) intinternalobj()        {}
func (x intInternNonblockPeek) intinternalobj()        {}
func (x intInternArrayLiteral) intinternalobj()        {}
func (x intInternHandshakeChannelVar) intinternalobj() {}
func (x intInternBufferedChannelVar) intinternalobj()  {}

// ========================================
// Steps
// Steps requried to determine the evaluated value of expression.
// TODO: This should be checked beforehand.

func (x intInternPrimitiveVar) Steps() int { return 0 }
func (x intInternArrayVar) Steps() int     { panic("ArrayVar cannot directly be expressed in NuSMV") }
func (x intInternLiteral) Steps() int      { return 0 }
func (x intInternNot) Steps() int          { return x.Sub.Steps() }
func (x intInternUnarySub) Steps() int     { return x.Sub.Steps() }
func (x intInternParen) Steps() int        { return x.Sub.Steps() }
func (x intInternBinOp) Steps() int        { return x.LHS.Steps() + x.RHS.Steps() }
func (x intInternTimeoutRecv) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x intInternTimeoutPeek) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x intInternNonblockRecv) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x intInternNonblockPeek) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x intInternArrayLiteral) Steps() int {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternHandshakeChannelVar) Steps() int { return 0 }
func (x intInternBufferedChannelVar) Steps() int  { return 0 }

// ========================================
// String
// Used for converting internal objects to NuSMV expression.

var operatorConversionTable = map[string]string{
	"+":  "+",
	"-":  "-",
	"*":  "*",
	"/":  "/",
	"%":  "mod",
	"&":  "&",
	"|":  "|",
	"^":  "xor",
	"<<": "<<",
	">>": ">>",
	"&&": "&",
	"||": "|",
	"==": "=",
	"<":  "<",
	">":  ">",
	"!=": "!=",
	"<=": "<=",
	">=": ">=",
}

func (x intInternPrimitiveVar) String() string { return x.RealName }
func (x intInternArrayVar) String() string     { panic("ArrayVar cannot directly be expressed in NuSMV") }
func (x intInternLiteral) String() string      { return x.Lit }
func (x intInternNot) String() string          { return "!" + x.Sub.String() }
func (x intInternUnarySub) String() string     { return "-" + x.Sub.String() }
func (x intInternParen) String() string        { return "(" + x.Sub.String() + ")" }
func (x intInternBinOp) String() string {
	// TODO: this cannot encode nonblock_recv(...) && nonblock_recv(...)
	return x.LHS.String() + operatorConversionTable[x.Op] + x.RHS.String()
}
func (x intInternTimeoutRecv) String() string {
	panic("timeout_recv cannot directly be expressed in NuSMV")
}
func (x intInternTimeoutPeek) String() string {
	panic("timeout_peek cannot directly be expressed in NuSMV")
}
func (x intInternNonblockRecv) String() string {
	panic("nonblock_recv cannot directly be expressed in NuSMV")
}
func (x intInternNonblockPeek) String() string {
	panic("nonblock_recv cannot directly be expressed in NuSMV")
}
func (x intInternArrayLiteral) String() string {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternHandshakeChannelVar) String() string { return x.RealName }
func (x intInternBufferedChannelVar) String() string  { return x.RealName }

func (x intInternArrayLiteral) ArgString() (ret []string) {
	for _, elem := range x.Elems {
		ret = append(ret, elem.String())
	}
	return ret
}

// ========================================
// Transition

func assignByString(x intInternExprObj, fromState, nextState intState, varName string) []intTransition {
	if varName == "" {
		return []intTransition{{FromState: fromState, NextState: nextState}}
	} else {
		return []intTransition{{
			FromState: fromState,
			NextState: nextState,
			Actions: []intAssign{
				{LHS: varName, RHS: x.String()},
			},
		}}
	}
}

func (x intInternPrimitiveVar) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternArrayVar) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("ArrayVar cannot directly be expressed in NuSMV")
}
func (x intInternLiteral) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternNot) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternUnarySub) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternParen) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternBinOp) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternTimeoutRecv) Transition(fromState, nextState intState, varName string) []intTransition {
	chType := x.Channel.GetType()

	recvedTrans := intTransition{FromState: fromState, NextState: nextState}
	timeoutTrans := intTransition{FromState: fromState, NextState: nextState}
	switch chType.(type) {
	case HandshakeChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.ready & !%s.received", x.Channel, x.Channel)
	case BufferedChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.ready", x.Channel)
	default:
		panic("unknown channel type")
	}
	recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
		LHS: fmt.Sprintf("%s.recv_received", x.Channel),
		RHS: "TRUE",
	})
	for i, arg := range x.Args {
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: fmt.Sprintf("next(%s)", arg),
			RHS: fmt.Sprintf("%s.value_%d", x.Channel, i),
		})
	}
	if varName != "" {
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: varName, RHS: "TRUE",
		})
		timeoutTrans.Actions = append(timeoutTrans.Actions, intAssign{
			LHS: varName, RHS: "FALSE",
		})
	}
	return []intTransition{recvedTrans, timeoutTrans}
}
func (x intInternTimeoutPeek) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Not Implemented")
}
func (x intInternNonblockRecv) Transition(fromState, nextState intState, varName string) []intTransition {
	chType := x.Channel.GetType()

	recvedTrans := intTransition{FromState: fromState, NextState: nextState}
	notRecvedTrans := intTransition{FromState: fromState, NextState: nextState}
	switch chType.(type) {
	case HandshakeChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.ready & !%s.received", x.Channel, x.Channel)
	case BufferedChannelType:
		recvedTrans.Condition = fmt.Sprintf("%s.ready", x.Channel)
	default:
		panic("unknown channel type")
	}
	notRecvedTrans.Condition = "!(" + recvedTrans.Condition + ")"
	recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
		LHS: fmt.Sprintf("%s.recv_received", x.Channel),
		RHS: "TRUE",
	})
	for i, arg := range x.Args {
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: fmt.Sprintf("next(%s)", arg),
			RHS: fmt.Sprintf("%s.value_%d", x.Channel, i),
		})
	}
	if varName != "" {
		recvedTrans.Actions = append(recvedTrans.Actions, intAssign{
			LHS: varName, RHS: "TRUE",
		})
		notRecvedTrans.Actions = append(notRecvedTrans.Actions, intAssign{
			LHS: varName, RHS: "FALSE",
		})
	}
	return []intTransition{recvedTrans, notRecvedTrans}
}
func (x intInternNonblockPeek) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Not Implemented")
}
func (x intInternArrayLiteral) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x intInternHandshakeChannelVar) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x intInternBufferedChannelVar) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}

// ========================================
// GetType

var operatorResultType = map[string]Type{
	"+":  NamedType{"int"},
	"-":  NamedType{"int"},
	"*":  NamedType{"int"},
	"/":  NamedType{"int"},
	"%":  NamedType{"int"},
	"&":  NamedType{"int"},
	"|":  NamedType{"int"},
	"^":  NamedType{"int"},
	"<<": NamedType{"int"},
	">>": NamedType{"int"},
	"&&": NamedType{"bool"},
	"||": NamedType{"bool"},
	"==": NamedType{"bool"},
	"<":  NamedType{"bool"},
	">":  NamedType{"bool"},
	"!=": NamedType{"bool"},
	"<=": NamedType{"bool"},
	">=": NamedType{"bool"},
}

func (x intInternPrimitiveVar) GetType() Type        { return x.Type }
func (x intInternArrayVar) GetType() Type            { return x.RealLiteral.GetType() }
func (x intInternLiteral) GetType() Type             { return x.Type }
func (x intInternNot) GetType() Type                 { return x.Sub.GetType() }
func (x intInternUnarySub) GetType() Type            { return x.Sub.GetType() }
func (x intInternParen) GetType() Type               { return x.Sub.GetType() }
func (x intInternBinOp) GetType() Type               { return operatorResultType[x.Op] }
func (x intInternTimeoutRecv) GetType() Type         { return NamedType{"bool"} }
func (x intInternTimeoutPeek) GetType() Type         { return NamedType{"bool"} }
func (x intInternNonblockRecv) GetType() Type        { return NamedType{"bool"} }
func (x intInternNonblockPeek) GetType() Type        { return NamedType{"bool"} }
func (x intInternArrayLiteral) GetType() Type        { return ArrayType{x.Elems[0].GetType()} }
func (x intInternHandshakeChannelVar) GetType() Type { return x.Type }
func (x intInternBufferedChannelVar) GetType() Type  { return x.Type }

// ========================================

func resolveRealObj(obj intInternExprObj) intInternExprObj {
	for {
		if primVarObj, isPrimVarObj := obj.(intInternPrimitiveVar); isPrimVarObj {
			if primVarObj.RealObj != nil {
				obj = primVarObj.RealObj
			} else {
				return obj
			}
		} else {
			return obj
		}
	}
}
