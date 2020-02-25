# Install

- Install Go
- Install Beego framework: `go get github.com/astaxie/beego`
- Install Bee tool: `go get github.com/beego/bee`

# Create project

`bee new beegoing`

# Run

1. `cd beegoing`
2. `bee run watchall`

# Tutorials

## Go

https://www.youtube.com/playlist?list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX

## Beego

- https://www.youtube.com/playlist?list=PLw3Nw6TNDAgUBG1Mh6fNacjXcxngYtM48 
- https://beego.me/docs/quickstart/new.md 
- https://github.com/astaxie/build-web-application-with-golang

# Scaffolding from DB

### MySQL

1. Install ORM `go get github.com/astaxie/beego/orm`
2. Install driver `go get github.com/go-sql-driver/mysql`
3. Make DB
```sql
CREATE DATABASE beegoing;
USE beegoing;
CREATE TABLE `places` (
	`id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` varchar(50) NOT NULL,
	`review` varchar(200) NOT NULL,
	PRIMARY KEY (id)
);
```
4. In `~GOPATH/src` directory: `bee api beegoing -conn=username:password@tcp\(127.0.0.1:3306\)/beegoing`
5. `bee run watchall`
6. Visit: http://localhost:8082/v1/places

### PostgreSQL

1. Install ORM `go get github.com/astaxie/beego/orm`
2. Install driver `go get github.com/lib/pq`
3. Make DB 
```sql
CREATE DATABASE beegoing;
\c beegoing;
CREATE TABLE places (
	id      BIGSERIAL PRIMARY KEY,
	name    VARCHAR(100) NOT NULL,
	review  VARCHAR(100) NOT NULL
);
```
4. In `~GOPATH/src` directory: `bee api beegoing -driver=postgres -conn="user=username password=password host=localhost dbname=beegoing sslmode=disable"`
5. `bee run watchall`
6. Visit: http://localhost:8082/v1/places

## Relevant tutorials

- https://beego.me/docs/mvc/model/overview.md
- https://beego.me/docs/mvc/model/orm.md
