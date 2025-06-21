// package routes

// import (
// 	"inventory-api/controllers"

// 	"github.com/gin-gonic/gin"
// )

// // Fungsi yg bertugas untuk menyusun rute-rute (endpoint) di aplikasi.
// // *gin.Engine: Tipe parameternya gin engine, artinya var router merupakan pointer ke objek gin.Default() milik framework Gin.
// // gin.Engine adalah "mesin utama" Gin yang menangani HTTP request/response, routing, middleware, dll.
// // * (pointer) artinya kamu mengirim alamat memori objek dari instant main (bukan salinannya)
// // supaya semua route langsung terpasang ke web server Gin yang sama.
// func SetupRoutes(router *gin.Engine) {
// 	// Endpoint ke controller
// 	router.GET("/external-products", controllers.GetExternalProducts)
// }

package routes

import (
	"inventory-api/controllers"

	"github.com/gin-gonic/gin"
)

// Struktur Quote
// type Quote struct {
// 	ID     int    `json:"id,omitempty"`
// 	Quote  string `json:"quote"`
// 	Author string `json:"author"`
// }

// Daftarkan semua route
func SetupRoutes(router *gin.Engine) {
	router.GET("/external-comments", controllers.GetExternalComments)
	router.POST("/external-comments/add", controllers.PostExternalComment)
	// router.POST("/external-quotes", PostExternalQuote)
	// router.DELETE("/external-quotes/:id", DeleteExternalQuote)
}

// Kirim quote ke dummyjson
// func PostExternalQuote(c *gin.Context) {
// 	var newQuote Quote

// 	// Ambil dari body request
// 	if err := c.BindJSON(&newQuote); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
// 		return
// 	}

// 	jsonData, _ := json.Marshal(newQuote)

// 	resp, err := http.Post("https://dummyjson.com/quotes/add", "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal kirim ke API eksternal"})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)

// 	var result map[string]interface{}
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal parsing hasil response"})
// 		return
// 	}

// 	c.JSON(resp.StatusCode, result)
// }

// // Hapus quote (simulasi)
// func DeleteExternalQuote(c *gin.Context) {
// 	id := c.Param("id")
// 	url := "https://dummyjson.com/quotes/" + id

// 	req, _ := http.NewRequest(http.MethodDelete, url, nil)
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengirim DELETE"})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)

// 	var result map[string]interface{}
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal parsing response"})
// 		return
// 	}

// 	c.JSON(resp.StatusCode, result)
// }
