# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Common Development Commands

### Building and Running
```bash
# Build the application
go build -o swadiq main.go

# Run directly with Go
go run main.go

# Run with hot reloading using Air
air

# Run with local database
LOCAL_DB=true go run main.go
```

### Database Operations
```bash
# Connect to remote PostgreSQL database
psql -h 129.80.199.242 -p 5432 -U imaad -d swadiq

# Apply schema to database
psql -h 129.80.199.242 -p 5432 -U imaad -d swadiq -f schema.sql

# For local development with PostgreSQL
createdb swadiq
psql -d swadiq -f schema.sql
export LOCAL_DB=true
```

### Testing and Development
```bash
# Install dependencies
go mod tidy

# Format code
go fmt ./...

# Vet code for issues
go vet ./...

# Run a specific test (example)
go test ./app/routes/auth -v
```

## Architecture Overview

### Technology Stack
- **Backend**: Go 1.21+ with Fiber web framework
- **Database**: PostgreSQL 18 with UUID support
- **Templates**: HTML templates with Tailwind CSS
- **Session Management**: Cookie-based with database storage
- **Password Security**: Bcrypt hashing

### Project Structure

The codebase follows a modular architecture pattern:

```
app/
├── config/         # Database configuration and connection management
├── database/       # Database queries and operations (queries.go)
├── models/         # Data structures (25+ models for school entities)
├── routes/         # Feature modules organized by domain
│   ├── auth/       # Authentication (login, sessions, middleware)
│   ├── dashboard/  # Main dashboard
│   ├── students/   # Student management
│   ├── teachers/   # Teacher management
│   ├── classes/    # Class management
│   ├── subjects/   # Subject management
│   ├── attendance/ # Attendance tracking
│   └── parents/    # Parent management
└── templates/      # HTML templates organized by module
```

### Key Architectural Patterns

**Module Structure**: Each feature module follows the same pattern:
- `API.go` - HTTP handlers and business logic
- `routes.go` - Route definitions and middleware setup
- `utils.go` - Helper functions specific to the module

**Database Layer**: Centralized in `app/database/queries.go` with prepared statements to prevent SQL injection.

**Authentication Flow**: 
- Session-based authentication with UUIDs
- Role-based access control (admin, head_teacher, class_teacher, subject_teacher)
- Middleware for both web pages and API endpoints
- Custom error handling for 401/403/404/500 errors

**Database Configuration**: 
- Automatic fallback from remote to local PostgreSQL
- Connection pooling configured (25 max open, 5 max idle)
- Environment variable `LOCAL_DB=true` for local development

### Data Model Relationships

The system uses a comprehensive relational model:
- **Users** have **Roles** (many-to-many via user_roles)
- **Students** belong to **Classes** and have **Parents** (many-to-many via student_parents)
- **Classes** have **Subjects** (many-to-many via class_subjects)
- **Subjects** belong to **Departments**
- **Attendance** tracks daily student presence
- **Sessions** manage user authentication

### Development Patterns

**Route Registration**: All routes are registered in `main.go` using module setup functions:
```go
auth.SetupAuthRoutes(app)
students.SetupStudentsRoutes(app)
```

**Middleware Usage**: Authentication middleware is applied at route group level, with role-based middleware for specific permissions.

**Template Rendering**: Fiber's HTML template engine with automatic reloading in development mode.

**Error Handling**: Custom error handler in `main.go` provides consistent error pages and API responses.

## Database Schema Notes

- Uses PostgreSQL UUID extension (`uuid_generate_v4()`)
- Soft deletes with `deleted_at` timestamps
- Audit trails with `created_at` and `updated_at`
- Foreign key constraints ensure data integrity
- Default values and check constraints for data validation

## Default Authentication

Test accounts are available:
- Admin: `admin@swadiq.com` / `password123`
- Head Teacher: `headteacher@swadiq.com` / `password123`
- Class Teacher: `teacher1@swadiq.com` / `password123`
- Subject Teacher: `teacher2@swadiq.com` / `password123`

## Development Server

The application runs on port 8080 by default. Static files are served from `./static/`.

## Hot Reloading

Air is configured for development with automatic rebuilding on Go file changes. Configuration is in `.air.toml`.
