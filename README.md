## Microservice to store and manage statistics for companies

### How to run
#### Example of config.yaml file
```yaml
authToken: "5672139asdaw"
port: ":8080"
timeLayout: "2006-01-02 15:04:05"
dateLayout: "2006-01-02"
mysqlURL: "cobrareviews:password@/cobrareviews"
```
Add `config.yaml` to `config` directory. <br/>
(If you need some help with MySQL connection URL, you could read <a href="https://github.com/go-sql-driver/mysql">https://github.com/go-sql-driver/mysql</a>). <br/>
And finally, execute `go run server.go` from directory, where `server.go` is located

### How to run tests
They are testing Rest API, so you need to have a working server. <br/>
Firstly, setup the server by running `go run server.go` <br/>
Then in new terminal run `cd tests/(stats or company)` and `go test` 

### How to run benchmark
`cd tests/stats/benchmark` && 
`go test -bench . -count 10`