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
	AddRecipes(recipes []recipe.Recipe) error
	ListRecipes(ctx context.Context, limit, offset uint64) ([]recipe.Recipe, error)
	DescribeRecipe(ctx context.Context, recipeId uint64) (*recipe.Recipe, error)
	RemoveRecipe(ctx context.Context, recipeId uint64) error
}

func New(dsn string) (RecipeRepo, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &repo{db: db}, nil
}

type repo struct {
	db *sqlx.DB
}

func (r *repo) AddRecipe(ctx context.Context, newRecipe recipe.Recipe) (uint64, error) {
	row := r.db.QueryRowxContext(
		ctx,
		"INSERT INTO recipe(user_id, name, description, actions) VALUES ($1, $2, $3, $4) RETURNING id",
		newRecipe.UserId(), newRecipe.Name(), newRecipe.Description(), pq.Array(newRecipe.Actions()),
	)
	var newRecipeId uint64
	if err := row.Scan(&newRecipeId); err != nil {
		return 0, err
	}
	return newRecipeId, nil
}

func (r *repo) AddRecipes(_ []recipe.Recipe) error {
	panic("Not implemented")
}

func (r *repo) ListRecipes(ctx context.Context, limit, offset uint64) ([]recipe.Recipe, error) {
	result := make([]recipe.Recipe, 0, limit)
	rows, err := r.db.QueryxContext(
		ctx, "SELECT * FROM recipe ORDER BY id LIMIT $1 OFFSET $2", limit, offset,
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
	result, err := r.db.ExecContext(ctx, "DELETE FROM recipe WHERE id = ?", recipeId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return NotFoundError
	}
	return nil
}
