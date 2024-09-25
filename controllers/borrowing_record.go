package controllers

import (
    "encoding/json"
    "net/http"
    "Library-Management-System/database"
    "Library-Management-System/models"
    "github.com/gorilla/mux"
)

// GetBorrowingRecords retrieves all borrowing records
func GetBorrowingRecords(w http.ResponseWriter, r *http.Request) {
    var records []models.BorrowingRecord
    studentID := r.URL.Query().Get("student_id")

    baseQuery := `SELECT * FROM Borrowing_Records`
    var args []interface{}

    if studentID != "" {
        baseQuery += ` WHERE Student_ID = ?`
        args = append(args, studentID)
    }

    rows, err := database.DB.Query(baseQuery, args...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var record models.BorrowingRecord
        err := rows.Scan(&record.BorrowingID, &record.StudentID, &record.BookID, &record.BorrowingDate, &record.ReturnDate, &record.Fine)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        records = append(records, record)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(records)
}

// AddBorrowingRecord adds a new borrowing record
func AddBorrowingRecord(w http.ResponseWriter, r *http.Request) {
    var record models.BorrowingRecord
    if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "INSERT INTO Borrowing_Records (Student_ID, Book_ID, Borrowing_Date, Return_Date, Fine) VALUES (?, ?, ?, ?, ?)"
    _, err := database.DB.Exec(query, record.StudentID, record.BookID, record.BorrowingDate, record.ReturnDate, record.Fine)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// UpdateBorrowingRecord updates borrowing record details
func UpdateBorrowingRecord(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    borrowingID := vars["borrowing_id"]
    var record models.BorrowingRecord
    if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "UPDATE Borrowing_Records SET Student_ID = ?, Book_ID = ?, Borrowing_Date = ?, Return_Date = ?, Fine = ? WHERE Borrowing_ID = ?"
    _, err := database.DB.Exec(query, record.StudentID, record.BookID, record.BorrowingDate, record.ReturnDate, record.Fine, borrowingID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// DeleteBorrowingRecord deletes a borrowing record
func DeleteBorrowingRecord(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    borrowingID := vars["borrowing_id"]

    query := "DELETE FROM Borrowing_Records WHERE Borrowing_ID = ?"
    _, err := database.DB.Exec(query, borrowingID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
