# Student gRPC Server with PostgreSQL

## Overview
This project is a **gRPC server** built using **Go** that manages student records with full **CRUD** operations (Create, Read, Update, Delete). The data is stored in a **PostgreSQL** database.

## Features
- Add a new student
- Retrieve all students
- Get details of a specific student
- Update student information
- Delete a student
- Uses gRPC for efficient communication
- PostgreSQL for data persistence

## Technologies Used
- **Go** (Golang)
- **gRPC** (Protocol Buffers)
- **PostgreSQL**
- **pgx** (PostgreSQL driver for Go)
- **Protocol Buffers (.proto)** for defining service and messages

## Installation & Setup
### Prerequisites
Ensure you have the following installed:
- **Go (>=1.18)**
- **PostgreSQL**
- **protoc** (Protocol Buffers compiler)
- **protoc-gen-go** and **protoc-gen-go-grpc** plugins

### Steps to Run the Server
1. **Clone the repository**
   ```sh
   git clone https://github.com/al-mamun-bup/gRPC_Student_Server.git
   cd gRPC_Student_Server
   ```
2. **Install dependencies**
   ```sh
   go mod tidy
   ```
3. **Set up PostgreSQL database**
   ```sql
   CREATE DATABASE student_db;
   CREATE TABLE students (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL,
       age INT NOT NULL,
       grade TEXT NOT NULL
   );
   ```
4. **Update the database connection details** in `server/main.go`
   ```go
   const (
       host     = "localhost"
       port     = 5432
       user     = "mamun"
       password = "1234"
       dbname   = "student_db"
   )
   ```
5. **Generate gRPC code from .proto file**
   ```sh
   protoc --go_out=. --go-grpc_out=. proto/student.proto
   ```
6. **Run the gRPC server**
   ```sh
   go run server/main.go
   ```
7. **Run the client (to test the API)**
   ```sh
   go run client/main.go
   ```

## gRPC Endpoints

### 1. Add a Student
- **RPC Method:** `AddStudent`
- **Request:**
```proto
message Student {
  string id = 1;
  string name = 2;
  int32 age = 3;
  string grade = 4;
}
```
- **Response:** `Success message`

### 2. Get All Students
- **RPC Method:** `GetAllStudents`
- **Response:** Returns a list of students

### 3. Get a Student by ID
- **RPC Method:** `GetStudent`
- **Response:** Student details or "Student not found"

### 4. Update a Student
- **RPC Method:** `UpdateStudent`
- **Request:** Updated student details
- **Response:** "Student updated successfully"

### 5. Delete a Student
- **RPC Method:** `DeleteStudent`
- **Response:** "Student deleted successfully"

## Repository
[GitHub Repository](https://github.com/al-mamun-bup/gRPC_Student_Server)

## License
This project is licensed under the MIT License.

---

Feel free to contribute or report issues. Happy coding! 🚀

