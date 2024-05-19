# NOTES

## Setup

## Tests

## API Docs




## Plan
- Wed : Task 1
- Thur : Task 2
- Fri : Task 3
- Sat : Task 4
- Sun : Cleanup, documentation ,testing, iteration 2

- First get all functionality to work
- Then tests
- Then cleanup (refactor, documentation)

## Sentiment Analysis
- Use LLM Gemini
- https://aistudio.google.com/app/prompts/sentiment-analysis-chat
- Use google api


## For async tasks
- Use https://github.com/hibiken/asynq
- Rate sentiments

## Setup
- Install go, checkout solution
- Install docker
- Compose and run containers: `docker compose up -d`
- To restart the containers run: `docker compose restart`
- Resolve project dependencies: `go mod tidy`
- Build packages: `go build ./...`
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

### config
- use .env for env configs
- database
- gemini api key

### database
- using UUID instead of ints
- add timestamps
- use db pool
- use transactions

### tests
- property based tests with rapid

### metrics
- run: `curl -v --silent http://localhost:8080/metrics 2>&1 | grep GET_/api/v1/restaurants`

### swagger docs
- swag init -g api/restaurant.go

### readings
- addresses/pointers/value/copy in go
- system design googl cloud designs

### References
- https://go.dev/
- https://github.com/lemoncode21/golang-crud-gin-gorm
- https://medium.com/readytowork-org/secure-your-go-web-application-jwt-authentication-e65a5af7c049
- https://gin-gonic.com/docs/
- https://gorm.io/
- https://cacm.acm.org/research/the-go-programming-language-and-environment/
- https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
- https://github.com/redis/go-redis/