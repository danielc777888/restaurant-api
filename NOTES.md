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
- use db pool, common object
- use transactions

### tests
- property based tests with rapid

### swagger docs
- swag init -g api/restaurant.go