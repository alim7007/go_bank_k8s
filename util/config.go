package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment                string        `mapstructure:"ENVIRONMENT"`
	DBDriver                   string        `mapstructure:"DB_DRIVER"`
	DBSource                   string        `mapstructure:"DB_SOURCE"`
	MigrationUrl               string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress          string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GATEAWAY_HTTPServerAddress string        `mapstructure:"GATEAWAY_HTTP_SERVER_ADDRESS"`
	GRPCServerAddress          string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	Redis_Address              string        `mapstructure:"REDIS_ADDRESS"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration       time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName            string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress         string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword        string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

// LoadConfig reads configuration from environment file or variables
func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	// viper.AutomaticEnv()

	// for directly from .env but not working with test
	// viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	accessTokenDurationStr := viper.GetString("ACCESS_TOKEN_DURATION")
	accessTokenDuration, err := time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		return
	}
	refreshTokenDurationStr := viper.GetString("REFRESH_TOKEN_DURATION")
	refreshTokenDuration, err := time.ParseDuration(refreshTokenDurationStr)
	if err != nil {
		return
	}
	config = Config{
		Environment:                viper.GetString("ENVIRONMENT"),
		DBDriver:                   viper.GetString("DB_DRIVER"),
		DBSource:                   viper.GetString("DB_SOURCE"),
		MigrationUrl:               viper.GetString("MIGRATION_URL"),
		Redis_Address:              viper.GetString("REDIS_ADDRESS"),
		HTTPServerAddress:          viper.GetString("HTTP_SERVER_ADDRESS"),
		GATEAWAY_HTTPServerAddress: viper.GetString("GATEAWAY_HTTP_SERVER_ADDRESS"),
		GRPCServerAddress:          viper.GetString("GRPC_SERVER_ADDRESS"),
		TokenSymmetricKey:          viper.GetString("TOKEN_SYMMETRIC_KEY"),
		AccessTokenDuration:        accessTokenDuration,
		RefreshTokenDuration:       refreshTokenDuration,
		EmailSenderName:            viper.GetString("EMAIL_SENDER_NAME"),
		EmailSenderAddress:         viper.GetString("EMAIL_SENDER_ADDRESS"),
		EmailSenderPassword:        viper.GetString("EMAIL_SENDER_PASSWORD"),
	}
	return
}
