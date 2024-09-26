package protobuf_test

import (
	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/dogmatiq/marshalkit/codec/internal/fixtures"
	. "github.com/dogmatiq/marshalkit/codec/protobuf"
	. "github.com/jmalloc/gomegax"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Codec (configured for JSON format)", func() {
	var codec Codec

	BeforeEach(func() {
		codec = DefaultJSONCodec
	})

	Describe("func BasicMediaType()", func() {
		It("returns the expected basic media-type", func() {
			Expect(codec.BasicMediaType()).To(Equal("application/vnd.google.protobuf+json"))
		})
	})

	Describe("func Marshal()", func() {
		It("marshals the value", func() {
			data, err := codec.Marshal(
				&ProtoMessage{
					Value: "<value>",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).To(Equal(`{"value":"<value>"}`))
		})

		It("returns an error if the type is not a protocol buffers message", func() {
			_, err := codec.Marshal(CommandA1)
			Expect(err).To(MatchError(
				"'stubs.CommandStub[github.com/dogmatiq/enginekit/enginetest/stubs.TypeA]' is not a protocol buffers message",
			))
		})
	})

	Describe("func Unmarshal()", func() {
		It("unmarshals the data", func() {
			data := []byte(`{"value":"<value>"}`)

			m := &ProtoMessage{}
			err := codec.Unmarshal(data, m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m).To(EqualX(
				&ProtoMessage{
					Value: "<value>",
				},
			))
		})

		It("returns an error if the type is not a protocol buffers message", func() {
			var m CommandStub[TypeA]
			err := codec.Unmarshal(nil, &m)
			Expect(err).To(MatchError(
				"'*stubs.CommandStub[github.com/dogmatiq/enginekit/enginetest/stubs.TypeA]' is not a protocol buffers message",
			))
		})
	})
})
