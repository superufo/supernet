package bj_gm

type RoleHasPermissions struct {
	PermissionId uint64 `xorm:"not null pk UNSIGNED BIGINT"`
	RoleId       uint64 `xorm:"not null pk index UNSIGNED BIGINT"`
}

func (m *RoleHasPermissions) TableName() string {
	return "role_has_permissions"
}
