package main

import (
	"log"
	"os"

	"github.com/abhishekshree/cdn/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CORS())

	initFolder()

	fs := r.Group("/api/")
	{
		fs.POST("/upload", handlers.UploadFileHandler)
		fs.DELETE("/delete", handlers.HelloHandler)
		fs.GET("/view", handlers.ViewFileHandler)
		fs.GET("/view/all", handlers.ViewAllHandler)
		fs.POST("/zip", handlers.HelloHandler)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	f, err := os.OpenFile("cdn.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(f)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func initFolder() {
	x := "cdn"
	if _, err := os.Stat(x); os.IsNotExist(err) {
		os.Mkdir(x, 0755)
	}
}
