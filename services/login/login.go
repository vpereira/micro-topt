package main

import (
  "fmt"
  "github.com/vpereira/micro-topt/libs/utils"
  "github.com/vpereira/micro-topt/models"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"
  )

// Binding from JSON
type LoginJSON struct {
    User     string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}
func main() {
    m := martini.Classic()
    user := model.User{}
    m.Use(render.Renderer())
    db,_ := utils.SetupDB("sqlite3", fmt.Sprintf("%s/users.db","db"))
    defer db.Close()
    // to test it:
    // curl -H "Content-Type: application/json"  -X POST -d '{"user":"foo","password":"123"}'  http://localhost:8080/login
    m.Post("/login", binding.Json(LoginJSON{}), binding.ErrorHandler, func(json LoginJSON,r render.Render) {
      // By this point, I assume that my own middleware took care of any errors
      pwd := utils.GetHash(json.Password)
      if db.Where(&model.User{Login: json.User, Password: pwd}).First(&user).RecordNotFound() {
        r.JSON(401,map[string]interface{}{"bad": "world"})
      } else {
        r.JSON(200,map[string]interface{}{"hello": "world"})
      }
    })
    // Listen and server on 0.0.0.0:8080
    m.Run()
}
