# Sandal language notes

Sandal has three built-in injectable faults.

- Abrupt Termination Faults
- Message Drop Faults
- Timeout Faults

TODO:

- Byzantine Generals Problem
  - Omission fault
  - Comission fault

Compile steps:

- Parsing
- TypeCheck
- Conversion

## Abrupt Termination Faults

Abrupt Termination is a fault which is injected to process or channel.

```
init {
  unstableProc : Boss(chQueue, chResp) @unstable,
}

init {
  unstableCh : channel { bool } @unstable,
}
```

### Parsing

```go
// Input
init { a : M(b) @unstable }

// Output
[]Definition{
	InitBlock{
		Pos{1, 1},
		[]InitVar{
			InstanceVar{
				Pos{1, 8},
				"a",
				"M",
				[]Expression{IdentifierExpression{Pos{1, 14}, "b"}},
				[]string{"unstable"},
			},
		},
	},
},
```

```go
// Input
init { a : channel { bool } @unstable }

// Output
[]Definition{
	InitBlock{
		Pos{1, 1},
		[]InitVar{
			ChannelVar{
				Pos{1, 8},
				"a",
				HandshakeChannelType{[]Type{NamedType{"bool"}}},
				[]string{"unstable"},
			},
		},
	},
},
```

### TypeCheck

TBD

### Conversion

TBD

## Message Drop Faults

```
init {
  chQueue : channel { Task } @drop,
}
```

## Timeout faults

```
received = timeout_recv(chResp, doneTask)
```
