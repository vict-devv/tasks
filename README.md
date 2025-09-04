# Tasks Application - WIP
This is a task management application built with Go and PostgreSQL.

## Database Setup
1. Install PostgreSQL on your machine and setup the access manually.
Or use [**Docker**](https://www.docker.com/) by running the following command:
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
[here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation).
4. Run the following command to apply the migrations:
   ```bash
   make migrate-local-up
   ```
5. To rollback all the migration and revert the database to its initial state, run:
   ```bash
   make migrate-local-down
   ```
6. [Optional] You can check the `Makefile` for other available database related commands.

## Tools
- [Go](https://golang.org/) - Programming language
- [PostgreSQL](https://www.postgresql.org/) - Database
- [go-migrate](https://github.com/golang-migrate/migrate) - Database Migration tool

---
<small>Authored by **Victor Ferreira**.</small>