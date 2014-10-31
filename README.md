# Sandal [![wercker status](https://app.wercker.com/status/2f1d23a1e9e7678d0fc89a70afbb63a0/s "wercker status")](https://app.wercker.com/project/bykey/2f1d23a1e9e7678d0fc89a70afbb63a0)

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

## Testing

```bash
$ ./run_test.sh
```

## Planned features

- Fault statements are unified as `@` prefixed, i.e. `recv(ch) @timeout`
- You can define orignal faults for any processes, channels or functions
