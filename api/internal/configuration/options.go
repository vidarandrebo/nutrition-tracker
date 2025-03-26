package configuration

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"os"
)

type Options struct {
	JwtIssuer            string
	JwtAudience          string
	ListenAddress        string
	StaticFilesDirectory string
	LogFile              string
	DBConnectionString   string
	SystemUsers          map[string]SystemUserCredentials
}

func ParseOptions(fileName string) (*Options, error) {
	file, err := os.Open("./appsettings.json")
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
	Email    string
	Password string
}
