package model

type RequestFeatureView struct {
	Id          string `gorm:"column:id; not null; primary_key"`
	ProjectId   string `gorm:"column:project_id; not null; unique_index:unique_request_fv"`
	Name        string `gorm:"column:name; not null; unique_index:unique_request_fv; size:255"`
	DataSource  []byte `gorm:"column:data_source; size:65535"`
	Description string `gorm:"column:description; size:65535"`
	Tags        []byte `gorm:"column:tags; size:65535"`
	Owner       string `gorm:"column:owner; size:255"`
	ProjectName string `gorm:"-"`
}
