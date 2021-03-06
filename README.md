# User Management System
## By: Dicky Novanto

## How to install the project
##### I. User Management System project Installation
3. Install MySQL server
4. Install Redis: create new Redis container by `docker-compose up -d`
5. Generate 10 millions of user data and insert all the data into table (see section III)
6. For generating the user data, python script is used. 
    - Hash algorithm for password: pbkdf2.
    - To get pbkdf2 hash algorithm on Go: `go get -u golang.org/x/crypto/pbkdf2`
7. For importing the generated user data into SQL into test_user user and test database to table User:
    - `mysql --local-infile test -u test_user -p`, with password: `test_user`

    - `LOAD DATA LOCAL INFILE 'user.txt' INTO TABLE User;`
8. Clone this project’s repository
11. Setup go driver sql: `go get -u github.com/go-sql-driver/mysql`
12. Setup profiling library: `go get -u github.com/pkg/profile`
15. Build and run these following servers:
    - picture server (in `pictureserver` directory)
    - HTTP server (in `web` directory)
    - TCP server (in `tcpserver` directory)
16. How to build and run each server:
    - Go to each server directory and enter these commands:
        - `go install` (this will create an executable binary file with the name of the server's directory and place the binary file to GOBIN)
        - `directory name`
17. Browse: `localhost:8080` to show home page, and the system can be used

##### II. WRK (HTTP benchmarking tool) Installation:
1. Install WRK: `brew install wrk`
2. Refer to section `How to run performance test` to run WRK
3. Details of how to use WRK is provided in the README of WRK Repository (https://github.com/wg/wrk)

##### III. Generating 10 millions of user data
1. Go to folder `scripts`
2. Run python script with command: `python generateUser.py`

## How to start User Management System
1. Run TCP server first by entering command: `tcpserver`
2. Run Picture server with command: `pictureserver`
3. Run HTTP server with command: `httpserver`

## How to run performance test
1. Go to scripts directory, open terminal
2. Enter this command: `wrk -t100 -c1000 -d10s -s post.lua http://localhost:8080/authenticate `
    - WRK will be running with the number of thread = 100, the number of keep-alive connection = 1000, duration = 10 seconds and running script post.lua that exist inside the directory, and the endpoint that is tested is http://localhost:8080/authenticate
3. The result provides many information about the performance test, for example the number of request per second.