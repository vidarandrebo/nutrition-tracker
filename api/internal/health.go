package internal

import (
	"database/sql"
	"log/slog"
	"net/http"
	"reflect"
)

type HealthEndpoint struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewHealthEndpoint(db *sql.DB, logger *slog.Logger) *HealthEndpoint {
	he := HealthEndpoint{db: db}
	he.logger = logger.With(slog.Any("module", reflect.TypeOf(he)))
	return &he
}

func (he *HealthEndpoint) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if he.db == nil {
		he.logger.Error("database connection is nil")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var result int
	err := he.db.QueryRowContext(request.Context(), "SELECT 1").Scan(&result)
	if err != nil {
		he.logger.Error("health check failed", slog.Any("err", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
