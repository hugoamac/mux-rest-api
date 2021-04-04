package infra

import "github.com/spf13/viper"

// LoadVars - This method provides the loading of the environment variables.
func LoadVars() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
