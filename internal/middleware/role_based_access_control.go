package middleware

type IAccessControl interface {
	RBAC(roles []interface{}, requiredRoles []string) bool
}

type roleBasedAccessControl struct{}

func NewRBAC() *roleBasedAccessControl {
	return &roleBasedAccessControl{}
}

func (rbac *roleBasedAccessControl) RBAC(roles []interface{}, requiredRoles []string) bool {
	if len(roles) == 0 {
		return false
	}

	//check if the user has all of the required roles
	for _, requiredRole := range requiredRoles {
		if !contains(roles, requiredRole) {
			return false
		}
	}
	return true
}

func contains(roles []interface{}, requiredRole string) bool {
	for _, role := range roles {
		if role == requiredRole {
			return true
		}
	}
	return false
}
