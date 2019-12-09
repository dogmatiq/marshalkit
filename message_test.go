package marshalkit_test

import (
	. "github.com/dogmatiq/dogma/fixtures"
	. "github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/fixtures" // can't dot-import due to conflicts
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func MarshalMessage()", func() {
	It("marshals the message using the marshaler", func() {
		p, err := MarshalMessage(
			fixtures.Marshaler,
			MessageA{
				Value: "<value>",
			},
		)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(p.MediaType).To(Equal("application/json; type=MessageA"))
		Expect(p.Data).To(Equal([]byte(`{"Value":"\u003cvalue\u003e"}`)))
	})

	It("returns an error if the type is not registered", func() {
		_, err := MarshalMessage(
			fixtures.Marshaler,
			"<scalar message>",
		)
		Expect(err).Should(MatchError("no codecs support the 'string' type"))
	})
})

var _ = Describe("func UnmarshalMessage()", func() {
	It("unmarshals the message using the marshaler", func() {
		m, err := UnmarshalMessage(
			fixtures.Marshaler,
			Packet{
				"application/json; type=MessageA",
				[]byte(`{"Value":"\u003cvalue\u003e"}`),
			},
		)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(m).To(Equal(
			MessageA{
				Value: "<value>",
			},
		))
	})

	It("returns an error if the type is not registered", func() {
		_, err := UnmarshalMessage(
			fixtures.Marshaler,
			Packet{
				"application/json; type=Unsupported",
				[]byte(`{"Value":"\u003cvalue\u003e"}`),
			},
		)
		Expect(err).To(MatchError("the portable type name 'Unsupported' is not recognized"))
	})
})

var _ = Describe("func MustMarshalMessage()", func() {
	It("marshals the message using the marshaler", func() {
		p := MustMarshalMessage(
			fixtures.Marshaler,
			MessageA{
				Value: "<value>",
			},
		)
		Expect(p.MediaType).To(Equal("application/json; type=MessageA"))
		Expect(p.Data).To(Equal([]byte(`{"Value":"\u003cvalue\u003e"}`)))
	})

	It("panics if the type is not registered", func() {
		Expect(func() {
			MustMarshalMessage(
				fixtures.Marshaler,
				"<scalar message>",
			)
		}).To(Panic())
	})
})

var _ = Describe("func MustUnmarshalMessage()", func() {
	It("unmarshals the message using the marshaler", func() {
		m := MustUnmarshalMessage(
			fixtures.Marshaler,
			Packet{
				"application/json; type=MessageA",
				[]byte(`{"Value":"\u003cvalue\u003e"}`),
			},
		)
		Expect(m).To(Equal(
			MessageA{
				Value: "<value>",
			},
		))
	})

	It("returns an error if the type is not registered", func() {
		Expect(func() {
			MustUnmarshalMessage(
				fixtures.Marshaler,
				Packet{
					"application/json; type=Unsupported",
					[]byte(`{"Value":"\u003cvalue\u003e"}`),
				},
			)
		}).To(Panic())
	})
})
