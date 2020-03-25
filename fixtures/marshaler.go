package fixtures

import (
	"reflect"

	"github.com/dogmatiq/dogma/fixtures"
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
			reflect.TypeOf(&fixtures.AggregateRoot{}),
			reflect.TypeOf(&fixtures.ProcessRoot{}),
			reflect.TypeOf(fixtures.MessageA{}),
			reflect.TypeOf(fixtures.MessageB{}),
			reflect.TypeOf(fixtures.MessageC{}),
			reflect.TypeOf(fixtures.MessageD{}),
			reflect.TypeOf(fixtures.MessageE{}),
			reflect.TypeOf(fixtures.MessageF{}),
			reflect.TypeOf(fixtures.MessageG{}),
			reflect.TypeOf(fixtures.MessageH{}),
			reflect.TypeOf(fixtures.MessageI{}),
			reflect.TypeOf(fixtures.MessageJ{}),
			reflect.TypeOf(fixtures.MessageK{}),
			reflect.TypeOf(fixtures.MessageL{}),
			reflect.TypeOf(fixtures.MessageM{}),
			reflect.TypeOf(fixtures.MessageN{}),
			reflect.TypeOf(fixtures.MessageO{}),
			reflect.TypeOf(fixtures.MessageP{}),
			reflect.TypeOf(fixtures.MessageQ{}),
			reflect.TypeOf(fixtures.MessageR{}),
			reflect.TypeOf(fixtures.MessageS{}),
			reflect.TypeOf(fixtures.MessageT{}),
			reflect.TypeOf(fixtures.MessageU{}),
			reflect.TypeOf(fixtures.MessageV{}),
			reflect.TypeOf(fixtures.MessageW{}),
			reflect.TypeOf(fixtures.MessageX{}),
			reflect.TypeOf(fixtures.MessageY{}),
			reflect.TypeOf(fixtures.MessageZ{}),
		},
		[]codec.Codec{
			&json.Codec{},
		},
	)
	if err != nil {
		panic(err)
	}

	Marshaler = m

	AggregateRootPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(&fixtures.AggregateRoot{}))
	ProcessRootPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(&fixtures.ProcessRoot{}))

	MessageA1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageA1)
	MessageA2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageA2)
	MessageA3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageA3)
	MessageAPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageA{}))

	MessageB1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageB1)
	MessageB2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageB2)
	MessageB3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageB3)
	MessageBPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageB{}))

	MessageC1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageC1)
	MessageC2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageC2)
	MessageC3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageC3)
	MessageCPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageC{}))

	MessageD1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageD1)
	MessageD2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageD2)
	MessageD3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageD3)
	MessageDPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageD{}))

	MessageE1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageE1)
	MessageE2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageE2)
	MessageE3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageE3)
	MessageEPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageE{}))

	MessageF1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageF1)
	MessageF2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageF2)
	MessageF3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageF3)
	MessageFPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageF{}))

	MessageG1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageG1)
	MessageG2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageG2)
	MessageG3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageG3)
	MessageGPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageG{}))

	MessageH1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageH1)
	MessageH2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageH2)
	MessageH3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageH3)
	MessageHPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageH{}))

	MessageI1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageI1)
	MessageI2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageI2)
	MessageI3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageI3)
	MessageIPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageI{}))

	MessageJ1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageJ1)
	MessageJ2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageJ2)
	MessageJ3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageJ3)
	MessageJPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageJ{}))

	MessageK1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageK1)
	MessageK2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageK2)
	MessageK3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageK3)
	MessageKPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageK{}))

	MessageL1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageL1)
	MessageL2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageL2)
	MessageL3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageL3)
	MessageLPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageL{}))

	MessageM1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageM1)
	MessageM2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageM2)
	MessageM3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageM3)
	MessageMPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageM{}))

	MessageN1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageN1)
	MessageN2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageN2)
	MessageN3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageN3)
	MessageNPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageN{}))

	MessageO1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageO1)
	MessageO2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageO2)
	MessageO3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageO3)
	MessageOPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageO{}))

	MessageP1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageP1)
	MessageP2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageP2)
	MessageP3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageP3)
	MessagePPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageP{}))

	MessageQ1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageQ1)
	MessageQ2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageQ2)
	MessageQ3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageQ3)
	MessageQPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageQ{}))

	MessageR1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageR1)
	MessageR2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageR2)
	MessageR3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageR3)
	MessageRPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageR{}))

	MessageS1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageS1)
	MessageS2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageS2)
	MessageS3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageS3)
	MessageSPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageS{}))

	MessageT1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageT1)
	MessageT2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageT2)
	MessageT3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageT3)
	MessageTPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageT{}))

	MessageU1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageU1)
	MessageU2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageU2)
	MessageU3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageU3)
	MessageUPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageU{}))

	MessageV1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageV1)
	MessageV2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageV2)
	MessageV3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageV3)
	MessageVPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageV{}))

	MessageW1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageW1)
	MessageW2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageW2)
	MessageW3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageW3)
	MessageWPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageW{}))

	MessageX1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageX1)
	MessageX2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageX2)
	MessageX3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageX3)
	MessageXPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageX{}))

	MessageY1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageY1)
	MessageY2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageY2)
	MessageY3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageY3)
	MessageYPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageY{}))

	MessageZ1Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageZ1)
	MessageZ2Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageZ2)
	MessageZ3Packet = marshalkit.MustMarshalMessage(m, fixtures.MessageZ3)
	MessageZPortableName = marshalkit.MustMarshalType(m, reflect.TypeOf(fixtures.MessageZ{}))
}
