# Santalum language design

Santalum has three built-in injectable faults.

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

```
init {
  bossProcess : Boss(chQueue, chResp) @unstable,
}
```

### Parsing

TBD

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
