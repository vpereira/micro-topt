package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "github.com/vpereira/micro-topt/backend"
  )
func main() {
  db, _ := gorm.Open("sqlite3", fmt.Sprintf("%s/users.db","db"))
  // Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
  db.DB()


  // move to somewhere else, maybe lib or utils
  pwd := backend.GetHash("foobarmar")

  // Create table
  db.CreateTable(&backend.User{})
    user := backend.User {
      Login: "foobar",
      Password:  pwd,
    }
    db.Create(&user)
    // Drop table
    //db.DropTable(&User{})
  }
