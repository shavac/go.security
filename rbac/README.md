##go.security rbac
==================

Go Security rbac provides Role Based Access Control.

API list:
=========

type User interface{}

type Role interface{}

type Perm interface{}

func HasRole(username string, role string) bool

func HasPerm(perm string) bool

func AddRole(username string, role Role) error

func DelRole(username string, role Role) error

func CreateRole(role string, ...perm Perm) error

func DropRole(role string) error

func GetRoles(username string) ([]Role, error)

func GetPerms(username string) ([]Perm, error)

func GetPermsByRole(rolename string) ([]Perm, error)

func GetUserByName(username string) *user
