package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "student_grpc/proto"

	_ "github.com/jackc/pgx/v5/stdlib"

	"google.golang.org/grpc"
)

// StudentServiceServer struct with database connection
type StudentServiceServer struct {
	pb.UnimplementedStudentServiceServer
	db *sql.DB
}

// Connect to PostgreSQL database
func connectDB() (*sql.DB, error) {
	connStr := "postgres://mamun:1234@localhost:5432/student_db?sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// AddStudent stores student data in the database
func (s *StudentServiceServer) AddStudent(ctx context.Context, student *pb.Student) (*pb.Response, error) {
	_, err := s.db.Exec("INSERT INTO students (id, name, age, class) VALUES ($1, $2, $3, $4)",
		student.Id, student.Name, student.Age, student.Class)
	if err != nil {
		return &pb.Response{Message: "Failed to add student", Success: false}, err
	}
	return &pb.Response{Message: "Student added successfully", Success: true}, nil
}

// GetStudent retrieves student details
func (s *StudentServiceServer) GetStudent(ctx context.Context, id *pb.StudentID) (*pb.Student, error) {
	var student pb.Student
	err := s.db.QueryRow("SELECT id, name, age, class FROM students WHERE id=$1", id.Id).
		Scan(&student.Id, &student.Name, &student.Age, &student.Class)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// UpdateStudent modifies student data
func (s *StudentServiceServer) UpdateStudent(ctx context.Context, student *pb.Student) (*pb.Response, error) {
	_, err := s.db.Exec("UPDATE students SET name=$1, age=$2, class=$3 WHERE id=$4",
		student.Name, student.Age, student.Class, student.Id)
	if err != nil {
		return &pb.Response{Message: "Failed to update student", Success: false}, err
	}
	return &pb.Response{Message: "Student updated successfully", Success: true}, nil
}

// DeleteStudent removes a student record
func (s *StudentServiceServer) DeleteStudent(ctx context.Context, id *pb.StudentID) (*pb.Response, error) {
	_, err := s.db.Exec("DELETE FROM students WHERE id=$1", id.Id)
	if err != nil {
		return &pb.Response{Message: "Failed to delete student", Success: false}, err
	}
	return &pb.Response{Message: "Student deleted successfully", Success: true}, nil
}

func main() {
	// Connect to database
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create gRPC server
	server := grpc.NewServer()
	pb.RegisterStudentServiceServer(server, &StudentServiceServer{db: db})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	fmt.Println("Starting gRPC server on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
