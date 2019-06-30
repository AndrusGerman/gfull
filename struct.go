package gfull

// Config database
type Config struct {
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
