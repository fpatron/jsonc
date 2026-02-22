package betterjson

import (
	"testing"
)

func TestDecommenter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single line comment",
			input:    `{"name": "Alice"} // Comment`,
			expected: `{"name": "Alice"}           `,
		},
		{
			name:     "Multi-line comment",
			input:    `{"name": "Alice" /* Comment */}`,
			expected: `{"name": "Alice"              }`,
		},
		{
			name:     "Nested comments",
			input:    `{"name": "Alice" /* Comment // Nested comment */}`,
			expected: `{"name": "Alice"                                }`,
		},
		{
			name:     "Comment at the end of file without newline",
			input:    `{"name": "Alice"} // Comment`,
			expected: `{"name": "Alice"}           `,
		},
		{
			name: "Comment at the end of file with newline",
			input: `{"name": "Alice"} // Comment
`,
			expected: `{"name": "Alice"}           
`,
		},
		{
			name: "Multiple single line comments",
			input: `// Comment 1
			// Comment 2
			{"name": "Alice"}`,
			expected: `            
			            
			{"name": "Alice"}`,
		},
		{
			name: "Multiple multi-line comments",
			input: `/* Comment 1 */
			/* Comment 2 */
			{"name": "Alice"}`,
			expected: `               
			               
			{"name": "Alice"}`,
		},
		{
			name: "Mixed comments",
			input: `/* Comment 1 */
			// Comment 2
			{"name": "Alice"}`,
			expected: `               
			            
			{"name": "Alice"}`,
		},
		{
			name:     "No comments",
			input:    `{"name": "Alice"}   `,
			expected: `{"name": "Alice"}   `,
		},
		{
			name:     " Single comment slashes in string",
			input:    `{"name": "//Alice"}   `,
			expected: `{"name": "//Alice"}   `,
		},
		{
			name:     " Multi comment slashes in string",
			input:    `{"name": "/*Alice*/"}   `,
			expected: `{"name": "/*Alice*/"}   `,
		},
		{
			name:     "Comment slashes inside string followed by actual comment",
			input:    `{"name": "//Alice"} // Comment`,
			expected: `{"name": "//Alice"}           `,
		},
		{
			name:     "Multi-line comment delimiters inside string",
			input:    `{"name": "/*Alice*/"} /* Comment */`,
			expected: `{"name": "/*Alice*/"}              `,
		},
		{
			name:     "Empty input",
			input:    ``,
			expected: ``,
		},
		{
			name:     "Only single line comment",
			input:    `// Comment`,
			expected: `          `,
		},
		{
			name:     "Only multi-line comment",
			input:    `/* Comment */`,
			expected: `             `,
		},
		{
			name:     "Comment after a comma",
			input:    `{"name": "Alice", /* Comment */ "age": 30}`,
			expected: `{"name": "Alice",               "age": 30}`,
		},
		{
			name:     "Comment inside array",
			input:    `["Alice", /* Comment */ "Bob"]`,
			expected: `["Alice",               "Bob"]`,
		},
		{
			name: "Newline inside string",
			input: `{"message": "Hello
			World"}`,
			expected: `{"message": "Hello
			World"}`,
		},
		{
			name: "Newline inside string followed by actual comment",
			input: `{"message": "Hello
			World"} // Comment`,
			expected: `{"message": "Hello
			World"}           `,
		},
		{
			name: "Newline inside string with multi-line comment after",
			input: `{"message": "Hello
			World"} /* Comment */`,
			expected: `{"message": "Hello
			World"}              `,
		},
		{
			name:     "Multiple newlines inside string & single comment at the end",
			input:    "{\"message\": \"Hello\n\nWorld\"}//   ",
			expected: "{\"message\": \"Hello\n\nWorld\"}     ",
		},
		{
			name: "Newline inside string with actual newline after",
			input: `{"message": "Hello
World"}
// Comment`,
			expected: `{"message": "Hello
World"}
          `,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decomment([]byte(tt.input))
			if string(got) != tt.expected {
				t.Errorf("\ntest: \"%s\"\ngot   \"%s\"\nwant  \"%s\"", tt.input, string(got), tt.expected)
			} else {
				t.Logf("test \"%s\" passed", tt.name)
			}
		})
	}

}
