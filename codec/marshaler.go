package codec

import (
	"fmt"
	"reflect"

	"github.com/dogmatiq/marshalkit"
	"github.com/dogmatiq/marshalkit/internal/mimex"
)

// Marshaler uses a set of priority-ordered codecs to marshal and unmarshal
// types and values.
type Marshaler struct {
	types map[reflect.Type]struct {
		defaultEncoder      Codec
		defaultPortableName string
		mediaTypes          []string
	}
	codecByBasicMediaType map[string]Codec
	typeByPortableName    map[string]reflect.Type
}

// NewMarshaler returns a new marshaler that uses the given set of codecs to
// marshal and unmarshal values.
//
// The codecs are given in order of preference.
func NewMarshaler(
	types []reflect.Type,
	codecs []Codec,
) (*Marshaler, error) {
	m := &Marshaler{
		types: map[reflect.Type]struct {
			defaultEncoder      Codec
			defaultPortableName string
			mediaTypes          []string
		}{},
		codecByBasicMediaType: map[string]Codec{},
		typeByPortableName:    map[string]reflect.Type{},
	}

	// build a list of all of the "unsupported" types
	unsupported := map[reflect.Type]struct{}{}
	for _, rt := range types {
		unsupported[rt] = struct{}{}
	}

	// scan the codecs in order of preference
	for _, c := range codecs {
		caps := c.Query(types)

		if len(caps.Types) > 0 {
			for rt, n := range caps.Types {
				if x, ok := m.typeByPortableName[n]; ok && x != rt {
					return nil, fmt.Errorf(
						"the type name '%s' is used by both '%s' and '%s'",
						n,
						x,
						rt,
					)
				}

				m.typeByPortableName[n] = rt

				bt, ok := m.types[rt]

				if !ok {
					bt.defaultEncoder = c
					bt.defaultPortableName = n
				}

				bt.mediaTypes = append(
					bt.mediaTypes,
					mimex.FormatMediaType(c.BasicMediaType(), n),
				)

				m.types[rt] = bt

				delete(unsupported, rt)
			}

			if _, ok := m.codecByBasicMediaType[c.BasicMediaType()]; ok {
				return nil, fmt.Errorf(
					"multiple codecs use the '%s' media-type",
					c.BasicMediaType(),
				)
			}

			m.codecByBasicMediaType[c.BasicMediaType()] = c
		}
	}

	for rt := range unsupported {
		return nil, fmt.Errorf(
			"no codecs support the '%s' type",
			rt,
		)
	}

	return m, nil
}

// MarshalType marshals a type to its portable representation.
func (m *Marshaler) MarshalType(rt reflect.Type) (string, error) {
	if bt, ok := m.types[rt]; ok {
		return bt.defaultPortableName, nil
	}

	return "", fmt.Errorf(
		"no codecs support the '%s' type",
		rt,
	)
}

// UnmarshalType unmarshals a type from its portable representation.
func (m *Marshaler) UnmarshalType(n string) (reflect.Type, error) {
	if rt, ok := m.typeByPortableName[n]; ok {
		return rt, nil
	}

	return nil, fmt.Errorf(
		"the portable type name '%s' is not recognized",
		n,
	)
}

// Marshal returns a binary representation of v.
func (m *Marshaler) Marshal(v any) (marshalkit.Packet, error) {
	rt := reflect.TypeOf(v)

	if bt, ok := m.types[rt]; ok {
		data, err := bt.defaultEncoder.Marshal(v)
		if err != nil {
			return marshalkit.Packet{}, err
		}

		return marshalkit.NewPacket(
			bt.defaultEncoder.BasicMediaType(),
			bt.defaultPortableName,
			data,
		), nil
	}

	return marshalkit.Packet{}, fmt.Errorf(
		"no codecs support the '%s' type",
		rt,
	)
}

// MarshalAs returns a binary representation of v encoded using a format
// associated with one of the supplied media-types.
//
// mediaTypes is a list of acceptible media-types, in order of preference.
// If none of the media-types are supported, ok is false.
func (m *Marshaler) MarshalAs(
	v any,
	mediaTypes []string,
) (p marshalkit.Packet, ok bool, err error) {
	if len(mediaTypes) == 0 {
		panic("at least one media-type must be provided")
	}

	for _, mt := range mediaTypes {
		rt := reflect.TypeOf(v)

		basic, n, err := mimex.ParseMediaType(mt)
		if err != nil {
			return marshalkit.Packet{}, false, err
		}

		if c, ok := m.codecByBasicMediaType[basic]; ok && m.typeByPortableName[n] == rt {
			data, err := c.Marshal(v)
			if err != nil {
				return marshalkit.Packet{}, false, err
			}

			return marshalkit.NewPacket(
				c.BasicMediaType(),
				n,
				data,
			), true, nil
		}
	}

	return marshalkit.Packet{}, false, nil
}

// Unmarshal produces a value from its binary representation.
func (m *Marshaler) Unmarshal(p marshalkit.Packet) (any, error) {
	c, rt, err := m.unpackMediaType(p)
	if err != nil {
		return nil, err
	}

	// If the type is already a pointer, we wan't to construct the element that
	// it points to, otherwise construct a new pointer to the actual type.
	var v reflect.Value
	if rt.Kind() == reflect.Ptr {
		v = reflect.New(rt.Elem())
	} else {
		v = reflect.New(rt)
	}

	if err := c.Unmarshal(p.Data, v.Interface()); err != nil {
		return nil, err
	}

	// Unwrap the pointer we created just to allow for unmarshalling.
	if rt.Kind() != reflect.Ptr {
		v = v.Elem()
	}

	return v.Interface(), nil
}

// MediaTypesFor returns the media-types that the marshaler can use to
// represent the given type, in order of preference.
//
// It returns an empty slice if the type is not supported.
func (m *Marshaler) MediaTypesFor(rt reflect.Type) []string {
	return m.types[rt].mediaTypes
}

func (m *Marshaler) unpackMediaType(p marshalkit.Packet) (Codec, reflect.Type, error) {
	mt, n, err := p.ParseMediaType()
	if err != nil {
		return nil, nil, err
	}

	c, ok := m.codecByBasicMediaType[mt]
	if !ok {
		return nil, nil, fmt.Errorf(
			"no codecs support the '%s' media-type",
			mt,
		)
	}

	rt, ok := m.typeByPortableName[n]
	if !ok {
		return nil, nil, fmt.Errorf(
			"the portable type name '%s' is not recognized",
			n,
		)
	}

	return c, rt, nil
}
