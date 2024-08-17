package marshalkit

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/dogma"
	"github.com/dogmatiq/interopspec/envelopespec"
)

// MustMarshalMessageIntoEnvelope marshals a Dogma message into an
// [envelopespec.Envelope].
//
// Deprecated: Use [MustMarshalCommandIntoEnvelope],
// [MustMarshalEventIntoEnvelope], or [MustMarshalTimeoutIntoEnvelope] instead.
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

// MustMarshalCommandIntoEnvelope marshals a [dogma.Command] into an
// [envelopespec.Envelope].
func MustMarshalCommandIntoEnvelope(
	vm ValueMarshaler,
	m dogma.Command,
	env *envelopespec.Envelope,
) {
	MustMarshalMessageIntoEnvelope(vm, m, env)
}

// MustMarshalEventIntoEnvelope marshals a [dogma.Event] into an
// [envelopespec.Envelope].
func MustMarshalEventIntoEnvelope(
	vm ValueMarshaler,
	m dogma.Event,
	env *envelopespec.Envelope,
) {
	MustMarshalMessageIntoEnvelope(vm, m, env)
}

// MustMarshalTimeoutIntoEnvelope marshals a [dogma.Timeout] into an
// [envelopespec.Envelope].
func MustMarshalTimeoutIntoEnvelope(
	vm ValueMarshaler,
	m dogma.Timeout,
	env *envelopespec.Envelope,
) {
	MustMarshalMessageIntoEnvelope(vm, m, env)
}

// UnmarshalMessageFromEnvelope unmarshals a Dogma message from an
// [envelopespec.Envelope].
//
// Deprecated: Use [UnmarshalCommandFromEnvelope], [UnmarshalEventFromEnvelope],
// or [UnmarshalTimeoutFromEnvelope] instead.
func UnmarshalMessageFromEnvelope(
	vm ValueMarshaler,
	env *envelopespec.Envelope,
) (dogma.Message, error) {
	p := Packet{
		MediaType: env.MediaType,
		Data:      env.Data,
	}

	v, err := vm.Unmarshal(p)
	if err != nil {
		return nil, err
	}

	m, ok := v.(dogma.Message)
	if !ok {
		return nil, fmt.Errorf(
			"'%s' is not a message",
			reflect.TypeOf(v),
		)
	}

	return m, nil
}

// UnmarshalCommandFromEnvelope unmarshals a [dogma.Command] from an
// [envelopespec.Envelope].
func UnmarshalCommandFromEnvelope(
	vm ValueMarshaler,
	env *envelopespec.Envelope,
) (dogma.Command, error) {
	return UnmarshalMessageFromEnvelope(vm, env)
}

// UnmarshalEventFromEnvelope unmarshals a [dogma.Event] from an
// [envelopespec.Envelope].
func UnmarshalEventFromEnvelope(
	vm ValueMarshaler,
	env *envelopespec.Envelope,
) (dogma.Event, error) {
	return UnmarshalMessageFromEnvelope(vm, env)
}

// UnmarshalTimeoutFromEnvelope unmarshals a [dogma.Timeout] from an
// [envelopespec.Envelope].
func UnmarshalTimeoutFromEnvelope(
	vm ValueMarshaler,
	env *envelopespec.Envelope,
) (dogma.Timeout, error) {
	return UnmarshalMessageFromEnvelope(vm, env)
}

// MustMarshalEnvelopeIdentity marshals id to its protocol buffers
// representation, as used within envelopespec.Envelope.
func MustMarshalEnvelopeIdentity(id configkit.Identity) *envelopespec.Identity {
	return &envelopespec.Identity{
		Name: id.Name,
		Key:  id.Key,
	}
}

// UnmarshalEnvelopeIdentity unmarshals id from its protocol buffers
// representation, as used within envelopespec.Envelope.
func UnmarshalEnvelopeIdentity(id *envelopespec.Identity) (configkit.Identity, error) {
	return configkit.NewIdentity(id.Name, id.Key)
}

// MustMarshalEnvelopeTime marshals t to its RFC-3339 representation, as used
// within envelopespec.Envelope.
func MustMarshalEnvelopeTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(time.RFC3339Nano)
}

// UnmarshalEnvelopeTime unmarshals t from its RFC-3339 representation, as
// used within envelopespec.Envelope.
func UnmarshalEnvelopeTime(t string) (time.Time, error) {
	if len(t) == 0 {
		return time.Time{}, nil
	}

	return time.Parse(time.RFC3339Nano, t)
}
