// https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/#overview

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	// host     = "localhost"
	host     = "192.168.133.38"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "mydb"
)

var db *sql.DB

func main() {
	// "Signin" and "Signup" are handler that we will implement
	// http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	// Init databsae connection
	initDB()
	// start server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func initDB() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)

	// Connect to the postgres db
	// you might have to change the connectio nstring to add your database crednetials
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
}

type Credentials struct {
	Password string `json: "password", db: "password"`
	Username string `json: "username", db: "username"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Signup Request")
	// Parase and decdoe the request body into a new `Credentials` instance
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is somethign wrong with the request body, return a 400 status
		fmt.Println("Encountered Error!")
		fmt.Println(err)
		fmt.Println("Sending Status 400")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Salt and hash the PW using the bcrypt algorithm
	// The second arument is the cost of hashing, which we arbitrarily set as 8
	// (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	// Insert the username, along with the hashed password into the database
	_, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword))
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We reach this point if the credientials we correclty stored in the DB
	// and the defautl status of 200 is sent back
}
