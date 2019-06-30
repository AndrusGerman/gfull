package gfull

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

// SetConfigFlag set configuration
func SetConfigFlag(DB *FlagCFG) {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		println("GFULL: Error to read config file")
		rangeClose()
	}

	if err := viper.UnmarshalKey("ConfigDB", &DB.ConfigDB); err != nil {
		println("GFULL: Error to parse config, DB")
		rangeClose()
	}

	if err := viper.UnmarshalKey("ConfigServer", &DB.ConfigServer); err != nil {
		println("GFULL: Error to parse config, Server")
		rangeClose()
	}

	flag.BoolVar(&DB.Production, "prod", false, "Enable production mode")
	// On Close App
	AddOnClose(func() {
		if DB.DB != nil {
			DB.DB.Close()
		}
	})
}

// ArrError add
type ArrError []error

// Add element array errors
func (ct *ArrError) Add(err error) {
	if err != nil {
		*ct = append(*ct, err)
	}
}

// Contain : Contain errors
func (ct *ArrError) Contain() bool {
	return len(*ct) != 0
}

// Error : Contain errors
func (ct ArrError) Error() error {
	// not error
	if len(ct) == 0 {
		return nil
	}
	// Add error in array
	var err []string
	for _, val := range ct {
		err = append(err, val.Error())
	}
	// Parse error
	var b, errM = json.Marshal(&err)
	if errM != nil {
		return errM
	}
	// return error
	return fmt.Errorf("%s", b)
}
