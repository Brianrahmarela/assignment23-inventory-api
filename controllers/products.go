package controllers

import (
	"encoding/json" // encoding/json untuk parsing JSON.
	"fmt"
	"io"       // io untuk membaca body dari response.
	"log"      // log untuk log error.
	"net/http" // net/http untuk melakukan HTTP request ke API eksternal.

	"github.com/gin-gonic/gin" // gin agar bisa kirim response ke client.
)

// *gin.Context: Pointer ke objek Context.
// gin.Context: Objek yang menyimpan semua informasi dari HTTP request (misal: body, header, params),
// dan bisa mengatur response.
// Context ini adalah jembatan komunikasi antara client dan server yg bisa:
// Mengambil data request, Mengirim response, Mengatur status code, Mengatur headers, Menangani middleware, dll.
func GetExternalProducts(c *gin.Context) {
	fmt.Println("Memanggil GetExternalProducts")
	fmt.Println("c:", c)
	resp, err := http.Get("https://dummyjson.com/products")
	if err != nil {
		log.Println("Gagal mengambil data:", err)
		//c.JSON: Fungsi untuk mengirim response berformat JSON ke client.
		//http.StatusInternalServerError: Status code 500 (artinya server error).
		//gin.H{...}: Alias untuk map[string]interface{}, yaitu JSON response-nya:
		// { "error": "Gagal mengambil data eksternal" }
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data eksternal"})
		return
	}
	// 	defer dijalankan di akhir fungsi utama sebelum fungsi keluar.
	// resp.Body.Close(): Menutup response body agar koneksi tidak bocor/macet.
	// Tujuannya: setelah kita selesai membaca isi resp.Body, kita wajib menutupnya agar resource sistem tidak terbuang sia-sia.
	defer resp.Body.Close()
	// io.ReadAll(...): Membaca semua isi body dari response ke dalam variabel body (tipe []byte atau array byte).
	// resp.Body: Objek berisi body dari response HTTP.
	// body: Isi dari response dalam bentuk data mentah (bytes).
	// err: Kalau ada error saat membaca.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Gagal membaca response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca response"})
		return
	}
	// var data map[string]interface{}:
	// Membuat variabel data bertipe map dari string ke interface{}.
	// Karena kita belum tahu bentuk JSON-nya, kita pakai interface{} (tipe bebas).
	var data map[string]interface{}
	// json.Unmarshal(body, &data):
	// Mengubah (unmarshal) JSON dalam body []byte JSON ke bentuk map.
	// &data: Kita kirim alamat dari data, karena fungsi Unmarshal mengisi nilai ke variabel tersebut.
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Gagal parsing JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal parsing JSON"})
		return
	}
	//c.JSON: Mengirim response ke client dengan format JSON.
	// http.StatusOK: Status code 200, artinya sukses.
	// data: JSON hasil dari dummyjson tadi, sekarang dikirim kembali ke client.
	c.JSON(http.StatusOK, data)
}
