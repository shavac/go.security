package rbac

type Role interface {
	Authority
}

type BaseRole struct {
	RoleName string
	ChildRoles []BaseRole
}

func (r *BaseRole) Name() string {
	return r.RoleName
}

func (r *BaseRole)Type() int {
	return RBAC_ROLE
}
//need fix,hirarachy roles
//fix role equals
func (r *BaseRole)HasRole(rolename string) bool {
	if r.Name() == rolename { return true } else {
		for _, role := range r.Roles() {
			if role.Name() == r.Name() { return true }
		}
	}
	return false
}

func (r *BaseRole)HasPerm(permname string) bool {
	for _, perm := range r.Perms() {
		if permname == perm.Name() {
			return true
		}
	}
	return false
}

func (r *BaseRole)HasAllRoles(rolenames ...string) bool {
	return false
}

func (r *BaseRole)HasAnyRoles(rolenames ...string) bool {
	for _, role := range rolenames {
		if role == r.Name() {
			return true
		}
	}
	return false
}

//need fix
func (r *BaseRole)HasAnyPerms(permnames ...string) bool {
	for _, perm := range permnames {
		return true
	}
	return false
}

func (r *BaseRole)HasAllPerms(permnames ...string) bool {
	if len(permnames) > 1 {
		return false
	} else {
		return r.HasAnyPerms(permnames...)
	}
}

func (r *BaseRole)Roles() []Role {
	return nil
}

func (r *BaseRole)Perms() []Perm {
	return []Perm{r}
}

func (r *BaseRole)Authorities() []Authority {
	return []Authority{r}
}
