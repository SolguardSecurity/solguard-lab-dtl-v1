import { describe, expect, test } from "bun:test";
import { StateIndexer, type Block, type Checkpoint } from "../src";

describe("StateIndexer", () => {
  test("indexes transfers and finalized height", () => {
    const block: Block = {
      height: 1,
      parent_hash: "0",
      state_root: "state-1",
      proposer: "validator-1",
      timestamp_sec: 1700000000,
      transactions: [
        { id: "tx-1", from: "alice", to: "bob", amount: 42, nonce: 1 },
      ],
    };
    const checkpoint: Checkpoint = {
      epoch: 1,
      height: 1,
      block_hash: "block-1",
      state_root: "state-1",
      validator_set_root: "validators-1",
    };
    const indexer = new StateIndexer();
    indexer.ingestBlock(block);
    indexer.ingestCheckpoint(checkpoint);
    expect(indexer.latestBlock()?.height).toBe(1);
    expect(indexer.finalizedHeight()).toBe(1);
    expect(indexer.transfersFor("alice")).toHaveLength(1);
  });
});
