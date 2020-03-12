package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

var LIFF_Data struct {
	LIFFID      string
	RedirectURL string
}

func main() {
	LIFF_Data.LIFFID = os.Getenv("YOUR_LIFF_ID")
	LIFF_Data.RedirectURL = os.Getenv("YOUR_REDIRECT_URL")

	//Web APIs
	http.HandleFunc("/", liff)

	//provide by Heroku
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

//WEB: For account link
func liff(w http.ResponseWriter, r *http.Request) {
	//5. The user accesses the linking URL.
	tmpl := template.Must(template.ParseFiles("_index.html"))
	//6. The web server displays the login screen.
	if err := tmpl.Execute(w, LIFF_Data); err != nil {
		log.Println("Template err:", err)
	}
}
