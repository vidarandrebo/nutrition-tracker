package configuration

import (
	"os"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type JwtOptions struct {
	Secret         string `mapstructure:"Secret"`
	Issuer         string `mapstructure:"Issuer"`
	Audience       string `mapstructure:"Audience"`
	ExpirationTime int64  `mapstructure:"ExpirationTime"`
}

type Options struct {
	Jwt                  JwtOptions                       `mapstructure:"Jwt"`
	ListenAddress        string                           `mapstructure:"ListenAddress"`
	StaticFilesDirectory string                           `mapstructure:"StaticFilesDirectory"`
	LogFile              string                           `mapstructure:"LogFile"`
	DBConnectionString   string                           `mapstructure:"DBConnectionString"`
	SystemUsers          map[string]SystemUserCredentials `mapstructure:"SystemUsers"`
	DataImporterTarget   string                           `mapstructure:"DataImporterTarget"`
}

func ParseOptions(fileName string) (*Options, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	options, err := utils.ParseJson[Options](file)
	if err != nil {
		return nil, err
	}

	return &options, nil
}

type SystemUserCredentials struct {
	Email    string `mapstructure:"Email"`
	Password string `mapstructure:"Password"`
}
