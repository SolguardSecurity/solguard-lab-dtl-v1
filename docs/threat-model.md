# Threat Model

## Assets

- Ledger balances and account nonces.
- Checkpoint finality decisions.
- Validator set membership and power.
- Bridge route execution guarantees.
- Indexed views consumed by downstream agents.

## Trust Boundaries

- RPC clients may submit malformed transactions.
- Network peers may gossip stale or conflicting envelopes.
- Relayers may receive messages from multiple domains and routes.
- Validator sets may rotate between epochs.

## Review Targets

- State transition validation.
- Quorum accounting.
- Validator set lookup and rotation behavior.
- Replay and idempotency handling.
- Cross-domain route validation.
- Indexer conflict handling.

The fixture focuses on deterministic code paths so automated tools can produce stable results across repeated runs.
