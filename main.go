package main

import (
	"database/sql"
	"os"

	"github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/api"
	db "github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/util"

	// _ "github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/doc/statik"
	// "github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/gapi"
	// "github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/pb"
	// "github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/util"
	// "github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/worker"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)

	runGinServer(config, store)

}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
