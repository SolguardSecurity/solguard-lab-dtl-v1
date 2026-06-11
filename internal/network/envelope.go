package network

import (
	"errors"
	"time"

	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

var ErrExpiredEnvelope = errors.New("expired envelope")

type Topic string

const (
	TopicBlock      Topic = "block"
	TopicCheckpoint Topic = "checkpoint"
	TopicBridge     Topic = "bridge"
)

type Envelope struct {
	Topic       Topic         `json:"topic"`
	PeerID      string        `json:"peer_id"`
	PayloadHash types.Hash    `json:"payload_hash"`
	Sequence    uint64        `json:"sequence"`
	ExpiresAt   time.Time     `json:"expires_at"`
	Signer      types.Address `json:"signer"`
}

func (e Envelope) ID() types.Hash {
	return types.HashJSON(struct {
		Topic       Topic         `json:"topic"`
		PeerID      string        `json:"peer_id"`
		PayloadHash types.Hash    `json:"payload_hash"`
		Sequence    uint64        `json:"sequence"`
		Signer      types.Address `json:"signer"`
	}{e.Topic, e.PeerID, e.PayloadHash, e.Sequence, e.Signer})
}

func (e Envelope) Validate(now time.Time) error {
	if now.After(e.ExpiresAt) {
		return ErrExpiredEnvelope
	}
	return nil
}
