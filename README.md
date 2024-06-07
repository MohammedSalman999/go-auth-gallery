# go-auth-gallery
Go Auth Gallery is a web application built with Go, featuring user authentication and a personal gallery. Users can sign up, log in, and access personalized gallery pages. The app uses session-based authentication and secure password hashing, demonstrating basic auth and session management in Go.

## Features

- User Sign-Up and Log-In
- Session-based Authentication
- User-Specific Gallery Page
- Password Hashing for Secure Authentication

## Requirements

- Go 1.16 or higher
- MySQL database
- Go modules

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/yourusername/go-auth-gallery.git
cd go-auth-gallery
```

### Set Up the Database

1. Ensure you have MySQL installed and running.
2. Create a database named `authentication`.

```sql
CREATE DATABASE authentication;
USE authentication;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARBINARY(255) NOT NULL,
    firstname VARCHAR(255),
    lastname VARCHAR(255)
);
```

### Configuration

Update the `dsn` variable in `main.go` with your MySQL credentials:

```go
dsn := "yourusername:yourpassword@tcp(127.0.0.1:3306)/authentication"
```

### Install Dependencies

```bash
go mod tidy
```

### Run the Application

```bash
go run .
```

### Access the Application

Open your web browser and navigate to `http://localhost:8080`.

## Project Structure

```
go-auth-gallery/
│
├── assets/                 # Static assets (CSS, JS, images)
│
├── templates/              # HTML templates
│   ├── index.tmpl
│   ├── gallery.tmpl
│   ├── login.tmpl
│   ├── signin.tmpl
│   └── person.tmpl
│
├── main.go                 # Entry point of the application
├── db.go                   # Database connection and initialization
├── handlers.go             # HTTP handlers for different routes
├── user.go                 # User-related functions and types
└── sessions.go             # Session management functions
```

## Routes

- `/` - Home Page
- `/gallery` - Gallery Page (requires login)
- `/login` - Login Page
- `/logout` - Logout
- `/person` - User's Personal Page (requires login)
- `/signin` - Sign-Up Page

## Contributing

Contributions are welcome! Please create a pull request or open an issue to discuss any changes.

## License

This project is licensed under the MIT License.

## Acknowledgments

- [Go](https://golang.org/)
- [MySQL](https://www.mysql.com/)
- [Satori UUID](https://github.com/satori/go.uuid)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

