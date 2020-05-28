package handler

import (
	"fmt"
	"go-web/model"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Welcome to my page")
	// fmt.Fprintln(w, "This is my first page")

	t, err := template.ParseFiles("./html/template_index.html")
	if err != nil {
		fmt.Println(err)
	}

	cus := model.Customer{"Martha", 10}

	t.Execute(w, cus)
}
