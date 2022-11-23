package main

import (
	"belajar-golang-database/controller"
	"belajar-golang-database/database"
	"belajar-golang-database/repository"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func PingDB() {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err := db.PingContext(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Koneksi berhasil")
}

func Insert(name, email string, balance int, rating float32, birth_date string, married bool) {
	var db *sql.DB = database.GetConnection()
	defer db.Close()

	var ctx context.Context = context.Background()
	// query := fmt.Sprintf("INSERT INTO customer(name,email,balance,rating,birth_date,married) VALUES('%s','%s','%d','%.1f','%s',%t)", name, email, balance, rating, birth_date, married)
	query := ("INSERT INTO customer(name,email,balance,rating,birth_date,married) VALUES(?,?,?,?,?,?)")
	res, err := db.ExecContext(ctx, query, name, email, balance, rating, birth_date, married)

	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())
	fmt.Println("Insert Berhasil")
}

func InsertUser(username string, password string) {
	var db *sql.DB = database.GetConnection()
	defer db.Close()

	var ctx context.Context = context.Background()
	var query = ("INSERT INTO user(username,password) VALUES(?,?)")
	var res, err = db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())
	fmt.Println("Insert Berhasil")
}

func SelectCustomer() {
	var db *sql.DB = database.GetConnection()
	defer db.Close()

	var ctx context.Context = context.Background()
	// var rows, err = db.QueryContext(ctx, "SELECT id,name,email,balance, rating, birth_date, married,created_at FROM customer")
	var query = ("SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer")
	var rows, err = db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email sql.NullString
		var balance int
		var rating float32
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool

		var err error = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}

		var emailString string
		if email.Valid {
			emailString = email.String
		}

		var birthDate time.Time
		if birth_date.Valid {
			birthDate = birth_date.Time
		}
		fmt.Printf("Id: %d Name: %s Email:%v Balance:%d Rating:%.1f Birth Date:%v Married:%t Created At:%v\n", id, name, emailString, balance, rating, birthDate, married, created_at)
	}
}

func GetUser() {
	var username string = "admin"
	// var username string = "' or ''='"
	var password string = "admin"
	// var password string = "' or ''='"
	var db *sql.DB = database.GetConnection()
	defer db.Close()

	var ctx context.Context = context.Background()
	var query string = "SELECT username FROM user WHERE username=? AND password=? LIMIT 1"
	fmt.Println(query)
	var rows, err = db.QueryContext(ctx, query, username, password)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var resUsername string
		rows.Scan(&resUsername)
		fmt.Println("Username: ", resUsername)
	}
}

func PreparedStatInsertComment() {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()
	stmt, err := db.PrepareContext(ctx, "INSERT INTO comment(email,comment) VALUES(?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("Eko%d@gmail.com", i)
		comment := fmt.Sprintf("Ini komen ke-%d", i)
		res, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsert, _ := res.LastInsertId()
		fmt.Printf("Ini merupakan data ke %d\n", lastInsert)
	}
}

func Transaction() {
	db := database.GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// ctx := context.Background()
	stmt, err := tx.Prepare("INSERT INTO comment(email,comment) VALUES(?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 30; i++ {
		email := fmt.Sprintf("Eko%d@gmail.com", i)
		comment := fmt.Sprintf("Ini komen ke-%d", i)
		res, err := stmt.Exec(email, comment)
		if err != nil {
			panic(err)
		}
		// if i == 24 {
		// 	tx.Rollback()
		// 	return
		// }
		lastInsert, _ := res.LastInsertId()
		fmt.Printf("Ini merupakan data ke %d\n", lastInsert)
	}

	tx.Commit()
}

func main() {
	// Insert("Luthfi", "luthfiarsyad68@gmail.com", 12000000, 5.0, "2000-05-04", false)
	// Insert("Umar", "umar@gmail.com", 2000000, 4.0, "1999-09-09", false)
	// Insert("Aan", "Aan@gmail.com", 100000, 3.0, "1999-09-09", false)
	// Insert("Wira", "Wira@gmail.com", 100000000, 5.0, "1999-09-09", false)
	// Insert("Angga", "Angga@gmail.com", 10000000, 4.5, "1999-09-09", true)
	// Insert("Agres", "Agres@gmail.com", 100000000, 5.0, "2021-09-13", false)
	// Insert("Riyan", "Riyan@gmail.com", 1000000, 3.0, "2021-09-13", false)
	// SelectCustomer()
	// GetUser()
	// InsertUser("Luthfi", "123")

	// PreparedStatInsertComment()
	// Transaction()
	db := database.GetConnection()
	commentImpl := repository.NewCommentRepository(db)
	// comment, err := controller.CreateComment(commentImpl, "Dulloh@gmail.com", "Dulloh comment")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(comment.Id, comment.Email, comment.Comment)

	comments, err := controller.GetComment(commentImpl)
	if err != nil {
		panic(err)
	}
	fmt.Println(comments)

	comment, err := controller.GetCommentById(commentImpl, 28)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}
