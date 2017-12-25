package main

import (
	"html/template"
	"fmt"
	"net/http"

)
const (
	Port = ":8080"
)
func main() {
		http.HandleFunc("/",serveStatic)
		http.ListenAndServe(Port, nil)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("test.html")
	if err != nil {
		fmt.Println(err)
	}
	data := struct {
		Title string
		Header string
		ArCode string
		ArName string
		ArAddress string
		Items []string
	}{
		Title: "GO HMTL Template",
		Header:"Notice for Billing Payment",
		ArCode:"41054",
		ArName:"นายสาธิต โฉมวัฒนา",
		ArAddress:"666666",
		Items: []string{
			"My photos",
			"My blog",
		},
	}
	t.Execute(w, data)

}