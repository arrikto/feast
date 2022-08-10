package model

type InfraObject struct {
	Id          string `gorm:"column:id; not null; primary_key"`
	ProjectId   string `gorm:"column:project_id; not null"`
	ClassType   string `gorm:"column:class_type; size:65535"`
	Object      []byte `gorm:"column:object; size:65535"`
	ProjectName string `gorm:"-"`
}
