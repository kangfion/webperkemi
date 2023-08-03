package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"webperkemi/config"
	"webperkemi/entities"
	"webperkemi/models"
)

var WargaModel = models.NewWargaModel()

func WargaList(response http.ResponseWriter, request *http.Request) {
	session, _ := config.Store.Get(request, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(response, request, "/login", http.StatusSeeOther)
	} else {
		warga, _ := WargaModel.FindAll()

		log.Println("Data warga ", warga)
		fmt.Println("Data warga ", warga)

		data := map[string]interface{}{
			"warga": warga,
		}

		temp, err := template.ParseFiles("views/warga.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	}
}

func GetUpload(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}
	var warga entities.Warga

	if err != nil {
		data = map[string]interface{}{
			"title": "Upload Data Warga",
			"warga": warga,
		}
	} else {

		err := WargaModel.Find(id, &warga)
		if err != nil {
			panic(err)
		}

		//log.Println("Dapat data warga id warga ", id)
		//log.Println("Dapat", warga)

		data = map[string]interface{}{
			"title": "Upload Gambar Warga : ",
			"warga": warga,
		}
	}

	temp, _ := template.ParseFiles("views/formupload.html")
	temp.Execute(w, data)
}

// func GetForm(response http.ResponseWriter, request *http.Request) {
func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}
	var warga entities.Warga

	if err != nil {
		data = map[string]interface{}{
			"title": "Tambah Data Warga",
			"warga": warga,
		}
	} else {

		err := WargaModel.Find(id, &warga)
		if err != nil {
			panic(err)
		}

		//log.Println("Dapat data warga id warga ", id)
		//log.Println("Dapat", warga)

		data = map[string]interface{}{
			"title": "Edit Data Warga",
			"warga": warga,
		}
	}

	temp, _ := template.ParseFiles("views/wargaadd.html")
	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		r.ParseForm()
		var warga entities.Warga

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		warga.Id = id
		warga.Name = r.Form.Get("name")
		warga.Alamat = r.Form.Get("alamat")

		var data map[string]interface{}

		if err != nil {
			log.Println("Masuk sini insert")
			log.Println("Insert ", warga)
			// insert data
			err := WargaModel.Create(&warga)
			if err != nil {
				panic(err)
			}
			data = map[string]interface{}{
				"message": "Data berhasil disimpan",
				"data":    warga,
			}

			temp, err := template.ParseFiles("views/warga.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, data)
		} else {
			// mengupdate data
			err := WargaModel.Update(id, &warga)
			if err != nil {
				panic(err)
			}
			data = map[string]interface{}{
				"message": "Data berhasil diubah",
				"data":    warga,
			}

			log.Println("data dirubah di update")
			log.Println(id)
			log.Println(warga)

			temp, err := template.ParseFiles("views/index.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, data)
		}

	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	err = WargaModel.Delete(id)
	if err != nil {
		panic(err)
	}

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
