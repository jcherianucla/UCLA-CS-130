# GradePortal Frontend

This is the entire frontend for GradePortal written with React.js

# Setup
- Install [Docker](https://docs.docker.com/engine/installation/)

- Inside the /frontend directory, run `./start.sh` to build and run the project
  with Docker

# Directory Structure
```
.
├── build
├── config
├── docs
|   └── DOCUMENTATION.md
├── node_modules
├── scripts
├── src
|   ├── containers
|   |   ├── professor
|   |   |   ├── UpsertClass.js
|   |   |   └── UpsertProject.js
|   |   ├── Login.js
|   |   ├── Classes.js
|   |   ├── FAQ.js
|   |   ├── Landing.js
|   |   ├── Project.js
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
```

# Directory/File Descriptions:
- build: Contains build configuration data for project creation
- config: Contains configuration details for development and production environments
- docs: Contains documentation for the project
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

