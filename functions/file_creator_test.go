package functions

import (
	"os"
	"testing"
)

// TestCreateFile tests the CreateFile function.
func TestCreateFile(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		content  string
		wantErr  bool
	}{
		{"Valid file creation", "testfile.txt", "Hello, World!", false},
		{"Empty file creation", "emptyfile.txt", "", false},
		{"Invalid path", "invalid/testfile.txt", "This should fail", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := CreateFile(tc.filePath, tc.content)

			if err != nil {
				if !tc.wantErr {
					t.Errorf("expected no error, but got %v", err)
				}
			} else {
				if tc.wantErr {
					t.Errorf("expected an error, but got none")
				} else {
					// Check if the file was created and validate content
					data, err := os.ReadFile(tc.filePath)
					if err != nil {
						t.Errorf("failed to read file: %v", err)
					} else if string(data) != tc.content {
						t.Errorf("expected content %q, but got %q", tc.content, string(data))
					}

					// Clean up
					os.Remove(tc.filePath)
				}
			}
		})
	}
}
