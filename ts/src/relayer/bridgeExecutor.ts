import type { BridgeMessage, Route } from "../types";

export interface ExecutionReceipt {
  message_id: string;
  route_id: string;
  executed: boolean;
}

export class BridgeExecutor {
  private readonly routes = new Map<string, Route>();
  private readonly executedMessages = new Set<string>();

  constructor(routes: Route[] = []) {
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
    const messageId = executionKey(message);
    if (this.executedMessages.has(messageId)) {
      return { message_id: messageId, route_id: message.route_id, executed: false };
    }
    this.executedMessages.add(messageId);
    return { message_id: messageId, route_id: message.route_id, executed: true };
  }
}

export function executionKey(message: BridgeMessage): string {
  return [
    message.route_id,
    message.source_domain,
    message.destination_domain,
    message.nonce,
    message.sender,
    message.recipient,
    message.payload_hash,
  ].join(":");
}
