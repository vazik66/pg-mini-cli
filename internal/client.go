package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func connectToDb(username string, password string, host string, port string) (*pgx.Conn, error) {
    uri := fmt.Sprintf("postgres://%s:%s@%s:%s", username, password, host, port)

	conn, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
        log.Printf("Error trying to ping database: %v\n", err)
		return nil, err
	}

	return conn, nil
}

type client struct {
	Conn *pgx.Conn
}

func NewClient(username string, password string, host string, port string) (*client, error) {
    conn, err := connectToDb(username, password, host, port)
    if err != nil {
        return nil, err
    }
	return &client{Conn: conn}, nil
}

func (c *client) ListDatabases() ([]string, error) {
	q := "SELECT datname FROM pg_database"

	rows, err := c.Conn.Query(context.Background(), q)
	if err != nil {
		log.Printf("Error while executing statement: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	databases := make([]string, 0)
	for rows.Next() {
		var database string
		if err = rows.Scan(&database); err != nil {
			log.Printf("Error while parsing rows: %v\n", err)
			return nil, err
		}
		databases = append(databases, database)
	}
	return databases, nil
}

func (c *client) DeleteDatabase(database string) error {
    q := fmt.Sprintf("DROP DATABASE %s;", database)
	rows, err := c.Conn.Query(context.Background(), q)
	if err != nil {
		log.Printf("Error while executing statement: %v\n", err)
		return err
	}
    defer rows.Close() 
    return nil
}

func (c *client) CreateBackup(database string, backupName string) error {
    return nil
}

func (c *client) RestoreFromBackup(backupName string) (error){
    return nil
}
