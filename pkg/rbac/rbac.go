package rbac

import "slices"

var (
	RolePermissions = map[uint][]string{
		UserRoleID: {},
	}
)

func HasPermission(roleID uint, perm string) bool {
	if roleID == AdminRoleID {
		return true
	}

	perms, ok := RolePermissions[roleID]
	if !ok {
		return false
	}

	if slices.Contains(perms, perm) {
		return true
	}

	return false
}
