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
|   |   └── users.go
|   ├── models
|   |   ├── common.go
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
|   └── log.go
├── tests 
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
├── Makefile
├── Dockerfile
├── docs.sh
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
- Makefile: Contains build system for running and testing
- Dockerfile: Specifies Docker configuration to support machine cross-compatibility and ease of setup
- docs.sh: Recursively downloads webpages from a locally running godoc server
- README.md: Describes the backend component of the GradePortal portal
- start.sh: Starts the Golang project through Docker configuration

# Tests:
The tests that were created are located in the ./tests directory.

## TestUserInsert:

Tests the Insert method which is used to add users to the database. The Insert function takes a User object which has an ID, first and last name, email, password, check if it is a professor. Once the Insert function receives a User object, it checks to see if all the fields are non-empty and valid (proper email etc.) and also checks if the User already exists in the database. If all is well, then the User is added to the database. Upon success, it returns a User and no error. If there is something wrong, it returns no User and an error describing what went wrong.

Our test function will have no output if it passes. If it fails, the command line output is shown by the Failure indicator.

## TestUserLogin:

Tests the Login method which allows Users to login provided they are already in the database. Login fails if the provided email or password are not in the database or do not match.

Our test function will have no output if it passes. If it fails, the command line output is shown by the Failure indicator.

## TestUserGet:

Tests the Get function which provides a general search function to search the User table. It returns a list of users that match the query and also an error if something goes wrong.

TestUsersGet tests for an empty query which when called with Get will return an error, “Get query preparation failed”. This error is checked for and if it is returned, we know the test is working and there is no output from this function. If the Get function isn’t able to catch that error, there will be command line output of FAIL: TestUsersGet “Errors don’t match up”.

## TestUserUpdate:

This tests the Update function which takes the ID of a user and the updated User object and performs the valid updates on the database. 

In addition, Update also has error-checks for when a query fails to be executed or if a duplicate user (same ID) is added to the table and thus avoids redundant records.

## TestUserDelete:

Tests the Delete function which deletes a record based on user ID and returns an error if something went wrong. 

We plan on expanding this testing from covering usage-based test scenarios to include internal error reporting checks. 

