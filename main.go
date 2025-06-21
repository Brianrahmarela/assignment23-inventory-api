package main

import (
	"fmt"
	"inventory-api/routes"

	// "os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Membuat instance router  dari alamat objek *gin.Engine. gin.Default() menyediakan:
	// Logger (untuk mencetak log ke terminal)
	// Recovery (agar server tidak crash saat panic)
	fmt.Println("tes!")
	router := gin.Default()
	// fmt.Println("router:", router)
	// log.Println("log router: ", router)

	// Memanggil fungsi SetupRoutes dari file routes/api.go, dan mengoper router agar route bisa disusun
	// dari luar file main.go.
	routes.SetupRoutes(router)

	router.Run(":8080")
}
