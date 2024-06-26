package test_utils

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	cPostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"

	DB "github.com/newtri-science/synonym-tool/db"
)

type TestDatabase struct {
	DB        *sql.DB
	Url       string
	Container testcontainers.Container
}

func CreateTestContainer(ctx context.Context) TestDatabase {
	var logger *zap.Logger
	logger, _ = zap.NewDevelopment()

	globalEnv := SetupEnvironment()

	container, err := cPostgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16.2-alpine"),
		cPostgres.WithDatabase(globalEnv.db.Name),
		cPostgres.WithUsername(globalEnv.db.User),
		cPostgres.WithPassword(globalEnv.db.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	dbURL, err := container.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("failed to get container connection string: %s", err)
	}

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}

	migrator := DB.NewMigrator(db, "testdata", logger.Sugar())
	if err := migrator.Up(); err != nil {
		log.Fatalf("failed to run migrations: %s", err)
	}

	return TestDatabase{
		DB:        db,
		Url:       dbURL,
		Container: container,
	}
}
