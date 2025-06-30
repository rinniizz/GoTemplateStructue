package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig   `mapstructure:"server"`
	Database  DatabaseConfig `mapstructure:"database"`
	Redis     RedisConfig    `mapstructure:"redis"`
	JWT       JWTConfig      `mapstructure:"jwt"`
	LogLevel  string         `mapstructure:"log_level"`
	LogFormat string         `mapstructure:"log_format"`
}

type ServerConfig struct {
	Host    string `mapstructure:"host"`
	Port    string `mapstructure:"port"`
	GinMode string `mapstructure:"gin_mode"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret     string        `mapstructure:"secret"`
	Expiration time.Duration `mapstructure:"expiration"`
}

func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set defaults
	setDefaults()

	// Bind environment variables
	bindEnvVars()

	// Read config file (optional)
	if err := viper.ReadInConfig(); err != nil {
		// Config file not found, use environment variables and defaults
		fmt.Println("Config file not found, using environment variables and defaults")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Parse JWT expiration
	if expStr := viper.GetString("JWT_EXPIRATION"); expStr != "" {
		if exp, err := time.ParseDuration(expStr); err == nil {
			config.JWT.Expiration = exp
		}
	}

	return &config, nil
}

func setDefaults() {
	// Server defaults
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.gin_mode", "debug")

	// Database defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.name", "gotemplate")
	viper.SetDefault("database.sslmode", "disable")

	// Redis defaults
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)

	// JWT defaults
	viper.SetDefault("jwt.secret", "your-super-secret-jwt-key")
	viper.SetDefault("jwt.expiration", 24*time.Hour)

	// Logging defaults
	viper.SetDefault("log_level", "info")
	viper.SetDefault("log_format", "json")
}

func bindEnvVars() {
	// Server
	viper.BindEnv("server.host", "SERVER_HOST")
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("server.gin_mode", "GIN_MODE")

	// Database
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.sslmode", "DB_SSLMODE")

	// Redis
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.port", "REDIS_PORT")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")

	// JWT
	viper.BindEnv("jwt.secret", "JWT_SECRET")

	// Logging
	viper.BindEnv("log_level", "LOG_LEVEL")
	viper.BindEnv("log_format", "LOG_FORMAT")
}
