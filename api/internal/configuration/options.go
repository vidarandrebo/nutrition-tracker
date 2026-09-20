package configuration

import (
	"strings"

	"github.com/spf13/viper"
)

type JwtOptions struct {
	Secret         string `mapstructure:"Secret"`
	Issuer         string `mapstructure:"Issuer"`
	Audience       string `mapstructure:"Audience"`
	ExpirationTime int64  `mapstructure:"ExpirationTime"`
}

type DbOptions struct {
	ConnectionString string `mapstructure:"ConnectionString"`
}

type Options struct {
	Jwt                  JwtOptions                       `mapstructure:"Jwt"`
	ListenAddress        string                           `mapstructure:"ListenAddress"`
	StaticFilesDirectory string                           `mapstructure:"StaticFilesDirectory"`
	LogFile              string                           `mapstructure:"LogFile"`
	Database             DbOptions                        `mapstructure:"Database"`
	SystemUsers          map[string]SystemUserCredentials `mapstructure:"SystemUsers"`
	DataImporterTarget   string                           `mapstructure:"DataImporterTarget"`
}

func ParseOptions() (*Options, error) {
	viper.SetConfigName("appsettings")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	options := Options{}
	viper.SetConfigType("json")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetEnvPrefix("NT_")
	viper.AutomaticEnv()

	err = viper.Unmarshal(&options)
	if err != nil {
		return nil, err
	}

	return &options, nil
}

type SystemUserCredentials struct {
	Email    string `mapstructure:"Email"`
	Password string `mapstructure:"Password"`
}
