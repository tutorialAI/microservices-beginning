package main

import (
	"database/sql"
	"fmt"
)

type DSN struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func main() {
	env := DSN{
		Host:     "clickhouse",
		Port:     8123,
		Username: "root",
		Password: "1234",
		Database: "analytics",
	}

	_, err := connect(&env)
	if err != nil {
		panic((err))
	}

	// ctx := context.Background()
	// rows, err := conn.Query(ctx, "SELECT name,toString(uuid) as uuid_str FROM system.tables LIMIT 5")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for rows.Next() {
	// 	var (
	// 		name, uuid string
	// 	)
	// 	if err := rows.Scan(
	// 		&name,
	// 		&uuid,
	// 	); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("name: %s, uuid: %s",
	// 		name, uuid)
	// }

}

func connect(env *DSN) (*sql.DB, error) {
	conn, err := sql.Open("clickhouse", "http://127.0.0.1:8123/default")

	if err != nil {
		fmt.Println(err)
	}

	return conn, nil
}
