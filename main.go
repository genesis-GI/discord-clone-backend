package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    r.GET("/", func(c *gin.Context){
        c.String(200, "Hello World!")
    })
    
    environment(gin.ReleaseMode)
    success("Server running on http://127.0.0.1:8050!")
    r.Run(":8050")
}