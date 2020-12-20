package codec_test

import (
	"reflect"

	. "github.com/dogmatiq/dogma/fixtures"
	"github.com/dogmatiq/marshalkit"
	. "github.com/dogmatiq/marshalkit/codec"
	. "github.com/dogmatiq/marshalkit/codec/internal/fixtures"
	"github.com/dogmatiq/marshalkit/codec/json"
	"github.com/dogmatiq/marshalkit/codec/protobuf"
	. "github.com/jmalloc/gomegax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ marshalkit.Marshaler = (*Marshaler)(nil)

var _ = Describe("type Marshaler", func() {
	var marshaler *Marshaler

	BeforeEach(func() {
		var err error
		marshaler, err = NewMarshaler(
			[]reflect.Type{
				reflect.TypeOf(&ProtoMessage{}),
				reflect.TypeOf(MessageA{}),
				reflect.TypeOf(MessageB{}),
			},
			[]Codec{
				&protobuf.NativeCodec{},
				&protobuf.JSONCodec{},
				&json.Codec{},
			},
		)
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("func NewMarshaler()", func() {
		It("returns an error if multiple codecs used the same media-type", func() {
			_, err := NewMarshaler(
				[]reflect.Type{
					reflect.TypeOf(MessageA{}),
				},
				[]Codec{
					&json.Codec{},
					&json.Codec{},
				},
			)
			Expect(err).To(MatchError(
				"multiple codecs use the 'application/json' media-type",
			))
		})

		It("returns an error if there conflicting portable type names", func() {
			_, err := NewMarshaler(
				[]reflect.Type{
					reflect.TypeOf(MessageA{}),
					reflect.TypeOf(&MessageA{}),
				},
				[]Codec{
					&json.Codec{},
				},
			)
			Expect(err).To(Or(
				MatchError(
					"the type name 'MessageA' is used by both 'fixtures.MessageA' and '*fixtures.MessageA'",
				),
				MatchError(
					"the type name 'MessageA' is used by both '*fixtures.MessageA' and 'fixtures.MessageA'",
				),
			))
		})

		It("returns an error if there are unsupported types", func() {
			_, err := NewMarshaler(
				[]reflect.Type{
					reflect.TypeOf(&ProtoMessage{}),
					reflect.TypeOf(MessageA{}),
				},
				[]Codec{
					&protobuf.JSONCodec{},
				},
			)
			Expect(err).To(MatchError(
				"no codecs support the 'fixtures.MessageA' type",
			))
		})
	})

	Describe("func MarshalType()", func() {
		It("returns the portable type name", func() {
			n, err := marshaler.MarshalType(
				reflect.TypeOf(MessageA{}),
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(n).To(Equal("MessageA"))
		})

		It("returns an error if the type is not supported", func() {
			_, err := marshaler.MarshalType(
				reflect.TypeOf(MessageC{}),
			)
			Expect(err).To(MatchError(
				"no codecs support the 'fixtures.MessageC' type",
			))
		})
	})

	Describe("func UnmarshalType()", func() {
		It("returns the reflection type", func() {
			t, err := marshaler.UnmarshalType("MessageA")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(t).To(Equal(reflect.TypeOf(MessageA{})))
		})

		It("returns an error if the type name is not recognized", func() {
			_, err := marshaler.UnmarshalType("MessageC")
			Expect(err).To(MatchError(
				"the portable type name 'MessageC' is not recognized",
			))
		})
	})

	Describe("func Marshal()", func() {
		It("marshals using the first suitable codec", func() {
			p, err := marshaler.Marshal(MessageA{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/json; type=MessageA"))

			p, err = marshaler.Marshal(&ProtoMessage{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage"))
		})

		It("returns an error if the codec fails", func() {
			_, err := marshaler.Marshal(
				&ProtoMessage{
					Value: string([]byte{0xfe}),
				},
			)
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the type is not supported", func() {
			_, err := marshaler.Marshal(MessageC{})
			Expect(err).To(MatchError(
				"no codecs support the 'fixtures.MessageC' type",
			))
		})
	})

	Describe("func MarshalAs()", func() {
		It("marshals using the codec associated with the given media type", func() {
			expected := []byte("{\"Value\":null}")
			p, err := marshaler.MarshalAs(
				MessageA{},
				"application/json; type=MessageA",
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/json; type=MessageA"))
			Expect(p.Data).To(Equal(expected))

			expected = []byte{10, 7, 60, 118, 97, 108, 117, 101, 62}
			p, err = marshaler.MarshalAs(
				&ProtoMessage{
					Value: "<value>",
				},
				"application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage"))
			Expect(p.Data).To(Equal(expected))
		})

		It("returns an error if the media-type is malformed", func() {
			_, err := marshaler.MarshalAs(
				MessageA{},
				"<malformed>",
			)
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the codec fails", func() {
			_, err := marshaler.MarshalAs(
				&ProtoMessage{
					Value: string([]byte{0xfe}),
				},
				"application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
			)
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the type is not supported", func() {
			_, err := marshaler.MarshalAs(
				MessageC{},
				"application/json; type=MessageC",
			)
			Expect(err).To(MatchError(
				"no codecs support marshaling the 'fixtures.MessageC' type as application/json; type=MessageC",
			))
		})

		It("returns an error if the portable name in the media-type does not match the value's type", func() {
			_, err := marshaler.MarshalAs(
				MessageA{},
				"application/json; type=MessageC",
			)
			Expect(err).To(MatchError(
				"no codecs support marshaling the 'fixtures.MessageA' type as application/json; type=MessageC",
			))
		})
	})

	Describe("func Unmarshal()", func() {
		It("unmarshals using the first suitable codec", func() {
			v, err := marshaler.Unmarshal(
				marshalkit.Packet{
					MediaType: "application/json; type=MessageA",
					Data:      []byte("{}"),
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).To(Equal(MessageA{}))

			v, err = marshaler.Unmarshal(
				marshalkit.Packet{
					MediaType: "application/vnd.google.protobuf+json; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
					Data:      []byte("{}"),
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).To(EqualX(&ProtoMessage{}))
		})

		It("returns an error if the media-type is not supported", func() {
			_, err := marshaler.Unmarshal(marshalkit.Packet{
				MediaType: "text/plain; type=MessageA",
			})
			Expect(err).To(MatchError(
				"no codecs support the 'text/plain' media-type",
			))
		})

		It("returns an error if the media-type is malformed", func() {
			_, err := marshaler.Unmarshal(marshalkit.Packet{
				MediaType: "<malformed>",
			})
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the media-type does not specify a type parameter", func() {
			_, err := marshaler.Unmarshal(marshalkit.Packet{
				MediaType: "application/json",
			})
			Expect(err).Should(MatchError(
				"the media-type 'application/json' does not specify a 'type' parameter",
			))
		})

		It("returns an error if the type is not supported", func() {
			_, err := marshaler.Unmarshal(marshalkit.Packet{
				MediaType: "application/json; type=MessageC",
			})
			Expect(err).Should(MatchError(
				"the portable type name 'MessageC' is not recognized",
			))
		})

		It("returns an error if the codec fails", func() {
			_, err := marshaler.Unmarshal(
				marshalkit.Packet{
					MediaType: "application/json; type=MessageA",
					Data:      []byte("{"),
				},
			)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("func MediaTypesFor()", func() {
		It("returns media types in order of codec priority", func() {
			mt := marshaler.MediaTypesFor(reflect.TypeOf(&ProtoMessage{}))
			Expect(mt).To(Equal([]string{
				"application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
				"application/vnd.google.protobuf+json; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
				"application/json; type=ProtoMessage",
			}))
		})

		It("returns an empty slice when given an unsupported message type", func() {
			mt := marshaler.MediaTypesFor(reflect.TypeOf(&MessageC{}))
			Expect(mt).To(BeEmpty())
		})
	})
})
