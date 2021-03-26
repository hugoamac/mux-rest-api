package infra

import "github.com/spf13/viper"

// LoadVars - This method provides the load environments vars by application
func LoadVars() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}