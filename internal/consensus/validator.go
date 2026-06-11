package consensus

import (
	"sort"

	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

type Validator struct {
	Address types.Address `json:"address"`
	Power   uint64        `json:"power"`
}

type ValidatorSet struct {
	Epoch      uint64      `json:"epoch"`
	Validators []Validator `json:"validators"`
}

func NewValidatorSet(epoch uint64, validators []Validator) ValidatorSet {
	copied := append([]Validator(nil), validators...)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i].Address < copied[j].Address
	})
	return ValidatorSet{Epoch: epoch, Validators: copied}
}

func (s ValidatorSet) Root() types.Hash {
	return types.HashJSON(s)
}

func (s ValidatorSet) TotalPower() uint64 {
	var total uint64
	for _, validator := range s.Validators {
		total += validator.Power
	}
	return total
}

func (s ValidatorSet) PowerOf(address types.Address) uint64 {
	for _, validator := range s.Validators {
		if validator.Address == address {
			return validator.Power
		}
	}
	return 0
}

func (s ValidatorSet) QuorumThreshold() uint64 {
	total := s.TotalPower()
	return (total * 2 / 3) + 1
}
