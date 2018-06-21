package controller

type Karya struct {
	id    string
	judul string
	isi   string
}

// create or insert a record to table siswa
func CreateDataKarya(id string, judul string, isi string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO karya VALUES ( ?, ?, ? )")
	stmt.Exec(id, judul, isi)
	defer stmt.Close()
}

// read all data karya
func ReadDataKarya() []Karya {
	db := Conn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM karya")
	if err != nil {
		panic(err.Error())
	}

	all_karya := []Karya{}
	for rows.Next() {
		k := Karya{}
		err = rows.Scan(&k.id, &k.judul, &k.isi)
		if err != nil {
			panic(err.Error())
		}
		all_karya = append(all_karya, k)
	}
	return all_karya
}

// read only one rows
func ReadDataKaryaById(id string) Karya {
	db := Conn()
	defer db.Close()
	k := Karya{}
	db.QueryRow("SELECT id, judul, isi FROM karya WHERE id=?", id).Scan(&k.id, &k.judul, &k.isi)
	return k
}

// update recore data karya by id
func UpdateDataKaryaById(id string, judul string, isi string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE karya SET judul=?, isi=? WHERE id=?")
	stmt.Exec(judul, isi, id)
	defer stmt.Close()
}

// remove data karya by id
func DeleteDataKarya(id string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("DELETE FROM karya WHERE id=?")
	stmt.Exec(id)
	defer stmt.Close()
}
