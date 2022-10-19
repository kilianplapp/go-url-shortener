package main

import (
	"fmt"
	"log"
	"net/http"
)

var db []URL

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(db); i++ {
		if string(r.URL.Path[1:]) == db[i].bullet {
			w.Header().Set("Location", db[0].url)
			w.WriteHeader(302)
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Example URL Shortener</h1><br><a href="/create">Create URL</a>`)
}

func form(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `	
		<form method="POST">
		<input name="bullet" placeholder="bullet">
		<input name="url" placeholder="url">
		<button type="submit">Submit</button>
		</form>
		`)
	case "POST":
		r.ParseForm()
		db = append(db, URL{r.FormValue("bullet"), r.FormValue("url")})
		fmt.Fprintf(w, "URL Created!")
	}
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/create", form)
	fmt.Println("Build sucessful! Running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
