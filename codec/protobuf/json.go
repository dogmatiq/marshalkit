package protobuf

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// JSONBasicMediaType is the type and subtype portion of the media-type used to
// identify data encoded using JSON.
const JSONBasicMediaType = "application/vnd.google.protobuf+json"

// DefaultJSONMarshaler is the text marshaler used by JSONCodec if none is
// provided.
var DefaultJSONMarshaler = protojson.MarshalOptions{
	UseProtoNames: true,
}

// DefaultJSONUnmarshaler is the text marshaler used by JSONCodec if none is
// provided.
var DefaultJSONUnmarshaler = protojson.UnmarshalOptions{
	DiscardUnknown: true,
}

// DefaultJSONCodec is a marshalkit.Codec that marshals protocol buffers
// messages in JSON format.
var DefaultJSONCodec = Codec{
	MediaType:   JSONBasicMediaType,
	Marshaler:   DefaultJSONMarshaler,
	Unmarshaler: DefaultJSONUnmarshaler,
}
