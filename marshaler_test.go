package marshalkit_test

import (
	. "github.com/dogmatiq/dogma/fixtures"
	. "github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/fixtures" // can't dot-import due to conflicts
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func MustMarshal()", func() {
	It("marshals the value using the marshaler", func() {
		p := MustMarshal(
			fixtures.Marshaler,
			&AggregateRoot{
				Value: "<value>",
			},
		)
		Expect(p.MediaType).To(Equal("application/json; type=AggregateRoot"))
		Expect(p.Data).To(Equal([]byte(`{"Value":"\u003cvalue\u003e"}`)))
	})

	It("panics if the type is not registered", func() {
		Expect(func() {
			MustMarshal(
				fixtures.Marshaler,
				"<scalar value>",
			)
		}).To(Panic())
	})
})

var _ = Describe("func MustUnmarshal()", func() {
	It("unmarshals the value using the marshaler", func() {
		v := MustUnmarshal(
			fixtures.Marshaler,
			Packet{
				"application/json; type=AggregateRoot",
				[]byte(`{"Value":"\u003cvalue\u003e"}`),
			},
		)
		Expect(v).To(Equal(
			&AggregateRoot{
				Value: "<value>",
			},
		))
	})

	It("returns an error if the type is not registered", func() {
		Expect(func() {
			MustUnmarshal(
				fixtures.Marshaler,
				Packet{
					"application/json; type=Unsupported",
					[]byte(`{"Value":"\u003cvalue\u003e"}`),
				},
			)
		}).To(Panic())
	})
})
