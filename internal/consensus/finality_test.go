package consensus

import (
	"testing"

	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

func TestFinalityVerifierAcceptsQuorum(t *testing.T) {
	set := NewValidatorSet(1, []Validator{
		{Address: "validator-1", Power: 40},
		{Address: "validator-2", Power: 35},
		{Address: "validator-3", Power: 25},
	})
	checkpoint := Checkpoint{
		Epoch:            1,
		Height:           10,
		BlockHash:        types.Hash("block"),
		StateRoot:        types.Hash("state"),
		ValidatorSetRoot: set.Root(),
	}
	verifier := NewFinalityVerifier()
	verifier.RegisterValidatorSet(set)
	signatures := []Signature{
		SignCheckpoint("validator-1", checkpoint),
		SignCheckpoint("validator-2", checkpoint),
	}
	if err := verifier.Verify(checkpoint, signatures); err != nil {
		t.Fatalf("verify checkpoint: %v", err)
	}
}

func TestFinalityVerifierRejectsUnknownSet(t *testing.T) {
	checkpoint := Checkpoint{Epoch: 7, Height: 1, ValidatorSetRoot: types.Hash("missing")}
	verifier := NewFinalityVerifier()
	if err := verifier.Verify(checkpoint, nil); err == nil {
		t.Fatal("expected unknown validator set")
	}
}
