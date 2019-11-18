package marshalkit

import (
	"github.com/dogmatiq/dogma"
)

// MarshalMessage returns a binary representation of a message.
func MarshalMessage(ma *Marshaler, m dogma.Message) (Packet, error) {
	return ma.Marshal(m)
}

// MustMarshalMessage returns a binary representation of a message.
// It panics if marshaling fails.
func MustMarshalMessage(ma *Marshaler, m dogma.Message) Packet {
	p, err := ma.Marshal(m)
	Must(err)
	return p
}

// UnmarshalMessage returns a message from its binary representation.
func UnmarshalMessage(ma *Marshaler, p Packet) (dogma.Message, error) {
	// Note: Unmarshal() returns interface{}, which works at the moment because
	// dogma.Message is also empty.
	//
	// If this fails to compile in the future, a branch needs to be added to
	// return a meaningful error if the unmarshaled value does not implement
	// dogma.Message.
	return ma.Unmarshal(p)
}

// MustUnmarshalMessage returns a message from its binary representation.
// It panics if unmarshaling fails.
func MustUnmarshalMessage(ma *Marshaler, p Packet) dogma.Message {
	m, err := UnmarshalMessage(ma, p)
	Must(err)
	return m
}
