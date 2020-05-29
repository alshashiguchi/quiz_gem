package configurations

import "os"

//PortServerConfig - Port server up
type PortServerConfig struct {
	Port string
}

//DataBaseConfig - Config DataBase
type DataBaseConfig struct {
	Drive          string
	URL            string
	PathMigrations string
}

//Config - configurations
type Configurations struct {
	PortServer PortServerConfig
	DataBase   DataBaseConfig
}

func New() *Configurations {
	return &Configurations{
		PortServer: PortServerConfig{
			Port: getEnv("PORT_SERVER", "8080"),
		},
		DataBase: DataBaseConfig{
			Drive:          getEnv("DRIVE_DATABASE", ""),
			URL:            getEnv("URL_DATABASE", ""),
			PathMigrations: getEnv("PATH_MIGRATION_DATABASE", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
