package main

import (
    "html/template"
    "Library-Management-System/database"
    "Library-Management-System/routes"
    "log"
    "net/http"
)

func main() {
    // Connect to the database
    database.Connect()

    // Set up API routes
    router := routes.SetupRoutes()

    // Serve the index.html file at the root path
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("static/index.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
    })

    log.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
