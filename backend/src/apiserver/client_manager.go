package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/feast-dev/feast/backend/src/apiserver/storage"

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
)

// Container for all service clients
type ClientManager struct {
	db                       *storage.DB
	dataSourceStore          storage.DataSourceStoreInterface
	entityStore              storage.EntityStoreInterface
	featureServiceStore      storage.FeatureServiceStoreInterface
	featureViewStore         storage.FeatureViewStoreInterface
	infraObjectStore         storage.InfraObjectStoreInterface
	miStore                  storage.MIStoreInterface
	onDemandFeatureViewStore storage.OnDemandFeatureViewStoreInterface
	projectStore             storage.ProjectStoreInterface
	requestFeatureViewStore  storage.RequestFeatureViewStoreInterface
	savedDatasetStore        storage.SavedDatasetStoreInterface
	time                     util.TimeInterface
	uuid                     util.UUIDGeneratorInterface
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

func (c *ClientManager) Time() util.TimeInterface {
	return c.time
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
