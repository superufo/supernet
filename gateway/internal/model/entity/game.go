package entity

type Games struct {
	ID        int64  `gorm:"column:id" xorm:"id" json:"id"`
	Name      string `gorm:"column:name" xorm:"name" json:"name"`
	Icon      string `gorm:"column:icon" xorm:"icon" json:"icon"`
	Code      string `gorm:"column:code" xorm:"code" json:"code"`
	IsDelete  int64  `gorm:"column:is_delete" xorm:"is_delete" json:"is_delete"`
	CreatedAt int64  `gorm:"column:created_at" xorm:"created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at" xorm:"updated_at" json:"updated_at"`
}

const TABLE_GAME = "games"
