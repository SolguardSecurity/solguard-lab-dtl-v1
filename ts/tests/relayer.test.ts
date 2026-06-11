import { describe, expect, test } from "bun:test";
import { BridgeExecutor, executionKey, type BridgeMessage, type Route } from "../src";

describe("BridgeExecutor", () => {
  const route: Route = {
    id: "eth-mainnet:solguard-rollup",
    source_domain: "eth-mainnet",
    destination_domain: "solguard-rollup",
    executor: "executor-1",
    enabled: true,
  };

  test("executes a message once", () => {
    const message: BridgeMessage = {
      route_id: route.id,
      source_domain: route.source_domain,
      destination_domain: route.destination_domain,
      nonce: 1,
      sender: "alice",
      recipient: "bob",
      payload_hash: "payload-1",
    };
    const executor = new BridgeExecutor([route]);
    expect(executor.execute(message)).toEqual({
      message_id: executionKey(message),
      route_id: route.id,
      executed: true,
    });
    expect(executor.execute(message).executed).toBe(false);
  });

  test("rejects disabled routes", () => {
    const executor = new BridgeExecutor([{ ...route, enabled: false }]);
    expect(() =>
      executor.execute({
        route_id: route.id,
        source_domain: route.source_domain,
        destination_domain: route.destination_domain,
        nonce: 1,
        sender: "alice",
        recipient: "bob",
        payload_hash: "payload-1",
      }),
    ).toThrow("route disabled");
  });
});
