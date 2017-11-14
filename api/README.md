# GradePortal Backend

This is the entire backend for GradePortal written in Golang. It is a pure RESTful API backend, serving data through JSON requests made by the frontend.

# Setup
- Install [Docker](https://docs.docker.com/engine/installation/)

- Inside the /backend directory, run `./start.sh` to build and run the project
  with Docker

# Directory Structure
```
.
├── app
|   ├── controllers
|   |   ├── users.go
|   ├── models
|   |   ├── common.go
|   |   ├── user.go
├── config
|   ├── db
|   |   ├── db.go
|   ├── router
|   |   ├── routes.go
├── middleware
|   ├── jwt.go
|   ├── log.go
├── utilities
|   ├── api.go
|   ├── common.go
|   ├── constants.go
|   ├── db.go
|   ├── err.go
|   ├── jwt.go
|   ├── log.go
├── app.go
├── Dockerfile
├── README.md
└── start.sh
```

# Directory/File Descriptions:
- app: Contains the Model and Controller layers for the over application MVC architecture.
- config: Contains independent services required to connect the application server to the frontend service through a router and the lowest layer for persistence store (DB)
- middleware: Contains the intermediate functionality that wraps each route
- utilities: Contains the disconnected components from the application to be used in any layer within the application for convenient functionality
- app.go: Contains the main server code to spin up the backend service
- Dockerfile: Specifies Docker configuration to support machine cross-compatibility and ease of setup
- README.md: Describes the backend component of the GradePortal portal
- start.sh: Starts the Golang project through Docker configuration
