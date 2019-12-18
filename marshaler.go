package marshalkit

// Marshaler is a marshaler that can marshal types and values.
type Marshaler interface {
	TypeMarshaler
	ValueMarshaler
}
