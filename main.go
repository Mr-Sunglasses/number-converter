package main

import (
	"github.com/Bhupesh-V/roman-to-no-api/controllers"
	"github.com/gin-gonic/gin"
)


func main() {
    setupServer().Run()
}

// The engine with all endpoints is now extracted from the main function
func setupServer() *gin.Engine {
    r := gin.Default()

	r.POST("/roman", controllers.ConvertRoman)

    return r
}

// func main() {
// r := gin.Default()
// r.POST("/roman", controllers.ConvertRoman)
// r.Run()
// }
