package goblet

import (
    "fmt"
    "database/sql"
    "github.com/lib/pq"
    // k "time"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME     = "gobletserver"
)

type UsernameBusyError struct { msg string }
func (e *UsernameBusyError) Error() string { return e.msg }
type EmailBusyError struct {msg string}
func (e *EmailBusyError) Error() string { return e.msg }

func connect() (db *sql.DB, err error) {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err = sql.Open("postgres", dbinfo)
    return db, err
}

func AddUser(userName, email string, passw int) error {
    db, err := connect()
    defer db.Close()
    checkErr(err)
    var lastInsertId int
    err = db.QueryRow("INSERT INTO users (username, email, passw) VALUES ($1, $2, $3) RETURNING id;",
                        userName, email, passw).Scan(&lastInsertId)
    if err == nil {
        fmt.Println("last inserted id =", lastInsertId)
        return nil
    } else if pqErr := err.(*pq.Error); pqErr.Code.Name() == "unique_violation"{
        if pqErr.Constraint == "unique_email" {
            return &EmailBusyError{"That email has already been signed up."}
        } else if pqErr.Constraint == "unique_user" {
            return &UsernameBusyError{"That username is already taken."}
        }
    }
    return err
}

func Connect(){
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    /* fmt.Println("# Inserting values")

    var lastInsertId int
    err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
    checkErr(err)
    fmt.Println("last inserted id =", lastInsertId)

    fmt.Println("# Updating")
    stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
    checkErr(err)

    res, err := stmt.Exec("astaxieupdate", lastInsertId)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")*/

    fmt.Println("# Querying")
    rows, err := db.Query("SELECT * FROM users")
    checkErr(err)

    for rows.Next() {
        var id int
        var username string
        var email string
        var passw int
        err = rows.Scan(&id, &username, &email, &passw)
        checkErr(err)
        fmt.Println("id | username |          email          | passw ")
        fmt.Printf("%2v | %8v | %23v | %5v\n", id, username, email, passw)
    }
/*
    fmt.Println("# Deleting")
    stmt, err = db.Prepare("delete from userinfo where uid=$1")
    checkErr(err)

    res, err = stmt.Exec(lastInsertId)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")*/
}

func checkErr(err error) {
    if err != nil {
    }
}
