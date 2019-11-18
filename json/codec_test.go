package json_test

import (
	"reflect"

	"github.com/dogmatiq/marshalkit/internal/fixtures"
	. "github.com/dogmatiq/marshalkit/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Codec", func() {
	var codec *Codec

	BeforeEach(func() {
		codec = &Codec{}
	})

	Describe("func Query()", func() {
		It("uses the unqualified type-name as the portable type", func() {
			rt := reflect.TypeOf(fixtures.PlainMessageA{})

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("PlainMessageA"))
		})

		It("uses the user-defined type name", func() {
			type LocalMessage fixtures.PlainMessageA
			rt := reflect.TypeOf(LocalMessage{})

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("LocalMessage"))
		})

		It("uses the element name for pointer types", func() {
			var m **fixtures.PlainMessageA
			rt := reflect.TypeOf(m)

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("PlainMessageA"))
		})
	})

	Describe("func MediaType()", func() {
		It("returns the expected media-type", func() {
			Expect(codec.MediaType()).To(Equal("application/json"))
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
			Expect(string(data)).To(Equal(`{"value":"\u003cvalue\u003e"}`))
		})
	})

	Describe("func Unmarshal()", func() {
		It("unmarshals the data", func() {
			data := []byte(`{"value":"\u003cvalue\u003e"}`)

			m := &fixtures.ProtoMessage{}
			err := codec.Unmarshal(data, m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m).To(Equal(
				&fixtures.ProtoMessage{
					Value: "<value>",
				},
			))
		})
	})
})
