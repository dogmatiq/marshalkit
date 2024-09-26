package fixtures

import (
	"reflect"

	"github.com/dogmatiq/dogma/fixtures"
	"github.com/dogmatiq/enginekit/enginetest/stubs"
	"github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/codec"
	"github.com/dogmatiq/marshalkit/codec/cbor"
	"github.com/dogmatiq/marshalkit/codec/json"
)

// Marshaler is a marshaler that is aware of the message and aggregate/process
// root fixture types. It uses the JSON codec.
var Marshaler marshalkit.Marshaler

func init() {
	m, err := codec.NewMarshaler(
		[]reflect.Type{
			reflect.TypeFor[*fixtures.AggregateRoot](),
			reflect.TypeFor[*fixtures.ProcessRoot](),

			reflect.TypeFor[fixtures.MessageA](),
			reflect.TypeFor[fixtures.MessageB](),
			reflect.TypeFor[fixtures.MessageC](),
			reflect.TypeFor[fixtures.MessageD](),
			reflect.TypeFor[fixtures.MessageE](),
			reflect.TypeFor[fixtures.MessageF](),
			reflect.TypeFor[fixtures.MessageG](),
			reflect.TypeFor[fixtures.MessageH](),
			reflect.TypeFor[fixtures.MessageI](),
			reflect.TypeFor[fixtures.MessageJ](),
			reflect.TypeFor[fixtures.MessageK](),
			reflect.TypeFor[fixtures.MessageL](),
			reflect.TypeFor[fixtures.MessageM](),
			reflect.TypeFor[fixtures.MessageN](),
			reflect.TypeFor[fixtures.MessageO](),
			reflect.TypeFor[fixtures.MessageP](),
			reflect.TypeFor[fixtures.MessageQ](),
			reflect.TypeFor[fixtures.MessageR](),
			reflect.TypeFor[fixtures.MessageS](),
			reflect.TypeFor[fixtures.MessageT](),
			reflect.TypeFor[fixtures.MessageU](),
			reflect.TypeFor[fixtures.MessageV](),
			reflect.TypeFor[fixtures.MessageW](),
			reflect.TypeFor[fixtures.MessageX](),
			reflect.TypeFor[fixtures.MessageY](),
			reflect.TypeFor[fixtures.MessageZ](),

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
			&cbor.Codec{},
		},
	)
	if err != nil {
		panic(err)
	}

	Marshaler = m

	AggregateRootPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(&fixtures.AggregateRoot{}))
	ProcessRootPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(&fixtures.ProcessRoot{}))

	MessageA1Packet = marshalkit.MustMarshal(m, fixtures.MessageA1)
	MessageA2Packet = marshalkit.MustMarshal(m, fixtures.MessageA2)
	MessageA3Packet = marshalkit.MustMarshal(m, fixtures.MessageA3)
	MessageAPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageA{}))

	MessageB1Packet = marshalkit.MustMarshal(m, fixtures.MessageB1)
	MessageB2Packet = marshalkit.MustMarshal(m, fixtures.MessageB2)
	MessageB3Packet = marshalkit.MustMarshal(m, fixtures.MessageB3)
	MessageBPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageB{}))

	MessageC1Packet = marshalkit.MustMarshal(m, fixtures.MessageC1)
	MessageC2Packet = marshalkit.MustMarshal(m, fixtures.MessageC2)
	MessageC3Packet = marshalkit.MustMarshal(m, fixtures.MessageC3)
	MessageCPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageC{}))

	MessageD1Packet = marshalkit.MustMarshal(m, fixtures.MessageD1)
	MessageD2Packet = marshalkit.MustMarshal(m, fixtures.MessageD2)
	MessageD3Packet = marshalkit.MustMarshal(m, fixtures.MessageD3)
	MessageDPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageD{}))

	MessageE1Packet = marshalkit.MustMarshal(m, fixtures.MessageE1)
	MessageE2Packet = marshalkit.MustMarshal(m, fixtures.MessageE2)
	MessageE3Packet = marshalkit.MustMarshal(m, fixtures.MessageE3)
	MessageEPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageE{}))

	MessageF1Packet = marshalkit.MustMarshal(m, fixtures.MessageF1)
	MessageF2Packet = marshalkit.MustMarshal(m, fixtures.MessageF2)
	MessageF3Packet = marshalkit.MustMarshal(m, fixtures.MessageF3)
	MessageFPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageF{}))

	MessageG1Packet = marshalkit.MustMarshal(m, fixtures.MessageG1)
	MessageG2Packet = marshalkit.MustMarshal(m, fixtures.MessageG2)
	MessageG3Packet = marshalkit.MustMarshal(m, fixtures.MessageG3)
	MessageGPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageG{}))

	MessageH1Packet = marshalkit.MustMarshal(m, fixtures.MessageH1)
	MessageH2Packet = marshalkit.MustMarshal(m, fixtures.MessageH2)
	MessageH3Packet = marshalkit.MustMarshal(m, fixtures.MessageH3)
	MessageHPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageH{}))

	MessageI1Packet = marshalkit.MustMarshal(m, fixtures.MessageI1)
	MessageI2Packet = marshalkit.MustMarshal(m, fixtures.MessageI2)
	MessageI3Packet = marshalkit.MustMarshal(m, fixtures.MessageI3)
	MessageIPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageI{}))

	MessageJ1Packet = marshalkit.MustMarshal(m, fixtures.MessageJ1)
	MessageJ2Packet = marshalkit.MustMarshal(m, fixtures.MessageJ2)
	MessageJ3Packet = marshalkit.MustMarshal(m, fixtures.MessageJ3)
	MessageJPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageJ{}))

	MessageK1Packet = marshalkit.MustMarshal(m, fixtures.MessageK1)
	MessageK2Packet = marshalkit.MustMarshal(m, fixtures.MessageK2)
	MessageK3Packet = marshalkit.MustMarshal(m, fixtures.MessageK3)
	MessageKPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageK{}))

	MessageL1Packet = marshalkit.MustMarshal(m, fixtures.MessageL1)
	MessageL2Packet = marshalkit.MustMarshal(m, fixtures.MessageL2)
	MessageL3Packet = marshalkit.MustMarshal(m, fixtures.MessageL3)
	MessageLPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageL{}))

	MessageM1Packet = marshalkit.MustMarshal(m, fixtures.MessageM1)
	MessageM2Packet = marshalkit.MustMarshal(m, fixtures.MessageM2)
	MessageM3Packet = marshalkit.MustMarshal(m, fixtures.MessageM3)
	MessageMPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageM{}))

	MessageN1Packet = marshalkit.MustMarshal(m, fixtures.MessageN1)
	MessageN2Packet = marshalkit.MustMarshal(m, fixtures.MessageN2)
	MessageN3Packet = marshalkit.MustMarshal(m, fixtures.MessageN3)
	MessageNPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageN{}))

	MessageO1Packet = marshalkit.MustMarshal(m, fixtures.MessageO1)
	MessageO2Packet = marshalkit.MustMarshal(m, fixtures.MessageO2)
	MessageO3Packet = marshalkit.MustMarshal(m, fixtures.MessageO3)
	MessageOPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageO{}))

	MessageP1Packet = marshalkit.MustMarshal(m, fixtures.MessageP1)
	MessageP2Packet = marshalkit.MustMarshal(m, fixtures.MessageP2)
	MessageP3Packet = marshalkit.MustMarshal(m, fixtures.MessageP3)
	MessagePPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageP{}))

	MessageQ1Packet = marshalkit.MustMarshal(m, fixtures.MessageQ1)
	MessageQ2Packet = marshalkit.MustMarshal(m, fixtures.MessageQ2)
	MessageQ3Packet = marshalkit.MustMarshal(m, fixtures.MessageQ3)
	MessageQPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageQ{}))

	MessageR1Packet = marshalkit.MustMarshal(m, fixtures.MessageR1)
	MessageR2Packet = marshalkit.MustMarshal(m, fixtures.MessageR2)
	MessageR3Packet = marshalkit.MustMarshal(m, fixtures.MessageR3)
	MessageRPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageR{}))

	MessageS1Packet = marshalkit.MustMarshal(m, fixtures.MessageS1)
	MessageS2Packet = marshalkit.MustMarshal(m, fixtures.MessageS2)
	MessageS3Packet = marshalkit.MustMarshal(m, fixtures.MessageS3)
	MessageSPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageS{}))

	MessageT1Packet = marshalkit.MustMarshal(m, fixtures.MessageT1)
	MessageT2Packet = marshalkit.MustMarshal(m, fixtures.MessageT2)
	MessageT3Packet = marshalkit.MustMarshal(m, fixtures.MessageT3)
	MessageTPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageT{}))

	MessageU1Packet = marshalkit.MustMarshal(m, fixtures.MessageU1)
	MessageU2Packet = marshalkit.MustMarshal(m, fixtures.MessageU2)
	MessageU3Packet = marshalkit.MustMarshal(m, fixtures.MessageU3)
	MessageUPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageU{}))

	MessageV1Packet = marshalkit.MustMarshal(m, fixtures.MessageV1)
	MessageV2Packet = marshalkit.MustMarshal(m, fixtures.MessageV2)
	MessageV3Packet = marshalkit.MustMarshal(m, fixtures.MessageV3)
	MessageVPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageV{}))

	MessageW1Packet = marshalkit.MustMarshal(m, fixtures.MessageW1)
	MessageW2Packet = marshalkit.MustMarshal(m, fixtures.MessageW2)
	MessageW3Packet = marshalkit.MustMarshal(m, fixtures.MessageW3)
	MessageWPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageW{}))

	MessageX1Packet = marshalkit.MustMarshal(m, fixtures.MessageX1)
	MessageX2Packet = marshalkit.MustMarshal(m, fixtures.MessageX2)
	MessageX3Packet = marshalkit.MustMarshal(m, fixtures.MessageX3)
	MessageXPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageX{}))

	MessageY1Packet = marshalkit.MustMarshal(m, fixtures.MessageY1)
	MessageY2Packet = marshalkit.MustMarshal(m, fixtures.MessageY2)
	MessageY3Packet = marshalkit.MustMarshal(m, fixtures.MessageY3)
	MessageYPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageY{}))

	MessageZ1Packet = marshalkit.MustMarshal(m, fixtures.MessageZ1)
	MessageZ2Packet = marshalkit.MustMarshal(m, fixtures.MessageZ2)
	MessageZ3Packet = marshalkit.MustMarshal(m, fixtures.MessageZ3)
	MessageZPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageZ{}))
}
