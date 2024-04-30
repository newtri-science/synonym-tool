package services_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/newtri-science/synonym-tool/services"
	"github.com/newtri-science/synonym-tool/test_utils"
)

var DB *sql.DB

func TestMain(m *testing.M) {
	// Setup test environment
	ctx := context.Background()
	testDb := test_utils.CreateTestContainer(ctx)
	container := testDb.Container

	DB = testDb.Db

	// Run the actual tests
	exitCode := m.Run()

	// Perform tear down
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	// Exit with the exit code from the tests
	os.Exit(exitCode)

}

func TestGetAllFoods(t *testing.T) {
	repo := services.NewFoodEntryService(DB, nil)
	expectedSize, err := repo.Count()
	if err != nil {
		t.Errorf("Error while trying to count foods: %s", err)
	}

	foods, err := repo.GetAllFoodEntries()
	if err != nil {
		t.Errorf("Error while trying to get all foods: %s", err)

	}

	actualSize := len(foods)
	if actualSize != expectedSize {
		t.Errorf("actual size %v is not expectedSize %v", actualSize, expectedSize)
	}
}



func TestCountFoods(t *testing.T) {
	repo := services.NewFoodEntryService(DB, nil)
	count, err := repo.Count()
	if err != nil {
		t.Errorf("Error while trying to count foods: %s", err)
	}
	if count == 0 {
		t.Errorf("No foods found")
	}
}
