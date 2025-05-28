package main

import (
	"fmt"
)

// User represents a user in the system
type User struct {
	Username string
	Role     string
}

// Resource represents a resource in the system
type Resource struct {
	Name string
}

// Authorizer represents the authorization service
type Authorizer struct {
	roles map[string][]string // Map of roles to their corresponding permissions
}

// NewAuthorizer creates a new instance of Authorizer
func NewAuthorizer() *Authorizer {
	roles := make(map[string][]string)
	roles["admin"] = []string{"read", "write", "delete"}
	roles["user"] = []string{"read"}

	return &Authorizer{roles}
}

// IsAuthorized checks if a user has permission to perform a certain action on a resource
func (a *Authorizer) IsAuthorized(user *User, action string, resource *Resource) bool {
	permissions, exists := a.roles[user.Role]
	if !exists {
		return false // User's role doesn't exist, so they are not authorized
	}

	for _, perm := range permissions {
		if perm == action {
			return true // User has the required permission
		}
	}

	return false // User doesn't have the required permission
}

func main() {
	// Create a new Authorizer
	authorizer := NewAuthorizer()

	// Create a user
	user := &User{
		Username: "alice",
		Role:     "user",
	}

	// Create a resource
	resource := &Resource{
		Name: "example.txt",
	}

	// Check if the user is authorized to read the resource
	if authorizer.IsAuthorized(user, "read", resource) {
		fmt.Printf("%s is authorized to read %s\n", user.Username, resource.Name)
	} else {
		fmt.Printf("%s is not authorized to read %s\n", user.Username, resource.Name)
	}
}	
