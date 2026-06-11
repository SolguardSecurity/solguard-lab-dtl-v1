package tests

import (
	"testing"

	"github.com/solguard-labs/solguard-dtl-test-v1/internal/consensus"
	"github.com/solguard-labs/solguard-dtl-test-v1/internal/ledger"
	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

func TestBlockCanBeFinalized(t *testing.T) {
	book := ledger.New(map[types.Address]uint64{"alice": 1000, "bob": 100})
	tx := types.NewTransaction("alice", "bob", 100, 1, "settlement")
	block, err := book.BuildBlock([]types.Transaction{tx}, "validator-1", 1700000000)
	if err != nil {
		t.Fatalf("build block: %v", err)
	}
	if err := book.ApplyBlock(block); err != nil {
		t.Fatalf("apply block: %v", err)
	}
	set := consensus.NewValidatorSet(1, []consensus.Validator{
		{Address: "validator-1", Power: 50},
		{Address: "validator-2", Power: 30},
		{Address: "validator-3", Power: 20},
	})
	checkpoint := consensus.NewCheckpoint(1, block, set)
	verifier := consensus.NewFinalityVerifier()
	verifier.RegisterValidatorSet(set)
	err = verifier.Verify(checkpoint, []consensus.Signature{
		consensus.SignCheckpoint("validator-1", checkpoint),
		consensus.SignCheckpoint("validator-2", checkpoint),
	})
	if err != nil {
		t.Fatalf("verify checkpoint: %v", err)
	}
}
