package model

import (
  "fmt"
  "crypto/sha256"
  )


//here we should define our structures and so on
type User struct {
  Id  int64
  Login string  `sql:"size:256"`
  Password string  `sql:"size:256"`
  Key string  `sql:"size:256"`
}


//it should be a function from the user struct
func GetHash(str string) string {
  return fmt.Sprintf("%x",sha256.Sum256([]byte(str)))
}
