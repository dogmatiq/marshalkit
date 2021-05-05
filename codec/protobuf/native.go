package protobuf

import (
	"google.golang.org/protobuf/proto"
)

// NativeBasicMediaType is the type and subtype portion of the media-type used
// to identify data encoded using the native protocol buffers wire format.
const NativeBasicMediaType = "application/vnd.google.protobuf"

// DefaultNativeMarshaler is the text marshaler used by JSONCodec if none is
// provided.
var DefaultNativeMarshaler = proto.MarshalOptions{}

// DefaultNativeUnmarshaler is the text marshaler used by JSONCodec if none is
// provided.
var DefaultNativeUnmarshaler = proto.UnmarshalOptions{}

// DefaultNativeCodec is a marshalkit.Codec that marshals protocol buffers
// messages using the native protocol buffers wire format.
var DefaultNativeCodec = Codec{}
