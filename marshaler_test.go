package marshalkit_test

import (
	"reflect"

	"github.com/dogmatiq/marshalkit/internal/fixtures"

	. "github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/json"
	"github.com/dogmatiq/marshalkit/protobuf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Marshaler", func() {
	var marshaler *Marshaler

	BeforeEach(func() {
		var err error
		marshaler, err = NewMarshaler(
			[]reflect.Type{
				reflect.TypeOf(&fixtures.ProtoMessage{}),
				reflect.TypeOf(fixtures.PlainMessageA{}),
				reflect.TypeOf(fixtures.PlainMessageB{}),
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
					reflect.TypeOf(fixtures.PlainMessageA{}),
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
					reflect.TypeOf(fixtures.PlainMessageA{}),
					reflect.TypeOf(&fixtures.PlainMessageA{}),
				},
				[]Codec{
					&json.Codec{},
				},
			)
			Expect(err).To(Or(
				MatchError(
					"the type name 'PlainMessageA' is used by both 'fixtures.PlainMessageA' and '*fixtures.PlainMessageA'",
				),
				MatchError(
					"the type name 'PlainMessageA' is used by both '*fixtures.PlainMessageA' and 'fixtures.PlainMessageA'",
				),
			))
		})

		It("returns an error if there are unsupported types", func() {
			_, err := NewMarshaler(
				[]reflect.Type{
					reflect.TypeOf(&fixtures.ProtoMessage{}),
					reflect.TypeOf(fixtures.PlainMessageA{}),
				},
				[]Codec{
					&protobuf.JSONCodec{},
				},
			)
			Expect(err).To(MatchError(
				"no codecs support the 'fixtures.PlainMessageA' type",
			))
		})
	})

	Describe("func MarshalType()", func() {
		It("returns the portable type name", func() {
			n, err := marshaler.MarshalType(
				reflect.TypeOf(fixtures.PlainMessageA{}),
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(n).To(Equal("PlainMessageA"))
		})

		It("returns an error if the type is not supported", func() {
			_, err := marshaler.MarshalType(
				reflect.TypeOf(fixtures.PlainMessageC{}),
			)
			Expect(err).To(MatchError(
				"no codecs support the 'fixtures.PlainMessageC' type",
			))
		})
	})

	Describe("func UnmarshalType()", func() {
		It("returns the reflection type", func() {
			t, err := marshaler.UnmarshalType("PlainMessageA")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(t).To(Equal(reflect.TypeOf(fixtures.PlainMessageA{})))
		})

		It("returns an error if the type name is not recognized", func() {
			_, err := marshaler.UnmarshalType("PlainMessageC")
			Expect(err).To(MatchError(
				"the portable type name 'PlainMessageC' is not recognized",
			))
		})
	})

	Describe("func UnmarshalTypeFromMediaType()", func() {
		It("returns the reflection type", func() {
			t, err := marshaler.UnmarshalTypeFromMediaType(
				"application/vnd.google.protobuf; type=PlainMessageA",
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(t).To(Equal(reflect.TypeOf(fixtures.PlainMessageA{})))
		})

		It("returns an error if the type name is not recognized", func() {
			_, err := marshaler.UnmarshalTypeFromMediaType(
				"application/vnd.google.protobuf; type=PlainMessageC",
			)
			Expect(err).To(MatchError(
				"the portable type name 'PlainMessageC' is not recognized",
			))
		})
	})

	Describe("func Marshal()", func() {
		It("marshals using the first suitable codec", func() {
			p, err := marshaler.Marshal(fixtures.PlainMessageA{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/json; type=PlainMessageA"))

			p, err = marshaler.Marshal(&fixtures.ProtoMessage{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage"))
		})

		It("returns an error if the codec fails", func() {
			_, err := marshaler.Marshal(
				&fixtures.ProtoMessage{
					Value: string([]byte{0xfe}),
				},
			)
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the type is not supported", func() {
			_, err := marshaler.Marshal(fixtures.PlainMessageC{})
			Expect(err).To(MatchError(
				"no codecs support the 'fixtures.PlainMessageC' type",
			))
		})
	})

	Describe("func Unmarshal()", func() {
		It("marshals using the first suitable codec", func() {
			v, err := marshaler.Unmarshal(
				Packet{
					"application/json; type=PlainMessageA",
					[]byte("{}"),
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).To(Equal(fixtures.PlainMessageA{}))

			v, err = marshaler.Unmarshal(
				Packet{
					"application/vnd.google.protobuf+json; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
					[]byte("{}"),
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).To(Equal(&fixtures.ProtoMessage{}))
		})

		It("returns an error if the media-type is not supported", func() {
			_, err := marshaler.Unmarshal(Packet{"text/plain", nil})
			Expect(err).To(MatchError(
				"no codecs support the 'text/plain' media-type",
			))
		})

		It("returns an error if the media-type is malformed", func() {
			_, err := marshaler.Unmarshal(Packet{"<malformed>", nil})
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the media-type does not specify a type parameter", func() {
			_, err := marshaler.Unmarshal(Packet{"application/json", nil})
			Expect(err).Should(MatchError(
				"the media-type 'application/json' does not specify a 'type' parameter",
			))
		})

		It("returns an error if the type is not supported", func() {
			_, err := marshaler.Unmarshal(Packet{"application/json; type=PlainMessageC", nil})
			Expect(err).Should(MatchError(
				"the portable type name 'PlainMessageC' is not recognized",
			))
		})

		It("returns an error if the codec fails", func() {
			_, err := marshaler.Unmarshal(
				Packet{
					"application/json; type=PlainMessageA",
					[]byte("{"),
				},
			)
			Expect(err).Should(HaveOccurred())
		})
	})
})
