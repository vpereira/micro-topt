package main

import (
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
    m.Use(render.Renderer())
    // Example for binding JSON ({"user": "manu", "password": "123"})
    // TODO
    // get the information from db
    // to test it:
    // curl -H "Content-Type: application/json"  -X POST -d '{"user":"foo","password":"123"}'  http://localhost:8080/login
    m.Post("/login", binding.Json(LoginJSON{}), binding.ErrorHandler, func(json LoginJSON,r render.Render) {
      // By this point, I assume that my own middleware took care of any errors
      if json.User == "admin" && json.Password == "123" {
        r.JSON(200,map[string]interface{}{"hello": "world"})
      }else{
        r.JSON(401,map[string]interface{}{"bad": "world"})
      }
    })
    // Listen and server on 0.0.0.0:8080
    m.Run()
}
