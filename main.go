package main

import (
    "github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main(){
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    r.GET("/", func(c *gin.Context){
        c.String(200, "Hello World!")
    })
    color.Cyan("[INFO] Routes initialized...")

    color.Green("[SUCCESS] Server running on http://127.0.0.1:8050")
    r.Run(":8050")
}