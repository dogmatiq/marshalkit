package marshalkit_test

import (
	. "github.com/dogmatiq/marshalkit"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Packet", func() {
	Describe("func NewPacket()", func() {
		It("encodes the portable type name in the media-type", func() {
			p := NewPacket(
				"text/plain",
				"SomeType",
				[]byte("data"),
			)

			Expect(p).To(Equal(
				Packet{
					MediaType: "text/plain; type=SomeType",
					Data:      []byte("data"),
				},
			))
		})
	})

	Describe("func ParseMediaType()", func() {
		It("returns the mime-type and portable type name", func() {
			p := NewPacket(
				"text/plain",
				"SomeType",
				[]byte("data"),
			)

			mt, n, err := p.ParseMediaType()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(mt).To(Equal("text/plain"))
			Expect(n).To(Equal("SomeType"))
		})

		It("returns an error if the media-type has no type parameter", func() {
			p := Packet{
				MediaType: "text/plain",
			}

			_, _, err := p.ParseMediaType()
			Expect(err).To(MatchError("the media-type 'text/plain' does not specify a 'type' parameter"))
		})

		It("returns an error if the media-type is malformed", func() {
			p := Packet{
				MediaType: "<malformed>",
			}

			_, _, err := p.ParseMediaType()
			Expect(err).Should(HaveOccurred())
		})
	})
})
