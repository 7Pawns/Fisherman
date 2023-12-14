package obfuscator

import (
	"Fisherman/server/baits"
	"encoding/base64"
)

type Obfuscator struct {
	bait baits.Bait
}

func NewObfuscator(bait baits.Bait) *Obfuscator {
	obfuscator := new(Obfuscator)
	obfuscator.bait = bait
	return obfuscator
}

func (obfuscator Obfuscator) Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(obfuscator.bait.BaitToString()))
}
