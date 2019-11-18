package protobuf

import (
	"reflect"

	. "github.com/dogmatiq/dogma/fixtures"
	. "github.com/dogmatiq/marshalkit/internal/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type TextCodec", func() {
	var codec codec

	Describe("func Query()", func() {
		It("uses the protocol name as the portable type", func() {
			rt := reflect.TypeOf(&ProtoMessage{})

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("dogmatiq.marshalkit.fixtures.ProtoMessage"))
		})

		It("excludes non-protocol-buffers types", func() {
			rt := reflect.TypeOf(MessageA{})

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types).To(BeEmpty())
		})
	})
})
