package json_test

import (
	"reflect"

	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/dogmatiq/marshalkit/codec/internal/fixtures"
	. "github.com/dogmatiq/marshalkit/codec/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Codec", func() {
	var codec *Codec

	BeforeEach(func() {
		codec = &Codec{}
	})

	Describe("func Query()", func() {
		It("uses the unqualified type-name as the portable type", func() {
			rt := reflect.TypeFor[CommandStub[TypeA]]()

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("CommandStub[TypeA]"))
		})

		It("uses the user-defined type name", func() {
			type LocalMessage CommandStub[TypeA]
			rt := reflect.TypeFor[LocalMessage]()

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("LocalMessage"))
		})

		It("uses the element name for pointer types", func() {
			rt := reflect.TypeFor[**CommandStub[TypeA]]()

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("CommandStub[TypeA]"))
		})
	})

	Describe("func BasicMediaType()", func() {
		It("returns the expected basic media-type", func() {
			Expect(codec.BasicMediaType()).To(Equal("application/json"))
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
			Expect(string(data)).To(Equal(`{"value":"\u003cvalue\u003e"}`))
		})
	})

	Describe("func Unmarshal()", func() {
		It("unmarshals the data", func() {
			data := []byte(`{"value":"\u003cvalue\u003e"}`)

			m := &ProtoMessage{}
			err := codec.Unmarshal(data, m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m).To(Equal(
				&ProtoMessage{
					Value: "<value>",
				},
			))
		})
	})
})
