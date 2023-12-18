# GoFr Money Tracker

RESTful APIs created for serving a money tracker web application using GoFr.

## Setup and Installation

1. Download and install [Go](https://go.dev/doc/install).
2. Download [Docker](https://www.docker.com/products/docker-desktop/) and set it as a background service.
3. Start a mysql docker container with the following command: `docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30`
4. Connect to the bash of the created terminal:  `docker exec -it gofr-mysql bash`
5. Create `user` and `record` tables with the following schema:
   
   ~~~
   create table user(email varchar(255) NOT NULL);
   ~~~
   ~~~
   create table record(
   id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
   email varchar(255) NOT NULL, date varchar(255) NOT NULL,
   amount DECIMAL(10, 2) NOT NULL,
   description varchar(255)); 
   ~~~
6. Create a file with name `.env` inside `configs` folder with the following content:
   ~~~
    APP_NAME=gofr-moneytracker
    HTTP_PORT=8080
    
    DB_HOST=localhost
    DB_USER=root
    DB_PASSWORD=root123
    DB_NAME=test_db
    DB_PORT=3306
    DB_DIALECT=mysql
   ~~~
7. Execute the commands inside our repository's directory in order to install the dependencies and run the project:
   ~~~
   go mod tidy
   ~~~
   ~~~
   go run main.go
   ~~~

## API Reference

You can view and test the APIs using [this](https://www.postman.com/solar-station-253785/workspace/zopsmart-gofr/collection/14034300-f7f969d5-a5ab-4f51-845d-a2d622450375?action=share&creator=14034300) Postman collection or can import the collection using `/ZopSmart.postman_collection.json` inside Postman.
