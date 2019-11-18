package protobuf_test

import (
	"github.com/dogmatiq/marshalkit/internal/fixtures"
	. "github.com/dogmatiq/marshalkit/protobuf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type TextCodec", func() {
	var codec *TextCodec

	BeforeEach(func() {
		codec = &TextCodec{}
	})

	Describe("func MediaType()", func() {
		It("returns the expected media-type", func() {
			Expect(codec.MediaType()).To(Equal("text/vnd.google.protobuf"))
		})
	})

	Describe("func Marshal()", func() {
		It("marshals the value", func() {
			data, err := codec.Marshal(
				&fixtures.ProtoMessage{
					Value: "<value>",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).To(Equal(`value: "<value>"` + "\n"))
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
			data := []byte(`value: "<value>"` + "\n")

			m := &fixtures.ProtoMessage{}
			err := codec.Unmarshal(data, m)
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
