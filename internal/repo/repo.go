package repo

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"ova-recipe-api/internal/recipe"
)

type Error string

func (e Error) Error() string { return string(e) }

const (
	NotFoundError = Error("Recipe not found. ")
)

type RecipeRepo interface {
	AddRecipe(ctx context.Context, newRecipe recipe.Recipe) (uint64, error)
	AddRecipes(ctx context.Context, recipes []recipe.Recipe) error
	ListRecipes(ctx context.Context, limit, offset uint64) ([]recipe.Recipe, error)
	DescribeRecipe(ctx context.Context, recipeId uint64) (*recipe.Recipe, error)
	RemoveRecipe(ctx context.Context, recipeId uint64) error
	UpdateRecipe(ctx context.Context, newRecipe recipe.Recipe) error
}

func OpenDb(dsn string) (*sql.DB, error) {
	return sql.Open("pgx", dsn)
}

func New(db *sql.DB) (RecipeRepo, error) {
	newDb := sqlx.NewDb(db, "pgx")
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &repo{db: newDb}, nil
}

type repo struct {
	db *sqlx.DB
}

const insertQuery = "INSERT INTO recipe(user_id, name, description, actions) VALUES ($1, $2, $3, $4) RETURNING id"

func (r *repo) AddRecipe(ctx context.Context, newRecipe recipe.Recipe) (uint64, error) {
	row := r.db.QueryRowxContext(
		ctx,
		insertQuery,
		newRecipe.UserId(), newRecipe.Name(), newRecipe.Description(), pq.Array(newRecipe.Actions()),
	)
	var newRecipeId uint64
	if err := row.Scan(&newRecipeId); err != nil {
		return 0, err
	}
	return newRecipeId, nil
}

func (r *repo) AddRecipes(ctx context.Context, recipes []recipe.Recipe) error {
	tx, txErr := r.db.BeginTx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback()
	stmt, stmtErr := tx.PrepareContext(ctx, insertQuery)
	if stmtErr != nil {
		return stmtErr
	}
	defer stmt.Close()
	for _, rec := range recipes {
		_, execErr := stmt.ExecContext(ctx, rec.UserId(), rec.Name(), rec.Description(), pq.Array(rec.Actions()))
		if execErr != nil {
			return execErr
		}
	}
	return tx.Commit()
}

func (r *repo) ListRecipes(ctx context.Context, limit, offset uint64) ([]recipe.Recipe, error) {
	result := make([]recipe.Recipe, 0, limit)
	rows, err := r.db.QueryxContext(
		ctx, "SELECT id, user_id, name, description, actions FROM recipe ORDER BY id LIMIT $1 OFFSET $2", limit, offset,
	)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var dbRecipeId, userId uint64
		var name, description string
		var actions []string
		if scanErr := rows.Scan(&dbRecipeId, &userId, &name, &description, pq.Array(&actions)); err != nil {
			if scanErr == sql.ErrNoRows {
				return nil, NotFoundError
			}
			return nil, scanErr
		}
		result = append(result, recipe.New(dbRecipeId, userId, name, description, actions))
	}
	return result, nil
}

func (r *repo) DescribeRecipe(ctx context.Context, recipeId uint64) (*recipe.Recipe, error) {
	row := r.db.QueryRowxContext(
		ctx, "SELECT user_id, name, description, actions FROM recipe WHERE id = $1", recipeId,
	)
	var userId uint64
	var name, description string
	var actions []string
	if err := row.Scan(&userId, &name, &description, pq.Array(&actions)); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFoundError
		}
		return nil, err
	}
	newRecipe := recipe.New(recipeId, userId, name, description, actions)
	return &newRecipe, nil
}

func (r *repo) RemoveRecipe(ctx context.Context, recipeId uint64) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM recipe WHERE id = $1", recipeId)
	if err != nil {
		return err
	}
	rowsAffected, rowsAffectedErr := result.RowsAffected()
	if rowsAffectedErr != nil {
		return err
	}
	if rowsAffected == 0 {
		return NotFoundError
	}
	return nil
}

func (r *repo) UpdateRecipe(ctx context.Context, newRecipe recipe.Recipe) error {
	result, err := r.db.ExecContext(
		ctx,
		"UPDATE recipe SET user_id = $1, name = $2, description = $3, actions = $4 WHERE id = $5",
		newRecipe.UserId(),
		newRecipe.Name(),
		newRecipe.Description(),
		pq.Array(newRecipe.Actions()),
		newRecipe.Id())
	if err != nil {
		return err
	}
	rowsAffected, rowsAffectedErr := result.RowsAffected()
	if rowsAffectedErr != nil {
		return err
	}
	if rowsAffected == 0 {
		return NotFoundError
	}
	return nil
}
