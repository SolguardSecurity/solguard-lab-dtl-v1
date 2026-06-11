export type Hash = string;
export type Address = string;
export type Domain = string;

export interface Transaction {
  id: Hash;
  from: Address;
  to: Address;
  amount: number;
  nonce: number;
  memo?: string;
}

export interface Block {
  height: number;
  parent_hash: Hash;
  state_root: Hash;
  proposer: Address;
  timestamp_sec: number;
  transactions: Transaction[];
}

export interface Validator {
  address: Address;
  power: number;
}

export interface ValidatorSet {
  epoch: number;
  validators: Validator[];
}

export interface Checkpoint {
  epoch: number;
  height: number;
  block_hash: Hash;
  state_root: Hash;
  validator_set_root: Hash;
}

export interface Route {
  id: string;
  source_domain: Domain;
  destination_domain: Domain;
  executor: Address;
  enabled: boolean;
}

export interface BridgeMessage {
  route_id: string;
  source_domain: Domain;
  destination_domain: Domain;
  nonce: number;
  sender: Address;
  recipient: Address;
  payload_hash: Hash;
}

export interface NodeStatus {
  chain_id: string;
  height: number;
  tip_hash: Hash;
  state_root: Hash;
}
