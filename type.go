package marshalkit

import "reflect"

// A TypeMarshaler marshals and unmarshals Go types to and from "portable"
// string representations.
type TypeMarshaler interface {
	// MarshalType marshals a type to its portable representation.
	MarshalType(rt reflect.Type) (string, error)

	// UnmarshalType marshals a type from its portable representation.
	UnmarshalType(n string) (reflect.Type, error)
}

// MustMarshalType marshals rt to its portable representation.
// It panics if the type can not be marshaled.
func MustMarshalType(ma TypeMarshaler, rt reflect.Type) string {
	n, err := ma.MarshalType(rt)
	if err != nil {
		panic(PanicSentinel{err})
	}

	return n
}

// MustUnmarshalType unmarshals a typefrom its portable representation.
// It panics if the type can not be unmarshaled.
func MustUnmarshalType(ma TypeMarshaler, n string) reflect.Type {
	t, err := ma.UnmarshalType(n)
	if err != nil {
		panic(PanicSentinel{err})
	}

	return t
}
