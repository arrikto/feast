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
