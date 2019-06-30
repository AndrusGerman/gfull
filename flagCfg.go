package gfull

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// FlagCFG : Config var
type FlagCFG struct {
	DB           *gorm.DB
	ConfigDB     ConfigDB
	ConfigServer ConfigServer
	Production   bool
}

// ConnectDB to database
func (ctx *FlagCFG) ConnectDB() (*FlagCFG, error) {
	if ctx.DB == nil {
		db, err := gorm.Open(ctx.ConfigDB.DBDriver, ctx.getConectionData())
		if err != nil {
			fmt.Println("Error db :" + err.Error())
			return ctx, err
		}
		ctx.DB = db
	}
	return ctx, nil
}

// Migrate one model
func (ctx *FlagCFG) getConectionData() (dt string) {
	// Get Mode
	var dbcf DatabaseConnection
	if ctx.Production {
		dbcf = ctx.ConfigDB.Prod
	} else {
		dbcf = ctx.ConfigDB.Dev
	}
	// Set Data for return
	switch ctx.ConfigDB.DBDriver {
	case "postgres":
		dt = fmt.Sprintf("host=%s port=%v user=%v dbname=%v password=%v %v", dbcf.DBHost, 5432, dbcf.DBUser, dbcf.DBName, dbcf.DBPassword, dbcf.DBParam)
	case "sqlite3":
		dt = dbcf.DBName
	case "mysql":
		dt = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local%s", dbcf.DBUser, dbcf.DBPassword, dbcf.DBName, dbcf.DBParam)
	default:
		fmt.Println("Error driver not compatible")
		os.Exit(1)
	}
	return dt
}
