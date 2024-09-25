package models

// Book represents the structure of a book record in the database
type Book struct {
    ISBN             string   `json:"isbn"`
    Title            string   `json:"title"`
    PublicationDate  string   `json:"publication_date"`
    Genre            string   `json:"genre"`
    PublisherName    string   `json:"publisher_name"`     
    PublisherAddress string   `json:"publisher_address"`  
    PublisherContact string   `json:"publisher_contact"`  
    Authors          []Author `json:"authors"`            
}


// Author represents the structure of an author record in the database
type Author struct {
    AuthorID   string `json:"author_id"`
    Name       string `json:"name"`
    Biography  string `json:"biography"`
}

// Publisher represents the structure of a publisher record in the database
type Publisher struct {
    PublisherID         string `json:"publisher_id"`
    Name                string `json:"name"`
    Address             string `json:"address"`
    ContactInformation  string `json:"contact_information"`
}

// Student represents the structure of a student record in the database
type Student struct {
    StudentID      string `json:"student_id"`
    Name           string `json:"name"`
    Email          string `json:"email"`
    Address        string `json:"address"`
}

// BorrowingRecord represents the structure of a borrowing record in the database
type BorrowingRecord struct {
    BorrowingID    string `json:"borrowing_id"`
    StudentID      string `json:"student_id"`
    BookID         string `json:"book_id"`
    BorrowingDate  string `json:"borrowing_date"` 
    ReturnDate     string `json:"return_date"`    
    Fine            float64 `json:"fine"`
}
