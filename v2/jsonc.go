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
	hasComment := false
scanLoop:
	for i := 0; i < len(data); i++ {
		switch state {
		case OUTSIDE:
			if data[i] == '/' && i+1 < len(data) && (data[i+1] == '/' || data[i+1] == '*') {
				hasComment = true
				break scanLoop
			}
			if data[i] == '"' {
				state = IN_STRING
			}
		case IN_STRING:
			if data[i] == '\\' && i+1 < len(data) {
				i++
			} else if data[i] == '"' {
				state = OUTSIDE
			}
		}
	}

	if !hasComment {
		return data
	}

	state = OUTSIDE
	result := make([]byte, 0, len(data))

	for i := 0; i < len(data); i++ {
		switch state {
		case OUTSIDE:
			if data[i] == '/' && i+1 < len(data) {
				if data[i+1] == '/' {
					state = SINGLE_LINE
					i++
				} else if data[i+1] == '*' {
					state = MULTI_LINE
					i++
				} else {
					result = append(result, data[i])
				}
			} else {
				if data[i] == '"' {
					state = IN_STRING
				}
				result = append(result, data[i])
			}
		case SINGLE_LINE:
			if data[i] == '\n' {
				state = OUTSIDE
				result = append(result, '\n')
			}
		case MULTI_LINE:
			if data[i] == '*' && i+1 < len(data) && data[i+1] == '/' {
				state = OUTSIDE
				i++
			}
		case IN_STRING:
			result = append(result, data[i])
			if data[i] == '\\' && i+1 < len(data) {
				i++
				result = append(result, data[i])
			} else if data[i] == '"' {
				state = OUTSIDE
			}
		}
	}

	return result
}
