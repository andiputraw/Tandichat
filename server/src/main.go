package main

import (
	"andiputraw/Tandichat/src/config"
	"andiputraw/Tandichat/src/database"
	"andiputraw/Tandichat/src/routes"
	"andiputraw/Tandichat/src/websocket"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/olahol/melody"
)

// * PROTOTYPE>
type sendedData struct {
	Data   string
	Target string
}

func main() {
	err := godotenv.Load()
	config.InitConfig()
	if err != nil {
		log.Fatal("error loading .env files")
	}
	err = database.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	r := gin.Default()
	m := melody.New()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(config))
	//* API AUTH
	r.POST("/api/register", routes.Register)
	r.POST("/api/login", routes.Login)
	r.POST("/api/logout", routes.Logout)

	//* API AVATAR
	r.PATCH("/api/avatar", routes.ChangeAvatar)

	//* API FRIEND
	r.GET("/api/friends", routes.GetAllFriends)
	r.POST("/api/friends/request", routes.RequestAddFriend)
	r.POST("/api/friends/accept", routes.AcceptFriendRequest)
	r.POST("/api/friends/decline", routes.RejectFriendRequest)
	r.GET("/api/friends/pending", routes.GetPendingFriendRequests)

	//* API USER
	r.GET("/api/user", routes.GetUser)
	r.PATCH("/api/user/username", routes.ChangeUsername)
	r.PATCH("/api/user/about", routes.ChangeAbout)
	r.POST("/api/user/block", routes.BlockUser)
	r.DELETE("/api/user/block", routes.UnBlockUser)

	r.GET("/ws/auth", routes.GenerateWebsocketAuthCode)
	r.GET("/ws/connect", routes.ConnectWebSocket(m))

	//* API MESSAGE
	r.GET("/api/message", routes.GetMessage)

	r.StaticFS("/static", http.Dir("static"))

	m.HandleConnect(websocket.HandleConnect)
	m.HandleMessage(websocket.HandleMessage(m))

	r.Run(":5050")
}
