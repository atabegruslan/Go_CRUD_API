This is the backend part of https://github.com/atabegruslan/beegoing_front

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
5. `bee run`
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
5. `bee run`
6. Visit: http://localhost:8082/v1/places

## Relevant tutorials

- https://beego.me/docs/mvc/model/overview.md
- https://beego.me/docs/mvc/model/orm.md

# Generate Swagger Documentation

1. `bee run -downdoc=true -gendoc=true`
2. Visit http://localhost:8082/swagger/
