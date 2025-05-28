package main

import (
	"fmt"
	"net/http"
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

// AuthorizationService handles authorization logic
type AuthorizationService struct {
}

// IsAuthorized checks if a user has permission to access a resource
func (a *AuthorizationService) IsAuthorized(user *User, resource *Resource) bool {
	// Define role-based access control rules
	rolePermissions := map[string][]string{
		"admin": {"read", "write", "delete"},
		"user":  {"read"},
	}

	// Check if the user's role has permission to access the resource
	permissions, exists := rolePermissions[user.Role]
	if !exists {
		return false // Role doesn't exist, so user is not authorized
	}

	// Check if the resource's required action is allowed for the user's role
	requiredAction := "read" // Example action
	for _, perm := range permissions {
		if perm == requiredAction {
			return true // User is authorized to perform the action on the resource
		}
	}

	return false // User's role does not have permission for the required action
}

func main() {
	// Create an instance of AuthorizationService
	authService := &AuthorizationService{}

	// Create a user
	user := &User{
		Username: "john",
		Role:     "user",
	}

	// Create a resource
	resource := &Resource{
		Name: "example.txt",
	}

	// Check if the user is authorized to access the resource
	if authService.IsAuthorized(user, resource) {
		fmt.Printf("%s is authorized to access %s\n", user.Username, resource.Name)
	} else {
		fmt.Printf("%s is not authorized to access %s\n", user.Username, resource.Name)
	}

	// Output: john is authorized to access example.txt
}
