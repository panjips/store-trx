# Store Trx

Store Trx is a learning project backend application built with **Golang** that provides APIs for managing user accounts, stores, products, and transactions. This project follows the **Clean Architecture** pattern and uses **GORM** for database interactions, **JWT** for authentication, and **Gorilla Mux** for routing.

---

## Features

- **User Management**:
  - Register a new user.
  - Login with JWT authentication.
  - Get and update user profile.

- **Store Management**:
  - Create and update store information.
  - Get store details.

- **Product Management**:
  - Add, update, and delete products.
  - Get product details and list all products.

- **Transaction Management**:
  - Create and manage transactions.
  - View transaction history.

- **Authentication & Authorization**:
  - JWT-based authentication.
  - Role-based access control (e.g., admin vs. regular user).

---

## Technologies Used

- **Golang**: The primary programming language.
- **GORM**: ORM for database interactions.
- **MySQL**: Database for storing application data.
- **JWT**: JSON Web Tokens for authentication.
- **Gorilla Mux**: HTTP router for handling API routes.

---

## On Development
- **Product Management**
- **Transaction Management**
- **Upload File**
- **Integrating with Object Storage**

---

## Project Structure

```
store-trx/
└── app/
│   └── main.go
├── internal/
│   ├── entity/           # Database models (entities)
│   ├── repository/       # Database operations
│   ├── usecase/          # Business logic
│   ├── handler/          # HTTP handlers
│   │   ├── responses/    # Custom response structures
│   │   └── routes/       # Route definitions
│   ├── middleware/       # Middleware for HTTP requests
│   └── pkg/              # Shared utilities (e.g., JWT, config)
├── .env                  # Environment variables
├── go.mod                # Go module file
├── go.sum                # Go module checksum file
└── README.md             # Project documentation
```

---

## Getting Started

### Prerequisites

- **Go**: Install Go from [here](https://golang.org/dl/).
- **MySQL**: Install MySQL from [here](https://dev.mysql.com/downloads/).
- **Git**: Install Git from [here](https://git-scm.com/).

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/store-trx.git
   cd store-trx
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Start the server:
   ```bash
   go run cmd/api/main.go
   ```

--- 

Let me know if you need further customization! 😊
