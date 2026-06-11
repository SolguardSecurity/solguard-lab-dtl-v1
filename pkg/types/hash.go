package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Hash string

const ZeroHash Hash = "0000000000000000000000000000000000000000000000000000000000000000"

func HashBytes(parts ...[]byte) Hash {
	h := sha256.New()
	for _, part := range parts {
		h.Write(part)
	}
	return Hash(hex.EncodeToString(h.Sum(nil)))
}

func HashJSON(value any) Hash {
	payload, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return HashBytes(payload)
}

func (h Hash) Empty() bool {
	return h == "" || h == ZeroHash
}
