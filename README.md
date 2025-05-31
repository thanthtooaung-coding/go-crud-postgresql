# Go CRUD PostgreSQL API

A simple and robust RESTful API built with Go, Gin, and GORM for performing Create, Read, Update, and Delete (CRUD) operations on a PostgreSQL database. This project serves as a foundational example for building Go-based web services interacting with a relational database.

## ✨ Features

* **Create Todo:** Add new todo items to the database.
* **Retrieve All Todos:** Fetch a list of all existing todo items.
* **Retrieve Single Todo:** Get details of a specific todo item by its ID.
* **Update Todo:** Modify existing todo item's content and status.
* **Delete Todo:** Remove a todo item from the database.
* **Database Migrations:** Automated database schema management using GORM's `AutoMigrate`.
* **Environment Variables:** Secure configuration management using `.env` file.

## 🚀 Technologies Used

* **Go (Golang):** The primary programming language.
* **Gin Web Framework:** A high-performance HTTP web framework.
* **GORM:** An excellent ORM library for Go, simplifying database interactions.
* **PostgreSQL:** A powerful, open-source relational database system.

## 📋 Prerequisites

Before running this project, ensure you have the following installed:

* **Go:** [Download and Install Go](https://golang.org/doc/install) (version 1.18 or higher recommended).
* **PostgreSQL:** [Download and Install PostgreSQL](https://www.postgresql.org/download/)
* **Git:** For cloning the repository.

## 🛠️ Getting Started

Follow these steps to set up and run the project locally.

### 1. Clone the Repository

```bash
git clone https://github.com/thanthtooaung-coding/go-crud-postgresql.git
cd go-crud-postgresql