# Sandal [![wercker status](https://app.wercker.com/status/f1cc50a0555ac54c4fdbd31bb15c3fba/s "wercker status")](https://app.wercker.com/project/bykey/f1cc50a0555ac54c4fdbd31bb15c3fba)

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

You can dump transition (experimental)

```
$ sandal -g your_model.sandal | dotdot -o -Tpng -o trans.png
```

## Testing

```bash
$ go build
$ ./run_test.sh
```

## Planned features

- You can define orignal faults for any processes, channels or functions
