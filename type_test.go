package marshalkit_test

import (
	"reflect"

	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/fixtures" // can't dot-import due to conflicts
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func MustMarshalType()", func() {
	It("marshals the type using the marshaler", func() {
		n := MustMarshalType(
			fixtures.Marshaler,
			reflect.TypeFor[*AggregateRootStub](),
		)
		Expect(n).To(Equal("AggregateRootStub"))
	})

	It("panics if the type is not registered", func() {
		Expect(func() {
			MustMarshalType(
				fixtures.Marshaler,
				reflect.TypeFor[string](),
			)
		}).To(Panic())
	})
})

var _ = Describe("func MustUnmarshalType()", func() {
	It("unmarshals the type using the marshaler", func() {
		rt := MustUnmarshalType(
			fixtures.Marshaler,
			"AggregateRootStub",
		)
		Expect(rt).To(Equal(
			reflect.TypeFor[*AggregateRootStub](),
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
