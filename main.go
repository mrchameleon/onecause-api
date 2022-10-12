package main

// reference for building onto this
// https://betterprogramming.pub/how-to-create-a-simple-web-login-using-gin-for-golang-9ac46a5b0f89

import (
  "fmt"
  "time"
  "net/http"
  // "net/url"
  "github.com/gin-gonic/gin"
)

type user struct {
ID string `json:"id"`
Email string `json:"email"`
Password string `json:"password"`
}

var users = []user{
  {ID: "1", Email: "c137@onecause.com", Password: "#th@nH@rm#y#r!$100%D0p#"},
}

// I know this isn't needed for the app reqs, but kept it in as I was learning the basics of gin
func getUsers(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
  // for sake of time, this only works with 
  // request content-type multipart/form-data
  // if needed, could make changes to make it work with raw post data and a "application/json" content-type header

  // {"username": "", "password": "", "token": ""}
  // token should be formatted as HHMM using current time

  // bug concern: what if server token mismatch when compared to client provided token
  // somewhat common for the system time to be set incorrectly on users systems. 
  // (not as common since modern OS sync over the web with source of truth, but worth mentioning)

  // might add in some kind of ~2-3 min margin of error as long as the password provided is correct
  // or provide user feedback in an error to ensure their computer time is synced with current time

  username := c.PostForm("username")
  password := c.PostForm("password")
  client_token := c.PostForm("token")
  server_token := time.Now().Format("1504")

  // print login attempts in the server output (even better, could go into a log file)
    if server_token == client_token && username == "c137@onecause.com" && password == "#th@nH@rm#y#r!$100%D0p#" {
    fmt.Println("Login Success: Username: ", username)
    response := []byte(`{"status": "success", "detail": "login success!"}`)

    // cors header
    c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
    c.Data(http.StatusOK, "application/json", response)
    // location := url.URL{Path: "https://onecause.com",}
    // c.Redirect(http.StatusFound, location.RequestURI())
  } else {
    fmt.Println("Login Failure: Username: ", username, "Client Token: ", client_token, "Server Token: ", server_token)
    response := []byte(`{"status": "error", "detail": "Invalid Login Details - Please try again!"}`)

    // cors header
    c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
    c.Data(http.StatusUnauthorized, "application/json", response)
  }

}


func main() {
  router := gin.Default()
  router.GET("/api/users", getUsers)
  router.POST("/api/login", Login)

  router.Run("localhost:8080")
}
