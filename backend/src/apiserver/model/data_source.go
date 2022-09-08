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

type DataSource struct {
	Id                  string `gorm:"column:id; not null; primary_key"`
	ProjectId           string `gorm:"column:project_id; not null; unique_index:unique_data_source"`
	Name                string `gorm:"column:name; not null; unique_index:unique_data_source; size:255"`
	Description         string `gorm:"column:description; size:65535"`
	Tags                []byte `gorm:"column:tags; size:65535"`
	Owner               string `gorm:"column:owner; size:255"`
	Type                int64  `gorm:"column:type"`
	FieldMapping        []byte `gorm:"column:field_mapping; size:65535"`
	TimestampField      string `gorm:"column:timestamp_field; size:255"`
	DatePartitionCol    string `gorm:"column:date_partition_column; size:255"`
	CreatedTimestampCol string `gorm:"column:created_timestamp_column; size:255"`
	ClassType           string `gorm:"column:class_type; size:65535"`
	BatchSource         []byte `gorm:"column:batch_source; size:65535"`
	Options             []byte `gorm:"column:options; size:65535"`
	ProjectName         string `gorm:"-"`
}
