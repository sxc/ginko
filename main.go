package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sxc/ginko/api"
	db "github.com/sxc/ginko/db/sqlc"
)

const (
	dbDriver       = "postgres"
	dbSource       = "postgresql://root:password@localhost:5432/ginko?sslmode=disable"
	serverAddresss = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddresss)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
