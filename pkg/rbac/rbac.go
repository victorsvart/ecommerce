package rbac

import (
	"slices"
)

var (
	RolePermissions = map[uint][]string{
		UserRoleID: {
			GetUser,
			UpdateUser,
			ListProduct,
			GetProduct,
			CreateProduct,
			UpdateProduct,
			DeleteProduct,
		},
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
