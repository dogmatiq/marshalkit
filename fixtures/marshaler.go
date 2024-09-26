package fixtures

import (
	"reflect"

	"github.com/dogmatiq/enginekit/enginetest/stubs"
	"github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/codec"
	"github.com/dogmatiq/marshalkit/codec/json"
)

// Marshaler is a marshaler that is aware of the message and aggregate/process
// root fixture types. It uses the JSON codec.
var Marshaler marshalkit.Marshaler

func init() {
	m, err := codec.NewMarshaler(
		[]reflect.Type{
			reflect.TypeFor[*stubs.AggregateRootStub](),
			reflect.TypeFor[*stubs.ProcessRootStub](),

			reflect.TypeFor[stubs.CommandStub[stubs.TypeA]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeB]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeC]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeD]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeE]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeF]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeG]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeH]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeI]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeJ]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeK]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeL]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeM]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeN]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeO]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeP]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeQ]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeR]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeS]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeT]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeU]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeV]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeW]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeX]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeY]](),
			reflect.TypeFor[stubs.CommandStub[stubs.TypeZ]](),

			reflect.TypeFor[stubs.EventStub[stubs.TypeA]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeB]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeC]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeD]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeE]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeF]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeG]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeH]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeI]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeJ]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeK]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeL]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeM]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeN]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeO]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeP]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeQ]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeR]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeS]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeT]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeU]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeV]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeW]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeX]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeY]](),
			reflect.TypeFor[stubs.EventStub[stubs.TypeZ]](),

			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeA]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeB]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeC]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeD]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeE]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeF]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeG]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeH]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeI]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeJ]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeK]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeL]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeM]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeN]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeO]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeP]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeQ]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeR]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeS]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeT]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeU]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeV]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeW]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeX]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeY]](),
			reflect.TypeFor[stubs.TimeoutStub[stubs.TypeZ]](),
		},
		[]codec.Codec{
			&json.Codec{},
		},
	)
	if err != nil {
		panic(err)
	}

	Marshaler = m
}
