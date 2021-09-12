# Search Movie Service

## Overview
Search Movie Service.

### High Level Features
* Search through OMDB api
* Log all the access search to postgres db

### Compiling Requirements
* Docker with support for Linux containers
* Docker compose
* GNU Make
* Internet connection

### Development Requirements
* See _Compiling Requirements_
* Go IDE (we like GoLand)
* Go compiler (latest)
* Docker Compose

### Building
* make build - Create project binaries and production docker image
* make test - Run the unit test of the service
* make run - Run service locally for deployment. Please make sure that you perform _make build_ in order to install dependencies and create project binaries

### Initial Database Setup
* DB Migration automatically at deployment.

### Applying DB Schema Changes
* Make changes in database model on model package ({project_dir}/pkg/model)
* Deploy the service
* Migration automatically at service deployed

### Compilation Options
There are compile-time options that can be set via environment variables:

* REVISION_ID

  Revision of the build, e.g. Git SHA (default: latest)

### Dependencies
* PostgreSQL - GoServlet will save data in to PostgreSQL database

### Deploying
#### Environment Variables
* OMDB_KEY

  The OMDB API Key

* DB_HOST

  Database host to connect to database

* DB_PORT

  Database port

* DB_NAME

  Database name

* DB_USERNAME

  Username to connect to database

* DB_PASSWORD

  Password to connect to database

#### Optional Environment Variable
* SERVER_HTTP_PORT

  The http port which the service will listen to (default: 8085)

* SERVER_GRPC_PORT

  The grpc port which the service will listen to (default: 9080)

* DB_SSL_ENABLED

  Use SSL for database connection (default: false)


## API
### Search

`[GET] /search`

#### Parameters
`{
    "pagination": int64,
    "searchWord": string
}`

### Health Check

`[GET] /healthz`

#### Parameters
`{ }`