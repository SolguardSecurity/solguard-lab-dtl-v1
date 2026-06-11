package ledger

import (
	"errors"
	"fmt"
	"sort"

	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidNonce        = errors.New("invalid nonce")
	ErrInvalidBlock        = errors.New("invalid block")
)

type Ledger struct {
	accounts map[types.Address]types.Account
	blocks   []types.Block
}

func New(genesis map[types.Address]uint64) *Ledger {
	accounts := make(map[types.Address]types.Account, len(genesis))
	for address, balance := range genesis {
		accounts[address] = types.Account{Address: address, Balance: balance}
	}
	return &Ledger{accounts: accounts}
}

func (l *Ledger) Clone() *Ledger {
	accounts := make(map[types.Address]types.Account, len(l.accounts))
	for address, account := range l.accounts {
		accounts[address] = account
	}
	blocks := append([]types.Block(nil), l.blocks...)
	return &Ledger{accounts: accounts, blocks: blocks}
}

func (l *Ledger) Balance(address types.Address) uint64 {
	return l.accounts[address].Balance
}

func (l *Ledger) Nonce(address types.Address) uint64 {
	return l.accounts[address].Nonce
}

func (l *Ledger) Height() uint64 {
	return uint64(len(l.blocks))
}

func (l *Ledger) TipHash() types.Hash {
	if len(l.blocks) == 0 {
		return types.ZeroHash
	}
	return l.blocks[len(l.blocks)-1].Hash()
}

func (l *Ledger) Blocks() []types.Block {
	return append([]types.Block(nil), l.blocks...)
}

func (l *Ledger) BuildBlock(txs []types.Transaction, proposer types.Address, timestampSec int64) (types.Block, error) {
	next := l.Clone()
	for _, tx := range txs {
		if err := next.applyTransaction(tx); err != nil {
			return types.Block{}, err
		}
	}
	return types.Block{
		Height:       l.Height() + 1,
		ParentHash:   l.TipHash(),
		StateRoot:    next.StateRoot(),
		Proposer:     proposer,
		TimestampSec: timestampSec,
		Transactions: append([]types.Transaction(nil), txs...),
	}, nil
}

func (l *Ledger) ApplyBlock(block types.Block) error {
	if block.Height != l.Height()+1 {
		return fmt.Errorf("%w: expected height %d got %d", ErrInvalidBlock, l.Height()+1, block.Height)
	}
	if block.ParentHash != l.TipHash() {
		return fmt.Errorf("%w: parent hash mismatch", ErrInvalidBlock)
	}
	next := l.Clone()
	for _, tx := range block.Transactions {
		if err := next.applyTransaction(tx); err != nil {
			return err
		}
	}
	if next.StateRoot() != block.StateRoot {
		return fmt.Errorf("%w: state root mismatch", ErrInvalidBlock)
	}
	l.accounts = next.accounts
	l.blocks = append(l.blocks, block)
	return nil
}

func (l *Ledger) StateRoot() types.Hash {
	accounts := make([]types.Account, 0, len(l.accounts))
	for _, account := range l.accounts {
		accounts = append(accounts, account)
	}
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].Address < accounts[j].Address
	})
	return types.HashJSON(accounts)
}

func (l *Ledger) Snapshot() []types.Account {
	accounts := make([]types.Account, 0, len(l.accounts))
	for _, account := range l.accounts {
		accounts = append(accounts, account)
	}
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].Address < accounts[j].Address
	})
	return accounts
}

func (l *Ledger) applyTransaction(tx types.Transaction) error {
	if tx.Amount == 0 {
		return fmt.Errorf("%w: zero amount", ErrInvalidBlock)
	}
	from := l.accounts[tx.From]
	if from.Address == "" {
		from.Address = tx.From
	}
	if tx.Nonce != from.Nonce+1 {
		return fmt.Errorf("%w: expected %d got %d", ErrInvalidNonce, from.Nonce+1, tx.Nonce)
	}
	if from.Balance < tx.Amount {
		return ErrInsufficientBalance
	}
	to := l.accounts[tx.To]
	if to.Address == "" {
		to.Address = tx.To
	}
	from.Balance -= tx.Amount
	from.Nonce++
	to.Balance += tx.Amount
	l.accounts[tx.From] = from
	l.accounts[tx.To] = to
	return nil
}
