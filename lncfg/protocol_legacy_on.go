// +build dev

package lncfg

// LegacyProtocol is a struct that we use to be able to test backwards
// compatibility of protocol additions, while defaulting to the latest within
// lnd.
type LegacyProtocol struct {
	// Onion if set to true, then we won't signal TLVOnionPayloadOptional.
	// As a result, nodes that include us in the route won't use the new
	// modern onion framing.
	Onion bool `long:"onion" description:"force node to not advertise the new modern TLV onion format"`
}

// LegacyOnion returns true if the old legacy onion format should be used when
// we're an intermediate or final hop. This controls if we set the
// TLVOnionPayloadOptional bit or not.
func (l *LegacyProtocol) LegacyOnion() bool {
	return l.Onion
}
