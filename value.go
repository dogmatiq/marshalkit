package marshalkit

import (
	"reflect"
)

// A ValueMarshaler marshals and unmarshals arbitrary Go values.
type ValueMarshaler interface {
	// Marshal returns a binary representation of v.
	Marshal(v interface{}) (Packet, error)

	// Unmarshal produces a value from its binary representation.
	Unmarshal(p Packet) (interface{}, error)

	// MediaTypesFor returns the list of supported media types for unmarshaling
	// the given type.
	MediaTypesFor(reflect.Type) []string
}

// MustMarshal returns a binary representation of v.
// It panics if v can not be marshalled.
func MustMarshal(ma ValueMarshaler, v interface{}) Packet {
	p, err := ma.Marshal(v)
	if err != nil {
		panic(PanicSentinel{err})
	}

	return p
}

// MustUnmarshal produces a value from its binary representation.
// It panics if p can not be unmarshalled.
func MustUnmarshal(ma ValueMarshaler, p Packet) interface{} {
	v, err := ma.Unmarshal(p)
	if err != nil {
		panic(PanicSentinel{err})
	}

	return v
}
