package betterjson

import "encoding/json/v2"

// Unmarshal parses the JSON-encoded data with comment support and stores the result
// in the value pointed to by v. Options may be passed to configure the behavior of
// encoding/json/v2's Unmarshal.
func Unmarshal(data []byte, v any, opts ...json.Options) error {
	processedData := decomment(data)
	return json.Unmarshal(processedData, v, opts...)
}
