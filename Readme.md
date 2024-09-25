# Library Management System

This project is a web-based application designed to manage a library's books, authors, publishers, students, and borrowing records. It provides a user-friendly interface for searching, adding, and deleting records.

## Features

- Search for books by title, author, or genre.
- Add new books, authors, publishers, and students.
- Add and manage borrowing records.
- Delete records for books, authors, publishers, students, and borrowing records.

## Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.15 or higher).
- [MySQL](https://dev.mysql.com/downloads/mysql/) installed and running.

## Steps to Install

1. **Set Up Go Modules**
   Initialize your Go module (if not already done):
   ```bash
   go mod init Library-Management-System
   ```

2. **Install Dependencies**
   Use the following command to fetch necessary packages:
   ```bash
   go get
   ```

## Database Setup


1. **Create Database and User**
   - Log in to your MySQL server:
     ```bash
     mysql -u root -p
     ```
   - Create the database:
     ```sql
     CREATE DATABASE library;
     ```
   - (Optional) Create a dedicated user and grant permissions:
     ```sql
     CREATE USER 'library_user'@'localhost' IDENTIFIED BY 'yourpassword';
     GRANT ALL PRIVILEGES ON library.* TO 'library_user'@'localhost';
     FLUSH PRIVILEGES;
     ```

2. **Update Database Connection**
   - Open `database.go` and update the connection string with your MySQL credentials:
   ```go
   connStr := "root:12345678@tcp(localhost:3306)/library" // Adjust username, password, and database name
   ```

## Starting the Server

1. **Run the Application**
   - Start the server using:
   ```bash
   go run main.go
   ```
   - The application should now be running on `http://localhost:8080`.