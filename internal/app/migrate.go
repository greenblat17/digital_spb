package app

import (
	"errors"
	"os"
	"time"

	// migrate tools
	"github.com/golang-migrate/migrate/v4"
	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	log "github.com/sirupsen/logrus"
)

const (
	defaultAttempts = 3
	defaultTimeout  = time.Second
)

func init() {
	os.Setenv("PG_URL", "postgres://user:password@db:5432/spb")
	databaseURL := os.Getenv("PG_URL")
	if len(databaseURL) == 0 {
		log.Fatalf("migrate: environment variable not declared: PG_URL")
	}

	databaseURL += "?sslmode=disable"

	var (
		attempts = defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migration", databaseURL)
		if err == nil {
			break
		}

		log.Printf("error creating migration ", err.Error())

		log.Printf("Migrate: pgdb is trying to connect to the database. Attempts left: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: pgdb connect error: %s", err)
	}

	err = m.Up()
	defer func() {
		_, _ = m.Close()
	}()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}
