# Tasks Application - WIP
This is a task management application built with Go and PostgreSQL.
For now I'm working on the REST API, but it will contain a CLI as well.

**Work in Progress !!!**

## Configuration Setup
1. Copy the `config.example.yaml` file and rename it to `config.yaml`.
2. Update the configuration values in the `config.yaml` file as needed.
3. [*Optional*] You can set environment variables to override the values in the `config.yaml` file.
   For example, to override the server port, first you need to add the `TASKS` prefix, then you
   can set the `TASKS_SERVER_PORT` environment variable.
   ```bash
   export TASKS_SERVER_PORT=3000 # or TASKS_SERVER_PORT=3000 make run
   ```
4. [*Optional*] You can leave the config file in the root directory or move it inside the
`internal/config/` folder.

### Configuration Default Values
In the abssence of any configuration field, the application will use the following default values:

| Config Key          | Default Value        | Description                       |
|---------------------|----------------------|-----------------------------------|
| SERVER_HOST         | localhost            | Server host                       |
| SERVER_PORT         | 8080                 | Port on which the server runs     |
| DB_HOST             | localhost            | Database host                     |
| DB_PORT             | 5432                 | Database port                     |
| DB_USER             | admin                | Database user                     |
| DB_PASSWORD         | 123456               | Database password                 |
| DB_NAME             | tasks_db             | Database name                     |
| DB_SSLMODE          | disable              | Database SSL mode                 |

## Docker - Containerization
### Database Setup
1. You can install PostgreSQL on your machine and setup the access manually. But it's better use
[**Docker**](https://www.docker.com/) by running the following command:
    ```bash
      docker run -d --name postgres \
        -e POSTGRES_USER=admin \
        -e POSTGRES_PASSWORD=123456 \
        -e POSTGRES_DB=tasks_db \
        -p 5432:5432 \
        postgres:alpine
    ```
      *Remember that you can change the user, password, and database name as you prefer.*
2. Update the database connection strings in the `Makefile` if necessary.
3. You must have `go-migrate` installed. You can find the installation instructions
[**here**](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation).
4. Run the following command to apply the migrations:
   ```bash
   make migrate-local-up
   ```
5. To rollback all the migration and revert the database to its initial state, run:
   ```bash
   make migrate-local-down
   ```
6. [*Optional*] You can check the `Makefile` for other available database related commands.

### Application Setup
If you want to run this application locally, make sure to setup your database (and app's configuration)
properly, use the `Dockerfile` to build the image. You can specify an exposed port using the `SERVER_PORT`
build argument, the default one is **8080**.

See the command bellow:
```bash
docker build [--build-args SERVER_PORT=8080] -t mytaghere:myversion .
```

I've also provided a ready-to-use image, you can download it from the
[Docker Hub](https://hub.docker.com/repository/docker/victordevv/tasks/general), by running:
```bash
docker run --name="tasks-app" -p 8080:8080 victordevv/tasks:lastest 
```

It will download the image and start the container bound to the default port 8080 on your machine.

Access *http://localhost:8080/health* on your browser to check if the application is running healthy.
It will return a basic *Ok* status.

## Tools
- [Go](https://golang.org/) - Programming language
- [PostgreSQL](https://www.postgresql.org/) - Database
- [go-migrate](https://github.com/golang-migrate/migrate) - Database Migration tool
- [Docker](https://www.docker.com/) - Containerization platform
- [Make](https://www.gnu.org/software/make/) - Build automation tool

## Third-Party Libraries
- [viper](https://github.com/spf13/viper) - Configuration management
- [fiber](https://gofiber.io/) - Web framework

---
<small>Authored by **Victor Ferreira**.</small>