package controller

import "net/http"

type Karya struct {
	Id_karya  string
	Kategory  string
	Judul     string
	Deskripsi string
	// Gambar    string
}

/**
*untuk menambahkan data karya harus mengisikan semua field(termasuk gambar)
*Id_karya disini dianggap string karena Auto Increment maka data yang dimasukkan dari aplikasi harus kosong Id_karya = ""
*apabila tidak ada gambar maka data kosong/ Gambar = ""
 */
// create or insert a record to table tb_karya
func CreateDataKarya(Id_karya string, Kategory string, Judul string, Deskripsi string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO tb_karya(Id_karya, Kategory, Judul, Deskripsi) VALUES (?, ?, ?, ? )")
	stmt.Exec(Id_karya, Kategory, Judul, Deskripsi)
	defer stmt.Close()
}

// read all data karya
func ReadDataKarya() []Karya {
	db := Conn()
	defer db.Close()
	rows, err := db.Query("SELECT Id_karya, Kategory, Judul, Deskripsi FROM tb_karya  ORDER BY Id_karya DESC")
	if err != nil {
		panic(err.Error())
	}

	all_karya := []Karya{}
	for rows.Next() {
		s := Karya{}
		err = rows.Scan(&s.Id_karya, &s.Kategory, &s.Judul, &s.Deskripsi)
		if err != nil {
			panic(err.Error())
		}
		all_karya = append(all_karya, s)
	}
	return all_karya
}

// read only one row
func ReadDataKaryaById(Id_karya string) Karya {
	db := Conn()
	defer db.Close()
	s := Karya{}
	db.QueryRow("SELECT Id_karya, Kategory, Judul, Deskripsi FROM tb_karya WHERE Id_karya=?", Id_karya).Scan(&s.Id_karya, &s.Kategory, &s.Judul, &s.Deskripsi)
	return s
}

// upadate record karya by id
func UpdateDataKaryaById(Id_karya string, Kategory string, Judul string, Deskripsi string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE tb_karya SET Kategory=?, Judul=?, Deskripsi=? WHERE Id_karya=?")
	stmt.Exec(Kategory, Judul, Deskripsi, Id_karya)
	defer stmt.Close()
}

// remove data siswa by id
func DeleteDataKarya(Id_karya string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("DELETE FROM tb_karya WHERE Id_karya=?")
	stmt.Exec(Id_karya)
	defer stmt.Close()
}

// variabel tabel tb_user untuk cek login
type Userlogin struct {
	Username string
	Password string
}

// fungsi cek login
func CekLogin(w http.ResponseWriter, Uname string, Upass string) Userlogin {
	db := Conn()
	defer db.Close()

	s := Userlogin{}
	db.QueryRow("SELECT Username, Password FROM tb_user WHERE Username=? AND Password=?", Uname, Upass).Scan(&s.Username, &s.Password)
	users := s.Username
	pass := s.Password

	if Uname == users && Upass == pass {
		// fmt.Println("benar")
		// membuat tempat penampungan data karya supaya dapat dipanggil melalui template html
		data := make(map[string]interface{})
		// mengambil data karya dari database
		data["karya"] = ReadDataKarya()

		RenderTemplate(w, Dir_Name+"header.html", nil)
		RenderTemplate(w, Dir_Name+"karya_list.html", data)
		RenderTemplate(w, Dir_Name+"footer.html", nil)
	} else {
		// fmt.Println("salah")
		data := make(map[string]interface{})
		data["salah"] = "Maaf username atau password anda salah"

		RenderTemplate(w, Dir_Name+"Login.html", data)
	}
	return s
}
