# Post-Comment Service

This is a Go-based microservice for managing posts and comments. It provides a RESTful API for user registration, authentication, creating posts, and adding comments to posts.

## Project Structure

```
post-comment-service/
│
├── cmd/
│ └── main.go
| └── run_migrations.go
│
├── internal/
│ ├── adapters/
│ │ └── repositories/
│ │   ├── http/
│ │   │ ├── handlers.go
│ │   │ └── user_handler.go
│ │   ├── postgres/
│ │   │ ├── comment_repository.go
│ │   │ ├── post_repository.go
│ │   │ └── user_repository.go
│ │   └── router/
│ │   └── router.go
│ │ 
│ ├── application/
│ │ └── service.go
│ ├── config/
│ │ ├── config.go
│ │ └── config.yaml
│ ├── domain/
│ │ └── models.go
│ └── ports/
│ └── ports.go
│
├── migrations/
│ ├── 000001_create_users.up.sql
│ ├── 000001_create_users.down.sql
│ ├── 000002_create_posts.up.sql
│ ├── 000002_create_posts.down.sql
│ ├── 000003_create_comments.up.sql
│ └── 000003_create_comments.down.sql
│
├── pkg/
│ ├── errors/
│ │ └── errors.go
│ ├── logger/
│ │ └── logger.go
│ └── middleware/
│ └── auth.go
│
├── tools/
│ └── create_migration.go
│
├── go.mod
├── go.sum
└── Makefile
```

## Prerequisites

- Go 1.16 or later
- PostgreSQL 12 or later

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/post-comment-service.git
   cd post-comment-service
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up the PostgreSQL database:
   - Create a new database named `post_comments_db`
   - Create a user with appropriate permissions

4. Update the `config.yaml` file in the `internal/config/` directory with your database credentials:
   ```yaml
   database:
     url: "postgres://yourusername:yourpassword@localhost:5432/post_comments_db?sslmode=disable"
   ```

5. Run database migrations:
   ```
   go run tools/create_migration.go -name <your_migration_name>

   ```

## Running the Service

To start the service, run:

```
go run cmd/main.go
```

The service will start on `localhost:8080` by default.

## API Endpoints

- `POST /register`: Register a new user
- `POST /login`: Authenticate and receive a JWT token
- `GET /posts`: List all posts
- `POST /posts`: Create a new post (requires authentication)
- `GET /posts/{id}`: Get a specific post
- `POST /posts/{postID}/comments`: Add a comment to a post (requires authentication)
- `GET /posts/{postID}/comments`: Get all comments for a post

## Usage Examples

Here are some curl commands to interact with the API:

1. Register a user:
   ```
   curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d "{\"username\":\"testuser\",\"password\":\"testpassword\"}"
   ```

2. Login:
   ```
   curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{\"username\":\"testuser\",\"password\":\"testpassword\"}"
   ```

3. Create a post (replace `YOUR_JWT_TOKEN` with the token received from login):
   ```
   curl -X POST http://localhost:8080/posts -H "Content-Type: application/json" -H "Authorization: Bearer YOUR_JWT_TOKEN" -d "{\"title\":\"My First Post\",\"content\":\"This is the content of my first post.\"}"
   ```

4. Get all posts:
   ```
   curl -X GET http://localhost:8080/posts
   ```

5. Add a comment to a post (replace `POST_ID` and `YOUR_JWT_TOKEN`):
   ```
   curl -X POST http://localhost:8080/posts/POST_ID/comments -H "Content-Type: application/json" -H "Authorization: Bearer YOUR_JWT_TOKEN" -d "{\"content\":\"This is a comment on the post.\"}"
   ```
