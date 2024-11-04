# CMS Backend

This repository contains the starter code for a Content Management System (CMS) backend API built with Go (Golang), the Gin web framework, and GORM for ORM. This project is intended to help learners understand CRUD operations, database migrations, environment configuration, and testing through hands-on coding.

## Project Objectives

The objective of this project is to build a fully functional backend API with CRUD capabilities for managing pages, posts, and media content. Learners will:

- Complete the CRUD handlers in the `Page` and `Post` controllers.
- Implement database operations using GORM.
- Configure environment variables for different environments.
- Add tests to verify functionality, including unit and integration tests.

The final deliverable is a backend API with comprehensive test coverage that follows best practices in backend development.

## Features

- RESTful API endpoints for managing pages, posts, and media.
- Database schema management with migrations using golang-migrate.
- Environment-specific configurations for development and production.
- GORM's AutoMigrate used in development for rapid schema updates.
- Secure handling of environment variables.
- Middleware setup with Gin for logging and recovery.

## Prerequisites

- [Go 1.16 or higher](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Git](https://git-scm.com/)
- [golang-migrate](https://github.com/golang-migrate)
- [pgAdmin 4](https://www.pgadmin.org/download/) (Optional)

## Installation

1. **Clone the repository**

    `git clone https://github.com/udacity/backend-dev-C1-starter.git `
    `cd project/solution`

2. **Install Go dependencies** 

    `go mod download`

## Environment setup

This section guides you through setting up the environment variables needed for development and production environments. Properly managing these configurations ensures that sensitive information is protected and makes switching between environments seamless.

### Development environment

1. Create a `.env` file. In the project root directory, create a `.env` file to store your development environment variables:

`cp .env.example .env`

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
ENV=development
```

2. Replace `your_db_user`, `your_db_password`, and `your_db_name` with your PostgreSQL development database credentials.

3. Add `.env` to `.gitignore`. Ensure that your `.env` file is not committed to version control by adding it to your `.gitignore` file.

### Production environment

For production, set ENV=production in the environment variables. This will ensure that golang-migrate is used for migrations, and database credentials are set to the production database instance.

## Create Databases

You can set up your PostgreSQL database using either the PostgreSQL CLI or pgAdmin 4.

### Using PostgreSQL CLI

1.  Start PostgreSQL service using the following commands:

`brew install postgresql`
`brew services start postgresql`

2. Access PostgreSQL CLI by running the following command:

`psql -U postgres`

3.  Create a database called `your_db_name` and user called `your_db_user`.
   
```
CREATE DATABASE your_db_name;
CREATE USER your_db_user WITH ENCRYPTED PASSWORD 'your_db_password';
GRANT ALL PRIVILEGES ON DATABASE your_db_name TO your_db_user;
```

1.  Exit PostgreSQL CLI using the following command:

`\q`

### Using pgAdmin 4

pgAdmin 4 is a web-based GUI tool that simplifies database management. Follow these steps if you prefer managing the database via a GUI.

1. Download from [pgAdmin's official website](https://www.pgadmin.org/download/). Follow the installation instructions for your operating system.

2. Launch pgAdmin 4 by opening pgAdmin 4 and create a new server connection with the following information:
  - *Name*: Local PostgreSQL (or any name you prefer)
  - *Host*: localhost
  - *Port*: 5432
  - *Username*: postgres (or your PostgreSQL superuser)
  - *Password*: Your PostgreSQL password

3. Let's create the database. Right-click on **Databases** in the sidebar and select **Create** > **Database**... Then, input the following information in the corresponding fields:
- *Database Name*: `your_db_name`
- *Owner*: `your_db_user` (you may need to create this user first)

4. Create User (Role) by navigating to **Login/Group Roles** in the sidebar. Right-click and select **Create** > **Login/Group Role**... and input the following information:
- *Role Name*: `your_db_user`
- *Password*: `your_db_password`
- *Privileges*: Assign appropriate privileges, such as Can login, Create DB, etc.

1. After creating the user, ensure they have the necessary privileges on your database. You can do this by running SQL queries in pgAdmin's Query Tool:
`GRANT ALL PRIVILEGES ON DATABASE your_db_name TO your_db_user;`

## Migrations

This section explains how to manage database migrations, which are essential for schema management and version control.

### Using golang-migrate for Production Migrations

golang-migrate is used to manage database schema changes in a production environment.

1. Install `golang-migrate`. If you're using macOS with Homebrew run the following command in Terminal: 
` brew install golang-migrate ` 
or follow installation instructions from in the [GoLang repo](https://github.com/golang-migrate/migrate).

2. Run migrations using the following command:

` migrate -database postgres://your_db_user:your_db_password@host:port/your_db_name?sslmode=disable -path ./migrations up `

### AutoMigrate in Development

During development, GORM’s AutoMigrate feature can be used to automatically update the schema based on your models, which is enabled in main.go when ENV=development.

## Running the application

    1. Ensure PostgreSQL is running
    ` pg_ctl -D /usr/local/var/postgres start  # macOS example `

    2. Run the application

    ` go run main.go `

## API Documentation

The API documentation provides an overview of the available endpoints for the CMS Backend. The CMS Backend API includes endpoints for managing pages, posts, and media content.

Base URL: http://localhost:8080/api/v1

Endpoints
 - Pages

        `GET /pages`    - Get all pages
        `GET /pages/:id` - Get a single page by ID
        `POST /pages` - Create a new page
        `PUT /pages/:id` - Update a page
        `DELETE /pages/:id` - Delete a page

- Posts

        `GET /posts` - Get all posts
        `GET /posts/:id` - Get a single post by ID
        `POST /posts` - Create a new post
        `PUT /posts/:id` - Update a post
        `DELETE /posts/:id` - Delete a post

- Media

        `GET /media` - Get all media
        `GET /media/:id` - Get a single media by ID
        `POST /media` - Create a new media
        `PUT /media/:id` - Update a media
        `DELETE /media/:id` - Delete a media

## Testing

This section outlines the process for running unit and integration tests. Comprehensive testing is essential to verify functionality and reliability of the CMS backend API.

### Unit Testing

#### Overview

Unit tests verify the functionality of individual components, such as controllers and services, in isolation. The tests use mocks to simulate database interactions and focus on request/response handling and error handling.

 Our unit tests focus on:
- Controllers (Media, Post, Page)
- Request/Response handling
- Database interactions (using mocks)
- Error handling

#### Running Unit Tests
```bash
# Run all unit tests with verbose output
go test ./controllers -v

# Run specific controller tests
go test ./controllers -run TestGetMedia -v
go test ./controllers -run TestCreatePost -v
go test ./controllers -run TestUpdatePage -v

# Run tests with coverage
go test ./controllers -cover
```

#### Unit Test Coverage

Below are the unit test cases included for each controller, which validate CRUD operations and error handling:

#### Media Controller Tests
- TestGetMedia - List all media
- TestGetMediaByID - Get single media
- TestCreateMedia - Create new media
- TestDeleteMedia - Delete media

#### Post Controller Tests
- TestGetPosts - List all posts
- TestGetPostsWithFilters - Filter posts by title/author
- TestGetPost - Get single post
- TestCreatePost - Create new post
- TestUpdatePost - Update existing post
- TestDeletePost - Delete post

#### Page Controller Tests
- TestGetPages - List all pages
- TestGetPage - Get single page
- TestCreatePage - Create new page
- TestUpdatePage - Update existing page
- TestDeletePage - Delete page

### Integration Testing

#### Overview
Integration tests verify the functionality of the API as a whole, including interactions between components and the database.

#### Prerequisites
- Go 1.19 or higher
- PostgreSQL 12 or higher
- Access to create/modify databases

#### Setup

To perform integration testing, create a separate test database and update the .env.test file with appropriate credentials.

1. Create test database:
```bash
# Using psql
PGPASSWORD=postgres psql -h localhost -U postgres -c "DROP DATABASE IF EXISTS cms_test;"
PGPASSWORD=postgres psql -h localhost -U postgres -c "CREATE DATABASE cms_test;"
```

#or using Makefile

` make create-test-db `

2. Configure environment (optional):
```env
# .env.test
TEST_DB_HOST=localhost
TEST_DB_USER=postgres
TEST_DB_PASSWORD=postgres
TEST_DB_NAME=cms_test
TEST_DB_PORT=5432
```

#### Running Integration Tests

To run all tests, run the following command:

`make test`

To run only integration tests, run the following command:

`make test-integration`
To run integration tests with database setup, run the following command:

`make test-integration-full`

To run specific integration tests, run the following command:
```bash
go test ./tests/integration -run TestMediaIntegration -v
go test ./tests/integration -run TestPostIntegration -v
```


#### Integration Test Coverage

Integration tests focus on end-to-end functionality, verifying that the API behaves correctly under various conditions:

#### Media Integration Tests
```go
TestMediaIntegration
├── Create Media
│   └── Tests media creation and validation of created attributes.
└── Get All Media
    └── Verifies retrieval of all media assets.
```

#### Post Integration Tests
```go
TestPostIntegration
├── Create Post with Media
│   └── Tests post creation with associated media.
└── Get Posts with Filter
    └──  Filters posts by title or author and verifies results.
```

## Makefile Commands

The Makefile automates common tasks, such as running tests and managing the database.

```makefile
# Available commands:
make test              # Run all tests
make test-unit         # Run only unit tests
make test-integration  # Run only integration tests
make create-test-db    # Create test database
```

## Troubleshooting

This section addresses common issues encountered when setting up or running the CMS backend API.

### Common Issues

1. Unit Test Failures
```bash
# Issue: Mock expectations not met
Error: mock expectation not met [file.go:123]
Solution: Check mock setup in test file
```

2. Integration Test Database Issues
```bash
# Issue: Cannot connect to database
Error: dial tcp [::1]:5432: connect: connection refused
Solution: Ensure PostgreSQL is running

# Issue: Database permissions
Error: permission denied for database "cms_test"
Solution: Run: GRANT ALL PRIVILEGES ON DATABASE cms_test TO postgres;
```

### Test Structure

Organizing tests by functionality ensures clarity and ease of maintenance. The following structure is used:

```
├── controllers/
│   ├── media_controller_test.go
│   ├── page_controller_test.go
│   └── post_controller_test.go
└── tests/
    └── integration/
        ├── main_test.go
        ├── media_test.go
        └── post_test.go
```

Each test suite within controllers/ and integration/ includes individual tests covering CRUD operations and validations, contributing to a well-tested and robust API backend.