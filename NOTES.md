# NOTES

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

## Property based testig Rapid

## For async tasks
- Use https://github.com/hibiken/asynq
- Rate sentiments

## Setup
- Install go, checkout solution
- Install docker
- Compose and run containers: `docker compose up -d`
- Run unit tests: `go run test/unit/main.go`
- Reseed database: `go run reseed/main.go`
- Start api server: `go run api/main.go`
- Run integration tests: `go run test/unit/main.go`