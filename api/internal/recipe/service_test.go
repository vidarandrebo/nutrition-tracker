package recipe_test

import (
	"bytes"
	"database/sql"
	"errors"
	"log/slog"
	"strings"
	"testing"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/recipe"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type ownershipRepository struct {
	recipe.IRepository
	item recipe.TableRecipe
	err  error
}

func (r ownershipRepository) GetByID(int64) (recipe.TableRecipe, error) {
	return r.item, r.err
}

func TestServiceCheckOwnership(t *testing.T) {
	dbErr := errors.New("database unavailable")
	tests := []struct {
		name    string
		item    recipe.TableRecipe
		err     error
		wantErr error
	}{
		{name: "owned", item: recipe.TableRecipe{ID: 1, OwnerID: 2}},
		{name: "other owner", item: recipe.TableRecipe{ID: 1, OwnerID: 3}, wantErr: utils.ErrForbidden},
		{name: "missing", err: sql.ErrNoRows, wantErr: utils.ErrNotFound},
		{name: "database failure", err: dbErr, wantErr: utils.ErrUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var logs bytes.Buffer
			s := recipe.NewService(ownershipRepository{item: tt.item, err: tt.err}, slog.New(slog.NewTextHandler(&logs, nil)))
			if err := s.CheckOwnership(1, 2); !errors.Is(err, tt.wantErr) {
				t.Fatalf("CheckOwnership() = %v, want %v", err, tt.wantErr)
			}
			if tt.err == dbErr && (!strings.Contains(logs.String(), dbErr.Error()) || !strings.Contains(logs.String(), "recipeID=1")) {
				t.Fatalf("missing database cause and recipe ID in log: %s", logs.String())
			}
		})
	}
}
