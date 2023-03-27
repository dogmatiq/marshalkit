package protobuf

import (
	"reflect"

	. "github.com/dogmatiq/dogma/fixtures"
	. "github.com/dogmatiq/marshalkit/codec/internal/fixtures"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Codec", func() {
	var codec Codec

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
