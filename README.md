# Sandal

Sandal is a fault-aware model checker for message passing systems.  
This repository is a forked version of [draftcode/sandal](https://github.com/draftcode/sandal).

## Installation

Sandal depends on NuSMV.

```bash
$ go get github.com/k0kubun/sandal
$ brew install homebrew/science/nusmv
```

## Usage

To see model checking result, you have to pipe sandal's stdout to NuSMV

```bash
$ sandal your_model.sandal | nusmv
```

### Fault Definition

Fault definition is an additional feature of this fork.
You can inject an unusual behavior to some statements non-deteministically.

```go
fault send(ch channel { bool }, val bool) @omission {
  // omit to send
}

proc Worker(ch channel { bool }) {
  send(ch, true) @omission
}
```

### Debug

You can dump AST, IR1 and IR2 with this forked version.

```bash
# Dump AST
$ sandal -a your_model.sandal

# Dump IR1
$ sandal -1 your_model.sandal

# Dump IR2
$ sandal -2 your_model.sandal
```

## Testing

```bash
$ go build
$ ./run_test.sh
```
