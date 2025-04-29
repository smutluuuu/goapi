# School Management System API

[![Go](https://img.shields.io/badge/Go-v1.23.2-blue?logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/smutluuuu/goapi)](https://github.com/smutluuuu/goapi/stargazers)

This is a RESTful API for a School Management System, built using Go. It provides functionalities for managing executives, students, and teachers.

## Features

- **Executive Management:** Create, retrieve, update, and delete executive user accounts. Includes login, logout, forgot password, and reset password functionalities.
- **Student Management:** Add, retrieve, update, and delete student records. Supports batch operations for adding and deleting students.
- **Teacher Management:** Add, retrieve, update, and delete teacher records. Supports batch operations for adding and deleting teachers.
- **Relationships:** Retrieve students associated with a specific teacher and get the count of students for a teacher.
- **Authentication:** JWT-based authentication for securing executive user access.

## Technologies

- **Go:** v1.23.2
- **Database:** MySQL
- **Routing:** `net/http` (standard Go library)
- **Go Modules:**
  - `github.com/go-sql-driver/mysql`: For MySQL database interaction.
  - `github.com/joho/godotenv`: For loading environment variables from a `.env` file.
  - `github.com/golang-jwt/jwt/v5`: For JSON Web Token (JWT) generation and verification.

## Installation

1.  **Ensure Go is Installed:** If you don't have Go installed on your system, please follow the installation guide at https://go.dev/doc/install.

2.  **Clone the Repository:**

    ```bash
    git clone https://github.com/smutluuuu/goapi.git
    cd goapi
    ```

3.  **Download Dependencies:**

    ```bash
    go mod tidy
    ```

    This command will download all the necessary Go packages listed in the `go.mod` file.

4.  **Create the `.env` File:**
    In the root directory of the project, create a file named `.env` and copy the following content into it. This file contains the environment variables required for the API to connect to the database and configure other settings.

    ```env
    DB_USER=root
    DB_PASSWORD=root
    DB_NAME=school
    API_PORT=:3000
    DB_PORT=3306
    HOST=127.0.0.1
    CONNECTION_STRING=root:@tcp(127.0.0.1:3306)/school
    JWT_SECRET="jwtsecretstring"
    JWT_EXPIRES_IN=20m
    RESET_TOKEN_EXP_DURATION=10
    ```

    **Note:** You can modify these values according to your local database setup and desired API configuration.

## Running the API

1.  **Navigate to the API Directory:**

    ```bash
    cd goapi/cmd/api
    ```

2.  **Run the API:**

    ```bash
    go run server.go
    ```

    This command will start the School Management System API server on the port specified in your `.env` file (default is `:3000`).

## API Endpoints

### Executives (`/execs`)

- `GET /execs`: Retrieve a list of all executives.
- `POST /execs`: Create a new executive.
- `PATCH /execs`: Partially update multiple executives.
- `GET /execs/{id}`: Retrieve a specific executive by ID.
- `PATCH /execs/{id}`: Partially update a specific executive by ID.
- `DELETE /execs/{id}`: Delete a specific executive by ID.
- `POST /execs/{id}/updatepassword`: Update the password of a specific executive.
- `POST /execs/login`: Login for executives.
- `POST /execs/logout`: Logout for executives.
- `POST /execs/forgotpassword`: Initiate the forgot password process.
- `POST /execs/resetpassword/reset/{resetcode}`: Reset the password using a reset code.

### Students (`/students`)

- `GET /students`: Retrieve a list of all students (supports filtering and sorting).
- `POST /students`: Add multiple new students.
- `PATCH /students`: Partially update multiple students.
- `DELETE /students`: Delete multiple students.
- `PUT /students/{id}`: Update an existing student by ID.
- `GET /students/{id}`: Retrieve a specific student by ID.
- `PATCH /students/{id}`: Partially update a specific student by ID.
- `DELETE /students/{id}`: Delete a specific student by ID.

### Teachers (`/teachers`)

- `GET /teachers`: Retrieve a list of all teachers (supports filtering and sorting).
- `POST /teachers`: Add multiple new teachers.
- `PATCH /teachers`: Partially update multiple teachers.
- `DELETE /teachers`: Delete multiple teachers.
- `PUT /teachers/{id}`: Update an existing teacher by ID.
- `GET /teachers/{id}`: Retrieve a specific teacher by ID.
- `PATCH /teachers/{id}`: Partially update a specific teacher by ID.
- `DELETE /teachers/{id}`: Delete a specific teacher by ID.
- `GET /teachers/{id}/students`: Retrieve all students associated with a specific teacher ID.
- `GET /teachers/{id}/studentcount`: Retrieve the number of students associated with a specific teacher ID.

## Usage

[Provide examples of how to interact with your API using `curl` or other HTTP clients. Focus on key functionalities like creating a user, logging in, and accessing protected resources (if applicable).]

**Example (Login as an executive):**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"username": "your_username", "password": "your_password"}' http://localhost:3000/execs/login
```
