# GradePortal Frontend

This is the entire frontend for GradePortal written with React.js

# Setup
- Install docker
- Inside the /frontend directory, run `./start.sh` to build and run the project
  with Docker

# Directory Structure
.
├── build
├── config
├── node_modules
├── scripts
├── src
|   ├── containers
|   |   ├── professor
|   |   |   ├── Analytics.js
|   |   |   ├── UpsertClass.js
|   |   |   └── UpsertProject.js
|   |   ├── student
|   |   |   └── Submission.js
|   |   ├── Login.js
|   |   ├── Classes.js
|   |   ├── Landing.js
|   |   └── Projects.js
|   ├── fonts
|   ├── images
|   ├── shared
|   |   ├── Content.js
|   |   ├── Header.js
|   |   ├── ItemCard.js
|   |   └── SidePanel.js
|   ├── styles
|   ├── utils
|   └── index.js
├── static
├── .dockerignore
├── docker-compose.yml
├── index.html
├── package.json
├── README.md
└── start.sh

# Directory/File Descriptions:
- build: Contains build configuration data for project creation
- config: Contains configuration details for development and production environments
- node_modules: Contains all external NPM packages that are used within the project
- public: Contains public assets that can be served individually from the server
- scripts: Contains scripts to build, start, and test the React project
- src: Contains most frontend application content, including React components, stylsheets, images, fonts, and utilities
- static: Auto-generated static files to provide utility to React services
- .dockerignore: Specifies files to ignore when running the project in Docker
- docker-compose.yml: Specifies directories that should be mounted to the Docker container to support hot-reloading of code through the server
- Dockerfile: Specifies Docker configuration to support machine cross-compatibility and ease of setup
- index.html: Provides the root HTML structure of the React project
- README.md: Describes the frontend component of the GradePortal project
- start.sh: Starts the React project through Docker configuration