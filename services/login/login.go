package main

import (
  "fmt"
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/vpereira/micro-topt/libs/utils"
  "github.com/vpereira/micro-topt/models"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/sessions"
  )

// Binding from JSON
type LoginJSON struct {
    User     string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func RequireLogin(rw http.ResponseWriter, req *http.Request,
  s sessions.Session, db *gorm.DB, c martini.Context) {
    user := model.User{}

    v := s.Get("userId")
    if v != nil {
      iId := v.(int64)

      if db.Where(&model.User{Id: iId}).First(&user).RecordNotFound() {
        // user not found
        http.Redirect(rw,req,"/login",http.StatusFound)
      }
      c.Map(user)
    }

}

func PostData(db *gorm.DB,json LoginJSON,r render.Render,s sessions.Session) {
  // By this point, I assume that my own middleware took care of any errors
  user := model.User{}
  pwd := utils.GetHash(json.Password)
  if db.Where(&model.User{Login: json.User, Password: pwd}).First(&user).RecordNotFound() {
    r.JSON(401,map[string]interface{}{"bad": "world"})
  } else {
    s.Set("userId",user.Id)
    r.JSON(200,map[string]interface{}{"hello": "world"})
  }
}
func main() {
    m := martini.Classic()

    m.Use(render.Renderer())
    // Sessions
    store := sessions.NewCookieStore([]byte("0xd34db33f"))
    m.Use(sessions.Sessions("auth",store))
    // Connect with the db. Today we are using sqlite3. In the future, postgres
    db,_ := utils.SetupDB("sqlite3", fmt.Sprintf("%s/users.db","db"))
    defer db.Close()
    // to test it:
    // curl -H "Content-Type: application/json" \
    // -X POST -d '{"user":"foobar","password":"foobarmar"}' \
    // http://localhost:8080/login
    m.Post("/login", binding.Json(LoginJSON{}), binding.ErrorHandler, PostData)
    m.Get("/logout", func(s sessions.Session, r render.Render){
      s.Delete("userId")
      r.JSON(200,map[string]interface{}{"bye": "world"})
      })
    m.Get("/",RequireLogin,func(u *model.User, r render.Render){
       r.JSON(200,map[string]interface{}{"hello": u.Login})
      })

    // Listen and serve it
    m.Run()
}
