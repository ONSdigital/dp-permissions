package authv2

import (
	"context"

	"github.com/ONSdigital/log.go/log"
)

type PermissionsVerifier struct {
}

func NewPermissionsVerifier() *PermissionsVerifier {
	return &PermissionsVerifier{}
}

// CheckAuthorisation check the actual Permissions satisfy the required Permissions. Returns nil if requirements are satisfied
// returns callerForbiddenError otherwise.
func (verifier *PermissionsVerifier) CheckAuthorisation(ctx context.Context, actual *Permissions, required *Permissions) error {
	required = defaultIfBlank(required)
	actual = defaultIfBlank(actual)
	missingPermissions := make([]permissionType, 0)

	if required.Create && !actual.Create {
		missingPermissions = append(missingPermissions, Create)
	}
	if required.Read && !actual.Read {
		missingPermissions = append(missingPermissions, Read)
	}
	if required.Update && !actual.Update {
		missingPermissions = append(missingPermissions, Update)
	}
	if required.Delete && !actual.Delete {
		missingPermissions = append(missingPermissions, Delete)
	}

	if len(missingPermissions) > 0 {
		log.Event(ctx, "action forbidden caller does not process the required permissions", log.Data{
			"required_permissions": required,
			"caller_permissions":   actual,
			"missing_permissions":  missingPermissions,
		})
		return callerForbiddenError
	}

	log.Event(ctx, "caller authorised to perform the requested action")
	return nil
}

func defaultIfBlank(p *Permissions) *Permissions {
	if p == nil {
		return &Permissions{}
	}
	return p
}
