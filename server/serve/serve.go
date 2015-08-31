package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sleep2death/civ/server"
	"log"
	"net/http"
)

const (
	Base = "./"
	App  = "app/"

	Port = ":8080"

	DBusr  = "mgo"
	DBpwd  = "HelloMG0"
	DBport = 27017
	DBname = "civ"

	ErrorUserNotExist = "User Not Exist"
	ErrorLoginError   = "Login Error"
)

func main() {
	server.ConnectDB(DBname, DBusr, DBpwd, DBport)

	//set to release mode
	//gin.SetMode(gin.ReleaseMode)

	//set to debug mode
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Static("/", Base+"/app")

	log.Fatal(router.Run(Port))
}

func ServeHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
