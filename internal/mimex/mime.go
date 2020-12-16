package mimex

// ParseMediaType returns the media-type and the portable type name encoded in
// the packet's MIME media-type.
// func ParseMediaType(mt string) (string, string, error) {
// 	mt, params, err := mime.ParseMediaType(p.MediaType)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	if n, ok := params["type"]; ok {
// 		return mt, n, nil
// 	}

// 	return "", "", fmt.Errorf(
// 		"the media-type '%s' does not specify a 'type' parameter",
// 		p.MediaType,
// 	)
// }
