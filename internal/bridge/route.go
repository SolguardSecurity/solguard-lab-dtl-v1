package bridge

import (
	"errors"
	"fmt"

	"github.com/solguard-labs/solguard-dtl-test-v1/pkg/types"
)

var (
	ErrRouteDisabled = errors.New("route disabled")
	ErrWrongDomain   = errors.New("wrong domain")
)

type Domain string

type Route struct {
	ID                string        `json:"id"`
	SourceDomain      Domain        `json:"source_domain"`
	DestinationDomain Domain        `json:"destination_domain"`
	Executor          types.Address `json:"executor"`
	Enabled           bool          `json:"enabled"`
}

type Message struct {
	RouteID           string        `json:"route_id"`
	SourceDomain      Domain        `json:"source_domain"`
	DestinationDomain Domain        `json:"destination_domain"`
	Nonce             uint64        `json:"nonce"`
	Sender            types.Address `json:"sender"`
	Recipient         types.Address `json:"recipient"`
	PayloadHash       types.Hash    `json:"payload_hash"`
}

func (m Message) ID() types.Hash {
	return types.HashJSON(m)
}

func ValidateRoute(route Route, message Message) error {
	if !route.Enabled {
		return ErrRouteDisabled
	}
	if route.ID != message.RouteID || route.SourceDomain != message.SourceDomain || route.DestinationDomain != message.DestinationDomain {
		return fmt.Errorf("%w: route %s message %s", ErrWrongDomain, route.ID, message.RouteID)
	}
	return nil
}
