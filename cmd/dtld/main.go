package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/solguard-labs/solguard-dtl-test-v1/internal/ledger"
	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

func main() {
	book := ledger.New(map[types.Address]uint64{
		"validator-1": 1_000_000,
		"alice":       500_000,
		"bob":         100_000,
	})
	tx := types.NewTransaction("alice", "bob", 1_250, 1, "bootstrap transfer")
	block, err := book.BuildBlock([]types.Transaction{tx}, "validator-1", time.Now().Unix())
	if err != nil {
		exit(err)
	}
	if err := book.ApplyBlock(block); err != nil {
		exit(err)
	}
	encoded, err := json.MarshalIndent(struct {
		Height    uint64          `json:"height"`
		TipHash   types.Hash      `json:"tip_hash"`
		StateRoot types.Hash      `json:"state_root"`
		Accounts  []types.Account `json:"accounts"`
	}{book.Height(), book.TipHash(), book.StateRoot(), book.Snapshot()}, "", "  ")
	if err != nil {
		exit(err)
	}
	fmt.Println(string(encoded))
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
