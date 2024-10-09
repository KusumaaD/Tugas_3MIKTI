package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

// Database initialization
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlTable := `
	CREATE TABLE IF NOT EXISTS students(
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		phone_number TEXT,
		address TEXT,
		gpa REAL,
		is_graduate BOOLEAN
	);`

	_, err = db.Exec(sqlTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	// Initialize database
	db := InitDB()
	defer db.Close()

	// Initialize repository
	studentRepo := NewStudentRepository(db)

	// Initialize handler
	studentHandler := NewStudentHandler(studentRepo)

	// Initialize Echo
	e := echo.New()

	// Routes
	students := e.Group("/students")
	students.GET("", studentHandler.GetStudents)
	students.GET("/:id", studentHandler.GetStudent)
	students.POST("", studentHandler.CreateStudent)
	students.PUT("/:id", studentHandler.UpdateStudent)
	students.DELETE("/:id", studentHandler.DeleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}