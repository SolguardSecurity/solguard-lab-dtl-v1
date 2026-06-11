# Operations

## Local Checks

```bash
go test ./...
cd ts
bun install
bun test
```

## Running the Symbolic Node

```bash
go run ./cmd/dtld
```

The command applies one bootstrap transfer and prints the resulting ledger state.

## Fixture Usage

This repository is designed to be cloned as a target for audit automation:

- `solguard-map` should enumerate Go packages, TypeScript clients, bridge routes and consensus paths.
- `solguard-trace` should inspect critical functions such as checkpoint verification, block application and bridge execution.
- `solguard-diff` should rank recent changes that touch consensus, route execution or state transitions.
- `solguard-backend analyze` should combine deterministic tool outputs with knowledge retrieval and model-generated hypotheses.
