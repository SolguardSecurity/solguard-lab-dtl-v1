import { describe, expect, test } from "bun:test";
import { DtlRpcClient, type NodeStatus } from "../src";

describe("DtlRpcClient", () => {
  test("reads node status", async () => {
    const expected: NodeStatus = {
      chain_id: "solguard-dtl-test-v1",
      height: 7,
      tip_hash: "tip",
      state_root: "state",
    };
    const client = new DtlRpcClient("http://node.local", async (input) => {
      expect(input).toBe("http://node.local/v1/status");
      return Response.json(expected);
    });
    await expect(client.status()).resolves.toEqual(expected);
  });
});
