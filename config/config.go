package config

import "github.com/spf13/viper"

type DbConfig struct {
	Host     string `mapstructure:"APP_DB_HOST"`
	Port     string `mapstructure:"APP_DB_PORT"`
	Name     string `mapstructure:"APP_DB_NAME"`
	User     string `mapstructure:"APP_DB_USER"`
	Password string `mapstructure:"APP_DB_PASSWORD"`
	Driver   string `mapstructure:"APP_DB_DRIVER"`
}

type ApiConfig struct {
	ApiHost string `mapstructure:"APP_API_HOST"`
	ApiPort string `mapstructure:"APP_API_PORT"`
}

type FileConfig struct {
	FilePath string `mapstructure:"APP_FILE_PATH"`
}

type TokenConfig struct {
	ApplicationName string `mapstructure:"APP_TOKEN_NAME"`
	JwtSignatureKey []byte `mapstructure:"APP_TOKEN_KEY"`
	//JwtSigningMethod *jwt.SigningMethodHMAC
	ExpirationToken int `mapstructure:"APP_EXPIRATION_TOKEN"`
}

type Config struct {
	DbConfig
	ApiConfig
	FileConfig
	TokenConfig
}

// Constructor
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.LoadConfig(".")
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c)
	return err
}
