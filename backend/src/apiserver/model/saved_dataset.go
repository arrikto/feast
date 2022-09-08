// Copyright 2022 Arrikto Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "time"

type SavedDataset struct {
	Id                   string    `gorm:"column:id; not null; primary_key"`
	ProjectId            string    `gorm:"column:project_id; not null; unique_index:unique_dataset"`
	Name                 string    `gorm:"column:name; not null; unique_index:unique_dataset; size:255"`
	Features             []byte    `gorm:"column:features; size:65535"`
	JoinKeys             []byte    `gorm:"column:join_keys; size:65535"`
	FullFeatureNames     bool      `gorm:"column:full_feature_names"`
	Storage              []byte    `gorm:"column:storage; size:65535"`
	FeatureServiceName   string    `gorm:"column:feature_service_name; size:255"`
	Tags                 []byte    `gorm:"column:tags; size:65535"`
	CreatedTimestamp     time.Time `gorm:"column:created_timestamp"`
	LastUpdatedTimestamp time.Time `gorm:"column:last_updated_timestamp"`
	MinEventTimestamp    time.Time `gorm:"column:min_event_timestamp"`
	MaxEventTimestamp    time.Time `gorm:"column:max_event_timestamp"`
	ProjectName          string    `gorm:"-"`
}
