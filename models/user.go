package model


//here we should define our structures and so on
type User struct {
  Id  int64
  Login string  `sql:"size:256"`
  Password string  `sql:"size:256"`
  Key string  `sql:"size:256"`
}
