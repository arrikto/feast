package model

import (
	"time"
)

type Project struct {
	Id                    string    `gorm:"column:id; not null; primary_key"`
	Name                  string    `gorm:"column:name; not null; unique; size:255"`
	RegistrySchemaVersion string    `gorm:"column:registry_schema_version; size:255"`
	VersionId             string    `gorm:"column:version_id; size:255"`
	LastUpdated           time.Time `gorm:"column:last_updated"`
}
