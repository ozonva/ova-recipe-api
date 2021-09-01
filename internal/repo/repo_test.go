package repo_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/repo"
)

var _ = Describe("Repo", func() {
	var (
		ctx       context.Context
		dbMock    *sql.DB
		queryMock sqlmock.Sqlmock
		repoTest  repo.RecipeRepo
	)
	BeforeEach(func() {
		ctx = context.Background()
		db, sqlMock, err := sqlmock.New()
		Expect(err).To(BeNil())
		dbMock = db
		queryMock = sqlMock
		r, newRepoErr := repo.New(db)
		Expect(newRepoErr).To(BeNil())
		repoTest = r
	})
	AfterEach(func() {
		defer dbMock.Close()
	})
	Describe("AddRecipe test", func() {
		It("returns new id", func() {
			expectedR := recipe.New(1, 1, "test", "test", []string{"test"})
			queryMock.ExpectQuery("INSERT INTO recipe").
				WithArgs(expectedR.UserId(), expectedR.Name(), expectedR.Description(), pq.Array(expectedR.Actions())).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			newId, err := repoTest.AddRecipe(ctx, expectedR)
			Expect(err).To(BeNil())
			Expect(newId).To(Equal(expectedR.Id()))
		})
	})
	Describe("DescribeRecipe tests", func() {
		It("returns recipe", func() {
			expectedR := recipe.New(1, 1, "test", "test", []string{"test"})
			queryMock.ExpectQuery("SELECT user_id, name, description, actions").
				WithArgs(uint64(1)).
				WillReturnRows(sqlmock.NewRows([]string{"user_id", "name", "description", "actions"}).
					AddRow(expectedR.UserId(), expectedR.Name(), expectedR.Description(), pq.Array(expectedR.Actions())))
			newRecipe, err := repoTest.DescribeRecipe(ctx, expectedR.Id())
			Expect(err).To(BeNil())
			Expect(*newRecipe).To(Equal(expectedR))
		})
	})
})
