package conversion

import (
	"fmt"
	. "github.com/k0kubun/sandal/lang/data"
)

type (
	ir1Obj interface {
		intinternalobj()
	}
)

// ========================================
// ir1Obj

type (
	ir1ConstantDef struct {
		Type Type
		Expr Expr
	}

	ir1DataTypeDef struct {
		Elems []string
	}

	ir1ProcDef struct {
		Def ProcDef
	}

	ir1ProcVar struct {
		Name       string
		ModuleName string
		Def        ir1ProcDef
		Args       []ir1ExprObj
		Pid        int
	}

	ir1FaultDef struct {
		Def FaultDef
	}
)

func (x ir1ConstantDef) intinternalobj() {}
func (x ir1DataTypeDef) intinternalobj() {}
func (x ir1ProcDef) intinternalobj()     {}
func (x ir1ProcVar) intinternalobj()     {}
func (x ir1FaultDef) intinternalobj()    {}

// ========================================
// ir1ExprObj

type (
	ir1ExprObj interface {
		ir1Obj
		Steps() int
		Transition(fromState, nextState intState, varName string) []intTransition
		String() string
		GetType() Type
	}

	ir1PrimitiveVar struct {
		RealName string
		Type     Type
		RealObj  ir1ExprObj
	}

	ir1ArrayVar struct {
		RealName    string
		RealLiteral ir1ArrayLiteral
	}

	ir1Literal struct {
		Lit  string
		Type Type
	}

	ir1Not struct {
		Sub ir1ExprObj
	}

	ir1UnarySub struct {
		Sub ir1ExprObj
	}

	ir1Paren struct {
		Sub ir1ExprObj
	}

	ir1BinOp struct {
		LHS ir1ExprObj
		Op  string
		RHS ir1ExprObj
	}

	ir1TimeoutRecv struct {
		Channel ir1ExprObj
		Args    []ir1ExprObj
	}

	ir1TimeoutPeek struct {
		Channel ir1ExprObj
		Args    []ir1ExprObj
	}

	ir1NonblockRecv struct {
		Channel ir1ExprObj
		Args    []ir1ExprObj
	}

	ir1NonblockPeek struct {
		Channel ir1ExprObj
		Args    []ir1ExprObj
	}

	ir1ArrayLiteral struct {
		Elems []ir1ExprObj
	}

	ir1HandshakeChannelVar struct {
		ModuleName string
		RealName   string
		Type       HandshakeChannelType
		Tags       []string
		Pids       map[int]bool
	}

	ir1BufferedChannelVar struct {
		ModuleName string
		RealName   string
		Type       BufferedChannelType
		Tags       []string
		Pids       map[int]bool
	}
)

func (x ir1PrimitiveVar) intinternalobj()        {}
func (x ir1ArrayVar) intinternalobj()            {}
func (x ir1Literal) intinternalobj()             {}
func (x ir1Not) intinternalobj()                 {}
func (x ir1UnarySub) intinternalobj()            {}
func (x ir1Paren) intinternalobj()               {}
func (x ir1BinOp) intinternalobj()               {}
func (x ir1TimeoutRecv) intinternalobj()         {}
func (x ir1TimeoutPeek) intinternalobj()         {}
func (x ir1NonblockRecv) intinternalobj()        {}
func (x ir1NonblockPeek) intinternalobj()        {}
func (x ir1ArrayLiteral) intinternalobj()        {}
func (x ir1HandshakeChannelVar) intinternalobj() {}
func (x ir1BufferedChannelVar) intinternalobj()  {}

// ========================================
// Steps
// Steps requried to determine the evaluated value of expression.
// TODO: This should be checked beforehand.

func (x ir1PrimitiveVar) Steps() int { return 0 }
func (x ir1ArrayVar) Steps() int     { panic("ArrayVar cannot directly be expressed in NuSMV") }
func (x ir1Literal) Steps() int      { return 0 }
func (x ir1Not) Steps() int          { return x.Sub.Steps() }
func (x ir1UnarySub) Steps() int     { return x.Sub.Steps() }
func (x ir1Paren) Steps() int        { return x.Sub.Steps() }
func (x ir1BinOp) Steps() int        { return x.LHS.Steps() + x.RHS.Steps() }
func (x ir1TimeoutRecv) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x ir1TimeoutPeek) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x ir1NonblockRecv) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x ir1NonblockPeek) Steps() int {
	steps := 1 + x.Channel.Steps()
	for _, arg := range x.Args {
		steps += arg.Steps()
	}
	return steps
}
func (x ir1ArrayLiteral) Steps() int {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x ir1HandshakeChannelVar) Steps() int { return 0 }
func (x ir1BufferedChannelVar) Steps() int  { return 0 }

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

func (x ir1PrimitiveVar) String() string { return x.RealName }
func (x ir1ArrayVar) String() string     { panic("ArrayVar cannot directly be expressed in NuSMV") }
func (x ir1Literal) String() string      { return x.Lit }
func (x ir1Not) String() string          { return "!" + x.Sub.String() }
func (x ir1UnarySub) String() string     { return "-" + x.Sub.String() }
func (x ir1Paren) String() string        { return "(" + x.Sub.String() + ")" }
func (x ir1BinOp) String() string {
	// TODO: this cannot encode nonblock_recv(...) && nonblock_recv(...)
	return x.LHS.String() + operatorConversionTable[x.Op] + x.RHS.String()
}
func (x ir1TimeoutRecv) String() string {
	panic("timeout_recv cannot directly be expressed in NuSMV")
}
func (x ir1TimeoutPeek) String() string {
	panic("timeout_peek cannot directly be expressed in NuSMV")
}
func (x ir1NonblockRecv) String() string {
	panic("nonblock_recv cannot directly be expressed in NuSMV")
}
func (x ir1NonblockPeek) String() string {
	panic("nonblock_recv cannot directly be expressed in NuSMV")
}
func (x ir1ArrayLiteral) String() string {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x ir1HandshakeChannelVar) String() string { return x.RealName }
func (x ir1BufferedChannelVar) String() string  { return x.RealName }

func (x ir1ArrayLiteral) ArgString() (ret []string) {
	for _, elem := range x.Elems {
		ret = append(ret, elem.String())
	}
	return ret
}

// ========================================
// Transition

func assignByString(x ir1ExprObj, fromState, nextState intState, varName string) []intTransition {
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

func (x ir1PrimitiveVar) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x ir1ArrayVar) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("ArrayVar cannot directly be expressed in NuSMV")
}
func (x ir1Literal) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x ir1Not) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x ir1UnarySub) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x ir1Paren) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x ir1BinOp) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x ir1TimeoutRecv) Transition(fromState, nextState intState, varName string) []intTransition {
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
func (x ir1TimeoutPeek) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Not Implemented")
}
func (x ir1NonblockRecv) Transition(fromState, nextState intState, varName string) []intTransition {
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
func (x ir1NonblockPeek) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Not Implemented")
}
func (x ir1ArrayLiteral) Transition(fromState, nextState intState, varName string) []intTransition {
	panic("Array literals cannot directly be expressed in NuSMV")
}
func (x ir1HandshakeChannelVar) Transition(fromState, nextState intState, varName string) []intTransition {
	return assignByString(x, fromState, nextState, varName)
}
func (x ir1BufferedChannelVar) Transition(fromState, nextState intState, varName string) []intTransition {
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

func (x ir1PrimitiveVar) GetType() Type        { return x.Type }
func (x ir1ArrayVar) GetType() Type            { return x.RealLiteral.GetType() }
func (x ir1Literal) GetType() Type             { return x.Type }
func (x ir1Not) GetType() Type                 { return x.Sub.GetType() }
func (x ir1UnarySub) GetType() Type            { return x.Sub.GetType() }
func (x ir1Paren) GetType() Type               { return x.Sub.GetType() }
func (x ir1BinOp) GetType() Type               { return operatorResultType[x.Op] }
func (x ir1TimeoutRecv) GetType() Type         { return NamedType{"bool"} }
func (x ir1TimeoutPeek) GetType() Type         { return NamedType{"bool"} }
func (x ir1NonblockRecv) GetType() Type        { return NamedType{"bool"} }
func (x ir1NonblockPeek) GetType() Type        { return NamedType{"bool"} }
func (x ir1ArrayLiteral) GetType() Type        { return ArrayType{x.Elems[0].GetType()} }
func (x ir1HandshakeChannelVar) GetType() Type { return x.Type }
func (x ir1BufferedChannelVar) GetType() Type  { return x.Type }

// ========================================

func resolveRealObj(obj ir1ExprObj) ir1ExprObj {
	for {
		if primVarObj, isPrimVarObj := obj.(ir1PrimitiveVar); isPrimVarObj {
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
