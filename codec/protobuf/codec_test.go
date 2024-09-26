package protobuf_test

import (
	"reflect"

	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/dogmatiq/marshalkit/codec/internal/fixtures"
	. "github.com/dogmatiq/marshalkit/codec/protobuf"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Codec", func() {
	var codec Codec

	Describe("func Query()", func() {
		It("uses the protocol name as the portable type", func() {
			rt := reflect.TypeFor[*ProtoMessage]()

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("dogmatiq.marshalkit.fixtures.ProtoMessage"))
		})

		It("excludes non-protocol-buffers types", func() {
			rt := reflect.TypeFor[CommandStub[TypeA]]()

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types).To(BeEmpty())
		})
	})
})
