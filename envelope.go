package marshalkit

import (
	"time"

	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/dogma"
	"github.com/dogmatiq/interopspec/envelopespec"
)

// MustMarshalMessageIntoEnvelope marshals a Dogma message into an envelope.
func MustMarshalMessageIntoEnvelope(
	vm ValueMarshaler,
	m dogma.Message,
	env *envelopespec.Envelope,
) {
	p := MustMarshal(vm, m)

	_, n, err := p.ParseMediaType()
	if err != nil {
		// CODE COVERAGE: This branch would require the marshaler to violate its
		// own requirements on the format of the media-type.
		panic(err)
	}

	env.PortableName = n
	env.MediaType = p.MediaType
	env.Data = p.Data
}

// UnmarshalMessageFromEnvelope unmarshals a Dogma message from an envelope.
func UnmarshalMessageFromEnvelope(
	vm ValueMarshaler,
	env *envelopespec.Envelope,
) (dogma.Message, error) {
	p := Packet{
		MediaType: env.MediaType,
		Data:      env.Data,
	}

	return vm.Unmarshal(p)
}

// MustMarshalEnvelopeIdentity marshals id to its protocol buffers
// representation, as used within message envelopes.
func MustMarshalEnvelopeIdentity(id configkit.Identity) *envelopespec.Identity {
	return &envelopespec.Identity{
		Name: id.Name,
		Key:  id.Key,
	}
}

// UnmarshalEnvelopeIdentity unmarshals id from its protocol buffers
// representation, as used within message envelopes.
func UnmarshalEnvelopeIdentity(id *envelopespec.Identity) (configkit.Identity, error) {
	return configkit.NewIdentity(id.Name, id.Key)
}

// MustMarshalEnvelopeTime marshals t to its RFC-3339 representation, as used
// within envelopes.
func MustMarshalEnvelopeTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(time.RFC3339Nano)
}

// UnmarshalEnvelopeTime unmarshals a t from its RFC-3339 representation, as
// used within envelopes.
func UnmarshalEnvelopeTime(t string) (time.Time, error) {
	if len(t) == 0 {
		return time.Time{}, nil
	}

	return time.Parse(time.RFC3339Nano, t)
}
