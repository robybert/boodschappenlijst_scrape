package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"strings"

	"log"

	"net/http"

	"os"

	"regexp"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB




type Page struct {

	Title string
	Menu []string
	Body  []byte

}

type Product struct {
	ID		int64
	Naam string
	Gewicht int
	Prijs 	[]int
}

func getProductByName(name string) ([]Product, error){
	var products []Product

	rows, err := db.Query("SELECT id,product,gewicht FROM producten WHERE product = ?", name)
	if err != nil{
		return nil, fmt.Errorf("productsByName %q: %v", name, err)
	}
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var prod Product
        if err := rows.Scan(&prod.ID, &prod.Naam, &prod.Gewicht); err != nil {
            return nil, fmt.Errorf("productsByName %q: %v", name, err)
        }
        products = append(products, prod)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("productsByName %q: %v", name, err)
    }
    return products, nil
}



var defaultMenu = [3]string{"FrontPage", "test", "db"}

func NewPage(title string, body []byte) Page {
	page := Page{}
	page.Title = title
	page.Menu = defaultMenu[:]
	page.Body = body
	return page
}




func (p *Page) save() error {

	filename :="data/" + p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)

}


func loadPage(title string) (*Page, error) {

	filename := "data/" + title + ".txt"

	body, err := os.ReadFile(filename)

	if err != nil {

		return nil, err

	}
	page := NewPage(title, body)

	return &page, nil

}


func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {

		http.Redirect(w, r, "/edit/"+title, http.StatusFound)

		return

	}

	renderTemplate(w, "view", p)

}


func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {
		var b []byte
		page := NewPage(title, b)
		p = &page

	}

	renderTemplate(w, "edit", p)

}


func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return

	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}


func rootHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage("FrontPage")

	if err != nil {


		http.NotFound(w, r)

		return
	}

	renderTemplate(w, "view", p)

}

func DBHandler(w http.ResponseWriter, r *http.Request, title string) {

	fmt.Println("in DBHandler")
	p, err := loadPage("db")

	if err != nil {


		http.NotFound(w, r)

		return
	}

	renderTemplate(w, "db", p)

}

func searchHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := strings.TrimSpace(r.FormValue("body"))

	fmt.Printf("%s", body)

	products, err := getProductByName(body)

	if err != nil{
		//do err handling
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("products found %v\n", products)

}



var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html", "templates/db.html"))


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}


var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$|^/$|^/(search|db)/([a-zA-Z0-9]*)$")


func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			fmt.Println("not a valid path")

			http.NotFound(w, r)

			return

		}

		fn(w, r, m[2])

	}

}


func main() {
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
		AllowNativePasswords: true,
        DBName: "prices",
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	http.HandleFunc("/view/", makeHandler(viewHandler))

	http.HandleFunc("/edit/", makeHandler(editHandler))

	http.HandleFunc("/save/", makeHandler(saveHandler))

	http.HandleFunc("/", makeHandler(rootHandler))

	http.HandleFunc("/search/", makeHandler(searchHandler))

	http.HandleFunc("/db/", makeHandler(DBHandler))



	log.Fatal(http.ListenAndServe(":8080", nil))

}
