package fixtures

import (
	"reflect"

	"github.com/dogmatiq/dogma/fixtures"
	"github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/json"
)

// Marshaler is a marshaler that is aware of the message and aggregate/process
// root fixture types. It uses the JSON codec.
var Marshaler *marshalkit.Marshaler

func init() {
	m, err := marshalkit.NewMarshaler(
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
		[]marshalkit.Codec{
			&json.Codec{},
		},
	)
	if err != nil {
		panic(err)
	}

	Marshaler = m

	MessageA1Packet = must(m.MarshalMessage(fixtures.MessageA1))
	MessageA2Packet = must(m.MarshalMessage(fixtures.MessageA2))
	MessageA3Packet = must(m.MarshalMessage(fixtures.MessageA3))

	MessageB1Packet = must(m.MarshalMessage(fixtures.MessageB1))
	MessageB2Packet = must(m.MarshalMessage(fixtures.MessageB2))
	MessageB3Packet = must(m.MarshalMessage(fixtures.MessageB3))

	MessageC1Packet = must(m.MarshalMessage(fixtures.MessageC1))
	MessageC2Packet = must(m.MarshalMessage(fixtures.MessageC2))
	MessageC3Packet = must(m.MarshalMessage(fixtures.MessageC3))

	MessageD1Packet = must(m.MarshalMessage(fixtures.MessageD1))
	MessageD2Packet = must(m.MarshalMessage(fixtures.MessageD2))
	MessageD3Packet = must(m.MarshalMessage(fixtures.MessageD3))

	MessageE1Packet = must(m.MarshalMessage(fixtures.MessageE1))
	MessageE2Packet = must(m.MarshalMessage(fixtures.MessageE2))
	MessageE3Packet = must(m.MarshalMessage(fixtures.MessageE3))

	MessageF1Packet = must(m.MarshalMessage(fixtures.MessageF1))
	MessageF2Packet = must(m.MarshalMessage(fixtures.MessageF2))
	MessageF3Packet = must(m.MarshalMessage(fixtures.MessageF3))

	MessageG1Packet = must(m.MarshalMessage(fixtures.MessageG1))
	MessageG2Packet = must(m.MarshalMessage(fixtures.MessageG2))
	MessageG3Packet = must(m.MarshalMessage(fixtures.MessageG3))

	MessageH1Packet = must(m.MarshalMessage(fixtures.MessageH1))
	MessageH2Packet = must(m.MarshalMessage(fixtures.MessageH2))
	MessageH3Packet = must(m.MarshalMessage(fixtures.MessageH3))

	MessageI1Packet = must(m.MarshalMessage(fixtures.MessageI1))
	MessageI2Packet = must(m.MarshalMessage(fixtures.MessageI2))
	MessageI3Packet = must(m.MarshalMessage(fixtures.MessageI3))

	MessageJ1Packet = must(m.MarshalMessage(fixtures.MessageJ1))
	MessageJ2Packet = must(m.MarshalMessage(fixtures.MessageJ2))
	MessageJ3Packet = must(m.MarshalMessage(fixtures.MessageJ3))

	MessageK1Packet = must(m.MarshalMessage(fixtures.MessageK1))
	MessageK2Packet = must(m.MarshalMessage(fixtures.MessageK2))
	MessageK3Packet = must(m.MarshalMessage(fixtures.MessageK3))

	MessageL1Packet = must(m.MarshalMessage(fixtures.MessageL1))
	MessageL2Packet = must(m.MarshalMessage(fixtures.MessageL2))
	MessageL3Packet = must(m.MarshalMessage(fixtures.MessageL3))

	MessageM1Packet = must(m.MarshalMessage(fixtures.MessageM1))
	MessageM2Packet = must(m.MarshalMessage(fixtures.MessageM2))
	MessageM3Packet = must(m.MarshalMessage(fixtures.MessageM3))

	MessageN1Packet = must(m.MarshalMessage(fixtures.MessageN1))
	MessageN2Packet = must(m.MarshalMessage(fixtures.MessageN2))
	MessageN3Packet = must(m.MarshalMessage(fixtures.MessageN3))

	MessageO1Packet = must(m.MarshalMessage(fixtures.MessageO1))
	MessageO2Packet = must(m.MarshalMessage(fixtures.MessageO2))
	MessageO3Packet = must(m.MarshalMessage(fixtures.MessageO3))

	MessageP1Packet = must(m.MarshalMessage(fixtures.MessageP1))
	MessageP2Packet = must(m.MarshalMessage(fixtures.MessageP2))
	MessageP3Packet = must(m.MarshalMessage(fixtures.MessageP3))

	MessageQ1Packet = must(m.MarshalMessage(fixtures.MessageQ1))
	MessageQ2Packet = must(m.MarshalMessage(fixtures.MessageQ2))
	MessageQ3Packet = must(m.MarshalMessage(fixtures.MessageQ3))

	MessageR1Packet = must(m.MarshalMessage(fixtures.MessageR1))
	MessageR2Packet = must(m.MarshalMessage(fixtures.MessageR2))
	MessageR3Packet = must(m.MarshalMessage(fixtures.MessageR3))

	MessageS1Packet = must(m.MarshalMessage(fixtures.MessageS1))
	MessageS2Packet = must(m.MarshalMessage(fixtures.MessageS2))
	MessageS3Packet = must(m.MarshalMessage(fixtures.MessageS3))

	MessageT1Packet = must(m.MarshalMessage(fixtures.MessageT1))
	MessageT2Packet = must(m.MarshalMessage(fixtures.MessageT2))
	MessageT3Packet = must(m.MarshalMessage(fixtures.MessageT3))

	MessageU1Packet = must(m.MarshalMessage(fixtures.MessageU1))
	MessageU2Packet = must(m.MarshalMessage(fixtures.MessageU2))
	MessageU3Packet = must(m.MarshalMessage(fixtures.MessageU3))

	MessageV1Packet = must(m.MarshalMessage(fixtures.MessageV1))
	MessageV2Packet = must(m.MarshalMessage(fixtures.MessageV2))
	MessageV3Packet = must(m.MarshalMessage(fixtures.MessageV3))

	MessageW1Packet = must(m.MarshalMessage(fixtures.MessageW1))
	MessageW2Packet = must(m.MarshalMessage(fixtures.MessageW2))
	MessageW3Packet = must(m.MarshalMessage(fixtures.MessageW3))

	MessageX1Packet = must(m.MarshalMessage(fixtures.MessageX1))
	MessageX2Packet = must(m.MarshalMessage(fixtures.MessageX2))
	MessageX3Packet = must(m.MarshalMessage(fixtures.MessageX3))

	MessageY1Packet = must(m.MarshalMessage(fixtures.MessageY1))
	MessageY2Packet = must(m.MarshalMessage(fixtures.MessageY2))
	MessageY3Packet = must(m.MarshalMessage(fixtures.MessageY3))

	MessageZ1Packet = must(m.MarshalMessage(fixtures.MessageZ1))
	MessageZ2Packet = must(m.MarshalMessage(fixtures.MessageZ2))
	MessageZ3Packet = must(m.MarshalMessage(fixtures.MessageZ3))
}

func must(p marshalkit.Packet, err error) marshalkit.Packet {
	if err != nil {
		panic(err)
	}

	return p
}
