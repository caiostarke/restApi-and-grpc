# Fiber Go Blog REST API

## Tools & Libraries used
- [Golang 1.16+](https://golang.org/doc/go1.16)
- [MongoDB](https://www.mongodb.com/)
- [Docker](https://www.docker.com/get-started)
- [Fiber framework](https://github.com/gofiber/fiber)
- [JWT](https://github.com/form3tech-oss/jwt-go)
- [Swagger docs](https://github.com/swaggo/swag)
- [gosec](https://github.com/securego/gosec)

## 📦 Used packages

| Name                                                                  | Version   | Type       |
| --------------------------------------------------------------------- | --------- | ---------- |
| [gofiber/fiber](https://github.com/gofiber/fiber)                     | `v2.7.1`  | core       |
| [gofiber/jwt](https://github.com/gofiber/jwt)                         | `v2.2.1`  | middleware |
| [arsmn/fiber-swagger](https://github.com/arsmn/fiber-swagger)         | `v2.6.0`  | middleware |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.0`  | tests      |
| [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)               | `v3.2.2`  | auth       |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.3.0`  | config     |

## 🗄 Project structure

### /app

**Folder with business logic only**. This directory doesn't care about what database driver you're using.

- `/app/controller` folder for functional controller (used in routes)
- `/app/dto` Data Transfer Objects(DTO) folder for transform data before sent to API clients
- `/app/model` folder for describe business models and methods of your project
- `/app/repository` folder for perform database operations for models of your project

### /cmd
**Main applications for this project.**

The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/server` `/cmd/cron`).
Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects,
then it should live in the `/pkg` directory.

### /docs

**Folder with API Documentation.**

This directory contains config files for auto-generated API Docs by Swagger, screenshots
and any other documents related to this project.

### /pkg

**Folder with project-specific functionality.** This directory contains all the project-specific code tailored only for your business use case.
<!-- 
- `/pkg/config` folder for configuration functions
- `/pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `/pkg/route` folder for describe routes of your project
- `/pkg/validator` folder with validation functions -->

### /platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project,
like setting up the database, logger instance and storing migrations, seeds(demo data).

<!-- - `/platform/database` folder with database setup functions (by default, PostgreSQL)
- `/platform/logger` folder with better logger setup functions (by default, Logrus)
- `/platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)
- `/platform/seeds` folder with demo data for application rapid setup. mostly **sql** scripts -->

## ⚙️ Configuration

```ini
# .env
; # APP settings:
; APP_HOST="0.0.0.0"
; APP_PORT=5000
; APP_READ_TIMEOUT=30
; APP_DEBUG=false


# JWT settings:
JWT_SECRET_KEY="super_secret_here"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=120

# Database settings:
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=password
DOCKER_SERVICE=mongo
MONGO_PORT=27017
```

## 🔨 Docker development

- Install **`docker`**, **`docker-compose`**
- Rename `.env.example` to `.env`
- Start a MongoDB container exposing port 27017
<!-- - Run migrations `make migrate.up`
- Now start api server with hot reloading `make docker.dev`
- Visit **`http://localhost:5000`** or **`http://localhost:5000/swagger/`** -->

### Todo

- [ ] ADD JWT Middleware
- [ ] Create Tests
- [ ] Refactor the code
- [ ] Create dockerfile and docker-compose
- [ ] Create Documentation with Swagger

## ⚠️ License

[MIT](https://opensource.org/licenses/MIT)
