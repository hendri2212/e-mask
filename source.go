package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"database/sql" // untuk SQL
	// "strings" // untuk pencarian
)

var Dir_Name = "D:/CODE/GOPATH/src/e-mask/"
var BaseURL = "http://localhost:8080/"

func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/puisi/", CategoryPuisi)
    http.HandleFunc("/video/", CategoryVideo)
    http.HandleFunc("/news/", CategoryNews)
    http.HandleFunc("/gambar/", CategoryGambar)
    http.HandleFunc("/login/", Login)

    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
    fmt.Println("Starting webserver on port 8080...")
    http.ListenAndServe("127.0.0.1:8080", nil)
}

func ReadHtmlFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		panic(0)
	}
	return data
}

func RenderTemplate(w http.ResponseWriter, name string, data map[string] interface{}) {
	funcs := template.FuncMap{"Add": TemplateAdd, "ToHTML": TemplateHTML}
	t := template.New("t1")
	t, err := t.Funcs(funcs).Parse(string(ReadHtmlFile(name)))
	if err != nil {
		panic(0)
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	data["baseurl"] = BaseURL
	t.Execute(w, struct{ Data map[string]interface{} }{Data: data})
}

func TemplateAdd(x interface{}, y int) int {
	x1 := x.(int)
	return x1 + y
}

func TemplateHTML(x interface{}) template.HTML {
	x1 := x.(string)
	return template.HTML(x1)
}

// Control
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Dir_Name+"index.html", nil)
}

func CategoryPuisi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		KeySearch := r.PostFormValue("KeySearch")
		fmt.Println(KeySearch)


		data := make(map[string]interface{})
		data["puisi"] = DataSearch()
		RenderTemplate(w, Dir_Name+"CategoryPuisi.html", data)
	} else {
		data := make(map[string]interface{})

		data["puisi"] = DataPuisi()

		RenderTemplate(w, Dir_Name+"CategoryPuisi.html", data)
	}
}

func CategoryNews(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	data["news"] = DataNews()

	RenderTemplate(w, Dir_Name+"CategoryNews.html", data)
}

func CategoryVideo(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Dir_Name+"CategoryVideo.html", nil)
}

func CategoryGambar(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Dir_Name+"CategoryGambar.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		Username := r.PostFormValue("Username")
		Password := r.PostFormValue("Password")
		fmt.Println(Username, Password)


		data := make(map[string]interface{})
		data["puisi"] = DataSearch()
		if Username == "user" && Password == "user" {
			RenderTemplate(w, Dir_Name+"PageAdmin.html", nil)
		} else {
			RenderTemplate(w, Dir_Name+"Login.html", nil)
		}
	} else {
		RenderTemplate(w, Dir_Name+"Login.html", nil)
	}
}

func PageAdmin(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Dir_Name+"PageAdmin.html", nil)
}



// Connection
func Conn() *sql.DB {
	db, err := sql.Open("mysql", "root:@/e-mask")
	if err != nil {
		panic(0)
	}
	return db
}


// Model Puisi
type Puisi struct {
	Id_karya int
	Kategory string
	Judul string
	Deskripsi string
}

func DataPuisi() []Puisi {
	db := Conn()
	defer db.Close()
	rows, err := db.Query("SELECT Id_karya, Kategory, Judul, Deskripsi FROM tb_karya WHERE Kategory='Puisi'")
	if err != nil {
		panic(err.Error())
	}

	all_puisi := []Puisi{}
	for rows.Next() {
		s := Puisi{}
		err = rows.Scan(&s.Id_karya, &s.Kategory, &s.Judul, &s.Deskripsi)
		if err != nil {
			panic(err.Error())
		}
		all_puisi = append(all_puisi, s)
	}
	return all_puisi
}

func DataSearch() []Puisi {
	db := Conn()
	defer db.Close()

	// KeySearch := "Cinta"
	// rows, err := db.Query("SELECT Id_karya, Kategory, Judul, Deskripsi FROM tb_karya WHERE Judul LIKE '%?%'", KeySearch)
	rows, err := db.Query("SELECT Id_karya, Kategory, Judul, Deskripsi FROM tb_karya WHERE Judul LIKE '%cinta%'")
	if err != nil {
		panic(err.Error())
	}

	all_search := []Puisi{}
	for rows.Next() {
		s := Puisi{}
		err = rows.Scan(&s.Id_karya, &s.Kategory, &s.Judul, &s.Deskripsi)
		if err != nil {
			panic(err.Error())
		}
		all_search = append(all_search, s)
	}
	return all_search
}


// Model News
type News struct {
	Id_karya int
	Kategory string
	Judul string
	Deskripsi string
}

func DataNews() []News {
	db := Conn()
	defer db.Close()
	rows, err := db.Query("SELECT Id_karya, Kategory, Judul, Deskripsi FROM tb_karya WHERE Kategory='News'")
	if err != nil {
		panic(err.Error())
	}

	all_news := []News{}
	for rows.Next() {
		s := News{}
		err = rows.Scan(&s.Id_karya, &s.Kategory, &s.Judul, &s.Deskripsi)
		if err != nil {
			panic(err.Error())
		}
		all_news = append(all_news, s)
	}
	return all_news
}