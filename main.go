package main

import(
	"fmt"
	"net/http"	
	"github.com/gorilla/mux"
	"github.com/igorvinnicius/lenslocked-go-web/controllers"
	"github.com/igorvinnicius/lenslocked-go-web/models"	
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "0000"
	dbname = "lenslocked_dev"
)


func notFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry, page not found.</h1>")
}

func main() {
	
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(psqlinfo)
	must(err)

	defer us.Close()

	us.AutoMigrate()

	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers(us)

	r := mux.NewRouter()	
	r.Handle("/", staticController.HomeView).Methods("GET")
	r.Handle("/contact", staticController.ContactView).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notFound)
	
	fmt.Println("Starting the server on :3000...")
	
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil{
		panic(err)
	}
}