package main

import (
	"go-web/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/aaa", handler.Index)
	http.HandleFunc("/barang", handler.GetBarang)

	http.ListenAndServe(":5000", nil)
}
