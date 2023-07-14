package bj_agent

type ModelHasPermissions struct {
	PermissionId uint64 `xorm:"not null pk UNSIGNED BIGINT"`
	ModelType    string `xorm:"not null pk index(model_has_permissions_model_id_model_type_index) VARCHAR(255)"`
	ModelId      uint64 `xorm:"not null pk index(model_has_permissions_model_id_model_type_index) UNSIGNED BIGINT"`
}

func (m *ModelHasPermissions) TableName() string {
	return "model_has_permissions"
}
