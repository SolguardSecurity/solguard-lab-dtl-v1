import type { Block, Checkpoint, Hash, Transaction } from "../types";

export interface IndexedTransfer {
  height: number;
  tx_id: Hash;
  from: string;
  to: string;
  amount: number;
}

export class StateIndexer {
  private readonly blocks = new Map<number, Block>();
  private readonly checkpoints = new Map<number, Checkpoint>();
  private readonly transfers: IndexedTransfer[] = [];

  ingestBlock(block: Block): void {
    if (this.blocks.has(block.height)) {
      return;
    }
    this.blocks.set(block.height, block);
    for (const tx of block.transactions) {
      this.indexTransfer(block.height, tx);
    }
  }

  ingestCheckpoint(checkpoint: Checkpoint): void {
    const existing = this.checkpoints.get(checkpoint.height);
    if (existing && existing.block_hash !== checkpoint.block_hash) {
      throw new Error(`conflicting checkpoint at height ${checkpoint.height}`);
    }
    this.checkpoints.set(checkpoint.height, checkpoint);
  }

  latestBlock(): Block | undefined {
    const heights = [...this.blocks.keys()].sort((a, b) => b - a);
    return heights.length > 0 ? this.blocks.get(heights[0]!) : undefined;
  }

  finalizedHeight(): number {
    const heights = [...this.checkpoints.keys()].sort((a, b) => b - a);
    return heights[0] ?? 0;
  }

  transfersFor(address: string): IndexedTransfer[] {
    return this.transfers.filter((transfer) => transfer.from === address || transfer.to === address);
  }

  private indexTransfer(height: number, tx: Transaction): void {
    this.transfers.push({
      height,
      tx_id: tx.id,
      from: tx.from,
      to: tx.to,
      amount: tx.amount,
    });
  }
}
