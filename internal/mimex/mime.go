package mimex

import (
	"fmt"
	"mime"
)

// FormatMediaType returns the media-type as the base media type and the
// message's portable name as the 'type' parameter.
func FormatMediaType(base string, portableName string) string {
	return mime.FormatMediaType(
		base,
		map[string]string{"type": portableName},
	)
}

// ParseMediaType returns the media-type and the portable type name encoded in
// the packet's MIME media-type.
func ParseMediaType(mediatype string) (string, string, error) {
	mt, params, err := mime.ParseMediaType(mediatype)
	if err != nil {
		return "", "", err
	}

	if n, ok := params["type"]; ok {
		return mt, n, nil
	}

	return "", "", fmt.Errorf(
		"the media-type '%s' does not specify a 'type' parameter",
		mediatype,
	)
}
