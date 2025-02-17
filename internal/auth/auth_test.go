package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected string
		wantErr  string
	}{
		{
			wantErr: "no authorization header",
		},
		{
			key:     "Authorization",
			wantErr: "no authorization header",
		},
		{
			key:     "Authorization",
			value:   "-",
			wantErr: "malformed authorization header",
		},
		{
			key:      "Authorization",
			value:    "ApiKey xxxxxx",
			expected: "xxxxxx",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case: %v", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil && test.wantErr == "" {
				t.Fatalf("Unexpected Error: %s", err.Error())
			}
			if err != nil && !strings.Contains(err.Error(), test.wantErr) {
				t.Fatalf("Expected Error: %v, Got: %v", test.wantErr, err.Error())
			}
			if err == nil && test.wantErr != "" {
				t.Fatalf("Expected Error: %v, Got: <nil>", test.wantErr)
			}

			if test.expected != output {
				t.Fatalf("Expected: %v, Got: %v", test.expected, output)
			}
		})
	}
}
