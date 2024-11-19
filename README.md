# CMS Backend Project

Content management systems (CMS) are important for managing and delivering content across websites and applications. As a backend developer, being able to build a robust CMS backend is a highly valuable skill. In this project, you will create a CMS backend API using Go, Gin, and PostgreSQL, simulating real-world tasks you might encounter in your career. This project covers essential backend skills, including CRUD operations, database migrations, testing, and deploying environment-specific configurations.   

## Project Objectives

The objective of this project is to build a fully functional backend API with CRUD capabilities for managing pages, posts, and media content in a content management system (CMS). By completing this project, you will:

- Apply Go programming skills to build a RESTful API.
- Utilize the Gin web framework for efficient routing and middleware management.
- Implement database interactions using GORM with PostgreSQL.
- Manage database schema changes using migrations.
- Write unit and integration tests to ensure code quality and reliability.
- Configure environment variables for different deployment environments.

The final deliverable is a backend API with comprehensive test coverage that follows best practices in backend development and demonstrates your ability to:

- Build scalable backend services using Go, one of the most in-demand programming languages in the industry.
- Work with relational databases and ORMs, a critical skill for backend developers.
- Implement RESTful APIs, which are foundational to modern web development.
- Write robust tests, showcasing your commitment to code quality and reliability.
- Manage environment configurations and understand the importance of secure credential handling.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go 1.16 or higher](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Git](https://git-scm.com/)
- [golang-migrate](https://github.com/golang-migrate)
- [pgAdmin 4](https://www.pgadmin.org/download/) (Optional)

### Installation steps

1. **Clone the repository**

    ```bash
    git clone https://github.com/udacity/backend-dev-C1-starter.git
    cd project/solution
    ```
2. **Install Go dependencies** 

    ```bash
    go mod download
    ```

3. **Set Up Environment Variables**

    Create a `.env` file. In the project root directory, create a `.env` file to store your development environment variables:

    ```bash
    cp .env.example .env
    ```

    Open the `.env` file and replace the placeholder values with your actual database credentials:

    ```bash
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    ENV=development
    ```

    Note: Ensure that the .env file is included in your .gitignore to prevent sensitive information from being committed to version control. Add `.env` to `.gitignore`.


### Create Databases

You can set up your PostgreSQL database using either the PostgreSQL CLI or pgAdmin 4.

**Option 1. Using PostgreSQL CLI**

1.  Start PostgreSQL service using the following commands:

    ```bash
    brew install postgresql
    brew services start postgresql
    ```

2. Access PostgreSQL CLI by running the following command:

    ```bash
    psql -U postgres
    ```

3.  Create a database called `your_db_name` and user called `your_db_user`.
   
    ```bash
    CREATE DATABASE your_db_name;
    CREATE USER your_db_user WITH ENCRYPTED PASSWORD 'your_db_password';
    GRANT ALL PRIVILEGES ON DATABASE your_db_name TO your_db_user;
    ```

4.  Exit PostgreSQL CLI using the following command:

    ```bash
    \q
    ```

**Option 2. Using pgAdmin 4**

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


## Step-by-Step Implementation Guide
In the following step-by-step guide, we'll walk you through building your CMS backend. You'll learn how to define your models, create database migrations, implement CRUD operations, set up routing, and write tests.

### Step 1: Define the Models
Files: `models/page.go`, `models/post.go`, and `models/media.go`

Task: Define data structures for `Page` and `Post`.

Instructions:

- Open each model file and define struct fields with appropriate GORM tags, as described in // TODO comments.
- Ensure fields like `ID`, `Title`, `Content`, `Slug`, `CreatedAt`, and `UpdatedAt` are included in each model.

### Step 2: Create Migrations for Models

Files: `migrations/000001_create_media_table.up.sql`, `migrations/000001_create_media_table.down.sql`, `migrations/000002_create_posts_table.up.sql`, `migrations/000002_create_posts_table.down.sql`

Task: Create migrations for `Media` and `Post` models.

Instructions:

1. Create the media table migration files:
   - Create `migrations/000001_create_media_table.up.sql`:
     - Define table with columns: id, url, type, created_at, updated_at
     - Add appropriate data types and constraints
     - Consider adding indexes for performance
   - Create `migrations/000001_create_media_table.down.sql`:
     - Include cleanup logic to remove the table
     - Drop any created indexes

2. Create the posts table migration files:
   - Create `migrations/000002_create_posts_table.up.sql`:
     - Define posts table with columns: id, title, content, author, created_at, updated_at
     - Create post_media junction table for many-to-many relationship
     - Add foreign key constraints with cascade delete
   - Create `migrations/000002_create_posts_table.down.sql`:
     - Drop tables in correct order (post_media before posts)
     - Ensure clean removal of all related objects

 Migration Best Practices:
   - Use appropriate data types (SERIAL for IDs, VARCHAR with limits, TEXT for content)
   - Include NOT NULL constraints where needed
   - Add timestamps with timezone support
   - Consider adding indexes for frequently queried columns
   - Ensure down migrations can cleanly reverse all changes

Example: See `/migrations/000001_create_pages_table.down.sql` and `/migrations/000002_create_media_table.down.sql` for a complete implementation.

### Step 2: Run Initial Database Migrations
Task: Apply the initial database schema to your PostgreSQL database.

Instructions:

Option 1: Using GORM AutoMigrate (Development Environment)

If you're in the development environment (ENV=development), you can use GORM's AutoMigrate feature to automatically migrate your database schema based on your models.

1. Run the Application to AutoMigrate:

` go run main.go `

2. Verify the Tables:

Use a PostgreSQL client or GUI tool to check that the pages, posts, and media tables have been created.
Note: Ensure that your main.go includes the AutoMigrate calls:

```go
if env == "development" {
    log.Println("Running AutoMigrate...")
    if err := db.AutoMigrate(&models.Page{}, &models.Post{}, &models.Media{}); err != nil {
        log.Fatalf("Failed to automigrate database: %v", err)
    }
    }
```

Option 2: Using golang-migrate (Production Environment)

For a more controlled migration process, especially in production environments (ENV=production), use golang-migrate to apply migrations from SQL files.

Option 2: Using golang-migrate (Production Environment)

For a more controlled migration process, especially in production environments (ENV=production), use golang-migrate to apply migrations from SQL files.

1. Install golang-migrate:

If you haven't installed it yet, follow the installation instructions from the golang-migrate GitHub repository.
Prepare Migration Files:

Ensure your migration SQL files are correctly set up in the ./migrations directory.
Example migration files:
- `000001_create_pages_table.up.sql`
- `000001_create_pages_table.down.sql`

Run Migrations:

```bash
migrate -database "postgres://your_db_user:your_db_password@localhost:5432/your_db_name?sslmode=disable" -path ./migrations up
```
Replace `your_db_user`, `your_db_password`, and `your_db_name` with your database credentials.

Verify the Migrations:

Use a PostgreSQL client or GUI tool to confirm that the tables have been created according to your migration files.
Note: When running in production mode, ensure that your main.go does not perform AutoMigrate to prevent unintended schema changes.

### Step 3: Implement the GetPost Handler
File: `controllers/post_controller.go`

Task: Complete the `GetPost` function to retrieves a specific post by ID

Instructions:

- Query the posts table to retrieve a specific post by ID.
- If there’s an error, return a `500` status with an error message.
- If successful, return a `200` status and the list of pages as JSON.

Example: See `GetPosts` function in `controllers/post_controller.go` for a complete implementation.

### Step 4: Implement the CreatePost Handler
File: `controllers/post_controller.go`

Task: Complete the CreatePost function to create a new post.

Instructions:

- Bind JSON data from the request to the page struct.
- Validate required fields (Title and Content).
- Insert the new post into the database.
- Return a 201 status and the created post as JSON.

### Step 5: Implement UpdatePost and DeletePost Handlers
File: `controllers/post_controller.go`

Task: Implement `UpdatePost` and `DeletePost` functions.

Instructions:

- `UpdatePage`:
    - Bind the update data to the `post` struct.
    - Find the post by ID and update its fields.
- `DeletePost`:
    - Find the post by ID and delete it from the database.
    - If not found, return a `404` status. 

### Step 6: Implement the Page Controller

- `page_controller.go`: repeat the CRUD implementation as outlined in `page_controller.go`.

### Step 7: Implement the Media Controller
- `media_controller.go`: repeat the CRUD implementation as outlined in `media_controller.go`.

### Step 8: Implement the Routes
File: `routes/routes.go`

Task: Define the routes for the API using Gin and link them to the handlers.

Instructions:

- Add database middleware that: 
    - Stores the db instance in the context using a `db` key
    - Calls `Next()` to continue to the next handler
- Create API version group that:
    - Uses `router.Group()` to create a new route group
    - Sets the group prefix to `/api/v1`
    - Stores the group in a variable for adding routes
- Within the api group, use `router.GET()`, `router.POST()`, `router.PUT()`, and `router.DELETE()` methods to define routes for pages, posts, and media.
- Mount the routes under the `/api/v1` prefix.

### Step 9: Run the Application and Test Endpoints
Task: Start the server and test the API endpoints manually.

Instructions:

Run the Application:

```bash
go run main.go
```

Test Endpoints Using cURL or Postman:

Example: Test the GetPosts endpoint.

```bash
curl -X GET http://localhost:8080/api/v1/posts 
```
Create a New Post:

```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{"title": "First Post", "content": "This is the content of the first post.", "author": "Admin"}'
```
Verify Responses:

- Ensure that the API responds with the correct status codes and data.
- Check the database to verify that data is being saved correctly.

### Step 10: Write Unit Tests for Page Controller

Unit tests verify the functionality of individual components, such as controllers and services, in isolation. The tests use mocks to simulate database interactions and focus on request/response handling and error handling.

 Our unit tests focus on:
- Controllers (Media, Post, Page)
- Request/Response handling
- Database interactions (using mocks)
- Error handling

Files: `tests/controllers/page_controller_test.go`

Task: Write unit tests for each CRUD operation in `page_controller.go`.

Instructions:
- Use `httptest` to simulate HTTP requests and `testify` for assertions.
- Mock database calls to ensure tests are isolated from real data.

Example: See `TestGetPages` function in `controllers/page_controller_test.go` for a complete implementation.

### Step 11: Write Unit Tests for Post Controller
Files: `tests/controllers/post_controller_test.go`  

Task: Write unit tests for each CRUD operation in `post_controller.go`.

Instructions:
- Use `httptest` to simulate HTTP requests and `testify` for assertions.
- Mock database calls to ensure tests are isolated from real data.

### Step 12: Write Unit Tests for Media Controller
Files: `tests/controllers/media_controller_test.go`  

Task: Write unit tests for each CRUD operation in `media_controller.go`.

Instructions:
- Use `httptest` to simulate HTTP requests and `testify` for assertions.
- Mock database calls to ensure tests are isolated from real data.

#### Step 13: Run Unit Tests
Command:

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


Instructions:

This command runs all unit tests in the tests/unit directory.
Ensure your unit tests cover all CRUD operations and handle both success and error cases.
Use `go test -cover` to check test coverage.

### Step 14: Write Integration Tests

Integration tests verify the functionality of the API as a whole, including interactions between components and the database.

Files: `tests/integration/main_test.go`, `post_test.go`, `media_test.go`

Task: Write integration tests to validate the end-to-end functionality of each API endpoint.

Instructions:
- Implement the `setup()` and `cleanup()` functions to prepare the test environment and clean up after tests in `main_test.go`.
- Implement the `"Get All Media"` test case in `media_test.go`.
- Implement the `"Create Post with Media"` test case in `post_test.go`.
- Implement the `"Get Posts with Filter"` test case in `post_test.go`.

Guidelines:
- Integration tests should use a test database to validate that CRUD operations affect data as expected.
- Each test should cover full CRUD operations to ensure the endpoints interact with the database correctly.

### Step 15: Running Integration Tests
Prerequisites:

- Set up a test database (e.g., cms_test).
- Update your .env.test file with test database credentials.

To perform integration testing, create a separate test database and update the `.env.test` file with appropriate credentials.

1. Create test database:
```bash
# Using psql
PGPASSWORD=postgres psql -h localhost -U postgres -c "DROP DATABASE IF EXISTS cms_test;"
PGPASSWORD=postgres psql -h localhost -U postgres -c "CREATE DATABASE cms_test;"
```

#or using Makefile

```bash
make create-test-db
```

2. Configure environment (optional):
```bash
# .env.test
TEST_DB_HOST=localhost
TEST_DB_USER=postgres
TEST_DB_PASSWORD=postgres
TEST_DB_NAME=cms_test
TEST_DB_PORT=5432
```

Command:

To run all tests, run the following command:

```bash
make test
```

To run only integration tests, run the following command:

```bash
make test-integration
```

To run integration tests with database setup, run the following command:

```bash
make test-integration-full
```

To run specific integration tests, run the following command:
```bash
go test ./tests/integration -run TestMediaIntegration -v
go test ./tests/integration -run TestPostIntegration -v
```

Instructions:

This command runs all integration tests, which test your application end-to-end.
Ensure that the test database is clean before running tests to avoid data contamination.
Use `make create-test-db` to set up the test database if needed.

## Troubleshooting

If you encounter issues while setting up or running the application, consider the following tips:

### Cannot Connect to Database

- Ensure that PostgreSQL is running.
- Verify that your database credentials in the `.env` file are correct.
- Check that the database exists and that the user has appropriate permissions.

### Migrations Fail

- Ensure that the migration files are correctly formatted.
- Check that you have the correct version of `golang-migrate` installed.
- Review error messages for specific details.

### Tests Fail

- Ensure that your test database is properly set up.
- Check for issues in your test setup code.
- Review error messages and stack traces to identify the problem.

### Server Crashes or Does Not Start

- Check the console output for error messages.
- **Dependency Issues:**
  - Run `go mod download` to ensure all dependencies specified in your `go.mod` file are downloaded.
  - If you still face issues, run `go mod tidy` to synchronize your modules:
    ```bash
    go mod tidy
    ```
    - This command adds missing module requirements and removes unnecessary ones.
- Verify that the code compiles without errors:
  ```bash
  go build
