package protobuf

import (
	"reflect"

	"github.com/dogmatiq/marshalkit/codec"
	"github.com/golang/protobuf/proto"
)

type commonCodec struct{}

// Query returns the capabilities of the codec for the given types.
func (commonCodec) Query(types []reflect.Type) codec.Capabilities {
	caps := codec.Capabilities{
		Types: map[reflect.Type]string{},
	}

	for _, rt := range types {
		z := reflect.Zero(rt).Interface()

		if m, ok := z.(proto.Message); ok {
			if n := proto.MessageName(m); n != "" {
				caps.Types[rt] = n
			}
		}
	}

	return caps
}
