package controllers

import (
	"bytes"
	"encoding/json" // encoding/json untuk parsing JSON.
	"fmt"

	// "fmt"
	// bytes untuk membuat buffer dari jsonData.
	"io" // io untuk membaca body dari response.

	// "log"      // log untuk log error.
	"net/http" // net/http untuk melakukan HTTP request ke API eksternal.

	"github.com/gin-gonic/gin" // gin agar bisa kirim response ke client.
)

func GetExternalComments(c *gin.Context) {
	fmt.Println("Memanggil GetExternalComments")
	resp, err := http.Get("https://dummyjson.com/comments")
	fmt.Println("=== HTTP Response ===")
	fmt.Println("Status      :", resp.Status)
	fmt.Println("StatusCode  :", resp.StatusCode)
	fmt.Println("Proto       :", resp.Proto)
	fmt.Println("ContentLength:", resp.ContentLength)
	fmt.Println("Uncompressed:", resp.Uncompressed)
	fmt.Println("Request     :", resp.Request)
	fmt.Println("TLS         :", resp.TLS)
	fmt.Println("\n=== Headers ===")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%-25s: %s\n", key, value)
		}
	}
	fmt.Println("=======================")

	fmt.Println("err: ", err)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	// fmt.Println("body: ", string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal parsing JSON"})
		return
	}

	c.JSON(http.StatusOK, result)
}

type Comment struct {
	Body   string `json:"body"`
	PostId int    `json:"postId"`
	UserId int    `json:"userId"`
}

func PostExternalComment(c *gin.Context) {
	fmt.Println("Memanggil PostExternalComment")
	var newComment Comment

	// 	c -> context dari Gin yang berisi permintaan HTTP.
	// BindJSON(&newComment) -> ambil data JSON dari HTTP request body -> konversi ke bentuk struct Go.
	// dan mengupdate nilainya ke var newComment karena sudah diberi alamatnya dgn &
	// err := →r eturn-nya err, jika gagal parsing JSON.
	if err := c.BindJSON(&newComment); err != nil {
		// c.JSON(...) -> kirim response JSON ke client.
		// http.StatusBadRequest-> status code 400.
		// gin.H{"error": ...} -> gin.H adalah alias map[string]interface{}, dikirim sebagai isi body JSON.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}
	//fungsi bawaan Go (encoding/json) untuk mengubah struct jadi JSON string yang terformat (bukan satu baris).
	jsonPretty, _ := json.MarshalIndent(newComment, "", "  ")
	// fmt.Println(string(jsonPretty))
	fmt.Println("newComment", string(jsonPretty))

	//Mengubah struct newComment menjadi JSON string (dalam bentuk []byte, tapi isinya JSON).
	// json.Marshal(...) -> fungsi standar Go untuk mengubah data ke JSON satu baris (tanpa indentasi).
	jsonData, _ := json.Marshal(newComment)
	fmt.Println("jsonData", jsonData)
	fmt.Println("jsonData convert to string", string(jsonData))

	resp, err := http.Post("https://dummyjson.com/comments/add", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal kirim ke API eksternal"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal parsing hasil response"})
		return
	}

	c.JSON(resp.StatusCode, result)
}
