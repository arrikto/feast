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

package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/feast-dev/feast/backend/src/apiserver/storage"

	"github.com/feast-dev/feast/backend/src/apiserver/auth"
	"github.com/feast-dev/feast/backend/src/apiserver/common"
	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	"github.com/feast-dev/feast/backend/src/apiserver/client"

	"github.com/cenkalti/backoff"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	mysqlServiceHost       = "DBConfig.Host"
	mysqlServicePort       = "DBConfig.Port"
	mysqlUser              = "DBConfig.User"
	mysqlPassword          = "DBConfig.Password"
	mysqlDBName            = "DBConfig.DBName"
	mysqlGroupConcatMaxLen = "DBConfig.GroupConcatMaxLen"
	mysqlExtraParams       = "DBConfig.ExtraParams"
	dbConMaxLifeTime       = "DBConfig.ConMaxLifeTime"

	initConnectionTimeout = "InitConnectionTimeout"

	clientQPS   = "ClientQPS"
	clientBurst = "ClientBurst"
)

// Container for all service clients
type ClientManager struct {
	db                        *storage.DB
	authenticators            []auth.Authenticator
	dataSourceStore           storage.DataSourceStoreInterface
	entityStore               storage.EntityStoreInterface
	featureServiceStore       storage.FeatureServiceStoreInterface
	featureViewStore          storage.FeatureViewStoreInterface
	infraObjectStore          storage.InfraObjectStoreInterface
	miStore                   storage.MIStoreInterface
	onDemandFeatureViewStore  storage.OnDemandFeatureViewStoreInterface
	projectStore              storage.ProjectStoreInterface
	requestFeatureViewStore   storage.RequestFeatureViewStoreInterface
	savedDatasetStore         storage.SavedDatasetStoreInterface
	subjectAccessReviewClient client.SubjectAccessReviewInterface
	time                      util.TimeInterface
	tokenReviewClient         client.TokenReviewInterface
	uuid                      util.UUIDGeneratorInterface
}

func (c *ClientManager) Authenticators() []auth.Authenticator {
	return c.authenticators
}

func (c *ClientManager) DataSourceStore() storage.DataSourceStoreInterface {
	return c.dataSourceStore
}

func (c *ClientManager) EntityStore() storage.EntityStoreInterface {
	return c.entityStore
}

func (c *ClientManager) FeatureServiceStore() storage.FeatureServiceStoreInterface {
	return c.featureServiceStore
}

func (c *ClientManager) FeatureViewStore() storage.FeatureViewStoreInterface {
	return c.featureViewStore
}

func (c *ClientManager) InfraObjectStore() storage.InfraObjectStoreInterface {
	return c.infraObjectStore
}

func (c *ClientManager) MiStore() storage.MIStoreInterface {
	return c.miStore
}

func (c *ClientManager) OnDemandFeatureViewStore() storage.OnDemandFeatureViewStoreInterface {
	return c.onDemandFeatureViewStore
}

func (c *ClientManager) ProjectStore() storage.ProjectStoreInterface {
	return c.projectStore
}

func (c *ClientManager) RequestFeatureViewStore() storage.RequestFeatureViewStoreInterface {
	return c.requestFeatureViewStore
}

func (c *ClientManager) SavedDatasetStore() storage.SavedDatasetStoreInterface {
	return c.savedDatasetStore
}

func (c *ClientManager) SubjectAccessReviewClient() client.SubjectAccessReviewInterface {
	return c.subjectAccessReviewClient
}

func (c *ClientManager) Time() util.TimeInterface {
	return c.time
}

func (c *ClientManager) TokenReviewClient() client.TokenReviewInterface {
	return c.tokenReviewClient
}

func (c *ClientManager) UUID() util.UUIDGeneratorInterface {
	return c.uuid
}

func (c *ClientManager) init() {
	glog.Info("Initializing client manager")
	db := initDBClient(common.GetDurationConfig(initConnectionTimeout))
	db.SetConnMaxLifetime(common.GetDurationConfig(dbConMaxLifeTime))

	// time
	c.time = util.NewRealTime()

	// UUID generator
	c.uuid = util.NewUUIDGenerator()

	c.db = db
	c.dataSourceStore = storage.NewDataSourceStore(db, c.time, c.uuid)
	c.entityStore = storage.NewEntityStore(db, c.time, c.uuid)
	c.featureServiceStore = storage.NewFeatureServiceStore(db, c.time, c.uuid)
	c.featureViewStore = storage.NewFeatureViewStore(db, c.time, c.uuid)
	c.infraObjectStore = storage.NewInfraObjectStore(db, c.time, c.uuid)
	c.miStore = storage.NewMIStore(db, c.time, c.uuid)
	c.onDemandFeatureViewStore = storage.NewOnDemandFeatureViewStore(db, c.time, c.uuid)
	c.projectStore = storage.NewProjectStore(db, c.time, c.uuid)
	c.requestFeatureViewStore = storage.NewRequestFeatureViewStore(db, c.time, c.uuid)
	c.savedDatasetStore = storage.NewSavedDatasetStore(db, c.time, c.uuid)

	// Use default value of client QPS (5) & burst (10) defined in
	// k8s.io/client-go/rest/config.go#RESTClientFor
	clientParams := util.ClientParameters{
		QPS:   common.GetFloat64ConfigWithDefault(clientQPS, 5),
		Burst: common.GetIntConfigWithDefault(clientBurst, 10),
	}

	if common.IsMultiUserMode() {
		c.subjectAccessReviewClient = client.CreateSubjectAccessReviewClientOrFatal(common.GetDurationConfig(initConnectionTimeout), clientParams)
		c.tokenReviewClient = client.CreateTokenReviewClientOrFatal(common.GetDurationConfig(initConnectionTimeout), clientParams)
		c.authenticators = auth.GetAuthenticators(c.tokenReviewClient)
	}

	glog.Infof("Client manager initialized successfully")
}

func (c *ClientManager) Close() {
	c.db.Close()
}

func initDBClient(initConnectionTimeout time.Duration) *storage.DB {
	driverName := common.GetStringConfig("DBConfig.DriverName")
	var arg string

	switch driverName {
	case "mysql":
		arg = initMysql(driverName, initConnectionTimeout)
	default:
		glog.Fatalf("Driver %v is not supported", driverName)
	}

	// db is safe for concurrent use by multiple goroutines
	// and maintains its own pool of idle connections.
	db, err := gorm.Open(driverName, arg)
	util.TerminateIfError(err)

	// Create table
	response := db.AutoMigrate(
		&model.DataSource{},
		&model.Entity{},
		&model.FeatureService{},
		&model.FeatureViewProjection{},
		&model.FvpFeature{},
		&model.FeatureView{},
		&model.Feature{},
		&model.InfraObject{},
		&model.MaterializationInterval{},
		&model.OnDemandFeatureView{},
		&model.OnDemandFeature{},
		&model.Project{},
		&model.RequestFeatureView{},
		&model.SavedDataset{},
	)

	if response.Error != nil {
		glog.Fatalf("Failed to initialize the databases.")
	}

	return storage.NewDB(db.DB(), storage.NewMySQLDialect())
}

// Initialize the connection string for connecting to Mysql database
// Format would be something like root@tcp(ip:port)/dbname?charset=utf8&loc=Local&parseTime=True
func initMysql(driverName string, initConnectionTimeout time.Duration) string {
	mysqlConfig := client.CreateMySQLConfig(
		common.GetStringConfigWithDefault(mysqlUser, "root"),
		common.GetStringConfigWithDefault(mysqlPassword, ""),
		common.GetStringConfigWithDefault(mysqlServiceHost, "mysql"),
		common.GetStringConfigWithDefault(mysqlServicePort, "3306"),
		"",
		common.GetStringConfigWithDefault(mysqlGroupConcatMaxLen, "1024"),
		common.GetMapConfig(mysqlExtraParams),
	)

	var db *sql.DB
	var err error
	var operation = func() error {
		db, err = sql.Open(driverName, mysqlConfig.FormatDSN())
		if err != nil {
			return err
		}
		return nil
	}
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = initConnectionTimeout

	backoff.RetryNotify(operation, b, func(e error, duration time.Duration) {
		glog.Errorf("%v", e)
	})

	defer db.Close()
	util.TerminateIfError(err)

	// Create database if not exist
	dbName := common.GetStringConfig(mysqlDBName)
	operation = func() error {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
		if err != nil {
			return err
		}
		return nil
	}
	b = backoff.NewExponentialBackOff()
	b.MaxElapsedTime = initConnectionTimeout
	err = backoff.Retry(operation, b)

	util.TerminateIfError(err)
	mysqlConfig.DBName = dbName
	// When updating, return rows matched instead of rows affected. This counts rows that are being
	// set as the same values as before. If updating using a primary key and rows matched is 0, then
	// it means this row is not found.
	// Config reference: https://github.com/go-sql-driver/mysql#clientfoundrows
	mysqlConfig.ClientFoundRows = true

	return mysqlConfig.FormatDSN()
}

// newClientManager creates and Init a new instance of ClientManager
func newClientManager() ClientManager {
	clientManager := ClientManager{}
	clientManager.init()

	return clientManager
}
