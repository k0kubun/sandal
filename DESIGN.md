# Santalum language design

Santalum has three built-in injectable faults.

- Timeout Faults
- Abrupt Termination Faults
- Message Drop Faults

## Timeout faults

```
received = timeout_recv(chResp, doneTask)
```

## Abrupt Termination Faults

```
init {
  bossProcess : Boss(chQueue, chResp) @unstable,
}
```

## Message Drop Faults

```
init {
  chQueue : channel { Task } @drop,
}
```
