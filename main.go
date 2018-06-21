package main

import (
	"e-mask/controller"
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("Buktikan kamu bisa Pazrin !!!")
	// membuat file server

	dir_name := controller.Dir_Name
	fileServer := http.FileServer(http.Dir(dir_name))
	http.Handle("res/", http.StripPrefix("res/", fileServer))

	// menghendle path halaman yang akan di tampilkan
	http.HandleFunc("/", controller.HomePage)
	http.HandleFunc("/karya/", controller.PageKaryaShowAll)

	// menjalankan server dengan domain/ip dengan port tertentu
	fmt.Println("Starting webserver on port 8081...")
	http.ListenAndServe("127.0.0.1:8081", nil)
}
