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

type FeatureService struct {
	Id                     string    `gorm:"column:id; not null; primary_key"`
	ProjectId              string    `gorm:"column:project_id; not null; unique_index:unique_feature_service"`
	Name                   string    `gorm:"column:name; not null; unique_index:unique_feature_service; size:255"`
	Tags                   []byte    `gorm:"column:tags; size:65535"`
	Description            string    `gorm:"column:description; size:65535"`
	Owner                  string    `gorm:"column:owner; size:255"`
	LoggingConfig          []byte    `gorm:"logging_config; size:65535"`
	CreatedTimestamp       time.Time `gorm:"column:created_timestamp"`
	LastUpdatedTimestamp   time.Time `gorm:"column:last_updated_timestamp"`
	FeatureViewProjections []*FeatureViewProjection
	ProjectName            string `gorm:"-"`
}

type FeatureViewProjection struct {
	Id          string `gorm:"column:id; not null; primary_key"`
	FSId        string `gorm:"column:fsid; not null; primary_key"`
	FVName      string `gorm:"column:feature_view_name; size:255"`
	FVNameAlias string `gorm:"column:feature_view_name_alias; size:255"`
	JoinKeyMap  []byte `gorm:"column:join_key_map; size:65535"`
	FVPFeatures []*FvpFeature
}

type FvpFeature struct {
	Id        string `gorm:"column:id; not null; primary_key"`
	FVPId     string `gorm:"column:fvpid; not null; primary_key; unique_index:unique_featurevp"`
	Name      string `gorm:"column:name; not null; unique_index:unique_featurevp; size:255"`
	ValueType int64  `gorm:"column:value_type"`
	Tags      []byte `gorm:"column:tags; size:65535"`
}