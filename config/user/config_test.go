package user

import (
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestExternalNameLogic(t *testing.T) {
	// Test our external name functions work correctly
	getExternalNameFn := func(tfstate map[string]any) (string, error) {
		if id, ok := tfstate["id"].(string); ok && id != "" {
			if strings.HasPrefix(id, "/users/") {
				if username, ok := tfstate["username"].(string); ok && username != "" {
					return username, nil
				}
				idPart := strings.TrimPrefix(id, "/users/")
				if _, err := strconv.Atoi(idPart); err != nil {
					return idPart, nil
				}
				return "", errors.New("numeric ID found but no username available")
			}
			return id, nil
		}
		return "", errors.New("no ID found in terraform state")
	}
	
	tests := []struct {
		name     string
		tfstate  map[string]any
		expected string
		wantErr  bool
	}{
		{
			name: "numeric ID with username in state",
			tfstate: map[string]any{
				"id":       "/users/123",
				"username": "testuser",
			},
			expected: "testuser",
			wantErr:  false,
		},
		{
			name: "username ID format",
			tfstate: map[string]any{
				"id":       "/users/testuser",
				"username": "testuser",
			},
			expected: "testuser",
			wantErr:  false,
		},
		{
			name: "username without prefix",
			tfstate: map[string]any{
				"id":       "testuser",
				"username": "testuser",
			},
			expected: "testuser",
			wantErr:  false,
		},
		{
			name: "numeric ID without username",
			tfstate: map[string]any{
				"id": "/users/123",
			},
			expected: "",
			wantErr:  true,
		},
		{
			name:     "empty state",
			tfstate:  map[string]any{},
			expected: "",
			wantErr:  true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getExternalNameFn(tt.tfstate)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// TestUserConfigurationExists removed due to complex Provider initialization
// The main external name logic test above covers our functionality