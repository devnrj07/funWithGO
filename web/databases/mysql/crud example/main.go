package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {

	//This won't immediately connect to db thus error won't be thrown
	db, err = sql.Open("mysql", "user:password@localhost:3306/mydb")
	check(err)
	defer db.Close()

	//Important to test connection as a connection is only made when needed not when 'OPEN'
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/friends", friends)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err = http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "connection successful to db")
	check(err)
}
func friends(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`select names from friends;`)
	check(err)
	defer rows.Close()

	//data to be used in query rows is raw data
	var s, names string

	//Will keep looping until there's  no next row
	for rows.Next() {
		err = rows.Scan(&names) // to read the raw data as string
		check(err)
		s += names + "\n"
	}
	io.WriteString(w, s)

	//fmt.Fprintln(w,s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`create table customer (id int, name varchar(20));`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Table created", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`insert into customer values( 1, "Jimmy");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Record Inserted", n)
}
func read(w http.ResponseWriter, req *http.Request) {
	//Use perpare for where & conditional reading
	//stmt, err := db.Prepare(`select * from customer;`)

	//for simple reading just user Query
	rows, err := db.Query(`select * from customer;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "List of names :", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`update customer set name ="jhonny" where name = "jimmy";`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Updated records", n)
}

func delete(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`delete from customer where name = "jhonny";`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)

}
func drop(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
