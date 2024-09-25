package controllers

import (
    "encoding/json"
    "net/http"
    "Library-Management-System/database"
    "Library-Management-System/models"
    "github.com/gorilla/mux"
)

// GetStudents retrieves all students
func GetStudents(w http.ResponseWriter, r *http.Request) {
    var students []models.Student
    rows, err := database.DB.Query("SELECT * FROM Students")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var student models.Student
        err := rows.Scan(&student.StudentID, &student.Name, &student.Email, &student.Address)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        students = append(students, student)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(students)
}

// AddStudent adds a new student
func AddStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student
    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "INSERT INTO Students (Name, Email, Address) VALUES (?, ?, ?)"
    _, err := database.DB.Exec(query, student.Name, student.Email, student.Address)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// UpdateStudent updates student details
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    studentID := vars["student_id"]
    var student models.Student
    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "UPDATE Students SET Name = ?, Email = ?, Address = ?, WHERE Student_ID = ?"
    _, err := database.DB.Exec(query, student.Name, student.Email, student.Address, studentID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// DeleteStudent deletes a student
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    studentID := vars["student_id"]

    query := "DELETE FROM Students WHERE Student_ID = ?"
    _, err := database.DB.Exec(query, studentID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
