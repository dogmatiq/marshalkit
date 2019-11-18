package marshalkit_test

import (
	"reflect"

	. "github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/internal/fixtures"
	"github.com/dogmatiq/marshalkit/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Context("messages", func() {
	var marshaler *Marshaler

	BeforeEach(func() {
		var err error
		marshaler, err = NewMarshaler(
			[]reflect.Type{
				reflect.TypeOf(fixtures.PlainMessageA{}),
			},
			[]Codec{
				&json.Codec{},
			},
		)
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("func MarshalMessage()", func() {
		It("marshals the message using the marshaler", func() {
			p, err := MarshalMessage(
				marshaler,
				fixtures.PlainMessageA{
					Value: "<value>",
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(p.MediaType).To(Equal("application/json; type=PlainMessageA"))
			Expect(p.Data).To(Equal([]byte(`{"Value":"\u003cvalue\u003e"}`)))
		})
	})

	Describe("func MustMarshalMessage()", func() {
		It("marshals the message using the marshaler", func() {
			p := MustMarshalMessage(
				marshaler,
				fixtures.PlainMessageA{
					Value: "<value>",
				},
			)
			Expect(p.MediaType).To(Equal("application/json; type=PlainMessageA"))
			Expect(p.Data).To(Equal([]byte(`{"Value":"\u003cvalue\u003e"}`)))
		})

		It("panics if marshaling fails", func() {
			Expect(func() {
				MustMarshalMessage(
					marshaler,
					fixtures.PlainMessageC{},
				)

			}).To(Panic())
		})
	})

	Describe("func UnmarshalMessage()", func() {
		It("unmarshals the message using the marshaler", func() {
			m, err := UnmarshalMessage(
				marshaler,
				Packet{
					"application/json; type=PlainMessageA",
					[]byte(`{"Value":"\u003cvalue\u003e"}`),
				},
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m).To(Equal(
				fixtures.PlainMessageA{
					Value: "<value>",
				},
			))
		})

		It("returns an error if the type is not registered", func() {
			_, err := UnmarshalMessage(
				marshaler,
				Packet{
					"application/json; type=PlainMessageC",
					[]byte(`{"Value":"\u003cvalue\u003e"}`),
				},
			)
			Expect(err).To(MatchError("the portable type name 'PlainMessageC' is not recognized"))
		})
	})

	Describe("func MustUnmarshalMessage()", func() {
		It("unmarshals the message using the marshaler", func() {
			m := MustUnmarshalMessage(
				marshaler,
				Packet{
					"application/json; type=PlainMessageA",
					[]byte(`{"Value":"\u003cvalue\u003e"}`),
				},
			)
			Expect(m).To(Equal(
				fixtures.PlainMessageA{
					Value: "<value>",
				},
			))
		})

		It("panics if the type is not registered", func() {
			Expect(func() {
				MustUnmarshalMessage(
					marshaler,
					Packet{
						"application/json; type=PlainMessageC",
						[]byte(`{"Value":"\u003cvalue\u003e"}`),
					},
				)
			}).To(Panic())
		})
	})
})
