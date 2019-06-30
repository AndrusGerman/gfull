package gfull

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/inflection"
)

// ModelCFG base
type ModelCFG struct {
	Schema   string
	Database *FlagCFG
}

// Load : Get schema by model
func (ctx *ModelCFG) Load(val interface{}) *gorm.DB {
	md := reflect.ValueOf(val).Elem()
	if mdp := md.FieldByName("ModelsCFG"); mdp.IsValid() {
		mdp.Set(reflect.ValueOf(ctx))
	}
	if mdp := md.FieldByName("DeletedAt"); mdp.IsValid() {
		mdp.Set(reflect.Zero(reflect.TypeOf(new(time.Time))))
	}
	return ctx.getSchemaModel(val)
}

// Create element by model
func (ctx *ModelCFG) Create(val interface{}) *gorm.DB {
	return ctx.Load(val).Create(val)
}

// Find element by model
func (ctx *ModelCFG) Find(out interface{}, where ...interface{}) *gorm.DB {
	fmt.Println(reflect.ValueOf(out).Elem().NumField())
	os.Exit(0)
	return ctx.Load(out).Find(out, where...)
}

// Update one User
func (ctx *ModelCFG) Update(dt ...interface{}) error {
	vl := reflect.ValueOf(dt[0]).Elem()
	if vl.IsValid() {
		id := vl.FieldByName("ID").Elem()
		if id.IsValid() {
			if vl.Int() == 0 {
				return errors.New("ID is 0")
			}
		} else {
			return errors.New("ID is not valid")
		}
	} else {
		return errors.New("Element is not valid")
	}
	return ctx.Load(dt[0]).Update(dt...).Error
}

// Pre : Pase model for preload
func (ctx *ModelCFG) Pre(md interface{}) (string, func(db *gorm.DB) *gorm.DB) {
	a := func(db *gorm.DB) *gorm.DB {
		return ctx.getSchemaName(ctx.Sf(md))
	}
	return SnakeCamelC(ctx.Sf(md)), a
}

// Sf : Return sankeName for model
func (ctx *ModelCFG) Sf(md interface{}) string {
	sk := reflect.ValueOf("snakeName").Elem()
	if !sk.IsValid() {
		fmt.Println("SK no es valido")
		return ""
	}
	if sk.String() != "" {
		return sk.String()
	}
	sk.SetString(ModelToSnake(ctx.Database, md))
	return sk.String()
}

// Delete : add unique index constrains
func (ctx *ModelCFG) Delete(model interface{}, where ...interface{}) error {
	Rid := reflect.ValueOf(model).Elem()
	if len(where) > 0 {
		return ctx.Load(model).Delete(model, where...).Error
	}
	if Rid.Int() != 0 {
		return ctx.Load(model).Delete(model, Rid.Int()).Error
	}
	return errors.New("ID not valido")
}

// Connect to database
func (ctx *ModelCFG) Connect(dt *FlagCFG) *ModelCFG {
	var err error
	ctx.Database, err = dt.ConnectDB()
	if err != nil {
		os.Exit(1)
	}
	return ctx
}

// ModelToSnake convert model in snakeCase string
func ModelToSnake(d *FlagCFG, val interface{}) string {
	return d.DB.NewScope(val).GetModelStruct().TableName(d.DB)
}

// getSchemaModel get schema one model
func (ctx *ModelCFG) getSchemaModel(v interface{}) *gorm.DB {
	// Retorna los datos
	return ctx.getSchemaName(ModelToSnake(ctx.Database, v))
}

// AddForignKey : add Forign Key to element
// Example:  DC.AddForignKey( new(User) , "email_id" , new(md.Email),"id",md.DC.Schema )
func (ctx *ModelCFG) AddForignKey(FirstModel interface{}, FirstElement string, SencondModel interface{}, SecondElement string, SecondElementSchema string) {
	tabla := ModelToSnake(ctx.Database, SecondElementSchema)
	ctx.getSchemaModel(FirstModel).AddForeignKey(FirstElement, SecondElementSchema+"."+tabla+"("+SecondElement+")", "RESTRICT", "RESTRICT")
}

// AddUniqueIndex : add unique index constrains
func (ctx *ModelCFG) AddUniqueIndex(model interface{}, nameConstrains string, index ...string) {
	ctx.getSchemaModel(model).AddUniqueIndex(nameConstrains, index...)
	//ctx.getSchemaName(ModelToSnake(ctx.Database, model)).AddUniqueIndex(nameConstrains, index...)
}

// GetSchemaName for table
func (ctx *ModelCFG) getSchemaName(name string) *gorm.DB {
	if ctx.Schema == "" {
		return ctx.Database.DB.Table(name)
	}
	return ctx.Database.DB.Table(ctx.Schema + "." + name)
}

// SnakeCamelC Convert SnakeCase in CamelCase
func SnakeCamelC(valor string) (total string) {
	arr := strings.SplitAfter(valor, "_")
	for _, v := range arr {
		total += strings.Title(v)
	}
	return strings.Replace(inflection.Singular(total), "_", "", 5)
}

// Migrate one model
func (ctx *ModelCFG) Migrate(val ...interface{}) error {
	if len(val) == 0 {
		return errors.New("Not Migrate Element")
	}
	return ctx.getSchemaModel(val[0]).AutoMigrate(val...).Error
}
