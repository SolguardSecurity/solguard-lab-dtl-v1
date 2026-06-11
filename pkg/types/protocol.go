package types

type Address string

type Transaction struct {
	ID     Hash    `json:"id"`
	From   Address `json:"from"`
	To     Address `json:"to"`
	Amount uint64  `json:"amount"`
	Nonce  uint64  `json:"nonce"`
	Memo   string  `json:"memo,omitempty"`
}

func NewTransaction(from Address, to Address, amount uint64, nonce uint64, memo string) Transaction {
	tx := Transaction{
		From:   from,
		To:     to,
		Amount: amount,
		Nonce:  nonce,
		Memo:   memo,
	}
	tx.ID = HashJSON(struct {
		From   Address `json:"from"`
		To     Address `json:"to"`
		Amount uint64  `json:"amount"`
		Nonce  uint64  `json:"nonce"`
		Memo   string  `json:"memo,omitempty"`
	}{tx.From, tx.To, tx.Amount, tx.Nonce, tx.Memo})
	return tx
}

type Block struct {
	Height       uint64        `json:"height"`
	ParentHash   Hash          `json:"parent_hash"`
	StateRoot    Hash          `json:"state_root"`
	Proposer     Address       `json:"proposer"`
	TimestampSec int64         `json:"timestamp_sec"`
	Transactions []Transaction `json:"transactions"`
}

func (b Block) Hash() Hash {
	return HashJSON(struct {
		Height       uint64        `json:"height"`
		ParentHash   Hash          `json:"parent_hash"`
		StateRoot    Hash          `json:"state_root"`
		Proposer     Address       `json:"proposer"`
		TimestampSec int64         `json:"timestamp_sec"`
		Transactions []Transaction `json:"transactions"`
	}{
		Height:       b.Height,
		ParentHash:   b.ParentHash,
		StateRoot:    b.StateRoot,
		Proposer:     b.Proposer,
		TimestampSec: b.TimestampSec,
		Transactions: b.Transactions,
	})
}

type Account struct {
	Address Address `json:"address"`
	Balance uint64  `json:"balance"`
	Nonce   uint64  `json:"nonce"`
}
