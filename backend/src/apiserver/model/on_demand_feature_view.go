package model

import "time"

type OnDemandFeatureView struct {
	Id                   string    `gorm:"column:id; not null; primary_key"`
	ProjectId            string    `gorm:"column:project_id; not null; unique_index:unique_odfv"`
	Name                 string    `gorm:"column:name; not null; unique_index:unique_odfv; size:255"`
	Sources              []byte    `gorm:"column:sources; size:65535"`
	UdfName              string    `gorm:"column:udf_name; size:255"`
	UdfBody              []byte    `gorm:"column:udf_body; size:65535"`
	Description          string    `gorm:"column:description; size:65535"`
	Tags                 []byte    `gorm:"column:tags; size:65535"`
	Owner                string    `gorm:"column:owner; size:255"`
	CreatedTimestamp     time.Time `gorm:"column:created_timestamp"`
	LastUpdatedTimestamp time.Time `gorm:"column:last_updated_timestamp"`
	Features             []*OnDemandFeature
	ProjectName          string `gorm:"-"`
}

type OnDemandFeature struct {
	Id        string `gorm:"column:id; not null; primary_key"`
	ODFVId    string `gorm:"column:odfvid; not null; primary_key; unique_index:unique_odfeature"`
	Name      string `gorm:"column:name; not null; unique_index:unique_odfeature; size:255"`
	ValueType int64  `gorm:"column:value_type"`
	Tags      []byte `gorm:"column:tags; size:65535"`
}
