package jsonc

import "encoding/json/v2"

func Unmarshal(data []byte, v any, opts ...json.Options) error {
	processedData := stripComments(data)
	return json.Unmarshal(processedData, v, opts...)
}

func stripComments(data []byte) []byte {
	const (
		OUTSIDE = iota
		SINGLE_LINE
		MULTI_LINE
		IN_STRING
	)

	state := OUTSIDE
	result := make([]byte, len(data))
	copy(result, data)

	for i := 0; i < len(result); i++ {
		switch state {
		case OUTSIDE:
			if data[i] == '/' && i+1 < len(data) {
				if data[i+1] == '/' {
					state = SINGLE_LINE
					result[i] = ' '
					result[i+1] = ' '
					i++
				} else if data[i+1] == '*' {
					state = MULTI_LINE
					result[i] = ' '
					result[i+1] = ' '
					i++
				}
			} else if data[i] == '"' {
				state = IN_STRING
			}
		case SINGLE_LINE:
			if data[i] == '\n' {
				state = OUTSIDE
			} else {
				result[i] = ' '
			}
		case MULTI_LINE:
			if data[i] == '*' && i+1 < len(result) && data[i+1] == '/' {
				state = OUTSIDE
				result[i] = ' '
				result[i+1] = ' '
				i++
			} else if result[i] != '\n' {
				result[i] = ' '
			}
		case IN_STRING:
			if data[i] == '\\' && i+1 < len(data) {
				i++
			} else if data[i] == '"' {
				state = OUTSIDE
			}
		}
	}

	return result
}
