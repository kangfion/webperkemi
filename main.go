package main

import (
	"fmt"
	"net/http"
	"os"

	authcontroller "webperkemi/controllers"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("views/asset"))))

	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/wargastore", authcontroller.Store)
	http.HandleFunc("/formupload", authcontroller.GetUpload)
	//http.HandleFunc("/wargaupload", authcontroller.upload)
	http.HandleFunc("/anggotaadd", authcontroller.GetForm)
	http.HandleFunc("/anggota", authcontroller.WargaList)
	http.HandleFunc("/anggotadel", authcontroller.Delete)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	// Menggunakan direktori "uploads" sebagai tempat menyimpan gambar yang diunggah
	if _, err := os.Stat("views/asset/uploads"); os.IsNotExist(err) {
		os.Mkdir("views/asset/uploads", os.ModePerm)
	}

	fmt.Println("Server jalan di: http://localhost:3333")
	http.ListenAndServe(":3333", nil)
}
