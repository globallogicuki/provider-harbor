package user

import (
	"context"
	"strconv"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
)

// Configure harbor_user resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_user", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.Kind = "User"
		
		// Handle external-name to numeric ID conversion for Harbor API compatibility
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, name string) {
			// When creating, use the username from spec
			if username, ok := base["username"].(string); ok {
				base["username"] = username
			}
		}
		
		// Custom external-name getter that converts username to numeric ID
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			// If we have a numeric ID in state, extract the username
			if id, ok := tfstate["id"].(string); ok && id != "" {
				// Harbor IDs are in format "/users/123" - we want to return the username
				// for external-name, but the numeric ID is used internally
				if strings.HasPrefix(id, "/users/") {
					// Get the username from the state
					if username, ok := tfstate["username"].(string); ok && username != "" {
						return username, nil
					}
					// Fallback to extracting from ID if username not available
					idPart := strings.TrimPrefix(id, "/users/")
					if _, err := strconv.Atoi(idPart); err != nil {
						// ID part is not numeric, it's already a username
						return idPart, nil
					}
					// We have a numeric ID but no username in state - this shouldn't happen
					return "", errors.New("numeric ID found but no username available")
				}
				return id, nil
			}
			return "", errors.New("no ID found in terraform state")
		}
		
		// Custom external-name setter that handles username to ID conversion
		r.ExternalName.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
			// The external name is the username, but Harbor API expects numeric IDs
			// We need to look up the user by username to get the numeric ID
			
			// For now, we'll let the terraform provider handle the lookup
			// This assumes our terraform provider fix is in place
			return externalName, nil
		}
	})
}
