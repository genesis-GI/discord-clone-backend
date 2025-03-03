package main

import "github.com/gin-gonic/gin"

func connectionHandler(c *gin.Context){
	remote :=c.RemoteIP()
	client := c.ClientIP()
	method := c.Request.Method
	uri := c.Request.RequestURI
	protocol := c.Request.Proto
	reqBody := c.Request.Body
	reqHost := c.Request.Host	

	c.JSON(200, gin.H{
		"Request Host": reqHost,
		"URI": uri,
		"Remote IP": remote,
		"Client IP": client,
		"Request Method": method,
		"Request Protocol": protocol,
		"Request Body": reqBody,
		"URL": reqHost + uri,
	}) 
}