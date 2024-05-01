package database

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type DB struct {
	conn *pgx.Conn
}

func NewDbConnection(dbConnection string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), dbConnection)
	if err != nil {
		println("Unable to connect to the database.", err.Error())
		return nil, err
	}
	db := &DB{
		conn: conn,
	}

	return db, nil
}
