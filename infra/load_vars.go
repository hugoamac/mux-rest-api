package infra

import "github.com/spf13/viper"

type Config struct {
	MongoUri    string `mapstructure:"MONGO_URI"`
	MongoDbName string `mapstructure:"MONGO_DBNAME"`
	AppPort     string `mapstructure:"APP_PORT"`
}

// LoadVars - This method provides the loading of the environment variables.
func LoadVars() (config Config, err error) {

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
