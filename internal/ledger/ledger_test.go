package ledger

import (
	"testing"

	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

func TestLedgerAppliesBlock(t *testing.T) {
	book := New(map[types.Address]uint64{"alice": 100, "bob": 5})
	tx := types.NewTransaction("alice", "bob", 25, 1, "payment")
	block, err := book.BuildBlock([]types.Transaction{tx}, "alice", 1000)
	if err != nil {
		t.Fatalf("build block: %v", err)
	}
	if err := book.ApplyBlock(block); err != nil {
		t.Fatalf("apply block: %v", err)
	}
	if got := book.Balance("alice"); got != 75 {
		t.Fatalf("alice balance = %d", got)
	}
	if got := book.Balance("bob"); got != 30 {
		t.Fatalf("bob balance = %d", got)
	}
	if got := book.Nonce("alice"); got != 1 {
		t.Fatalf("alice nonce = %d", got)
	}
}

func TestLedgerRejectsBadStateRoot(t *testing.T) {
	book := New(map[types.Address]uint64{"alice": 100, "bob": 0})
	tx := types.NewTransaction("alice", "bob", 10, 1, "")
	block, err := book.BuildBlock([]types.Transaction{tx}, "alice", 1000)
	if err != nil {
		t.Fatalf("build block: %v", err)
	}
	block.StateRoot = types.Hash("bad-root")
	if err := book.ApplyBlock(block); err == nil {
		t.Fatal("expected bad state root to fail")
	}
}
