package controller

import (
	"fmt"
	"net/http"
	// "strconv"
)

// fungsi untuk menangani halaman utama
func HomePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Dir_Name+"header.html", nil)
	RenderTemplate(w, Dir_Name+"index.html", nil)
	RenderTemplate(w, Dir_Name+"footer.html", nil)
}

// ------ FUNGSI LOGIN------ /
// fungsi untuk menangani halaman utama admin (harus login dulu)
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		User := r.PostFormValue("Username")
		Pass := r.PostFormValue("Password")
		fmt.Println(User, Pass)

		// fungsi model Cek Login
		CekLogin(w, User, Pass)
	} else {
		RenderTemplate(w, Dir_Name+"Login.html", nil)
	}
}

// halaman list data karya
func PageKaryaShowAll(w http.ResponseWriter, r *http.Request) {
	// membuat tempat penampungan data karya supaya dapat dipanggil melalui template html
	data := make(map[string]interface{})

	// mengambil data karya dari database
	data["karya"] = ReadDataKarya()

	RenderTemplate(w, Dir_Name+"header.html", nil)
	RenderTemplate(w, Dir_Name+"karya_list.html", data)
	RenderTemplate(w, Dir_Name+"footer.html", nil)
}

// halaman tambah record karya
func PageKaryaInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		Kategory := r.PostFormValue("Kategory")
		Judul := r.PostFormValue("Judul")
		Deskripsi := r.PostFormValue("Deskripsi")

		// menampilkan di cmd
		fmt.Println(Kategory, Judul, Deskripsi)

		// model CreateDataKarya()
		CreateDataKarya("", Kategory, Judul, Deskripsi)
		data := make(map[string]interface{})
		data["info"] = "Data berhasil ditambahkan ke database"

		// fungsi RenderTemplate()
		RenderTemplate(w, Dir_Name+"header.html", nil)
		RenderTemplate(w, Dir_Name+"karya_insert.html", data)
		RenderTemplate(w, Dir_Name+"footer.html", nil)
		// PageKaryaShowAll(w, r)
	} else {
		RenderTemplate(w, Dir_Name+"header.html", nil)
		RenderTemplate(w, Dir_Name+"karya_insert.html", nil)
		RenderTemplate(w, Dir_Name+"footer.html", nil)
	}
}

// Halaman edit record karya
func PageKaryaEdit(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	// ambil Id_karya dari URL
	r.ParseForm()
	Id_karya := r.FormValue("Id_karya")

	// jika POST maka jalankan condition if
	if r.Method == "POST" {
		Kategory := r.PostFormValue("Kategory")
		Judul := r.PostFormValue("Judul")
		Deskripsi := r.PostFormValue("Deskripsi")

		// menampilkan di cmd
		fmt.Println(Kategory, Judul, Deskripsi)

		// model UpdateDataSiswaById()
		UpdateDataKaryaById(Id_karya, Kategory, Judul, Deskripsi)

		// info yang dicetak di html
		data["info"] = "Data berhasil diupdate"
	}

	data["karya"] = ReadDataKaryaById(Id_karya)
	RenderTemplate(w, Dir_Name+"header.html", nil)
	RenderTemplate(w, Dir_Name+"karya_edit.html", data)
	RenderTemplate(w, Dir_Name+"footer.html", nil)
}

// halaman hapus record Karya
func PageKaryaDelete(w http.ResponseWriter, r *http.Request) {
	// ambil Id_karya dari URL
	r.ParseForm()
	Id_karya := r.FormValue("Id_karya")

	// model remove karya by Id_karya
	DeleteDataKarya(Id_karya)
	data := make(map[string]interface{})
	data["info"] = "Data dengan id =" + Id_karya + " berhasil dihapus!"
	RenderTemplate(w, Dir_Name+"header.html", nil)
	RenderTemplate(w, Dir_Name+"karya_delete.html", data)
	RenderTemplate(w, Dir_Name+"footer.html", nil)
}
