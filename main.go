package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main(){
    gin.SetMode(gin.ReleaseMode)
    checkForParameters()
    rwEnv := os.Getenv("RAILWAY_ENVIRONMENT")
	isProduction := rwEnv == "production"
	isLocal := rwEnv == ""
    _=isLocal
    _=isProduction
    
     
    r := gin.Default()

    r.GET("/", func(c *gin.Context){
        c.String(200, "Hello World!")
    })

    r.GET("/connection/info", func(c *gin.Context){
        connectionHandler(c)
    })
    

    environment(rwEnv)
    success("Server running on http://127.0.0.1:8050!")
    r.Run(":8050")
}