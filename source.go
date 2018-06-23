package main

import (
	"fmt"
	"html/template"
	"net/http"
	"io/ioutil"
)

var Dir_Name = "D:/CODE/GOPATH/src/e-mask/"
var BaseURL = "http://localhost:8080/"

func main() {
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     var data = map[string]string{
    //         "Name":    "john wick",
    //         "Message": "have a nice day",
    //     }

    //     var t, err = template.ParseFiles("index.html")
    //     if err != nil {
    //         fmt.Println(err.Error())
    //         return
    //     }

    //     t.Execute(w, data)
    // })

    http.HandleFunc("/", Home)
    // http.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("./assets/img"))))
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

    fmt.Println("Starting webserver on port 8080...")
    http.ListenAndServe(":8080", nil)
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

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Dir_Name+"index.html", nil)
}

type Image struct {
    Title string
    URL   string
}

var images = map[string]*Image{
    "gambar":     {"The Go Gopher", "https://i1.wp.com/www.raparapa.com/wp-content/uploads/2017/12/keserasian-seni-rupa.jpg?fit=675%2C900&ssl=1"},
}