package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sxc/ginko/api"
	db "github.com/sxc/ginko/db/sqlc"
	"github.com/sxc/ginko/util"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:password@localhost:5432/ginko?sslmode=disable"
// 	ServerAddress = "0.0.0.0:3000"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	// conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	// err = server.Start(config.ServerAddress)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
