package client

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func CreateMySQLConfig(user, password string, mysqlServiceHost string,
	mysqlServicePort string, dbName string, mysqlGroupConcatMaxLen string, mysqlExtraParams map[string]string) *mysql.Config {

	params := map[string]string{
		"charset":              "utf8",
		"parseTime":            "True",
		"loc":                  "Local",
		"group_concat_max_len": mysqlGroupConcatMaxLen,
	}

	for k, v := range mysqlExtraParams {
		params[k] = v
	}

	return &mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", mysqlServiceHost, mysqlServicePort),
		Params:               params,
		DBName:               dbName,
		AllowNativePasswords: true,
	}
}
