package main
// import 4 important libraries 
// 1. "net/http" to access the core go http functionality
// 2. "fmt" for formatting text
// 3. html.template a library for interacting with html
// 4. time - for working with date and time
import (
	"net/http"
	"fmt"
	"time"
	"html/template"
)

// Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

// Go applicaiton entrypoint
func main() {
	// Instantiate a Welcome struct object and pass in random info. we get the name of the user as a query parameter from the URL
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	//instruct Go where to find our html file. have Go parse the html file (notice relative path).
	//wrap call to template.Must() which handles any errors and halts if there are fatal errors
	
	templates :=
	template.Must(template.ParseFiles("templates/welcome-template.html"))

	//Our HTML comes with CSS that go provides when we run the app.
	// tell go to creat a handle that looks in the static directory, go then uses the "/static/"
	// as a url that our html can refer to when looking for our css and other files.

	http.Handle("/static/", // final url can be anything
	http.StripPrefix("/static/",
	http.FileServer(http.Dir("static")))) //Go looks in the relative "static"
	// directory first using http.FileServer(), then matches it to a 
	// url of our choice as shown in http.Handle. with this url we reference our css files
	// Once the server begins. Our html code would be <link rel="stylesheet" href="/static/stylesheet/...">
	// The url in http.Handle can be whatever we like.

	//This method takes in the URL path "/" and a function that takes in a response writer
	// and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Takes the name from the URL query, will set welcome.Name to that name
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name;
		}
		// if errors show an internal server error message, pass welcome struct to html file
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} 
	})

	// Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	// Print any errors from starting the webserver using fmt
	fmt.Println("Listening");
	fmt.Println(http.ListenAndServe(":8080", nil));
}