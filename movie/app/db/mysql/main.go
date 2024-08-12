package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

var db *sql.DB

func Connect() {
	fmt.Println("trying to connect to mysql")

	// Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   "admin",
	// 	Passwd: "1234",
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "movie",
	// }
	var err error
	db, err = sql.Open("mysql", "admin:1234@tcp(mysql:3306)/movie?charset=utf8")

	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func SelectByName(name string) (User, error) {
	var user User

	row := db.QueryRow("SELECT * FROM users WHERE name = ?", name)
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("userByName %s: no such user", name)
		}
		return user, fmt.Errorf("userByName %s: %v", name, err)
	}
	return user, nil
}

func SelectUsers() ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("select all users: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("select all users: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("select all users: %v", err)
	}
	return users, nil
}

func AddUser(user User) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO users (name, email) VALUES (?, ?)",
		user.Name, user.Email,
	)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil
}
