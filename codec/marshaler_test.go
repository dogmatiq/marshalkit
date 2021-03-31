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
				&protobuf.DefaultNativeCodec,
				&protobuf.DefaultJSONCodec,
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
					&protobuf.DefaultJSONCodec,
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
			p, err := marshaler.Marshal(
				MessageA{
					Value: "<value>",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/json; type=MessageA"))
			Expect(p.Data).To(Equal([]byte(`{"Value":"\u003cvalue\u003e"}`)))

			p, err = marshaler.Marshal(
				&ProtoMessage{
					Value: "<value>",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage"))
			Expect(p.Data).To(Equal([]byte{10, 7, 60, 118, 97, 108, 117, 101, 62}))
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
			p, ok, err := marshaler.MarshalAs(
				MessageA{
					Value: "<value>",
				},
				[]string{
					"application/json; type=MessageA",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).To(BeTrue())
			Expect(p.MediaType).To(Equal("application/json; type=MessageA"))
			Expect(p.Data).To(Equal([]byte(`{"Value":"\u003cvalue\u003e"}`)))

			p, ok, err = marshaler.MarshalAs(
				&ProtoMessage{
					Value: "<value>",
				},
				[]string{
					"application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).To(BeTrue())
			Expect(p.MediaType).To(Equal("application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage"))
			Expect(p.Data).To(Equal([]byte{10, 7, 60, 118, 97, 108, 117, 101, 62}))

			p, ok, err = marshaler.MarshalAs(
				&ProtoMessage{
					Value: "<value>",
				},
				[]string{
					"application/json; type=ProtoMessage",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).To(BeTrue())
			Expect(p.MediaType).To(Equal("application/json; type=ProtoMessage"))
			Expect(p.Data).To(Equal([]byte(`{"value":"\u003cvalue\u003e"}`)))
		})

		It("marshals using the codec associated with the highest priority media-type", func() {
			p, ok, err := marshaler.MarshalAs(
				&ProtoMessage{
					Value: "<value>",
				},
				[]string{
					"application/json; type=ProtoMessage",
					"application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).To(BeTrue())
			Expect(p.MediaType).To(Equal("application/json; type=ProtoMessage"))
			Expect(p.Data).To(Equal([]byte(`{"value":"\u003cvalue\u003e"}`)))
		})

		It("ignores unsupported media-types", func() {
			p, ok, err := marshaler.MarshalAs(
				MessageA{},
				[]string{
					"application/binary; type=MessageA",
					"application/json; type=MessageA",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).To(BeTrue())
			Expect(p.MediaType).To(Equal("application/json; type=MessageA"))
		})

		It("returns an error if the media-type is malformed", func() {
			_, _, err := marshaler.MarshalAs(
				MessageA{},
				[]string{
					"<malformed>",
				},
			)
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the codec fails", func() {
			_, _, err := marshaler.MarshalAs(
				&ProtoMessage{
					Value: string([]byte{0xfe}),
				},
				[]string{
					"application/vnd.google.protobuf; type=dogmatiq.marshalkit.fixtures.ProtoMessage",
				},
			)
			Expect(err).Should(HaveOccurred())
		})

		It("returns false if the media-type is not supported", func() {
			_, ok, err := marshaler.MarshalAs(
				MessageC{},
				[]string{
					"application/json; type=MessageC",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).To(BeFalse())
		})

		It("returns false if the portable name in the media-type does not match the value's type", func() {
			_, ok, err := marshaler.MarshalAs(
				MessageA{},
				[]string{
					"application/json; type=MessageC",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).To(BeFalse())
		})

		It("panics if no media-types are provided", func() {
			Expect(func() {
				marshaler.MarshalAs(
					MessageA{},
					nil,
				)
			}).To(PanicWith("at least one media-type must be provided"))
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
