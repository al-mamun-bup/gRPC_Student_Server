package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "student_grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Add a new student
	addResp, err := client.AddStudent(ctx, &pb.Student{Id: "2", Name: "Mamun", Age: 23, Class: "CS101"})
	if err != nil {
		log.Fatalf("Could not add student: %v", err)
	}
	fmt.Printf("AddStudent Response: %v\n", addResp)

	// Get the student
	getResp, err := client.GetStudent(ctx, &pb.StudentID{Id: "1"})
	if err != nil {
		log.Fatalf("Could not get student: %v", err)
	}
	fmt.Printf("GetStudent Response: %v\n", getResp)

	// Update the student
	updateResp, err := client.UpdateStudent(ctx, &pb.Student{Id: "1", Name: "Efaz", Age: 24, Class: "CS102"})
	if err != nil {
		log.Fatalf("Could not update student: %v", err)
	}
	fmt.Printf("UpdateStudent Response: %v\n", updateResp)

	// Delete the student
	deleteResp, err := client.DeleteStudent(ctx, &pb.StudentID{Id: "1"})
	if err != nil {
		log.Fatalf("Could not delete student: %v", err)
	}
	fmt.Printf("DeleteStudent Response: %v\n", deleteResp)
}
