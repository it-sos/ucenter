package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"ucenter/datamodels"
	"ucenter/s/tests"
)

var db = tests.ConnectDb().Conn

func main() {
	if viper.Get("use-driver") == "sqlite3" {
		if _, e := os.Stat("sqlite3.db"); !os.IsExist(e) {
			log.Print("remove sqlite3.db")
			os.Remove("sqlite3.db")
		}
	}
	db.DropTables(
		new(datamodels.Role),
		new(datamodels.RolePermission),
		new(datamodels.AppRole),
		new(datamodels.Route),
		new(datamodels.UserRole),
		new(datamodels.UserPermission),
		new(datamodels.App),
		new(datamodels.User))
	db.Sync2(
		new(datamodels.Role),
		new(datamodels.RolePermission),
		new(datamodels.AppRole),
		new(datamodels.Route),
		new(datamodels.UserRole),
		new(datamodels.UserPermission),
		new(datamodels.App),
		new(datamodels.User))
}
