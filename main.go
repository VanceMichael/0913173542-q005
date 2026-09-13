package main
import "github.com/gin-gonic/gin"
func main(){r:=gin.Default();r.GET("/healthz",func(c *gin.Context){c.JSON(200,gin.H{"status":"ok"})});r.POST("/corrections",func(c *gin.Context){var v map[string]any;if c.BindJSON(&v)==nil{c.JSON(201,gin.H{"accepted":true,"data":v})}});r.Run(":8080")}
