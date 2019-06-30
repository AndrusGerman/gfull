package gfull

// ConfigDB database
type ConfigDB struct {
	// DBDriver : 'sqlite3','postgres'
	DBDriver string `mapstructure:"DBDriver"`
	Dev      DatabaseConnection
	Prod     DatabaseConnection
}

// DatabaseConnection config
type DatabaseConnection struct {
	DBName     string `mapstructure:"DBName"`
	DBPassword string `mapstructure:"DBPassword"`
	DBUser     string `mapstructure:"DBUser"`
	DBHost     string `mapstructure:"DBHost"`
	DBParam    string `mapstructure:"DBParam"`
}

// ConfigServer modes
type ConfigServer struct {
	Dev  ServerData
	Prod ServerData
}

// ServerData config structure
type ServerData struct {
	Addr string `mapstructure:"Addr"`
	SSL  bool   `mapstructure:"SSL"`
	Key  string `mapstructure:"Key"`
	Cert string `mapstructure:"Cert"`
}
