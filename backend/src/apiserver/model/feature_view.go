package model

import "time"

type FeatureView struct {
	Id                   string        `gorm:"column:id; not null; primary_key"`
	ProjectId            string        `gorm:"column:project_id; not null; unique_index:unique_feature_view"`
	Name                 string        `gorm:"column:name; not null; unique_index:unique_feature_view; size:255"`
	Entities             []byte        `gorm:"column:entities; size:65535"`
	Description          string        `gorm:"column:description; size:65535"`
	Tags                 []byte        `gorm:"column:tags; size:65535"`
	Owner                string        `gorm:"column:owner; size:255"`
	Ttl                  time.Duration `gorm:"column:ttl"`
	BatchSource          []byte        `gorm:"column:batch_source; size:65535"`
	StreamSource         []byte        `gorm:"column:stream_source; size:65535"`
	Online               bool          `gorm:"column:online"`
	CreatedTimestamp     time.Time     `gorm:"column:created_timestamp"`
	LastUpdatedTimestamp time.Time     `gorm:"column:last_updated_timestamp"`
	Features             []*Feature
	MIs                  []*MaterializationInterval
	ProjectName          string `gorm:"-"`
}

type Feature struct {
	Id        string `gorm:"column:id; not null; primary_key"`
	FVId      string `gorm:"column:fvid; not null; primary_key; unique_index:unique_feature"`
	Name      string `gorm:"column:name; not null; unique_index:unique_feature; size:255"`
	ValueType int64  `gorm:"column:value_type"`
	Tags      []byte `gorm:"column:tags; size:65535"`
}

type MaterializationInterval struct {
	FVId      string    `gorm:"column:fvid; not null; primary_key"`
	StartTime time.Time `gorm:"column:start_time; not null; primary_key"`
	EndTime   time.Time `gorm:"column:end_time; not null; primary_key"`
}
