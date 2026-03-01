package jsonc

import (
	"encoding/json/v2"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Config struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestUnmarshalWithComments(t *testing.T) {
	testFilesDir := "../test_files/"

	expectedFiles := make(map[string][]byte)

	files, err := os.ReadDir(testFilesDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		filename := file.Name()
		if strings.Contains(filename, "_expected.json") {
			content, err := os.ReadFile(filepath.Join(testFilesDir, filename))
			if err != nil {
				t.Errorf("Failed to read file %s: %v", filename, err)
				continue
			}
			testName := strings.TrimSuffix(filename, "_expected.json")
			expectedFiles[testName] = content
		}
	}

	for _, file := range files {
		filename := file.Name()
		if strings.Contains(filename, "_input.json") {
			content, err := os.ReadFile(filepath.Join(testFilesDir, filename))
			if err != nil {
				t.Errorf("Failed to read file %s: %v", filename, err)
				continue
			}

			testName := strings.TrimSuffix(filename, "_input.json")
			expectedContent, exists := expectedFiles[testName]
			if !exists {
				t.Errorf("No expected file found for %s", testName)
				continue
			}

			var inputConfig, expectedConfig Config
			if err := Unmarshal(content, &inputConfig); err != nil {
				t.Errorf("Failed to unmarshal input from %s: %v", filename, err)
				continue
			}
			if err := json.Unmarshal(expectedContent, &expectedConfig); err != nil {
				t.Errorf("Failed to unmarshal expected content for %s: %v", testName, err)
				continue
			}

			if inputConfig != expectedConfig {
				t.Errorf("Mismatch in %s. Expected: %+v, Got: %+v", testName, expectedConfig, inputConfig)
			}
		}
	}
}

func TestStripComments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single line comment",
			input:    `{"name": "Alice"} // Comment`,
			expected: `{"name": "Alice"} `,
		},
		{
			name:     "Multi-line comment",
			input:    `{"name": "Alice" /* Comment */}`,
			expected: `{"name": "Alice" }`,
		},
		{
			name:     "Nested comments",
			input:    `{"name": "Alice" /* Comment // Nested comment */}`,
			expected: `{"name": "Alice" }`,
		},
		{
			name:     "Comment at the end of file without newline",
			input:    `{"name": "Alice"} // Comment`,
			expected: `{"name": "Alice"} `,
		},
		{
			name: "Comment at the end of file with newline",
			input: `{"name": "Alice"} // Comment
`,
			expected: "{\"name\": \"Alice\"} \n",
		},
		{
			name: "Multiple single line comments",
			input: `// Comment 1
			// Comment 2
			{"name": "Alice"}`,
			expected: "\n\t\t\t\n\t\t\t{\"name\": \"Alice\"}",
		},
		{
			name: "Multiple multi-line comments",
			input: `/* Comment 1 */
			/* Comment 2 */
			{"name": "Alice"}`,
			expected: "\n\t\t\t\n\t\t\t{\"name\": \"Alice\"}",
		},
		{
			name: "Mixed comments",
			input: `/* Comment 1 */
			// Comment 2
			{"name": "Alice"}`,
			expected: "\n\t\t\t\n\t\t\t{\"name\": \"Alice\"}",
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
			expected: `{"name": "//Alice"} `,
		},
		{
			name:     "Multi-line comment delimiters inside string",
			input:    `{"name": "/*Alice*/"} /* Comment */`,
			expected: `{"name": "/*Alice*/"} `,
		},
		{
			name:     "Empty input",
			input:    ``,
			expected: ``,
		},
		{
			name:     "Only single line comment",
			input:    `// Comment`,
			expected: ``,
		},
		{
			name:     "Only multi-line comment",
			input:    `/* Comment */`,
			expected: ``,
		},
		{
			name:     "Comment after a comma",
			input:    `{"name": "Alice", /* Comment */ "age": 30}`,
			expected: `{"name": "Alice",  "age": 30}`,
		},
		{
			name:     "Comment inside array",
			input:    `["Alice", /* Comment */ "Bob"]`,
			expected: `["Alice",  "Bob"]`,
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
			World"} `,
		},
		{
			name: "Newline inside string with multi-line comment after",
			input: `{"message": "Hello
			World"} /* Comment */`,
			expected: `{"message": "Hello
			World"} `,
		},
		{
			name:     "Multiple newlines inside string & single comment at the end",
			input:    "{\"message\": \"Hello\n\nWorld\"}//   ",
			expected: "{\"message\": \"Hello\n\nWorld\"}",
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
			got := stripComments([]byte(tt.input))
			if string(got) != tt.expected {
				t.Errorf("\ntest: \"%s\"\ngot   \"%s\"\nwant  \"%s\"", tt.input, string(got), tt.expected)
			} else {
				t.Logf("test \"%s\" passed", tt.name)
			}
		})
	}

}
