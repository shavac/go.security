package rbac

type Authority interface {
	Name() string
	Type() int
	HasRole(rolename string) bool
	HasPerm(permname string) bool
	HasAllRoles(rolenames ...string) bool
	HasAllPerms(permnames ...string) bool
	HasAnyRoles(rolenames ...string) bool
	HasAnyPerms(permnames ...string) bool
	Roles() []Role
	Perms() []Perm
	Authorities() []Authority
}

