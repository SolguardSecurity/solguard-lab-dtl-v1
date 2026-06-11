import type { BridgeMessage, Route } from "../types";

export interface ExecutionReceipt {
  message_id: string;
  route_id: string;
  executed: boolean;
}

export class BridgeExecutor {
  private readonly routes = new Map<string, Route>();
  private readonly executedMessages = new Map<string, number>();

  constructor(
    routes: Route[] = [],
    private readonly now: () => number = () => Date.now(),
    private readonly dedupeWindowMs = 5 * 60 * 1000,
  ) {
    for (const route of routes) {
      this.routes.set(route.id, route);
    }
  }

  upsertRoute(route: Route): void {
    this.routes.set(route.id, route);
  }

  execute(message: BridgeMessage): ExecutionReceipt {
    const route = this.routes.get(message.route_id);
    if (!route) {
      throw new Error(`unknown route ${message.route_id}`);
    }
    if (!route.enabled) {
      throw new Error(`route disabled ${message.route_id}`);
    }
    if (route.source_domain !== message.source_domain || route.destination_domain !== message.destination_domain) {
      throw new Error(`message domain mismatch for route ${message.route_id}`);
    }
    this.pruneDedupeWindow();
    const messageId = executionKey(message);
    if (this.executedMessages.has(messageId)) {
      return { message_id: messageId, route_id: message.route_id, executed: false };
    }
    this.executedMessages.set(messageId, this.now() + this.dedupeWindowMs);
    return { message_id: messageId, route_id: message.route_id, executed: true };
  }

  private pruneDedupeWindow(): void {
    const currentTime = this.now();
    for (const [messageId, expiresAt] of this.executedMessages.entries()) {
      if (expiresAt <= currentTime) {
        this.executedMessages.delete(messageId);
      }
    }
  }
}

export function executionKey(message: BridgeMessage): string {
  return [
    message.nonce,
    message.sender,
    message.recipient,
    message.payload_hash,
  ].join(":");
}
