package main

import (
	"net/http"
)

func main() {
	// простейший сервер, которому доступны все файлы в поддиректории static
	// Handler that server http requests with the conent of the fs
	// fs := http.FileServer(http.Dir("./"))
	// http.Handle("/golang/", http.StripPrefix("/golang/", fs))
	http.HandleFunc("/golang/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./main.go")
	})
	http.Handle("/", http.RedirectHandler("/golang", http.StatusMovedPermanently))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// stop ordinary control flow
		// invoke defer function
		// continue up the stack
		// program crashes
		panic(err)
	}
}
