# NOTES


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
- Browse here to read the api docs and try it out: `http://localhost:8080/swagger/index.html`
- Browse here to view prometheus metrics: `http://localhost:8080/metrics`
- To see some curl examples see **CURL.md**
- To get a GEMINI API KEY go here: `https://ai.google.dev/gemini-api/docs/api-key`

## Swagger Docs
- To re-generate swagger docs run: `swag init -g api/*.go`

## References
- https://go.dev/
- https://github.com/lemoncode21/golang-crud-gin-gorm
- https://medium.com/readytowork-org/secure-your-go-web-application-jwt-authentication-e65a5af7c049
- https://gin-gonic.com/docs/
- https://gorm.io/
- https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
- https://github.com/redis/go-redis/
- https://aistudio.google.com/app/prompts/sentiment-analysis-chat