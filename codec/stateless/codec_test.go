package stateless_test

import (
	"reflect"

	"github.com/dogmatiq/dogma"
	. "github.com/dogmatiq/marshalkit/codec/stateless"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Codec", func() {
	var codec *Codec

	BeforeEach(func() {
		codec = &Codec{}
	})

	Describe("func Query()", func() {
		It("uses a short string as the portable type", func() {
			rt := reflect.TypeOf(dogma.StatelessProcessRoot)

			caps := codec.Query(
				[]reflect.Type{rt},
			)

			Expect(caps.Types[rt]).To(Equal("process"))
		})
	})

	Describe("func BasicMediaType()", func() {
		It("returns the expected basic media-type", func() {
			Expect(codec.BasicMediaType()).To(Equal("application/x-empty"))
		})
	})

	Describe("func Marshal()", func() {
		It("returns an empty byte-slice", func() {
			data, err := codec.Marshal(dogma.StatelessProcessRoot)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(BeEmpty())
		})

		It("returns an error if passed any value other than dogma.StatelessProcessRoot", func() {
			_, err := codec.Marshal(123)
			Expect(err).To(MatchError("int is not dogma.StatelessProcessRoot"))
		})
	})

	Describe("func Unmarshal()", func() {
		It("does nothing on success", func() {
			before := dogma.StatelessProcessRoot

			err := codec.Unmarshal(nil, &dogma.StatelessProcessRoot)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dogma.StatelessProcessRoot).To(Equal(before))
		})
		It("returns an error if passed a non-empty byte-slice", func() {
			data := []byte(` `)
			err := codec.Unmarshal(data, &dogma.StatelessProcessRoot)
			Expect(err).To(MatchError("expected empty data, got 1 byte(s)"))
		})

		It("returns an error if passed any value other than the address of dogma.StatelessProcessRoot", func() {
			err := codec.Unmarshal(nil, 123)
			Expect(err).To(MatchError("int is not a pointer to dogma.StatelessProcessRoot"))
		})
	})
})
