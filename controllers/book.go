package controllers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "Library-Management-System/database"
    "Library-Management-System/models"
    "github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
    var books []models.Book

    // Retrieve the search query from URL parameters
    searchQuery := r.URL.Query().Get("search")

    // Base query to get books along with publisher information
    baseQuery := `
        SELECT b.ISBN, b.Title, b.Publication_Date, b.Genre, 
               p.Name AS Publisher_Name, p.Address AS Publisher_Address, 
               p.Contact_Information AS Publisher_Contact, 
               a.Name AS Author_Name
        FROM Books b
        JOIN Publishers p ON b.Publisher_ID = p.Publisher_ID
        LEFT JOIN Books_Authors ba ON b.ISBN = ba.Book_ID
        LEFT JOIN Authors a ON ba.Author_ID = a.Author_ID
    `

    // Add the search condition if a search query is provided
    var args []interface{}
    if searchQuery != "" {
        searchQuery = "%" + searchQuery + "%"
        baseQuery += ` WHERE b.Title LIKE ? OR a.Name LIKE ? OR p.Name LIKE ?`
        args = append(args, searchQuery, searchQuery, searchQuery)
    }

    // Execute the query
    rows, err := database.DB.Query(baseQuery, args...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Temporary map to hold book details and associated authors
    bookMap := make(map[string]*models.Book)

    // Scan the results
    for rows.Next() {
        var (
            isbn               string
            title              string
            publicationDate    string
            genre              string
            publisherName      string
            publisherAddress   string
            publisherContact   string
            authorName         string
        )

        err := rows.Scan(&isbn, &title, &publicationDate, &genre,
            &publisherName, &publisherAddress, &publisherContact, &authorName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Check if the book already exists in the map
        book, exists := bookMap[isbn]
        if !exists {
            book = &models.Book{
                ISBN:             isbn,
                Title:            title,
                PublicationDate:  publicationDate,
                Genre:            genre,
                PublisherName:    publisherName,
                PublisherAddress: publisherAddress,
                PublisherContact: publisherContact,
                Authors:          []models.Author{},
            }
            bookMap[isbn] = book
        }

        // Append the author if not empty
        if authorName != "" {
            book.Authors = append(book.Authors, models.Author{Name: authorName})
        }
    }

    // Convert map to slice
    for _, book := range bookMap {
        books = append(books, *book)
    }

    // Set the response header and encode the books as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}


// AddBook adds a new book along with its author(s) and publisher if they are not present
func AddBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Add or retrieve the publisher
    var publisherID int64
    err := database.DB.QueryRow("SELECT Publisher_ID FROM Publishers WHERE Name = ?", book.PublisherName).Scan(&publisherID)
    if err != nil {
        if err == sql.ErrNoRows { // Publisher doesn't exist, insert new publisher
            result, err := database.DB.Exec("INSERT INTO Publishers (Name, Address, Contact_Information) VALUES (?, ?, ?)", 
                book.PublisherName, book.PublisherAddress, book.PublisherContact)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Get the newly inserted Publisher_ID
            publisherID, err = result.LastInsertId()
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    // Add the book to the database
    query := "INSERT INTO Books (ISBN, Title, Publication_Date, Genre, Publisher_ID) VALUES (?, ?, ?, ?, ?)"
    _, err = database.DB.Exec(query, book.ISBN, book.Title, book.PublicationDate, book.Genre, publisherID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Add authors if they are not already present
    for _, author := range book.Authors {
        var authorID int64
        err = database.DB.QueryRow("SELECT Author_ID FROM Authors WHERE Name = ?", author.Name).Scan(&authorID)
        
        if err != nil {
            if err == sql.ErrNoRows {
                // Insert the author if it doesn't exist
                result, err := database.DB.Exec("INSERT INTO Authors (Name, Biography) VALUES (?, ?)", 
                    author.Name, author.Biography)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                // Get the newly inserted Author_ID
                authorID, err = result.LastInsertId()
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
            } else {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        }

        // Link the book with its authors
        _, err = database.DB.Exec("INSERT INTO Books_Authors (Book_ID, Author_ID) VALUES (?, ?)", 
            book.ISBN, authorID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    // Respond with a success status
    w.WriteHeader(http.StatusCreated)
}




// UpdateBook updates book details
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    isbn := vars["isbn"]

    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Check if the publisher exists; if not, insert it
    var publisherID int64
    err := database.DB.QueryRow("SELECT Publisher_ID FROM Publishers WHERE Name = ?", book.PublisherName).Scan(&publisherID)
    if err != nil {
        if err == sql.ErrNoRows { // Publisher doesn't exist, insert new publisher
            result, err := database.DB.Exec("INSERT INTO Publishers (Name, Address, Contact_Information) VALUES (?, ?, ?)", 
            book.PublisherName, book.PublisherAddress, book.PublisherContact)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Get the newly inserted Publisher_ID
            publisherID, err = result.LastInsertId()
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    // Update the book with the new details and publisher ID
    query := "UPDATE Books SET Title = ?, Publication_Date = ?, Genre = ?, Publisher_ID = ? WHERE ISBN = ?"
    _, err = database.DB.Exec(query, book.Title, book.PublicationDate, book.Genre, publisherID, isbn)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}


// DeleteBook deletes a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    isbn := vars["isbn"]

    query1 := "DELETE FROM Books_Authors WHERE book_id = ?"
    _, err1 := database.DB.Exec(query1, isbn)
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }

    query2 := "DELETE FROM Books WHERE ISBN = ?"
    _, err2 := database.DB.Exec(query2, isbn)
    if err2 != nil {
        http.Error(w, err2.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
