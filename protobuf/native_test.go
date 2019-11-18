package protobuf_test

import (
	"github.com/dogmatiq/marshalkit/internal/fixtures"
	. "github.com/dogmatiq/marshalkit/protobuf"
	"github.com/golang/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type NativeCodec", func() {
	var codec *NativeCodec

	BeforeEach(func() {
		codec = &NativeCodec{}
	})

	Describe("func MediaType()", func() {
		It("returns the expected media-type", func() {
			Expect(codec.MediaType()).To(Equal("application/vnd.google.protobuf"))
		})
	})

	Describe("func Marshal()", func() {
		It("marshals the value", func() {
			m := &fixtures.ProtoMessage{
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
			_, err := codec.Marshal(
				fixtures.PlainMessageA{},
			)
			Expect(err).To(MatchError(
				"'fixtures.PlainMessageA' is not a protocol buffers message",
			))
		})
	})

	Describe("func Unmarshal()", func() {
		It("unmarshals the data", func() {
			m := &fixtures.ProtoMessage{
				Value: "<value>",
			}

			data, err := codec.Marshal(m)
			Expect(err).ShouldNot(HaveOccurred())

			m = &fixtures.ProtoMessage{}
			err = codec.Unmarshal(data, m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m).To(Equal(
				&fixtures.ProtoMessage{
					Value: "<value>",
				},
			))
		})

		It("returns an error if the type is not a protocol buffers message", func() {
			m := fixtures.PlainMessageA{}
			err := codec.Unmarshal(nil, m)
			Expect(err).To(MatchError(
				"'fixtures.PlainMessageA' is not a protocol buffers message",
			))
		})
	})
})
