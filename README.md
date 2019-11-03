## Microservice to store and manage statistics for companies

### How to run
Setup `cr_mysql_uri` MySQL connection URL to the environment. <br/>
It should look like `user:password@/dbname`. (If you need some help, you could read <a href="https://github.com/go-sql-driver/mysql">https://github.com/go-sql-driver/mysql</a>). <br/>
And finally, execute `go run server.go` from directory, where `server.go` is located

### How to run tests
They are testing Rest API, so you need to have a working server. <br/>
Firstly, setup the server by running `go run server.go` <br/>
Then in new terminal run `cd tests/(stats or company)` and `go test` 

### How to run benchmark
`cd tests/stats/benchmark` && 
`go test -bench . -count 10`