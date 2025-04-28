package utils

import "errors"

type ContextKey string

func AuthorizeUser(userRole string, allowedRoles ...string) (bool, error) {
	for _, allowedRole := range allowedRoles {
		if userRole == allowedRole {
			return true, nil
		}
	}
	return false, errors.New("user not authorized")
}
