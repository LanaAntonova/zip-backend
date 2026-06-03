package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	DB     DBConfig     `mapstructure:"db"`
	Redis  RedisConfig  `mapstructure:"redis"`
	NATS   NATSConfig   `mapstructure:"nats"`
	MinIO  MinIOConfig  `mapstructure:"minio"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	Google GoogleConfig `mapstructure:"google"`
}

type AppConfig struct {
	Env  string `mapstructure:"env"`
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	URL string `mapstructure:"url"`
}

type RedisConfig struct {
	URL string `mapstructure:"url"`
}

type NATSConfig struct {
	URL string `mapstructure:"url"`
}

type MinIOConfig struct {
	Endpoint  string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	UseSSL    bool   `mapstructure:"use_ssl"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	AccessTTL  string `mapstructure:"access_ttl"`
	RefreshTTL string `mapstructure:"refresh_ttl"`
}

type GoogleConfig struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectURL  string `mapstructure:"redirect_url"`
}

func Load(path string) (*Config, error) {
	v := viper.New()

	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	// env vars override yaml: APP_PORT → app.port
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}
