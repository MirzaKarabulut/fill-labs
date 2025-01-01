### User Management System
- This project is a User Management System built with a Go backend using the Fiber framework and a React frontend. The system allows for creating, reading, updating, and deleting users.

## Table of Contents:
Installation
Usage
API Endpoints
Frontend Components


# Installation
  ```bash
- Backend

Clone the repository:

git clone https://github.com/MirzaKarabulut/fill-labs.git
cd user-management/backend

Install dependencies:
go mod tidy

Run the server:
go run main.go

- Frontend

Navigate to the frontend directory:
cd ../frontend

Install dependencies:
npm install

Run the development server:
npm run dev
```
# Usage

- Backend
The backend server will be running on http://localhost:8080. It provides a RESTful API for managing users.

- Frontend
The frontend development server will be running on http://localhost:3000. It provides a user interface for managing users.

# API Endpoints

- GET /api/users:
Retrieve a list of users.

- GET /api/users/:id:
Retrieve a specific user by ID.

- POST /api/users:
Create a new user.

- PUT /api/users/:id:
Update an existing user by ID.

- DELETE /api/users/:id:
Delete a user by ID.

# Frontend Components

- UserTable
Displays a table of users with options to create, edit, and delete users.

- UserForm
Provides a form for creating, editing, and deleting users.
