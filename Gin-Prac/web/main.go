package main

import (
	"github.com/gin-gonic/gin"

	//"net/http"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("index.html")
	router.Static("/", "./" )


	//router.GET("/", getWeb)

	router.Run(":8080")
}

//func getWeb (c *gin.Context){
//	c.HTML(http.StatusOK, "index.html",nil)
//}