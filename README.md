# Todo List Application
<p align="left">
<a href="https://github.com/maurana"><img src="https://img.shields.io/badge/creator-%40thisismaulanaa-blueviolet" alt="Creator"></a>
<a href="https://echo.labstack.com/"><img src="https://img.shields.io/badge/echo v4-important" alt="Framework"></a>
<a href="https://go.dev/"><img src="https://img.shields.io/badge/Go v1.18-yellowgreen" alt="Programming"></a>
</p>

> Build & Run procedure

Build
```bash
> go build
```

Starting the server with PostgreSQL connection
```bash
> go run main.go start-postgres
```

Starting the server with MySQL connection
```bash
> go run main.go start-mysql
```

Migrate with PostgreSQL
```bash
> go run main.go migrate-postgres
```

Migrate with MySQL
```bash
> go run main.go migrate-mysql
```

Launch server with the migrated PostgreSQL
```bash
> go run main.go lauch-postgres
```

Launch server with the migrated MySQL
```bash
> go run main.go migrate-mysql
```

Unit test Coverage, result = 66.7%
```bash
> go test -cover
```

> Business specifications
* `List`
- Display all data list (include pagination, filter[Search by: title, description], sort)
- Display data list by id
- Create data list with single/multiple upload file (validation type: pdf & txt) or without upload
- Update data list
- Delete data list by id

* `Sublist`
- Display all data sublist (include pagination, filter[Search by: title, description], sort)
- Display data sublist by id
- Create data sublist with single/multiple upload file (validation type: pdf & txt) or without upload
- Update data sublist
- Delete data sublist by id

> Technical specifications
- Echo Framework Go Language
- Built in clean architecture & solid principle
- Unit Test coverage = 66.7%
- API Spesification (Postman Collection & Swagger)
- Capability to change DB Engine (PostgreSQL, MySQL)
- Capability to change File Storage (AWS, Azure, Google Cloud, Hadoop Storage)
