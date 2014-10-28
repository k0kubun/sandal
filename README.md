# Santalum

Extensible fault-aware model checking language forked from [Sandal](https://github.com/draftcode/sandal).

## Features

### Sandal-derived features

These features are derived from Sandal.

- Simple syntax like Go language
- Easy fault injection without changing semantics

### Extended features

Santalum has some extra features Sandal does not have.

- Fault statements are unified as `@` prefixed, like `recv @timeout`
- You can define orignal faults for any processes for functions

## Installation

Santalum depends on NuSMV.

```bash
$ go get github.com/k0kubun/santalum
$ brew install homebrew/science/nusmv
```

## Usage

To see model checking result, you have to pipe Santalum's stdout to NuSMV

```bash
$ santalum your_model.sant | nusmv
```

## Model examples

```go
proc ProcA(ch0 channel {bool}) {
  var b int
  send(ch0, true)
}

init {
  ch:    channel {bool},
  proc1: ProcA(ch),
}
```
