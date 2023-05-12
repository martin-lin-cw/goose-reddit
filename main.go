package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martin-lin-cw/goose-reddit/controller"
	"github.com/martin-lin-cw/goose-reddit/mysql"
)

func main() {
	server := gin.Default()

	store, err := mysql.NewStore("root:victor60628@tcp(localhost:3306)/goose-reddit?parseTime=true")
	if err != nil {
		panic(err)
	}
	threadController := controller.NewThreadController(store)

	server.GET("/thread/:id", func(ctx *gin.Context) {
		thread, err := threadController.GetThread(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"thread": thread})
	})

	server.Run(":8080")
}
