package routes

import (
    "github.com/gorilla/mux"
    "Library-Management-System/controllers"
)

// SetupRoutes initializes the API routes
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Book routes
    router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
    router.HandleFunc("/books", controllers.AddBook).Methods("POST")
    router.HandleFunc("/books/{isbn}", controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/books/{isbn}", controllers.DeleteBook).Methods("DELETE")

    // Author routes
    router.HandleFunc("/authors", controllers.GetAuthors).Methods("GET")
    router.HandleFunc("/authors", controllers.AddAuthor).Methods("POST")
    router.HandleFunc("/authors/{author_id}", controllers.UpdateAuthor).Methods("PUT")
    router.HandleFunc("/authors/{author_id}", controllers.DeleteAuthor).Methods("DELETE")

    // Publisher routes
    router.HandleFunc("/publishers", controllers.GetPublishers).Methods("GET")
    router.HandleFunc("/publishers", controllers.AddPublisher).Methods("POST")
    router.HandleFunc("/publishers/{publisher_id}", controllers.UpdatePublisher).Methods("PUT")
    router.HandleFunc("/publishers/{publisher_id}", controllers.DeletePublisher).Methods("DELETE")

    // Student routes
    router.HandleFunc("/students", controllers.GetStudents).Methods("GET")
    router.HandleFunc("/students", controllers.AddStudent).Methods("POST")
    router.HandleFunc("/students/{student_id}", controllers.UpdateStudent).Methods("PUT")
    router.HandleFunc("/students/{student_id}", controllers.DeleteStudent).Methods("DELETE")

    // Borrowing record routes
    router.HandleFunc("/borrowing_records", controllers.GetBorrowingRecords).Methods("GET")
    router.HandleFunc("/borrowing_records", controllers.AddBorrowingRecord).Methods("POST")
    router.HandleFunc("/borrowing_records/{borrowing_id}", controllers.UpdateBorrowingRecord).Methods("PUT")
    router.HandleFunc("/borrowing_records/{borrowing_id}", controllers.DeleteBorrowingRecord).Methods("DELETE")
    
    return router
}
