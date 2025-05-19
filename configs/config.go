package configs

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type ServerConfig struct {
	Port    string `mapstructure:"port"`
	RunMode string `mapstructure:"run_mode"`
}

type LoggerConfig struct {
	Level    string `mapstructure:"level"`
	FilePath string `mapstructure:"file_path"`
	Encoding string `mapstructure:"encoding"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	PoolTimeout  time.Duration    `mapstructure:"pool_timeout"`
	DialTimeout  time.Duration    `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration    `mapstructure:"read_timeout"`
	WriteTimeout time.Duration    `mapstructure:"write_timeout"`
	IdleCheckFrequency int    `mapstructure:"idle_check_frequency"`
}



func GetConfig() (*Config, error) {
	v, err := LoadConfig(getConfigPath(os.Getenv("ENV")), "yml")
	if err != nil {
		return nil, err
	}
	return ParseConfig(v)
}


func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v,nil
}

func getConfigPath(env string) string {
	switch env {
	case "development":
		return "../configs/config-development.yml"
	case "docker":
		return "./configs/config-docker.yml"
	case "production":
		return "./configs/config-production.yml"
	default:
		return "../configs/config-development.yml"
	}
}
