# Architecture

`solguard-dtl-test-v1` models a compact distributed ledger stack for audit tooling fixtures.

## Components

- `cmd/dtld` starts a symbolic node and prints the current ledger view.
- `internal/ledger` applies blocks, tracks balances and verifies state roots.
- `internal/consensus` tracks validator sets and verifies checkpoint finality.
- `internal/network` models signed gossip envelopes.
- `internal/bridge` models domain routes and cross-domain messages.
- `ts/src/rpc` provides a minimal HTTP client.
- `ts/src/relayer` executes bridge messages against configured routes.
- `ts/src/indexer` builds a queryable view from blocks and checkpoints.

## Data Flow

1. Transactions are built by clients and submitted to the node.
2. The ledger builds a block from ordered transactions.
3. Validators sign a checkpoint for a block and state root.
4. The finality verifier accepts a checkpoint when enough validator power signs it.
5. Indexers consume blocks and checkpoints for downstream queries.
6. Relayers execute bridge messages against enabled routes.

The project intentionally avoids external services so tooling can run it offline.
