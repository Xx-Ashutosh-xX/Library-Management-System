package controllers

import (
    "encoding/json"
    "net/http"
    "Library-Management-System/database"
    "Library-Management-System/models"
    "github.com/gorilla/mux"
)

// GetAuthors retrieves all authors
func GetAuthors(w http.ResponseWriter, r *http.Request) {
    var authors []models.Author
    rows, err := database.DB.Query("SELECT * FROM Authors")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var author models.Author
        err := rows.Scan(&author.AuthorID, &author.Name, &author.Biography)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        authors = append(authors, author)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(authors)
}

// AddAuthor adds a new author
func AddAuthor(w http.ResponseWriter, r *http.Request) {
    var author models.Author
    if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "INSERT INTO Authors (Name, Biography) VALUES (?, ?)"
    _, err := database.DB.Exec(query, author.Name, author.Biography)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// UpdateAuthor updates author details
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    authorID := vars["author_id"]
    var author models.Author
    if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "UPDATE Authors SET Name = ?, Biography = ? WHERE Author_ID = ?"
    _, err := database.DB.Exec(query, author.Name, author.Biography, authorID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// DeleteAuthor deletes an author
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    authorID := vars["author_id"]

    query := "DELETE FROM Authors WHERE Author_ID = ?"
    _, err := database.DB.Exec(query, authorID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
