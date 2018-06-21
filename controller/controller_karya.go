package controller

import (
	//"fmt"
	// "strconv"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Dir_Name+"index.html", nil)
}

func PageKaryaShowAll(w http.ResponseWriter, r *http.Request) {
	//membuat penampung data karya agar dapat dipanggil melalui template html
	data := make(map[string]interface{})

	//mengambil data karya
	data["karya"] = ReadDataKarya()

	// RenderTemplate(w, Dir_Name+"header.html", nil)
	RenderTemplate(w, Dir_Name+"karya_list.html", data)
	// RenderTemplate(w, Dir_Name+"footer.html", nil)
}
