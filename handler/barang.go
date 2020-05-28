package handler

import (
	"fmt"
	"net/http"
)

func GetBarang(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ini adalah halaman web barang")
}
