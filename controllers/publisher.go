package controllers

import (
    "encoding/json"
    "net/http"
    "Library-Management-System/database"
    "Library-Management-System/models"
    "github.com/gorilla/mux"    
)

// GetPublishers retrieves all publishers
func GetPublishers(w http.ResponseWriter, r *http.Request) {
    var publishers []models.Publisher
    rows, err := database.DB.Query("SELECT * FROM Publishers")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var publisher models.Publisher
        err := rows.Scan(&publisher.PublisherID, &publisher.Name, &publisher.Address, &publisher.ContactInformation)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        publishers = append(publishers, publisher)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(publishers)
}

// AddPublisher adds a new publisher
func AddPublisher(w http.ResponseWriter, r *http.Request) {
    var publisher models.Publisher
    if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "INSERT INTO Publishers (Name, Address, Contact_Information) VALUES (?, ?, ?)"
    _, err := database.DB.Exec(query, publisher.Name, publisher.Address, publisher.ContactInformation)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// UpdatePublisher updates publisher details
func UpdatePublisher(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    publisherID := vars["publisher_id"]
    var publisher models.Publisher
    if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "UPDATE Publishers SET Name = ?, Address = ?, Contact_Information = ? WHERE Publisher_ID = ?"
    _, err := database.DB.Exec(query, publisher.Name, publisher.Address, publisher.ContactInformation, publisherID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// DeletePublisher deletes a publisher
func DeletePublisher(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    publisherID := vars["publisher_id"]
    query := "DELETE FROM Publishers WHERE Publisher_ID = ?"
    _, err := database.DB.Exec(query, publisherID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
