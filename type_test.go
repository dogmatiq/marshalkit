package marshalkit_test

import (
	"reflect"

	. "github.com/dogmatiq/dogma/fixtures"
	. "github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/fixtures" // can't dot-import due to conflicts
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func MustMarshalType()", func() {
	It("marshals the type using the marshaler", func() {
		n := MustMarshalType(
			fixtures.Marshaler,
			reflect.TypeOf(&AggregateRoot{}),
		)
		Expect(n).To(Equal("AggregateRoot"))
	})

	It("panics if the type is not registered", func() {
		Expect(func() {
			MustMarshalType(
				fixtures.Marshaler,
				reflect.TypeOf("<scalar value>"),
			)
		}).To(Panic())
	})
})

var _ = Describe("func MustUnmarshalType()", func() {
	It("unmarshals the type using the marshaler", func() {
		rt := MustUnmarshalType(
			fixtures.Marshaler,
			"AggregateRoot",
		)
		Expect(rt).To(Equal(
			reflect.TypeOf(&AggregateRoot{}),
		))
	})

	It("returns an error if the type is not registered", func() {
		Expect(func() {
			MustUnmarshalType(
				fixtures.Marshaler,
				"Unsupported",
			)
		}).To(Panic())
	})
})
