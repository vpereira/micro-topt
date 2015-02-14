package main

import (
  "github.com/gin-gonic/gin"
  )

// Binding from JSON
type LoginJSON struct {
    User     string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}
func main() {
    r := gin.Default()

    // Example for binding JSON ({"user": "manu", "password": "123"})
    // TODO
    // get the information from db
    // to test it:
    // curl -H "Content-Type: application/json"  -X POST -d '{"user":"foo","password":"123"}'  http://localhost:8080/login
    r.POST("/login", func(c *gin.Context) {
      var json LoginJSON

      c.Bind(&json) // This will infer what binder to use depending on the content-type header.
      if json.User == "foo" && json.Password == "123" {
        c.JSON(200, gin.H{"status": "you are logged in"})
      } else {
          c.JSON(401, gin.H{"status": "unauthorized"})
      }
    })
    // Listen and server on 0.0.0.0:8080
    r.Run(":8080")
}
