package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "github.com/vpereira/micro-topt/models"
  )
func main() {
  db, _ := gorm.Open("sqlite3", fmt.Sprintf("%s/users.db","db"))
  // Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
  db.DB()


  // move to somewhere else, maybe lib or utils
  pwd := model.GetHash("foobarmar")

  // Create table
  db.CreateTable(&model.User{})
    user := model.User {
      Login: "foobar",
      Password:  pwd,
    }
    db.Create(&user)
    // Drop table
    //db.DropTable(&User{})
  }
