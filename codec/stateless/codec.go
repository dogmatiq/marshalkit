package stateless

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/dogmatiq/dogma"
	"github.com/dogmatiq/marshalkit/codec"
)

// Codec is an implementation of [marshalkit.Codec] that marshals
// [dogma.StatelessProcessRoot] values.
type Codec struct{}

// DefaultCodec is a marshalkit.Codec that marshals [dogma.StatelessProcessRoot]
// values.
var DefaultCodec = Codec{}

var processRootType = reflect.TypeOf(dogma.StatelessProcessRoot)

// Query returns the capabilities of the codec for the given types.
func (Codec) Query(types []reflect.Type) codec.Capabilities {
	caps := codec.Capabilities{}

	if slices.Contains(types, processRootType) {
		caps.Types = map[reflect.Type]string{
			processRootType: "process",
		}
	}

	return caps
}

// BasicMediaType returns the type and subtype portion of the media-type used to
// identify data encoded by this codec.
func (Codec) BasicMediaType() string {
	return "application/x-empty"
}

// Marshal returns the binary representation of v.
func (Codec) Marshal(v any) ([]byte, error) {
	if v == dogma.StatelessProcessRoot {
		return nil, nil
	}
	return nil, fmt.Errorf("%T is not dogma.StatelessProcessRoot", v)
}

// Unmarshal decodes a binary representation into v.
func (Codec) Unmarshal(data []byte, v any) error {
	if v != &dogma.StatelessProcessRoot {
		return fmt.Errorf("%T is not a pointer to dogma.StatelessProcessRoot", v)
	}
	if len(data) != 0 {
		return fmt.Errorf("expected empty data, got %d byte(s)", len(data))
	}
	return nil
}
