package main

import (
	"database/sql"

	"github.com/google/uuid"
)

// Repository
type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Create(student Student) error {
	query := `INSERT INTO students (id, name, email, phone_number, address, gpa, is_graduate) 
             VALUES (?, ?, ?, ?, ?, ?, ?)`
	
	_, err := r.db.Exec(query, student.ID, student.Name, student.Email, 
		student.PhoneNumber, student.Address, student.GPA, student.IsGraduate)
	return err
}

func (r *StudentRepository) GetAll() ([]Student, error) {
	query := `SELECT id, name, email, phone_number, address, gpa, is_graduate FROM students`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		var id string
		err := rows.Scan(&id, &student.Name, &student.Email, 
			&student.PhoneNumber, &student.Address, &student.GPA, &student.IsGraduate)
		if err != nil {
			return nil, err
		}
		student.ID, _ = uuid.Parse(id)
		students = append(students, student)
	}
	return students, nil
}

func (r *StudentRepository) GetByID(id string) (Student, error) {
	query := `SELECT id, name, email, phone_number, address, gpa, is_graduate 
             FROM students WHERE id = ?`
	
	var student Student
	var dbID string
	err := r.db.QueryRow(query, id).Scan(&dbID, &student.Name, &student.Email, 
		&student.PhoneNumber, &student.Address, &student.GPA, &student.IsGraduate)
	if err != nil {
		return Student{}, err
	}
	student.ID, _ = uuid.Parse(dbID)
	return student, nil
}

func (r *StudentRepository) Update(id string, student Student) error {
	query := `UPDATE students SET name=?, email=?, phone_number=?, address=?, 
             gpa=?, is_graduate=? WHERE id=?`
	
	_, err := r.db.Exec(query, student.Name, student.Email, student.PhoneNumber, 
		student.Address, student.GPA, student.IsGraduate, id)
	return err
}

func (r *StudentRepository) Delete(id string) error {
	query := `DELETE FROM students WHERE id=?`
	_, err := r.db.Exec(query, id)
	return err
}