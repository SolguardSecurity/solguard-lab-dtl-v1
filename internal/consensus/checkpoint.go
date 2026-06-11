package consensus

import "github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"

type Checkpoint struct {
	Epoch            uint64     `json:"epoch"`
	Height           uint64     `json:"height"`
	BlockHash        types.Hash `json:"block_hash"`
	StateRoot        types.Hash `json:"state_root"`
	ValidatorSetRoot types.Hash `json:"validator_set_root"`
}

func NewCheckpoint(epoch uint64, block types.Block, set ValidatorSet) Checkpoint {
	return Checkpoint{
		Epoch:            epoch,
		Height:           block.Height,
		BlockHash:        block.Hash(),
		StateRoot:        block.StateRoot,
		ValidatorSetRoot: set.Root(),
	}
}

func (c Checkpoint) SigningRoot() types.Hash {
	return types.HashJSON(struct {
		Epoch            uint64     `json:"epoch"`
		Height           uint64     `json:"height"`
		BlockHash        types.Hash `json:"block_hash"`
		StateRoot        types.Hash `json:"state_root"`
		ValidatorSetRoot types.Hash `json:"validator_set_root"`
	}{c.Epoch, c.Height, c.BlockHash, c.StateRoot, c.ValidatorSetRoot})
}

type Signature struct {
	Signer types.Address `json:"signer"`
	Digest types.Hash    `json:"digest"`
}

func SignCheckpoint(signer types.Address, checkpoint Checkpoint) Signature {
	return Signature{Signer: signer, Digest: checkpoint.SigningRoot()}
}
