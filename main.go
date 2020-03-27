package main

import(
	"fmt"
	"net/http"	
	"github.com/gorilla/mux"
	"github.com/igorvinnicius/lenslocked-go-web/views"
)

var(
	homeView *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Frequently Asked Questions</h1><p>Here is a list of questions our users commonly ask.</p>")
}

func notFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry, page not found.</h1>")
}

func main() {
	
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()	
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", contact)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", r)
}