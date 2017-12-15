# GradePortal Backend

This is the entire backend for GradePortal written in Golang. It is a pure RESTful API backend, serving data through JSON requests made by the frontend.

# Setup
- Install [PostgreSQL](http://postgresguide.com/setup/install.html)
- Create a database with DB_NAME in PostgreSQL
- Create the following environment variables
	- DB_USER: The user that is given access to the database
	- DB_PORT: The port that PostgreSQL runs on
	- DB_PASSWORD: The password needed to connect to the database
	- DB_HOST: The host machine (usually localhost) the database runs on
	- DB_NAME: The name of the created database to connect to
- Run `./api` to run the pre-built server executable

To build and run on your own:
	- Install [Golang](https://golang.org/doc/install)
	- Run `go get ./` to fetch all the go dependencies
	- Run `go build` to build the executable
	- Run `./api` to run the newly built executable

Note that running `./start.sh` isn't enough on a local machine because of the necessary connection needed between the api server and the database.

# Directory Structure
```
.
├── app
|   ├── controllers
|   |   ├── assignments.go
|   |   ├── classes.go
|   |   ├── submissions.go
|   |   └── users.go
|   ├── models
|   |   ├── assignment.go
|   |   ├── class.go
|   |   ├── common.go
|   |   ├── enrolled.go
|   |   ├── submission.go
|   |   └── user.go
├── config
|   ├── db
|   |   └── db.go
|   ├── router
|   |   └── routes.go
├── middleware
|   ├── jwt.go
|   └── log.go
├── utilities
|   ├── api.go
|   ├── common.go
|   ├── constants.go
|   ├── db.go
|   ├── err.go
|   ├── jwt.go
|   ├── lang.go
|   └── log.go
├── tests 
|   ├── files
|   |   ├── grade.sh
|   |   └── correct_submission.cpp
|   ├── assignment_model_test.go
|   ├── class_model_test.go
|   ├── submission_model_test.go
|   └── user_model_test.go
├── docs
|   ├── doc
|   ├── lib
|   ├── pkg
|   |   ├── github.com
|   |   |   ├── jcherianucla
|   |   |   |   ├── gradeportal
|   |   |   |   |   └── index.html
├── app.go
├── deploy.sh
├── Makefile
├── Dockerfile
├── docs.sh
├── Procfile
├── README.md
└── start.sh
```

# Directory/File Descriptions:
- app: Contains the Model and Controller layers for the over application MVC architecture.
- config: Contains independent services required to connect the application server to the frontend service through a router and the lowest layer for persistence store (DB)
- middleware: Contains the intermediate functionality that wraps each route
- utilities: Contains the disconnected components from the application to be used in any layer within the application for convenient functionality
- tests: Contains all the tests for the application.
- docs: Contains documentation for the backend code generated with godoc
- app.go: Contains the main server code to spin up the backend service
- deploy.sh: Builds the docker container and pushes to the Heroku Registry for the grade-portal-api application
- Makefile: Contains build system for running and testing
- Dockerfile: Specifies Docker configuration to support machine cross-compatibility, used primarily for deployment to Heroku with capability to run cross-compilation for gcc within container
- docs.sh: Recursively downloads webpages from a locally running godoc server
- Procfile: Defines what process to run as a resource on the Heroku application
- README.md: Describes the backend component of the GradePortal portal
- start.sh: Starts the Golang project through Docker configuration. Note that this is currently not useful locally as it has no method to connect to the PostgreSQL database

# Tests:

Run tests by running `make test`.

# Endpoints:

Find full API documentation in the [GradePortal Wiki](https://github.com/jcherianucla/UCLA-CS-130/wiki/)
