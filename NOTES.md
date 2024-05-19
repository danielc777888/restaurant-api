# NOTES

## For async tasks
- Use https://github.com/hibiken/asynq
- Rate sentiments

## Setup
- Install go, checkout solution
- Install docker
- Compose and run containers: `docker compose up -d`
- To restart the containers run: `docker compose restart`
- Resolve project dependencies: `go mod tidy`
- Build project: `go build ./...`
- Copy **.env_sample** to **.env** and set any private values.
- Reseed database: `go run reseed/main.go`
- Start api server: `go run api/main.go`
- Run all tests: `go test ./...`
- Browse here to read the api docs: `http://localhost:8080/swagger/index.html`

## TODO if enough time

### api perf
- db connection pool
- redis caching, daily eviction for restaurants + dishes
- hourly for ratings
- benchmark before and after

### prometheus for observability


### database
- add timestamps
- use db pool
- use transactions

## Tests
- property based tests with rapid

## Metrics
- run: `curl -v --silent http://localhost:8080/metrics 2>&1 | grep GET_/api/v1/restaurants`

### Swagger Docs
- To generate swagger docs run: `swag init -g api/*.go`

### References
- https://go.dev/
- https://github.com/lemoncode21/golang-crud-gin-gorm
- https://medium.com/readytowork-org/secure-your-go-web-application-jwt-authentication-e65a5af7c049
- https://gin-gonic.com/docs/
- https://gorm.io/
- https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
- https://github.com/redis/go-redis/
- https://aistudio.google.com/app/prompts/sentiment-analysis-chat