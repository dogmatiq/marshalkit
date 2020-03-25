package marshalkit

import (
	"fmt"
	"mime"
)

// Packet is a container of marshaled data and its related meta-data.
type Packet struct {
	// MediaType is a MIME media-type describing the content and encoding of the
	// binary data.
	MediaType string

	// Data is the marshaled binary data.
	Data []byte
}

// NewPacket returns a new packet.
//
// mt is the MIME media-type describing the content and encoding of the binary
// data. n is the marshaled value's portable type name.
func NewPacket(mt string, n string, data []byte) Packet {
	return Packet{
		mime.FormatMediaType(
			mt,
			map[string]string{"type": n},
		),
		data,
	}
}

// ParseMediaType returns the media-type and the portable type name encoded in
// the packet's MIME media-type.
func (p *Packet) ParseMediaType() (string, string, error) {
	mt, params, err := mime.ParseMediaType(p.MediaType)
	if err != nil {
		return "", "", err
	}

	if n, ok := params["type"]; ok {
		return mt, n, nil
	}

	return "", "", fmt.Errorf(
		"the media-type '%s' does not specify a 'type' parameter",
		p.MediaType,
	)
}
