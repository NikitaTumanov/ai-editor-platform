package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	pool, err := pgxpool.NewWithConfig(context.Background(), Config())
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
		return nil, err
	}

	fmt.Println("Connected to the database!!")

	return pool, nil
}
