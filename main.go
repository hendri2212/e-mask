package main

import (
	"e-mask-2/controller"
	"fmt"
	"net/http"
)

func main() {
	// file server
	dir_name := controller.Dir_Name + "res/"
	fileServer := http.FileServer(http.Dir(dir_name))
	http.Handle("/res/", http.StripPrefix("/res/", fileServer))

	// menghendel path halaman yang akan divisualisasikan
	http.HandleFunc("/", controller.Login)
	http.HandleFunc("/karya/", controller.PageKaryaShowAll)
	http.HandleFunc("/karya/insert", controller.PageKaryaInsert)
	http.HandleFunc("/karya/edit", controller.PageKaryaEdit)
	http.HandleFunc("/karya/delete", controller.PageKaryaDelete)

	//menjalankan server dengan domain/ip dengan port tertentu
	fmt.Println("Starting webserver on port 8080...")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
