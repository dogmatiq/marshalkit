package codec

import "reflect"

// Codec is an interface for encoding and decoding values.
type Codec interface {
	// Query returns the capabilities of the codec for the given types.
	Query(types []reflect.Type) Capabilities

	// BasicMediaType returns the type and subtype portion of the media-type
	// used to identify data encoded by this codec.
	BasicMediaType() string

	// Marshal returns the binary representation of v.
	Marshal(v any) ([]byte, error)

	// Unmarshal decodes a binary representation into v.
	Unmarshal(data []byte, v any) error
}

// Capabilities describes the capabilities of a codec as it relates to
// specific Go types.
type Capabilities struct {
	// Types is a map of the supported types to their portable type name.
	Types map[reflect.Type]string
}
