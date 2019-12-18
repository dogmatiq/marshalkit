package marshalkit

import "github.com/dogmatiq/dogma"

// MarshalMessage returns a binary representation of a message.
func MarshalMessage(ma ValueMarshaler, m dogma.Message) (Packet, error) {
	return ma.Marshal(m)
}

// UnmarshalMessage returns a message from its binary representation.
func UnmarshalMessage(ma ValueMarshaler, p Packet) (dogma.Message, error) {
	// Note: Unmarshal() returns interface{}, which works at the moment because
	// dogma.Message is also empty.
	//
	// If this fails to compile in the future, a branch needs to be added to
	// return a meaningful error if the unmarshaled value does not implement
	// dogma.Message.
	return ma.Unmarshal(p)
}

// MustMarshalMessage returns a binary representation of a message.
// It panics if the message can not be marshaled.
func MustMarshalMessage(ma ValueMarshaler, m dogma.Message) Packet {
	p, err := MarshalMessage(ma, m)
	if err != nil {
		panic(PanicSentinel{err})
	}

	return p
}

// MustUnmarshalMessage returns a message from its binary representation.
// It panics if the message can not be unmarshaled.
func MustUnmarshalMessage(ma ValueMarshaler, p Packet) dogma.Message {
	m, err := UnmarshalMessage(ma, p)
	if err != nil {
		panic(PanicSentinel{err})
	}

	return m
}
