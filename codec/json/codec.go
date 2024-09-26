package json

import (
	"encoding/json"
	"reflect"

	"github.com/dogmatiq/marshalkit/codec"
	"github.com/dogmatiq/marshalkit/codec/internal/portablename"
)

// Codec is an implementation of marshalkit.Codec that uses Go's standard JSON
// implementation.
type Codec struct{}

// DefaultCodec is a marshalkit.Codec that marshals values using Go's standard
// JSON implementation.
var DefaultCodec = Codec{}

// Query returns the capabilities of the codec for the given types.
func (Codec) Query(types []reflect.Type) codec.Capabilities {
	caps := codec.Capabilities{
		Types: map[reflect.Type]string{},
	}

	for _, rt := range types {
		if n, ok := portablename.FromReflect(rt); ok {
			caps.Types[rt] = n
		}
	}

	return caps
}

// BasicMediaType returns the type and subtype portion of the media-type used to
// identify data encoded by this codec.
func (Codec) BasicMediaType() string {
	return "application/json"
}

// Marshal returns the binary representation of v.
func (Codec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal decodes a binary representation into v.
func (Codec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
