package marshalkit

// A Marshaler marshals and unmarshals arbitrary Go values.
type Marshaler interface {
	// Marshal returns a binary representation of v.
	Marshal(v interface{}) (Packet, error)

	// Unmarshal produces a value from its binary representation.
	Unmarshal(p Packet) (interface{}, error)
}

// MustMarshal returns a binary representation of v.
// It panics if v can not be marshalled.
func MustMarshal(ma Marshaler, v interface{}) Packet {
	p, err := ma.Marshal(v)
	if err != nil {
		panic(PanicSentinel{err})
	}

	return p
}

// MustUnmarshal produces a value from its binary representation.
// It panics if p can not be marshalled.
func MustUnmarshal(ma Marshaler, p Packet) interface{} {
	v, err := ma.Unmarshal(p)
	if err != nil {
		panic(PanicSentinel{err})
	}

	return v
}
