package utils

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "crypto/sha256"
  "fmt"
  )


//it should be a function from the user struct
func GetHash(str string) string {
    return fmt.Sprintf("%x",sha256.Sum256([]byte(str)))
}

func SetupDB(driver string, connection_string string) (gorm.DB,error) {
  db, err := gorm.Open(driver,connection_string)
  db.DB()
  return db,err
}
