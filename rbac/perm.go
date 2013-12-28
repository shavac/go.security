package rbac

type Perm interface {
	Authority
}

type BasePerm struct {
	PermName string
}

func (p *BasePerm)Name() string {
	return p.PermName
}

func (p *BasePerm)Type() int {
	return RBAC_PERM
}

func (p *BasePerm)HasRole(rolename string) bool {
	return false
}

func (p *BasePerm)HasPerm(permname string) bool {
	return p.Name() == permname
}

func (p *BasePerm)HasAllRoles(rolenames ...string) bool {
	return false
}

func (p *BasePerm)HasAnyRoles(rolenames ...string) bool {
	return false
}

func (p *BasePerm)HasAnyPerms(permnames ...string) bool {
	for _, perm := range permnames {
		if p.Name() == perm {
			return true
		}
	}
	return false
}

func (p *BasePerm)HasAllPerms(permnames ...string) bool {
	if len(permnames) > 1 {
		return false
	} else {
		return p.HasAnyPerms(permnames...)
	}
}

func (p *BasePerm)Roles() []Role {
	return nil
}

func (p *BasePerm)Perms() []Perm {
	return []Perm{p}
}

func (p *BasePerm)Authorities() []Authority {
	return []Authority{p}
}
