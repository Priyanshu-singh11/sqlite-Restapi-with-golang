#A SIMPLE PROJECT FOR TECHING MY FRIENDS 
# SQLite REST API with Golang

A production-ready REST API built with **Go**, **Gin**, and **SQLite** demonstrating clean project architecture, CRUD operations, request validation, middleware, and database integration.

---

## Tech Stack

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-%2300ADD8.svg?style=for-the-badge&logo=gin&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405E.svg?style=for-the-badge&logo=sqlite&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-00ADD8?style=for-the-badge)
![REST API](https://img.shields.io/badge/REST_API-009688?style=for-the-badge)
![JSON](https://img.shields.io/badge/JSON-000000?style=for-the-badge&logo=json&logoColor=white)
![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

---

## Features

- Create Resource
- Get All Resources
- Get Resource by ID
- Update Resource
- Delete Resource
- RESTful API Design
- SQLite Database
- JSON Responses
- Clean Architecture
- Error Handling
- Request Validation

---

## Project Structure

```text
.
├── cmd/
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── models/
│   ├── routes/
│   └── repository/
├── go.mod
├── go.sum
└── main.go
```

---

## Installation

Clone the repository

```bash
git clone https://github.com/Priyanshu-singh11/sqlite-Restapi-with-golang.git
```

Go to the project directory

```bash
cd sqlite-Restapi-with-golang
```

Install dependencies

```bash
go mod tidy
```

Run the server

```bash
go run main.go
```

---

## API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | / | Home |
| GET | /api/... | Get All |
| GET | /api/.../:id | Get By ID |
| POST | /api/... | Create |
| PUT | /api/.../:id | Update |
| DELETE | /api/.../:id | Delete |




**Priyanshu Singh**

GitHub: https://github.com/Priyanshu-singh11
