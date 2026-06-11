# RPC

The TypeScript client expects a small JSON HTTP API.

## Endpoints

```text
GET  /v1/status
GET  /v1/blocks/:height
GET  /v1/checkpoints/:height
POST /v1/transactions
```

## Status Response

```json
{
  "chain_id": "solguard-dtl-test-v1",
  "height": 10,
  "tip_hash": "hex",
  "state_root": "hex"
}
```

## Transaction Submission

```json
{
  "id": "hex",
  "from": "alice",
  "to": "bob",
  "amount": 100,
  "nonce": 1,
  "memo": "settlement"
}
```

The fixture does not include a production HTTP server. The RPC client is a typed boundary for tools that need to inspect API-facing code.
