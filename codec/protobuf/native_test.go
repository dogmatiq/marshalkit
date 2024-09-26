package protobuf_test

import (
	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/dogmatiq/marshalkit/codec/internal/fixtures"
	. "github.com/dogmatiq/marshalkit/codec/protobuf"
	. "github.com/jmalloc/gomegax"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/proto"
)

var _ = Describe("type Codec (configured for the native wire protocol format)", func() {
	var codec Codec

	BeforeEach(func() {
		codec = DefaultNativeCodec
	})

	Describe("func BasicMediaType()", func() {
		It("returns the expected basic media-type", func() {
			Expect(codec.BasicMediaType()).To(Equal("application/vnd.google.protobuf"))
		})
	})

	Describe("func Marshal()", func() {
		It("marshals the value", func() {
			m := &ProtoMessage{
				Value: "<value>",
			}

			data, err := codec.Marshal(m)
			Expect(err).ShouldNot(HaveOccurred())

			expected, err := proto.Marshal(m)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expected))
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
			m := &ProtoMessage{
				Value: "<value>",
			}

			data, err := codec.Marshal(m)
			Expect(err).ShouldNot(HaveOccurred())

			m = &ProtoMessage{}
			err = codec.Unmarshal(data, m)
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
