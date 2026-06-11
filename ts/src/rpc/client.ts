import type { Block, Checkpoint, NodeStatus, Transaction } from "../types";

export interface FetchLike {
  (input: string, init?: RequestInit): Promise<Response>;
}

export class DtlRpcClient {
  private readonly baseUrl: string;
  private readonly fetcher: FetchLike;

  constructor(baseUrl: string, fetcher: FetchLike = fetch) {
    this.baseUrl = baseUrl.replace(/\/+$/, "");
    this.fetcher = fetcher;
  }

  async status(): Promise<NodeStatus> {
    return this.getJson<NodeStatus>("/v1/status");
  }

  async block(height: number): Promise<Block> {
    return this.getJson<Block>(`/v1/blocks/${height}`);
  }

  async checkpoint(height: number): Promise<Checkpoint> {
    return this.getJson<Checkpoint>(`/v1/checkpoints/${height}`);
  }

  async submitTransaction(tx: Transaction): Promise<{ accepted: boolean; tx_id: string }> {
    return this.postJson("/v1/transactions", tx);
  }

  private async getJson<T>(path: string): Promise<T> {
    const response = await this.fetcher(`${this.baseUrl}${path}`);
    return decodeJson<T>(response);
  }

  private async postJson<TBody, TResult>(path: string, body: TBody): Promise<TResult> {
    const response = await this.fetcher(`${this.baseUrl}${path}`, {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify(body),
    });
    return decodeJson<TResult>(response);
  }
}

async function decodeJson<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const text = await response.text();
    throw new Error(`rpc request failed: ${response.status} ${text}`);
  }
  return (await response.json()) as T;
}
