package main

import "net/http"

func main() {
	// Initialize the database connection and defer its closure.
	InitDB()
	defer db.Close()

	// Define HTTP request handlers for various routes.
	http.HandleFunc("/", index)
	http.HandleFunc("/gallery", gallery)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/person", person)
	http.HandleFunc("/signin", signin)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Serve static files from the assets directory.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Start the HTTP server and listen on port 8080.
	http.ListenAndServe(":8080", nil)
}
