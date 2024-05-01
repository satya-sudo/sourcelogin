package config

import "github.com/spf13/viper"

type Config struct {
	Cf *viper.Viper
}

func NewConfig(configFileName string, fileType string, path string) (*Config, error) {
	viperConfig := viper.New()
	viperConfig.SetConfigName(configFileName)
	viperConfig.SetConfigType(fileType)
	viperConfig.AddConfigPath(path)
	err := viperConfig.ReadInConfig()
	if err != nil {
		println("Error getting the configuration.", err.Error())
		return nil, err
	}

	config := &Config{
		Cf: viperConfig,
	}
	return config, nil
}
