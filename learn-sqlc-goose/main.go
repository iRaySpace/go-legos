package main

import (
	"context"
	"log"

	"github.com/irayspace/go-legos/learn-sqlc-goose/data/generated"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := generated.New(conn)
	users, err := queries.ListUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}
