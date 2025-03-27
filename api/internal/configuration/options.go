package configuration

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"os"
)

type Options struct {
	JwtSecret            string                           `json:"JwtSecret"`
	JwtIssuer            string                           `json:"JwtIssuer"`
	JwtAudience          string                           `json:"JwtAudience"`
	ListenAddress        string                           `json:"ListenAddress"`
	StaticFilesDirectory string                           `json:"StaticFilesDirectory"`
	LogFile              string                           `json:"LogFile"`
	DBConnectionString   string                           `json:"DBConnectionString"`
	SystemUsers          map[string]SystemUserCredentials `json:"SystemUsers"`
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

	return options, nil
}

type SystemUserCredentials struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
