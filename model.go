package gfull

import (
	"github.com/jinzhu/gorm"
)

// Model base
type Model struct {
	gorm.Model
	ModelsCFG *ModelCFG `gorm:"-"`
	snakeName string
}

// Sf get schema
func (ctx *Model) Sf(v interface{}) string {
	if ctx.snakeName != "" {
		return ctx.snakeName
	}
	ctx.snakeName = ModelToSnake(ctx.ModelsCFG.Database, v)
	return ctx.snakeName
}

// SetID get schema
func (ctx *Model) SetID(id string) (err error) {
	ctx.ID, err = StrToUint(id)
	return err
}

// NewModelCFG new model cfg
func NewModelCFG(schema string, dtbase FlagCFG) *ModelCFG {
	dt := new(ModelCFG).Connect(&dtbase)
	dt.Schema = schema
	return dt
}
