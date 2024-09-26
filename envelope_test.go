package marshalkit_test

import (
	"time"

	"github.com/dogmatiq/configkit"
	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/dogmatiq/interopspec/envelopespec"
	. "github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/fixtures"
	. "github.com/jmalloc/gomegax"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func MustMarshalMessageIntoEnvelope()", func() {
	It("marshals the message into the envelope", func() {
		var env Envelope

		MustMarshalMessageIntoEnvelope(
			fixtures.Marshaler,
			CommandA1,
			&env,
		)

		Expect(&env).To(EqualX(
			&Envelope{
				PortableName: "CommandStub[TypeA]",
				MediaType:    `application/json; type="CommandStub[TypeA]"`,
				Data:         []byte(`{"content":"A1"}`),
			},
		))
	})
})

var _ = Describe("func MustMarshalCommandIntoEnvelope()", func() {
	It("marshals the message into the envelope", func() {
		var env Envelope

		MustMarshalCommandIntoEnvelope(
			fixtures.Marshaler,
			CommandA1,
			&env,
		)

		Expect(&env).To(EqualX(
			&Envelope{
				PortableName: "CommandStub[TypeA]",
				MediaType:    `application/json; type="CommandStub[TypeA]"`,
				Data:         []byte(`{"content":"A1"}`),
			},
		))
	})
})

var _ = Describe("func MustMarshalEventIntoEnvelope()", func() {
	It("marshals the message into the envelope", func() {
		var env Envelope

		MustMarshalEventIntoEnvelope(
			fixtures.Marshaler,
			EventA1,
			&env,
		)

		Expect(&env).To(EqualX(
			&Envelope{
				PortableName: "EventStub[TypeA]",
				MediaType:    `application/json; type="EventStub[TypeA]"`,
				Data:         []byte(`{"content":"A1"}`),
			},
		))
	})
})

var _ = Describe("func MustMarshalTimeoutIntoEnvelope()", func() {
	It("marshals the message into the envelope", func() {
		var env Envelope

		MustMarshalTimeoutIntoEnvelope(
			fixtures.Marshaler,
			TimeoutA1,
			&env,
		)

		Expect(&env).To(EqualX(
			&Envelope{
				PortableName: "TimeoutStub[TypeA]",
				MediaType:    `application/json; type="TimeoutStub[TypeA]"`,
				Data:         []byte(`{"content":"A1"}`),
			},
		))
	})
})

var _ = Describe("func UnmarshalMessageFromEnvelope()", func() {
	It("unmarshals the message from the envelope", func() {
		env := &Envelope{
			PortableName: "CommandStub[TypeA]",
			MediaType:    `application/json; type="CommandStub[TypeA]"`,
			Data:         []byte(`{"content":"A1"}`),
		}

		m, err := UnmarshalMessageFromEnvelope(fixtures.Marshaler, env)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(m).To(Equal(CommandA1))
	})
})

var _ = Describe("func UnmarshalCommandFromEnvelope()", func() {
	It("unmarshals the message from the envelope", func() {
		env := &Envelope{
			PortableName: "CommandStub[TypeA]",
			MediaType:    `application/json; type="CommandStub[TypeA]"`,
			Data:         []byte(`{"content":"A1"}`),
		}

		m, err := UnmarshalCommandFromEnvelope(fixtures.Marshaler, env)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(m).To(Equal(CommandA1))
	})
})

var _ = Describe("func UnmarshalEventFromEnvelope()", func() {
	It("unmarshals the message from the envelope", func() {
		env := &Envelope{
			PortableName: "EventStub[TypeA]",
			MediaType:    `application/json; type="EventStub[TypeA]"`,
			Data:         []byte(`{"content":"A1"}`),
		}

		m, err := UnmarshalEventFromEnvelope(fixtures.Marshaler, env)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(m).To(Equal(EventA1))
	})
})

var _ = Describe("func UnmarshalTimeoutFromEnvelope()", func() {
	It("unmarshals the message from the envelope", func() {
		env := &Envelope{
			PortableName: "TimeoutStub[TypeA]",
			MediaType:    `application/json; type="TimeoutStub[TypeA]"`,
			Data:         []byte(`{"content":"A1"}`),
		}

		m, err := UnmarshalTimeoutFromEnvelope(fixtures.Marshaler, env)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(m).To(Equal(TimeoutA1))
	})
})

var _ = Describe("func MustMarshalEnvelopeIdentity()", func() {
	It("returns the protocol buffers identity", func() {
		in := configkit.MustNewIdentity("<name>", "81ff36f1-96ba-401a-8291-024a725cf60c")

		out := MustMarshalEnvelopeIdentity(in)
		Expect(out).To(EqualX(
			&Identity{
				Name: "<name>",
				Key:  "81ff36f1-96ba-401a-8291-024a725cf60c",
			},
		))
	})
})

var _ = Describe("func UnmarshalEnvelopeIdentity()", func() {
	It("returns the configkit identity", func() {
		in := &Identity{
			Name: "<name>",
			Key:  "ed555abf-d9fd-45e5-9725-542e47a61667",
		}

		out, err := UnmarshalEnvelopeIdentity(in)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(out).To(Equal(
			configkit.MustNewIdentity("<name>", "ed555abf-d9fd-45e5-9725-542e47a61667"),
		))
	})
})

var _ = Describe("func MustMarshalEnvelopeTime()", func() {
	It("returns the time formatted as per RFC-3339", func() {
		in := time.Date(2001, 02, 03, 04, 05, 06, 0, time.UTC)

		out := MustMarshalEnvelopeTime(in)
		Expect(out).To(Equal("2001-02-03T04:05:06Z"))
	})

	It("returns an empty string if the time is the zero-value", func() {
		var in time.Time

		out := MustMarshalEnvelopeTime(in)
		Expect(out).To(Equal(""))
	})
})

var _ = Describe("func UnmarshalEnvelopeTime()", func() {
	It("parses the time from RFC-3339 format", func() {
		in := "2001-02-03T04:05:06Z"

		out, err := UnmarshalEnvelopeTime(in)
		Expect(err).ShouldNot(HaveOccurred())

		expect := time.Date(2001, 02, 03, 04, 05, 06, 0, time.UTC)
		Expect(out).To(BeTemporally("==", expect))
	})

	It("returns the zero-value if the input is empty", func() {
		out, err := UnmarshalEnvelopeTime("")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(out.IsZero()).To(BeTrue())
	})
})
