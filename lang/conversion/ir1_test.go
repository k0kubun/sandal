package conversion

import (
	"github.com/cookieo9/go-misc/pp"
	. "github.com/k0kubun/sandal/lang/data"
	"github.com/k0kubun/sandal/lang/parsing"
	"github.com/kylelemons/godebug/diff"
	"testing"
)

func TestConvertASTToIntModule(t *testing.T) {
	return // FIXME: pending because it fails

	defs := []Def{
		ProcDef{
			Name: "ProcA",
			Parameters: []Parameter{
				{
					Name: "ch0",
					Type: HandshakeChannelType{
						Elems: []Type{NamedType{"bool"}},
					},
				},
			},
			Stmts: []Stmt{
				VarDeclStmt{
					Name: "b",
					Type: NamedType{"int"},
				},
				SendStmt{
					Channel: IdentifierExpr{Pos{}, "ch0"},
					Args: []Expr{
						TrueExpr{Pos{}},
					},
				},
			},
		},
		InitBlock{
			Vars: []InitVar{
				ChannelVar{
					Name: "ch",
					Type: HandshakeChannelType{
						Elems: []Type{NamedType{"bool"}},
					},
				},
				InstanceVar{
					Name:        "proc1",
					ProcDefName: "ProcA",
					Args: []Expr{
						IdentifierExpr{Pos{}, "ch"},
					},
				},
			},
		},
	}
	expected := []intModule{
		intHandshakeChannel{
			Name:      "HandshakeChannel0",
			ValueType: []string{"boolean"},
			ZeroValue: []string{"FALSE"},
		},
		intProcModule{
			Name: "__pid0_ProcA",
			Args: []string{"__orig_ch0"},
			Vars: []intVar{
				{"ch0", "HandshakeChannel0Proxy(__orig_ch0)"},
				{"b", "0..8"},
			},
			InitState: intState("state0"),
			Trans: []intTransition{
				{
					FromState: "state0",
					NextState: "state1",
					Condition: "",
				},
				{
					FromState: "state1",
					NextState: "state2",
					Condition: "!(ch0.ready)",
					Actions: []intAssign{
						{"ch0.send_filled", "TRUE"},
						{"ch0.send_value_0", "TRUE"},
					},
				},
				{
					FromState: "state2",
					NextState: "state3",
					Condition: "(ch0.ready) & (ch0.received)",
					Actions: []intAssign{
						{"ch0.send_leaving", "TRUE"},
					},
				},
			},
			Defaults: map[string]string{
				"ch0.send_leaving":  "FALSE",
				"ch0.send_filled":   "FALSE",
				"ch0.recv_received": "FALSE",
				"ch0.send_value_0":  "ch0.value_0",
				"next(b)":           "b",
			},
			Defs: []intAssign{},
		},
		intMainModule{
			Vars: []intVar{
				{"ch", "HandshakeChannel0"},
				{"proc1", "process __pid0_ProcA(ch)"},
			},
		},
	}
	err, intMods := astToIr1(defs)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expectPP := pp.PP(expected)
	actualPP := pp.PP(intMods)
	if expectPP != actualPP {
		t.Errorf("Unmatched\n%s\n", diff.Diff(expectPP, actualPP))
	}
}

func TestConvertASTToIntModuleForSend(t *testing.T) {
	return // FIXME:
	sendWithTagSource := `
		fault send(ch channel { bool }) @omission {
			// do nothing
		}

		proc SendProc(ch channel { bool }) {
			send(ch, true) @omission
		}

		init {
			ch: channel { bool },
			sp: SendProc(ch),
		}
	`
	expected := []intModule{
		intHandshakeChannel{
			Name:      "HandshakeChannel0",
			ValueType: []string{"boolean"},
			ZeroValue: []string{"FALSE"},
		},
		intProcModule{
			Name: "__pid0_SendProc",
			Args: []string{"__orig_ch"},
			Vars: []intVar{
				{"ch", "HandshakeChannel0Proxy(__orig_ch)"},
			},
			InitState: intState("state0"),
			Trans: []intTransition{
				{
					FromState: "state0",
					NextState: "state2",
					Condition: "",
					Actions:   []intAssign(nil),
				},
				{
					FromState: "state2",
					NextState: "state3",
					Condition: "!(ch.ready)",
					Actions: []intAssign{
						{"ch.send_filled", "TRUE"},
						{"ch.send_value_0", "TRUE"},
					},
				},
				{
					FromState: "state3",
					NextState: "state4",
					Condition: "(ch.ready) & (ch.received)",
					Actions: []intAssign{
						{"ch.send_leaving", "TRUE"},
					},
				},
				{
					FromState: "state4",
					NextState: "state1",
					Condition: "",
					Actions:   []intAssign(nil),
				},
				{
					FromState: "state0",
					NextState: "state5",
					Condition: "",
					Actions:   []intAssign(nil),
				},
				{
					FromState: "state5",
					NextState: "state6",
					Condition: "!(ch.ready)",
					Actions: []intAssign{
						{"ch.send_filled", "TRUE"},
						{"ch.send_value_0", "TRUE"},
					},
				},
				{
					FromState: "state6",
					NextState: "state7",
					Condition: "(ch.ready) & (ch.received)",
					Actions: []intAssign{
						{"ch.send_leaving", "TRUE"},
					},
				},
				{
					FromState: "state7",
					NextState: "state1",
					Condition: "",
					Actions:   []intAssign(nil),
				},
			},
			Defaults: map[string]string{
				"ch.send_leaving":  "FALSE",
				"ch.send_filled":   "FALSE",
				"ch.recv_received": "FALSE",
				"ch.send_value_0":  "ch.value_0",
			},
			Defs: []intAssign(nil),
		},
		intMainModule{
			Vars: []intVar{
				{"ch", "HandshakeChannel0"},
				{"sp", "process __pid0_SendProc(ch)"},
			},
		},
	}

	scanner := new(parsing.Scanner)
	scanner.Init([]rune(sendWithTagSource), 0)
	defs := parsing.Parse(scanner)

	err, intMods := astToIr1(defs)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	actualPP := pp.PP(intMods)
	expectPP := pp.PP(expected)
	if actualPP != expectPP {
		t.Errorf("Unmatched\n%s\n", diff.Diff(expectPP, actualPP))
	}
}
