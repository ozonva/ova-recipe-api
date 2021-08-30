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
		dbMock        *sql.DB
		queryMock sqlmock.Sqlmock
		repoTest repo.RecipeRepo
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
	Describe("AddRecipe tests", func() {
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
})
