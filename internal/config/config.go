package config

import "github.com/spf13/viper"

type Config struct {
	InputFile string `mapstructure:"INPUT_FILE"`

	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     int    `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"POSTGRES_DB"`

	RabbitHost     string `mapstructure:"RABBIT_HOST"`
	RabbitPort     int    `mapstructure:"RABBIT_PORT"`
	RabbitUser     string `mapstructure:"RABBIT_USER"`
	RabbitPassword string `mapstructure:"RABBIT_PASSWORD"`
}

func LoadConfig(path string) Config {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	return config
}
