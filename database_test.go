package go_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_db?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(16 * time.Minute)
	return db

}

func InsertData() {
	db := OpenConnection()
	defer db.Close()
	ctx := context.Background()
	// faker := faker.NAME
	_, err := db.ExecContext(ctx, "insert customer values('1', 'Moch Juang');")
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert Berhasil")
}
func Err(err interface{}) {
	if err != nil {
		panic(err)
	}
}
func GetData() {
	db := OpenConnection()
	defer db.Close()
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "select * from customer;")
	Err(err)
	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float32
		var createdAt, birthDay time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDay, &married)
		if err != nil {
			Err(err)
		}
		fmt.Println("id : ", id)
		fmt.Println("name : ", name)
		fmt.Println("email : ", email)
		fmt.Println("balance : ", balance)
		fmt.Println("rating : ", rating)
		fmt.Println("createdAt : ", createdAt)
		fmt.Println("birthDay : ", birthDay)
		fmt.Println("married : ", married)
		println("----------------------------------")
	}

	rows.Close()
}

func TestDatabase(t *testing.T) {
	// InsertData()
	GetData()
}
