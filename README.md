# Solguard Lab (DTL-v1) 

Symbolic distributed ledger infrastructure used as an audit fixture for Solguard tooling.

![banner](./assets/banner.png)

The project models a compact DLT stack with:

- Go ledger, consensus, network and bridge primitives.
- TypeScript RPC, relayer and indexer clients.
- Integration tests and operational documentation.

It is intentionally small enough to inspect, but structured like a real target repository so mapping, tracing, diffing and analysis tools have meaningful surfaces to enumerate.

## Commands

```bash
go test ./...
cd ts
bun install
bun test
```

## Layout

```text
cmd/dtld/           symbolic node entrypoint
internal/           Go implementation packages
pkg/types/          shared Go protocol types
tests/              Go integration tests
ts/                 TypeScript clients and tests
docs/               architecture and operations notes
```
