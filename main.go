package main


// TODO
// a cli to manage db, intialize stuffs and so

import (
  "fmt"
  "github.com/vpereira/micro-topt/models"
  "github.com/vpereira/micro-topt/libs/utils"
  )
func main() {
  // we gonna be using postgres, but for now sqlite3
  db,_ := utils.SetupDB("sqlite3", fmt.Sprintf("%s/users.db","db"))

  // get a sha256 from the string
  pwd := utils.GetHash("foobarmar")

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
