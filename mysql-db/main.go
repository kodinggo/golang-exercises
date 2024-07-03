package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // hanya menjalankan fungsi init()
)

type User struct {
	ID        int64     `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// db environment
	var (
		dbUsername = "root"
		dbPassword = ""
		dbName     = "kodinggo"
		dbHost     = "localhost:3306"
	)

	// prepare connection string
	// charset=utf8mb4 agar dapat menyimpan semua karakter unicode
	// parseTime=true agar dapat diparsing dari timestamp ke tipe time.Time
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		dbUsername,
		dbPassword,
		dbHost,
		dbName)
	connDB, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("failed to connect server db, error: %s", err.Error())
	}

	connDB.SetConnMaxLifetime(3 * time.Minute) // berapa lama koneksi mysql tetap open
	connDB.SetConnMaxIdleTime(3 * time.Minute) // berapa lama koneksi mysql tetap open dalam kondisi idle (tidak dipakai)
	connDB.SetMaxOpenConns(10)                 // maksimal jumlah user yang bukan koneksi secara bersamaan
	connDB.SetMaxIdleConns(10)                 // maksimal jumlah user yang bukan koneksi secara bersamaan dalam kondisi idle(tidak dikapai)

	// Verifikasi koneksi
	if err := connDB.Ping(); err != nil {
		log.Fatalf("failed to connect server db, error: %s", err.Error())
	}

	log.Println("successfully to connect server db")

	r := repo{db: connDB}

	// users, err := r.findAllUsers()
	// if err != nil {
	// 	log.Fatalf("failed find all users, error: %s", err.Error())
	// }
	// fmt.Println(users)

	// user, err := r.findUserByID(1)
	// if err != nil {
	// 	log.Fatalf("failed find user by id, error: %s", err.Error())
	// }
	// fmt.Println(user)

	insertedUser, err := r.insertUser(User{
		Fullname: "Joko",
		Email:    "joko@gmail.com",
	})
	if err != nil {
		log.Fatalf("failed insert new user, error: %s", err.Error())
	}
	fmt.Println(insertedUser)
}

type repo struct {
	db *sql.DB
}

func (r repo) findAllUsers() ([]User, error) {
	rows, err := r.db.Query("SELECT id, fullname, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Println("failed to scan rows")
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (r repo) findUserByID(id int64) (*User, error) {
	row := r.db.QueryRow("SELECT id, fullname, email, created_at FROM users WHERE id = ?", id)
	var user User
	err := row.Scan(&user.ID, &user.Fullname, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repo) insertUser(user User) (*User, error) {
	timeNow := time.Now().UTC()
	result, err := r.db.Exec("INSERT INTO users (fullname, email, created_at) VALUES (?, ?, ?)", user.Fullname, user.Email, timeNow)
	if err != nil {
		return nil, err
	}
	user.ID, _ = result.LastInsertId()
	user.CreatedAt = timeNow
	return &user, nil
}
