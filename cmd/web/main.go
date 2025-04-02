package main

import (
	"app/pkg/logic"
	"encoding/gob"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

const IP_ADDR string = "localhost:5000"

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("ui/html/*")
	router.StaticFile("style.css", "./ui/static/styles/style.css")
	router.StaticFile("icon.svg", "./ui/static/images/icon.svg")
	router.StaticFile("icon-light.svg", "./ui/static/images/icon-light.svg")
	router.StaticFile("wave.png", "./ui/static/images/wave.png")
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("DATA", store))
	initializeRoutes()
	gob.Register(logic.User{})
	log.Println("Сервер запускается http://" + IP_ADDR)
	router.Run(IP_ADDR)
}
