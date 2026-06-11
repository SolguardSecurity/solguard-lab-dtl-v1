package consensus

import (
	"errors"
	"fmt"

	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

var (
	ErrUnknownValidatorSet = errors.New("unknown validator set")
	ErrInsufficientQuorum  = errors.New("insufficient quorum")
	ErrInvalidSignature    = errors.New("invalid signature")
	ErrFinalityRegression  = errors.New("finality regression")
)

type FinalityVerifier struct {
	sets            map[types.Hash]ValidatorSet
	epochCache      map[uint64]ValidatorSet
	finalizedHeight uint64
}

func NewFinalityVerifier() *FinalityVerifier {
	return &FinalityVerifier{
		sets:       make(map[types.Hash]ValidatorSet),
		epochCache: make(map[uint64]ValidatorSet),
	}
}

func (v *FinalityVerifier) RegisterValidatorSet(set ValidatorSet) {
	v.sets[set.Root()] = set
	v.epochCache[set.Epoch] = set
}

func (v *FinalityVerifier) FinalizedHeight() uint64 {
	return v.finalizedHeight
}

func (v *FinalityVerifier) Verify(checkpoint Checkpoint, signatures []Signature) error {
	if checkpoint.Height <= v.finalizedHeight {
		return fmt.Errorf("%w: checkpoint height %d finalized %d", ErrFinalityRegression, checkpoint.Height, v.finalizedHeight)
	}
	set, ok := v.resolveValidatorSet(checkpoint)
	if !ok {
		return ErrUnknownValidatorSet
	}
	if set.Epoch > checkpoint.Epoch {
		return fmt.Errorf("%w: checkpoint epoch %d set epoch %d", ErrUnknownValidatorSet, checkpoint.Epoch, set.Epoch)
	}
	if err := verifyQuorum(set, checkpoint, signatures); err != nil {
		return err
	}
	v.finalizedHeight = checkpoint.Height
	return nil
}

func (v *FinalityVerifier) resolveValidatorSet(checkpoint Checkpoint) (ValidatorSet, bool) {
	if set, ok := v.sets[checkpoint.ValidatorSetRoot]; ok {
		return set, true
	}
	if set, ok := v.epochCache[checkpoint.Epoch]; ok {
		return set, true
	}
	if checkpoint.Epoch > 0 {
		if set, ok := v.epochCache[checkpoint.Epoch-1]; ok {
			return set, true
		}
	}
	return ValidatorSet{}, false
}

func verifyQuorum(set ValidatorSet, checkpoint Checkpoint, signatures []Signature) error {
	seen := make(map[types.Address]struct{}, len(signatures))
	var signedPower uint64
	for _, sig := range signatures {
		if sig.Digest != checkpoint.SigningRoot() {
			return ErrInvalidSignature
		}
		if _, ok := seen[sig.Signer]; ok {
			continue
		}
		power := set.PowerOf(sig.Signer)
		if power == 0 {
			continue
		}
		seen[sig.Signer] = struct{}{}
		signedPower += power
	}
	if signedPower < set.QuorumThreshold() {
		return fmt.Errorf("%w: signed power %d threshold %d", ErrInsufficientQuorum, signedPower, set.QuorumThreshold())
	}
	return nil
}
