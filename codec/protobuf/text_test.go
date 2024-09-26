package protobuf_test

import (
	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/dogmatiq/marshalkit/codec/internal/fixtures"
	. "github.com/dogmatiq/marshalkit/codec/protobuf"
	. "github.com/jmalloc/gomegax"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type TextCodec (configured for text format)", func() {
	var codec Codec

	BeforeEach(func() {
		codec = DefaultTextCodec
	})

	Describe("func BasicMediaType()", func() {
		It("returns the expected basic media-type", func() {
			Expect(codec.BasicMediaType()).To(Equal("text/vnd.google.protobuf"))
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

			// Note that we need to use a regex to match an arbitrary amount of
			// whitespace in between the key and value as a result of the
			// behavior described in
			// https://github.com/golang/protobuf/issues/1121.
			Expect(data).To(MatchRegexp(`value:\s+\"\<value\>\"\n`))
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
			data := []byte(`value: "<value>"` + "\n")

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
