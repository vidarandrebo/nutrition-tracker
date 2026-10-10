package meal_test

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"testing"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/meal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type ownershipRepository struct {
	meal.IRepository
	t    *testing.T
	item meal.TableMeal
	err  error
}

func (r ownershipRepository) GetById(id int64) (meal.TableMeal, error) {
	r.t.Helper()
	if id != 1 {
		r.t.Fatalf("GetById(%d), want 1", id)
	}
	return r.item, r.err
}

func TestServiceCheckOwnership(t *testing.T) {
	queryErr := errors.New("database unavailable")
	tests := []struct {
		name    string
		item    meal.TableMeal
		getErr  error
		ownerID int64
		wantErr error
	}{
		{name: "owned", item: meal.TableMeal{ID: 1, OwnerID: 2}, ownerID: 2},
		{name: "not owned", item: meal.TableMeal{ID: 1, OwnerID: 2}, ownerID: 3, wantErr: utils.ErrForbidden},
		{name: "missing", getErr: sql.ErrNoRows, ownerID: 2, wantErr: utils.ErrNotFound},
		{name: "wrapped missing", getErr: fmt.Errorf("lookup: %w", sql.ErrNoRows), ownerID: 2, wantErr: utils.ErrNotFound},
		{name: "query error", getErr: queryErr, ownerID: 2, wantErr: utils.ErrUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := ownershipRepository{t: t, item: tt.item, err: tt.getErr}
			var logs bytes.Buffer
			service := meal.NewService(repo, slog.New(slog.NewTextHandler(&logs, nil)))
			err := service.CheckOwnership(1, tt.ownerID)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("CheckOwnership() error = %v, want %v", err, tt.wantErr)
			}
			if tt.name == "query error" {
				if !strings.Contains(logs.String(), queryErr.Error()) || !strings.Contains(logs.String(), "mealID=1") {
					t.Fatalf("missing lookup error and meal ID in log: %s", logs.String())
				}
			}
		})
	}
}
