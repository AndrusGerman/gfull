package gfull

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// Database Database
type Database struct {
	DB         *gorm.DB
	Config     Config
	Production bool
}

// Connect to database
func (ctx *Database) Connect() (*Database, error) {
	if ctx.DB == nil {
		db, err := gorm.Open(ctx.Config.DBDriver, ctx.getConectionData())
		if err != nil {
			fmt.Println("Error db :" + err.Error())
			return ctx, err
		}
		ctx.DB = db
	}
	return ctx, nil
}

// Migrate one model
func (ctx *Database) getConectionData() (dt string) {
	// Get Mode
	var dbcf DatabaseConnection
	if ctx.Production {
		dbcf = ctx.Config.Prod
	} else {
		dbcf = ctx.Config.Dev
	}
	// Set Data for return
	switch ctx.Config.DBDriver {
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
